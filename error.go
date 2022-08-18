package main

import "fmt"

type UserNotFound struct {
	UserName string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User Not Found %v", e.UserName)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	} else {
		return &UserNotFound{"Mike"}
	}
}

func main() {
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}
