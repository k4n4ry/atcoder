package main

import "fmt"

func main() {
	uf := newUnionFind(5)
	// fmt.Println(uf)
	uf.unite(1, 3)
	// fmt.Println(uf)
	uf.unite(2, 4)
	// fmt.Println(uf)
	uf.unite(3, 0)

	// fmt.Println(uf)
	G := uf.groups()
	fmt.Println(G)
}

type UnionFind struct {
	n   int
	par []int
}

func newUnionFind(n int) *UnionFind {
	uf := new(UnionFind)
	uf.n = n
	uf.par = make([]int, n)
	for i := range uf.par {
		uf.par[i] = -1
	}
	return uf
}

func (uf UnionFind) root(x int) int {
	if uf.par[x] < 0 {
		return x
	}
	// ここで、根を取得するのにあわせて、xを根直下においてるんじゃない?
	uf.par[x] = uf.root(uf.par[x])
	return uf.par[x]
}

func (uf UnionFind) unite(x, y int) {
	rx, ry := uf.root(x), uf.root(y)
	// もし、違うグループだったら
	if rx != ry {
		if uf.size(rx) > uf.size(ry) {
			rx, ry = ry, rx
		}
		// rxの値をryに付与する(サイズ計算のため。ryが2とか3で、par[ry]が-(ノードの数)なので、)
		uf.par[ry] += uf.par[rx]
		// rxがryにぶら下がる
		uf.par[rx] = ry
	}
}

func (uf UnionFind) same(x, y int) bool {
	return uf.root(x) == uf.root(y)
}

func (uf UnionFind) size(x int) int {
	// uniteで、併合するたびに-1, -2・・・となっていくので、下記のとおりで求まる
	return -uf.par[uf.root(x)]
}

func (uf UnionFind) groups() [][]int {
	rootBuf, groupSize := make([]int, uf.n), make([]int, uf.n)
	// 各要素の根を取得し、groupごとのサイズを取得
	for i := 0; i < uf.n; i++ {
		rootBuf[i] = uf.root(i)
		groupSize[rootBuf[i]]++
	}
	// res初期化
	res := make([][]int, uf.n)
	for i := 0; i < uf.n; i++ {
		res[i] = make([]int, 0, groupSize[i])
	}
	// 同じ値をもつ要素をgroup化
	for i := 0; i < uf.n; i++ {
		res[rootBuf[i]] = append(res[rootBuf[i]], i)
	}
	// 要素ゼロは除外してreturn
	result := make([][]int, 0, uf.n)
	for i := 0; i < uf.n; i++ {
		if len(res[i]) != 0 {
			r := make([]int, len(res[i]))
			copy(r, res[i])
			result = append(result, r)
		}
	}
	return result
}
