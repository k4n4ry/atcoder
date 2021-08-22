package main

import (
	"fmt"
)

func main() {
	fmt.Println(map[bool]string{true: "AC", false: "WA"}[func() bool { var s string; fmt.Scan(&s); return "Hello,World!" == s }()])
}
