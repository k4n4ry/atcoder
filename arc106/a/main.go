package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type node struct {
	id     int
	val    int
	cap    int
	linked []int
	flg    bool
}

type hen struct {
	f int
	t int
}

func main() {
	var sc = newScanner()
	n := sc.readInt()
	m := sc.readInt()
	var a = getNums(sc, n)
	var b = getNums(sc, n)
	var hens []hen
	for i := 0; i < m; i++ {
		c := sc.readInt()
		d := sc.readInt()
		c--
		d--
		hens = append(hens, hen{
			f: c,
			t: d,
		})
	}
	var nodes []node
	for i := range a {
		nodes = append(nodes, node{
			id:  i + 1,
			val: a[i],
			cap: b[i],
		})
	}

	uf := newUnionFind(n)
	for i := range hens {
		uf.unite(hens[i].f, hens[i].t)
	}

	G := uf.groups()
	for _, group := range G {
		as, bs := 0, 0
		for _, id := range group {
			as += a[id]
			bs += b[id]
		}

		if as != bs {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")

}

// union find

type UnionFind struct {
	n   int
	par []int
}

func newUnionFind(n int) *UnionFind {
	uf := new(UnionFind)
	uf.n = n
	uf.par = make([]int, n)
	for i, _ := range uf.par {
		uf.par[i] = -1
	}
	return uf
}
func (uf UnionFind) root(x int) int {
	if uf.par[x] < 0 {
		return x
	}
	uf.par[x] = uf.root(uf.par[x])
	return uf.par[x]
}
func (uf UnionFind) unite(x, y int) {
	rx, ry := uf.root(x), uf.root(y)
	if rx != ry {
		if uf.size(rx) > uf.size(ry) {
			rx, ry = ry, rx
		}
		uf.par[ry] += uf.par[rx]
		uf.par[rx] = ry
	}
}
func (uf UnionFind) same(x, y int) bool {
	return uf.root(x) == uf.root(y)
}
func (uf UnionFind) size(x int) int {
	return -uf.par[uf.root(x)]
}
func (uf UnionFind) groups() [][]int {
	rootBuf, groupSize := make([]int, uf.n), make([]int, uf.n)
	for i := 0; i < uf.n; i++ {
		rootBuf[i] = uf.root(i)
		groupSize[rootBuf[i]]++
	}
	res := make([][]int, uf.n)
	for i := 0; i < uf.n; i++ {
		res[i] = make([]int, 0, groupSize[i])
	}
	for i := 0; i < uf.n; i++ {
		res[rootBuf[i]] = append(res[rootBuf[i]], i)
	}
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

/*
snipet--------------------------------------
*/
func getNums(sc *scanner, len int) (nums []int) {
	var a = make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = sc.readInt()
	}
	return a
}

func getNums64(sc *scanner, len int) (nums []int64) {
	var a = make([]int64, len)
	for i := 0; i < len; i++ {
		a[i] = sc.readInt64()
	}
	return a
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func pow(p, q int) int {
	return int(math.Pow(float64(p), float64(q)))
}

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

type modInt struct {
	x      int64
	modVal int64
}

func newModInt(x, modVal int64) *modInt {
	return &modInt{
		x:      x,
		modVal: modVal,
	}
}

// 和
func (m *modInt) add(n int64) *modInt {
	m.x = (m.x + n) % m.modVal
	return m
}

// 差
func (m *modInt) sub(n int64) *modInt {
	if m.x-n < 0 {
		m.x = (m.x - n) + m.modVal
	} else {
		m.x = m.x - n
	}
	return m
}

// 積
func (m *modInt) mul(n int64) *modInt {
	m.x = (m.x * n) % m.modVal
	return m
}

// 商
func (m *modInt) div(n int64) *modInt {
	var a, b int64 = n, m.modVal
	var u, v int64 = 1, 0
	for b > 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	if u < 0 {
		u += m.modVal
	}
	u %= m.modVal
	m.x = (m.x * u) % m.modVal
	return m
}

/*
io------------------------------------------
*/

type scanner struct {
	bufScanner *bufio.Scanner
}

func newScanner() *scanner {
	bufSc := bufio.NewScanner(os.Stdin)
	bufSc.Split(bufio.ScanWords)
	bufSc.Buffer(nil, 100000000)
	return &scanner{bufSc}
}

func (sc *scanner) readString() string {
	bufSc := sc.bufScanner
	bufSc.Scan()
	return bufSc.Text()
}

func (sc *scanner) readInt() int {
	bufSc := sc.bufScanner
	bufSc.Scan()
	text := bufSc.Text()

	num, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return num
}

func (sc *scanner) readInt64() int64 {
	bufSc := sc.bufScanner
	bufSc.Scan()
	text := bufSc.Text()

	num, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		panic(err)
	}
	return num
}

func (sc *scanner) readUint64() uint64 {
	bufSc := sc.bufScanner
	bufSc.Scan()
	text := bufSc.Text()

	num, err := strconv.ParseUint(text, 10, 64)
	if err != nil {
		panic(err)
	}
	return num
}
