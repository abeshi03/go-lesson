package main

import "fmt"

func main() {
	l := []string{"python", "go", "java"}

	for i, v := range l {
		fmt.Println(i, v)
	}

	for _, v := range l {
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "orange": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
