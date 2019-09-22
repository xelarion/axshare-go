package utils

import (
	"fmt"
	"os/exec"
)

func RunCommand(name string, arg ...string) (err error) {
	command := exec.Command(name, arg...)
	err = command.Run()
	if err != nil {
		fmt.Println(err)
	}
	return err
}
