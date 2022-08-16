package main

import "log"

func main() {
	log.Println("logging")
	log.Printf("%T %v", "test", "test")

	log.Fatalln("処理が終了する")
	log.Println("通らない")
}
