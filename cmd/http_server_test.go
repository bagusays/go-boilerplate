package cmd_test

import (
	"fmt"
	"go-boilerplate/cmd"
	"go-boilerplate/shared/config"
	"os"
	"syscall"
	"testing"
)

func TestHttpServer(t *testing.T) {
	cfg := config.New("../configs/")

	isReady := make(chan bool)
	isShutdown := make(chan bool)
	server := cmd.NewServer(cmd.ServerDependencies{
		IsReady:    isReady,
		IsShutdown: isShutdown,
		Config:     cfg,
	})
	go server.StartServer()
	<-isReady

	p, _ := os.FindProcess(syscall.Getpid())
	p.Signal(syscall.SIGINT)

	<-isShutdown

	fmt.Println("Service stopped")
}
