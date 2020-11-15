package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var mem []int
var nMem []int

func main() {
	sc := newScanner()
	var n = sc.readInt64()
	var x = sc.readInt()
	var m = sc.readInt()

	var ans int64 = int64(x)
	mem = make([]int, 0, 100000)
	mem = append(mem, x)
	var a int64 = 1
	for i := 0; i < 100000; i++ {
		a++
		x = x * x % m
		if i, ok := contains(x); ok {
			nMem = mem[i:]
			var zanLen = (n - int64(len(mem)))
			ans += (zanLen / int64(len(nMem))) * sumI()
			var g = zanLen % int64(len(nMem))
			for j := 0; j < int(g); j++ {
				ans += int64(nMem[j])
			}
			break
		}
		mem = append(mem, x)
		ans += int64(x)
		if a == n {
			break
		}
	}
	fmt.Println(ans)
}

//////////////////
//////////////////
//////////////////

func contains(num int) (int, bool) {
	for i := range mem {
		if mem[i] == num {
			return i, true
		}
	}
	return -1, false
}

func sumI() int64 {
	var sum int64 = 0
	for i := range nMem {
		sum += int64(nMem[i])
	}
	return sum
}

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
