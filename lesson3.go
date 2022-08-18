package main

import "fmt"

type Vertex struct {
	X, Y int
}

//func (e Vertex) Plus() int {
//	return e.X + e.Y
//}

func (p Vertex) String() string {
	return fmt.Sprintf("X is %d! Y is %d", p.X, p.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v)
}
