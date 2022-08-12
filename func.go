package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func cal(price int, item int) (result int) {
	result = price + item
	return
}

func main() {
	r := add(10, 20)
	fmt.Println(r)

	r3 := cal(1000, 2000)
	fmt.Println(r3)
}
