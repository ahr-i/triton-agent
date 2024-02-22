package main

import (
	"net/http"

	"github.com/ahr-i/triton-agent/handler"
	"github.com/ahr-i/triton-agent/schedulerCommunicator/healthPinger"
	"github.com/ahr-i/triton-agent/setting"
	"github.com/ahr-i/triton-agent/src/corsController"
	"github.com/ahr-i/triton-agent/src/logCtrlr"
	"github.com/ahr-i/triton-agent/tritonController"
	"github.com/urfave/negroni"
)

func Init() {
	logCtrlr.Log("Initialize the agent.")

	tritonController.Init(setting.ModelRepository)

	if setting.SchedulerActive {
		logCtrlr.Log("Use scheduler.")
		go healthPinger.Enter()
	}
}

func main() {
	Init()

	mux := handler.CreateHandler()
	handler := negroni.Classic()

	//defer mux.Close()

	handler.Use(corsController.SetCors("*", "GET, POST, PUT, DELETE", "*", true))
	handler.UseHandler(mux)

	// HTTP Server Start
	logCtrlr.Log("HTTP server start.")
	http.ListenAndServe(":"+setting.ServerPort, handler)
}
