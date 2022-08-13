package main

import (
	"github.com/dedekrnwan/go-clean/config"
	"github.com/dedekrnwan/go-clean/internal/deliveries/rest"
	"github.com/dedekrnwan/go-clean/internal/factory"
	"github.com/dedekrnwan/go-clean/pkg/utils"

	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

func init() {
	if os.Getenv("ENV") == "" {
		env := utils.NewEnv()
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
	f := factory.NewFactory()

	//rest
	restInstance := rest.New(f)
	starterEcho, stopperEcho := restInstance.PrepareEcho()

	//grpc

	wg := new(sync.WaitGroup)

	wg.Add(2)
	go func() {
		utils.StartProcessAtBackground(starterEcho)
		utils.GracefullStopProcessAtBackground(time.Second*10, stopperEcho)
		wg.Done()
	}()

	wg.Wait()
}
