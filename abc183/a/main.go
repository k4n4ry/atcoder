package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc := newScanner()
	var n = sc.readInt()
	var w = sc.readInt()
	var ms = make([]int, 220000)
	for i := 0; i < n; i++ {
		var s = sc.readInt()
		var t = sc.readInt()
		var p = sc.readInt()
		// imos
		ms[s] += p
		ms[t] -= p
	}
	// var rui = make([]int, 220000)
	for i := 1; i < len(ms); i++ {
		ms[i] += ms[i-1]
	}
	// get max
	var max = 0
	for i := 0; i < len(ms); i++ {
		if ms[i] > max {
			max = ms[i]
		}
	}
	if max > w {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}

/*
snipet--------------------------------------
*/

// int配列で取得
func getNums(sc *scanner, len int) (nums []int) {
	var a = make([]int, len)
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
