package main

import (
	"fmt"
	"time"
)

func main() {
	os := "mac"
	switch os {
	case "mac":
		fmt.Println("mac")
	case "Windows":
		fmt.Println("win")
	default:
		fmt.Println("default")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("afternoon")
	case t.Hour() < 24:
		fmt.Println("evening")

	}
}
