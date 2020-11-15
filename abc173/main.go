package main

import (
	"fmt"
)

var blackCount int
var MAX int
var HMAX int
var WMAX int

func main() {
	var h, w, k int
	fmt.Scan(&h, &w, &k)
	var p = make([][]int, h)
	for i := range p {
		p[i] = make([]int, w)
	}
	var s []string
	for i := 0; i < h; i++ {
		var m string
		// for j := 0; j < w; j++ {
		fmt.Scan(&m)
		s = append(s, m)
		// fmt.Println(i)
		// fmt.Println(j)
		// }
	}

	for i := 0; i < h; i++ {
		for j, v := range s[i] {
			switch v {
			case '.':
				p[i][j] = 0
			case '#':
				p[i][j] = 1
				blackCount++
			}

		}
	}
	fmt.Println(p)
	MAX = h + w
	HMAX = h
	WMAX = w
	dfs(0, 0, p)
}

func dfs(h, w int, p [][]int) {
	if h+w == MAX {
		return
	}
	if h < HMAX {
		for i := 0; i < WMAX; i++ {
			// p[h][i] == 9
		}
	}
}

func count(a [][]int) int {
	var count int
	for i := range a {
		for j := range a[i] {
			if a[i][j] == 1 {
				count++
			}
		}
	}
	return count
}
