package task

import (
	"fmt"
)

func HelloWorld(arg string) (string, error) {
	return "Hi, i'm worker@localhost", nil
}
func Add(args ...int64) error {
	fmt.Println("fffffff")
	return nil
}
