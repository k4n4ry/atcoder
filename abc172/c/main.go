package main

import (
	"fmt"
)

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)
	var a = make([]int, n)
	var b = make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}

	var sum, count int
	for {
		if a[0] < b[0] {
			sum += a[0]
			fmt.Println(a[0])
			a = a[1:]
		} else {
			sum += b[0]
			fmt.Println(b[0])
			b = b[1:]
		}
		if sum > k {
			break
		}
		count++
	}
	fmt.Println(count)
}
