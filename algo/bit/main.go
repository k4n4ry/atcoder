package main

import "fmt"

// https://qiita.com/3x8tacorice/items/0b8341d7fd3ff3779111
func main() {
	var n int = 3

	// bitsが右に3シフト(= 1000 = 8)するまでloop
	for bits := 0; bits < (1 << uint64(n)); bits++ {
		fmt.Println(bits)
		// 3桁のbitsを右から順に評価
		for i := 0; i < n; i++ {
			if (bits>>uint64(i))&1 == 1 {
				// fmt.Println(bits)
			}
		}
	}
}
