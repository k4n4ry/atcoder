package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
	a := 1
	fmt.Println(a)
	b := 2
	fmt.Println(b)
	type hoge struct {
		fuga string
		piyo string
		mama string
	}
	xx := hoge{
		fuga: "",
		piyo: "",
	}
	fmt.Println(xx)
}
