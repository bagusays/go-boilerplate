package main

import (
	"fmt"
	"go-boilerplate/cmd"
	"go-boilerplate/shared/config"
)

func main() {
	cfg := config.New("./configs/")

	isReady := make(chan bool)
	isShutdown := make(chan bool)
	server := cmd.NewServer(cmd.ServerDependencies{
		IsReady:    isReady,
		IsShutdown: isShutdown,
		Config:     cfg,
	})
	go server.StartServer()
	<-isReady
	<-isShutdown

	fmt.Println("Service stopped")
}
