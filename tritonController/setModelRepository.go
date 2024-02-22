package tritonController

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/ahr-i/triton-agent/setting"
	"github.com/ahr-i/triton-agent/src/logCtrlr"
	"github.com/ahr-i/triton-agent/tritonCommunicator"
	"golang.org/x/crypto/ssh"
)

func SetModelRepository(provider string, model string, version string) error {
	log.Println("========== Model Repository ==========")
	log.Println("Model repository:", provider)
	modelRepositoryPath := fmt.Sprintf("%s/%s", "/models", provider)
	makeFolder(modelRepositoryPath)

	config := &ssh.ClientConfig{
		User: setting.TritonUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(setting.TritonPassword),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	logCtrlr.Log("Attach Triton server.")
	client, err := ssh.Dial("tcp", setting.TritonSSH, config)
	if err != nil {
		return err
	}
	defer client.Close()

	// pkill 명령 실행
	logCtrlr.Log("Kill Triton server.")
	if err := executeCommand(client, "pkill -f /opt/tritonserver"); err != nil {
		return err
	}

	// tritonserver 시작 명령 실행
	logCtrlr.Log("Change the model repository and start the Triton server.")
	startCommand := fmt.Sprintf("nohup /opt/tritonserver/bin/tritonserver --model-repository %s > /dev/null 2>&1 & exit", modelRepositoryPath)
	if err := executeCommand(client, startCommand); err != nil {
		return err
	}

	if err := polling(model, version); err != nil {
		return err
	}

	return nil
}

func executeCommand(client *ssh.Client, command string) error {
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	session.Stderr = &b

	if err := session.Run(command); err != nil {
		return fmt.Errorf("failed to run command: %s, output: %s", err, b.String())
	}

	fmt.Println("Command output:", b.String())
	return nil
}

func polling(model string, version string) error {
	logCtrlr.Log("Polling start - Model repository.")
	time.Sleep(5 * time.Second)

	for {
		ready, err := tritonCommunicator.Ready(model, version)
		if err != nil {
			return err
		}

		if ready {
			break
		}
		time.Sleep(2 * time.Second)
	}

	return nil
}
