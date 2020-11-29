package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// ac

var h, w, n int
var s [][]string
var d [][]int
var start hw
var tgts []hw

func main() {
	sc := newScanner()
	h = sc.readInt()
	w = sc.readInt()
	n = sc.readInt()
	s = getStrs2div(sc, h, w)
	d = make([][]int, h)
	for i := range d {
		dd := make([]int, w)
		for j := range dd {
			dd[j] = math.MaxInt32
		}
		d[i] = dd
	}
	tgts = make([]hw, n)

	// get start and goal and tgts
	for i := range s {
		for j := range s[i] {
			if s[i][j] == "S" {
				start = hw{i, j}
			} else if s[i][j] != "." && s[i][j] != "X" {
				a, _ := strconv.Atoi(s[i][j])
				tgts[a-1] = hw{i, j}
			}
		}
	}
	var q que
	q.enqueue(start)
	d[start.h][start.w] = 0
	var tgt = 1
	for len(q) > 0 {
		now := q.dequeue()
		if s[now.h][now.w] != "S" && s[now.h][now.w] != "." {
			nn, _ := strconv.Atoi(s[now.h][now.w])
			if nn == tgt {
				if tgt == n {
					fmt.Println(d[now.h][now.w])
					break
				}
				// reset
				tgt++
				q = q[:0]
				for i := range d {
					dd := make([]int, w)
					for j := range dd {
						if i == now.h && j == now.w {
							dd[j] = d[i][j]
							continue
						}
						dd[j] = math.MaxInt32
					}
					d[i] = dd
				}
			}
		}
		if now.h != 0 && s[now.h-1][now.w] != "X" && d[now.h-1][now.w] == math.MaxInt32 {
			d[now.h-1][now.w] = d[now.h][now.w] + 1
			q.enqueue(hw{now.h - 1, now.w})
		}
		if now.h != h-1 && s[now.h+1][now.w] != "X" && d[now.h+1][now.w] == math.MaxInt32 {
			d[now.h+1][now.w] = d[now.h][now.w] + 1
			q.enqueue(hw{now.h + 1, now.w})
		}
		if now.w != 0 && s[now.h][now.w-1] != "X" && d[now.h][now.w-1] == math.MaxInt32 {
			d[now.h][now.w-1] = d[now.h][now.w] + 1
			q.enqueue(hw{now.h, now.w - 1})
		}
		if now.w != w-1 && s[now.h][now.w+1] != "X" && d[now.h][now.w+1] == math.MaxInt32 {
			d[now.h][now.w+1] = d[now.h][now.w] + 1
			q.enqueue(hw{now.h, now.w + 1})
		}
	}
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

func gcd(a, b int64) int64 {
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
