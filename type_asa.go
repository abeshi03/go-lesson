package main

import "fmt"

func do(i interface{}) {
	// タイプアサーション
	//ii := i.(int)
	//ii *= 2
	//fmt.Println(ii)

	// switch type文
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	case bool:
		fmt.Println(!v)
	}
}

func main() {
	do(10)
	do("Make")
	do(true)
}
