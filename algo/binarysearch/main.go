package main

import (
	"fmt"
	"sort"
)

func main() {
	// https://nishipy.com/archives/1616
	var data = []int{2, 5, 3, 9, 1, 6, 3}
	sort.Ints(data)
	tgt := 6
	idx := sort.Search(len(data), func(i int) bool { return data[i] >= tgt })
	if idx < len(data) && data[idx] == tgt {
		fmt.Println("tgt index is ", idx)
		fmt.Println("val is", data[idx])
	} else {
		fmt.Println("no values")
	}
}
