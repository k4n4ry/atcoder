package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(map[bool]string{true: "AC", false: "WA"}["Hello,World!" == s])
}
