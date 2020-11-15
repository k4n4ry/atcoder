package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc := newScanner()
	var h = sc.readInt()
	var w = sc.readInt()
	var m = sc.readInt()
	var d = make([][]int, h)
	// グリッドの初期化
	for i := 0; i < h; i++ {
		d[i] = make([]int, w)
		for j := 0; j < w; j++ {
			d[i][j] = 0
		}
	}

	for i := 0; i < m; i++ {
		h1 := sc.readInt()
		w1 := sc.readInt()
		d[h1-1][w1-1] = 1
	}

	// 縦横反転
	var dRev = make([][]int, w)
	for i := 0; i < w; i++ {
		dRev[i] = make([]int, h)
		for j := 0; j < h; j++ {
			dRev[i][j] = d[j][i]
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Printf("[%v]", d[i][j])
		}
		fmt.Println("")
	}

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			fmt.Printf("[%v]", dRev[i][j])
		}
		fmt.Println("")
	}

	var max = 0
	for i := 0; i < h; i++ {
		// 横の爆弾数
		var row = sumSlice(d[i])
		// 縦の最大を取得
		var maxCol = 0
		for j := 0; j < w; j++ {
			var col = sumSlice(dRev[j]) - dRev[j][i]
			if maxCol < col {
				maxCol = col
			}
		}
		if max < row+maxCol {
			max = row + maxCol
		}
	}
	fmt.Println(max)

}

func sumSlice(a []int) int {
	var sum int
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

//////////////////
//////////////////
//////////////////

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func min(values ...int) int {
	if len(values) == 0 {
		panic("no values")
	}

	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

func max(values ...int) int {
	if len(values) == 0 {
		panic("no values")
	}

	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

func pow(base, exponent int) int {
	if exponent < 0 {
		panic(fmt.Sprintf("exponent (%d) should not be a minus", exponent))
	}

	answer := 1
	for i := 0; i < exponent; i++ {
		answer *= base
	}
	return answer
}

func ceil(divident, dividor int) int {
	if dividor == 0 {
		panic("dividor should not be 0")
	}

	quo := divident / dividor
	rem := divident % dividor

	if rem != 0 {
		if (divident > 0 && dividor > 0) ||
			(divident < 0 && dividor < 0) {
			return quo + 1
		}
	}
	return quo
}

// -   sortutil

func reverseInts(a []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
}

func reverseStrings(a []string) {
	sort.Sort(sort.Reverse(sort.StringSlice(a)))
}

// -   io

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
