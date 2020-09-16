package utils

import (
	"os/exec"
)

func RunCommand(name string, arg ...string) (err error) {
	command := exec.Command(name, arg...)
	return command.Run()
}
