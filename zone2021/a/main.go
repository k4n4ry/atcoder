package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := newScanner()
	var zone = "ZONe"
	s := sc.readString()
	fmt.Println(strings.Count(s, zone))
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
type unionFind struct {
	n   int
	par []int
}

func newUnionFind(n int) *unionFind {
	uf := new(unionFind)
	uf.n = n
	uf.par = make([]int, n)
	for i := range uf.par {
		uf.par[i] = -1
	}
	return uf
}

func (uf unionFind) root(x int) int {
	if uf.par[x] < 0 {
		return x
	}
	// ここで、根を取得するのにあわせて、xを根直下においてるんじゃない?
	uf.par[x] = uf.root(uf.par[x])
	return uf.par[x]
}

func (uf unionFind) unite(x, y int) {
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

func (uf unionFind) same(x, y int) bool {
	return uf.root(x) == uf.root(y)
}

func (uf unionFind) size(x int) int {
	// uniteで、併合するたびに-1, -2・・・となっていくので、下記のとおりで求まる
	return -uf.par[uf.root(x)]
}

func (uf unionFind) groups() [][]int {
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

// primenumber
// Get all prime factors of a given number n
func getPrimeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}
