package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ST359/rest-api-todo/internal/api"
)

func main() {
	api, err := api.New()
	if err != nil {
		panic(err)
	}
	go api.Run()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	_ = <-c
	fmt.Println("Gracefully shutiing down")
	_ = api.Shutdown()
	fmt.Println("Shut down succesfully")
}
