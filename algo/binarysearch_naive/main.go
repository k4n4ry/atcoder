package main

import (
	"fmt"
	"sort"
)

func main() {
	isOk := func(n, tgt int) bool {
		return n <= tgt
	}
	ns := []int{2, 7, 1, 5, 4, 3, 2}
	var mark int
	fmt.Scan(&mark)
	sort.Ints(ns)
	ok, ng := -1, len(ns)
	for {
		md := (ok + ng) / 2
		if isOk(ns[md], mark) {
			ok = md
		} else {
			ng = md
		}
		if ng-ok == 1 {
			break
		}
	}
	fmt.Println(ns[ok])
}
