package main

import (
	"net/http"
	"time"

	"github.com/ahr-i/triton-agent/handler"
	"github.com/ahr-i/triton-agent/schedulerCommunicator/healthPinger"
	"github.com/ahr-i/triton-agent/setting"
	"github.com/ahr-i/triton-agent/src/corsController"
	"github.com/ahr-i/triton-agent/src/logCtrlr"
	"github.com/ahr-i/triton-agent/tritonController"
	"github.com/urfave/negroni"
)

func initialization() {
	logCtrlr.Log("Initialize the agent.")

	tritonController.Init(setting.ModelRepository)

	if setting.ManagerActive {
		logCtrlr.Log("Use manager.")
		go healthPinger.Enter()
	}
}

func startServer() {
	mux := handler.CreateHandler()
	handler := negroni.Classic()
	defer mux.Close()

	handler.Use(corsController.SetCors("*", "GET, POST, PUT, DELETE", "*", true))
	handler.UseHandler(mux)

	// HTTP Server Start
	logCtrlr.Log("HTTP server start.")
	http.ListenAndServe(":"+setting.ServerPort, handler)
}

func logs() {
	logCtrlr.Log("Your models:")
	tritonController.PrintModelRepository()
}

func main() {
	initialization()

	logs()

	go test()

	startServer()
}

func test() {
	time.Sleep(time.Millisecond * 1000)

	type testModel struct {
		Provider string
		Name     string
		Version  string
	}

	testModels := []testModel{
		{"meta", "Llama-2-7B-Chat", "1"},
	}

	for _, model := range testModels {
		healthPinger.UpdateModel(model.Provider, model.Name, model.Version)
	}
}
