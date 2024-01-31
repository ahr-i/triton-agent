package main

import (
	"net/http"

	"github.com/ahr-i/triton-agent/handler"
	"github.com/ahr-i/triton-agent/setting"
	"github.com/ahr-i/triton-agent/src/corsController"
	"github.com/ahr-i/triton-agent/src/p2p"
	"github.com/urfave/negroni"
)

func Init() {
	go p2p.Enter()
}

func main() {
	Init()

	mux := handler.CreateHandler()
	handler := negroni.Classic()

	//defer mux.Close()

	handler.Use(corsController.SetCors("*", "GET, POST, PUT, DELETE", "*", true))
	handler.UseHandler(mux)

	// HTTP Server Start
	http.ListenAndServe(":"+setting.ServerPort, handler)
}
