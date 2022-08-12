package main

import "fmt"

func main() {
	m := map[string]int{"app": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["app"])
}
