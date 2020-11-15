package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var h int
var w int
var c [][]string
var b [][]bool
var findFlg bool

type hw struct {
	h int
	w int
}

func main() {
	// 1文字ずつ処理
	sc := newScanner()
	h = sc.readInt()
	w = sc.readInt()
	c = make([][]string, h)
	b = make([][]bool, h)
	for i := 0; i < h; i++ {
		var a = make([]string, w)
		c[i] = a
		var bb = make([]bool, w)
		b[i] = bb
	}

	for i := 0; i < h; i++ {
		cs := sc.readString()
		for j, cc := range cs {
			c[i][j] = string(cc)
		}
	}
	// fmt.Println(c)

	var start hw
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if c[i][j] == "s" {
				start = hw{i, j}
			}
		}
	}
	dfs(start)
	if findFlg {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func dfs(now hw) {
	b[now.h][now.w] = true
	if c[now.h][now.w] == "g" {
		findFlg = true
		return
	}
	if now.h != 0 && !b[now.h-1][now.w] && c[now.h-1][now.w] != "#" {
		dfs(hw{now.h - 1, now.w})
	}
	if now.h != h-1 && !b[now.h+1][now.w] && c[now.h+1][now.w] != "#" {
		dfs(hw{now.h + 1, now.w})
	}
	if now.w != 0 && !b[now.h][now.w-1] && c[now.h][now.w-1] != "#" {
		dfs(hw{now.h, now.w - 1})
	}
	if now.w != w-1 && !b[now.h][now.w+1] && c[now.h][now.w+1] != "#" {
		dfs(hw{now.h, now.w + 1})
	}
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

// binary_search
func binarySearch(array []int, target int) int {
	// 範囲start < endを探索する
	arrayLen := len(array)
	start := 0
	end := arrayLen - 1
	var index int
	for {
		if end < start {
			return -1
		}
		index = (start + end) / 2

		if array[index] == target {
			return index
		}

		if array[index] < target {
			start = index + 1
		} else {
			end = index - 1
		}
	}
}

// LowerBound ...
func lowerBound(array []int, target int) int {
	low, high, mid := 0, len(array)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if array[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

// UpperBound ...
func upperBound(array []int, target int) int {
	low, high, mid := 0, len(array)-1, 0

	for low <= high {
		mid = (low + high) / 2
		if array[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

// mod
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
