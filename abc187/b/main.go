package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type xy struct {
	x float64
	y float64
}

func main() {
	sc := newScanner()
	n := sc.readInt()
	var xys []xy
	for i := 0; i < n; i++ {
		x := sc.readInt()
		y := sc.readInt()
		xys = append(xys, xy{float64(x), float64(y)})
	}

	var cnt int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d := (xys[i].y - xys[j].y) / (xys[i].x - xys[j].x)
			if -1 <= d && d <= 1 {
				cnt++
			}
		}
	}
	fmt.Println(cnt)

}

/*
snipet--------------------------------------
*/

// int2次元配列を取得する
func getNums2div(sc *scanner, h, w int) [][]int {
	ret := make([][]int, h)
	for i := 0; i < h; i++ {
		tmpret := make([]int, w)
		for j := 0; j < w; j++ {
			a := sc.readInt()
			tmpret[j] = a
		}
		ret[i] = tmpret
	}
	return ret
}

// string2次元配列を取得する
func getStrs2div(sc *scanner, h, w int) [][]string {
	ret := make([][]string, h)
	for i := 0; i < h; i++ {
		tmp := sc.readString()
		tmpret := getStringSlice(tmp)
		ret[i] = tmpret
	}
	return ret
}

// int配列で取得
func getNums(sc *scanner, len int) (nums []int) {
	var a = make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = sc.readInt()
	}
	return a
}

// int配列で取得(Cap指定_dpのときなどで使う)
func getNumsCap(sc *scanner, len, cap int) (nums []int) {
	var a = make([]int, len, cap)
	for i := 0; i < len; i++ {
		a[i] = sc.readInt()
	}
	return a
}

// int配列で取得
func getNums64(sc *scanner, len int) (nums []int64) {
	var a = make([]int64, len)
	for i := 0; i < len; i++ {
		a[i] = sc.readInt64()
	}
	return a
}

// stringを1文字ずつの配列に変換
func getStringSlice(s string) []string {
	var cs = make([]string, len(s))
	for i, c := range s {
		cs[i] = string(c)
	}
	return cs
}

// stringを1文字ずつの数値に変換
func getIntSliceFromString(s string) []int {
	var cs = make([]int, len(s))
	for i, c := range s {
		ss := string(c)
		nt, _ := strconv.Atoi(ss)
		cs[i] = nt
	}
	return cs
}

// intを1桁ずつの配列に変換(下の桁から)
func getDigitSliceLow(n int) []int {
	var tmp = n
	var s = strconv.Itoa(tmp)
	var ret = make([]int, 0, len(s))
	for tmp > 0 {
		ret = append(ret, tmp%10)
		tmp = tmp / 10
	}
	return ret
}

// intを1桁ずつの配列に変換(上の桁から)
func getDigitSliceHigh(n int) []int {
	var tmp = n
	var s = strconv.Itoa(tmp)
	var tmpRet = make([]int, 0, len(s))
	for tmp > 0 {
		tmpRet = append(tmpRet, tmp%10)
		tmp = tmp / 10
	}
	var ret = make([]int, len(tmpRet))
	for i := len(tmpRet) - 1; i >= 0; i-- {
		ret[i] = tmpRet[len(tmpRet)-1-i]
	}
	return ret
}

// intスライスを逆順にする
func intRev(s []int) []int {
	var ret = make([]int, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		ret[i] = s[len(s)-1-i]
	}
	return ret
}

// stringスライスを逆順にする
func strRev(s []string) []string {
	var ret = make([]string, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		ret[i] = s[len(s)-1-i]
	}
	return ret
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func pow(p, q int) int {
	return int(math.Pow(float64(p), float64(q)))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// que
type hw struct {
	h int
	w int
}

type que []hw

func (q *que) enqueue(i hw) {
	*q = append(*q, i)
}

func (q *que) dequeue() hw {
	result := (*q)[0]
	*q = (*q)[1:]
	return result
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

func min(nums ...int) int {
	ret := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] < ret {
			ret = nums[i]
		}
	}
	return ret
}

func max(nums ...int) int {
	ret := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > ret {
			ret = nums[i]
		}
	}
	return ret
}

// アドレスで渡す必要があるんか。。。
func chmin(pa, pb *int) bool {
	a, b := *pa, *pb
	if a > b {
		*pa = *pb
		return true
	}
	return false
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

// 商, uが逆元
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

// nの逆元を取得する
func moddiv(n, modVal int) int {
	var a, b int = n, modVal
	var u, v int = 1, 0
	for b > 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	if u < 0 {
		u += modVal
	}
	u %= modVal
	return u
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

// unionfind
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
