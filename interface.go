package main

import "fmt"

type Human interface {
	Say() string
}

type Person struct {
	Name string
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	return p.Name
}

func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("run")
	} else {
		fmt.Println("むずい。。。")
	}
}

func main() {
	var mike Human = &Person{"Mike"}
	//mike.Say()
	DriveCar(mike)
}
