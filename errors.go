package common_plugin

import "fmt"

type Error struct {
	Status int
	Err    error
}

func (e Error) Error() string {
	fmt.Println("only debugging purpose")
	return e.Err.Error()
}
