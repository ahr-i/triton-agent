package main

import (
	"log"
	"net/http"

	"github.com/ahr-i/triton-agent/handler"
	"github.com/ahr-i/triton-agent/schedulerCommunicator/healthPinger"
	"github.com/ahr-i/triton-agent/schedulerCommunicator/taskTokenManager"
	"github.com/ahr-i/triton-agent/setting"
	"github.com/ahr-i/triton-agent/src/corsController"
	"github.com/urfave/negroni"
)

func Init() {
	log.Println("* (System) Initialize the agent.")

	taskTokenManager.Init()

	go healthPinger.Enter()
}

func main() {
	Init()

	mux := handler.CreateHandler()
	handler := negroni.Classic()

	//defer mux.Close()

	handler.Use(corsController.SetCors("*", "GET, POST, PUT, DELETE", "*", true))
	handler.UseHandler(mux)

	// HTTP Server Start
	log.Println("* (System) HTTP server start.")
	http.ListenAndServe(":"+setting.ServerPort, handler)
}
