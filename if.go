package main

import "fmt"

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func main() {
	num := 3
	if num%2 == 0 {
		fmt.Println("by 2")
	} else if num%3 == 0 {
		fmt.Println("else if")
	} else {
		fmt.Println("else")
	}

	x, y := 10, 10

	if x == 10 && y == 10 {
		fmt.Println("&&")
	}
	if x == 10 || y == 10 {
		fmt.Println("||")
	}
	if x == 11 || y == 12 {
		fmt.Println("false")
	}

	// if 文を続けて書く方法
	if result2 := by2(2); result2 == "ok" {
		fmt.Println("great")
	}
}
