package main

import (
	"fmt"
	"math"
)

// ac
func main() {
	var n int
	fmt.Scan(&n)
	var ans int = -1
	if n == 5 || n%2 == 0 {
		fmt.Println("-1")
		return
	}
	var num float64 = 0
	for i := 0; i < n; i++ {
		num += 7 * math.Pow(float64(10), float64(i))
		fmt.Println(num)
		if int(num)%n == 0 {
			ans = i + 1
			break
		}
	}
	fmt.Println(ans)
}
