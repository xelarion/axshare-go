package task

import (
	"fmt"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
)

var MachineryServer *machinery.Server

const configPath = "configs/machinery_config.yaml"

func startMachineryServer() (*machinery.Server, error) {
	cnf, err := config.NewFromYaml(configPath, true)

	//init server
	server, err := machinery.NewServer(cnf)
	if err != nil {
		panic(err)
	}

	err = server.RegisterTasks(registerTasks)

	return server, err
}

func RunMachineryServer() {
	consumerTag := "machinery_worker"

	server, err := startMachineryServer()
	if err != nil {
		panic(err)
	}
	MachineryServer = server

	worker := server.NewWorker(consumerTag, 0)
	err = worker.Launch()
	if err != nil {
		panic(err)
	}
}

func Send(taskSignature *tasks.Signature) error {
	_, err := MachineryServer.SendTask(taskSignature)
	if err != nil {
		panic(fmt.Errorf("Could not send task: %s \n", err.Error()))
	}
	return nil
}
