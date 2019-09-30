package task

import (
	"fmt"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
)

var MachineryServer *machinery.Server

func startMachineryServer() (*machinery.Server, error) {
	var cnf = &config.Config{
		Broker:        "amqp://guest:guest@localhost:5672/",
		DefaultQueue:  "machinery_tasks",
		ResultBackend: "redis://127.0.0.1:6379",
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

	// Register tasks
	tasks := map[string]interface{}{
		"add":         Add,
		"hello_world": HelloWorld,
	}

	err = server.RegisterTasks(tasks)

	return server, err
}

func RunMachineryServer() error {
	consumerTag := "machinery_worker"

	server, err := startMachineryServer()
	if err != nil {
		return err
	}
	MachineryServer = server

	worker := server.NewWorker(consumerTag, 0)
	return worker.Launch()
}

func Send(taskSignature *tasks.Signature) error {
	_, err := MachineryServer.SendTask(taskSignature)
	if err != nil {
		return fmt.Errorf("Could not send task: %s", err.Error())
	}
	return nil
}
