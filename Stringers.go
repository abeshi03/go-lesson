package main

import "fmt"

type Person struct {
	Name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("My name is %v", p.Name)
}

func main() {
	mike := Person{"Mike", 28}
	fmt.Println(mike)
}
