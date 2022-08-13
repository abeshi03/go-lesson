package main

import "fmt"

func main() {
	f := 1.11
	i := int(f)
	fmt.Println("%T %\n", i, i)

	s := []int{1, 2, 5, 6, 2, 3, 1}
	fmt.Println(s[2:4])
	//5

	m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	fmt.Printf("%T %v", m, m)
}
