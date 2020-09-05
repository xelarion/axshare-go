package task

import (
	"fmt"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
	"os"
)

var MachineryServer *machinery.Server

func startMachineryServer() (*machinery.Server, error) {
	var cnf = &config.Config{
		Broker:        os.Getenv("REDIS_URL"),
		DefaultQueue:  "machinery_tasks",
		ResultBackend: os.Getenv("REDIS_URL"),
		AMQP: &config.AMQPConfig{
			Exchange:      "machinery_exchange",
			ExchangeType:  "direct",
			BindingKey:    "machinery_task",
			PrefetchCount: 3,
		},
	}

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
