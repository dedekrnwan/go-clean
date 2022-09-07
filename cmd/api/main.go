package main

import (
	"github.com/dedekrnwan/go-clean/internal/config"
	"github.com/dedekrnwan/go-clean/internal/contract"
	"github.com/dedekrnwan/go-clean/internal/deliveries/api"
	"github.com/dedekrnwan/go-clean/pkg/utils/env"
	"github.com/dedekrnwan/go-clean/pkg/utils/graceful"

	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

func init() {
	if os.Getenv("ENV") == "" {
		env := env.NewEnv()
		env.Load()
	}

	err := config.Load("")
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func main() {
	//initialize all needs (db conn, adapter outbound)
	c := contract.New()

	//api
	restInstance := api.New(c)
	starterEcho, stopperEcho := restInstance.PrepareEcho()

	//grpc

	wg := new(sync.WaitGroup)

	wg.Add(2)
	go func() {
		graceful.StartProcessAtBackground(starterEcho)
		graceful.StopProcessAtBackground(time.Second*10, stopperEcho)
		wg.Done()
	}()

	wg.Wait()
}
