package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type nab struct {
	n int64
	a int64
	b int64
}

func main() {
	var ho int64 = 1000000007
	sc := newScanner()
	var t = sc.readInt()
	var tc []nab
	for i := 0; i < t; i++ {
		n := sc.readInt64()
		a := sc.readInt64()
		b := sc.readInt64()
		tc = append(tc, nab{n, a, b})
	}

	for i := 0; i < t; i++ {
		lnab := tc[i]
		var ans int64
		a := lnab.a
		b := lnab.b
		n := lnab.n
		if lnab.a < lnab.b {
			ans = ((n-a+1)*(n-a+1)*(n-b+1)*(n-b+1) - b*b*(n-b+1)*(n-b+1)) % ho
		} else if a > b {
			ans = ((n-a+1)*(n-a+1)*(n-b+1)*(n-b+1) - a*a*(n-a+1)*(n-a+1)) % ho
		} else {
			ans = ((n-a+1)*(n-a+1)*(n-b+1)*(n-b+1) - (4*a*a + (4 * (n - a - 1) * a * (2*a - 1)) + ((n-a+1)*(n-a+1)-4*(n-a-1)-4)*(2*a-1)*(2*a-1))) % ho
		}
		if ans < 0 {
			ans += ho
		}
		fmt.Println(ans)

	}
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
