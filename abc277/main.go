package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type node struct {
	lf     []int
	isPast bool
}

var nd map[int][]int
var run map[int]int

var maxi int

func main() {
	io := NewIo()
	defer io.Flush()

	n, m := io.NextInt(), io.NextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = io.NextInt()
	}

	sort.Ints(a)
	fmt.Println(a)
	var sum int
	for _, b := range a {
		sum += b
	}

	var sums []int
	var iszero, islast bool
	var sam int
	for i := range a {
		if i < len(a)-1 && (a[i] == a[i+1] || a[i]+1 == a[i+1]) || i == len(a)-1 && (a[i] == m-1 && a[0] == 0) {
			fmt.Println(a[i])
			sam += a[i]
			if a[i] == 0 {
				iszero = true
			}
			if a[i] == m-1 {
				islast = true
			}
			if i == len(a)-1 {
				sums = append(sums, sam)
				sam = 0
			}
		} else {
			sums = append(sums, sam)
			sam = 0
		}
	}
	fmt.Println(sums)
	fmt.Println(sum)
	if len(sums) == 1 {
		fmt.Println(sum - sums[0])
		return
	}
	var x []int
	if islast && iszero {
		x = append(x, sums[0]+sums[len(sums)-1])
		for i := 1; i < len(sums)-1; i++ {
			x = append(x, sums[i])
		}
	} else {
		for i := 0; i < len(sums); i++ {
			x = append(x, sums[i])
		}
	}

	sort.Ints(x)
	fmt.Println(x)
	fmt.Println(sum - x[len(x)-1])
}

//_________io_________//
type Io struct {
	reader    *bufio.Reader
	writer    *bufio.Writer
	tokens    []string
	nextToken int
}

func NewIo() *Io {
	return &Io{
		reader: bufio.NewReader(os.Stdin),
		writer: bufio.NewWriter(os.Stdout),
	}
}

func (io *Io) Flush() {
	err := io.writer.Flush()
	if err != nil {
		panic(err)
	}
}

func (io *Io) NextLine() string {
	var buffer []byte
	for {
		line, isPrefix, err := io.reader.ReadLine()
		if err != nil {
			panic(err)
		}
		buffer = append(buffer, line...)
		if !isPrefix {
			break
		}
	}
	return string(buffer)
}

func (io *Io) Next() string {
	for io.nextToken >= len(io.tokens) {
		line := io.NextLine()
		io.tokens = strings.Fields(line)
		io.nextToken = 0
	}
	r := io.tokens[io.nextToken]
	io.nextToken++
	return r
}

func (io *Io) NextInt() int {
	i, err := strconv.Atoi(io.Next())
	if err != nil {
		panic(err)
	}
	return i
}

func (io *Io) NextFloat() float64 {
	i, err := strconv.ParseFloat(io.Next(), 64)
	if err != nil {
		panic(err)
	}
	return i
}

func (io *Io) PrintLn(a ...interface{}) {
	fmt.Fprintln(io.writer, a...)
}

func (io *Io) Printf(format string, a ...interface{}) {
	fmt.Fprintf(io.writer, format, a...)
}

func (io *Io) PrintIntLn(a []int) {
	b := []interface{}{}
	for _, x := range a {
		b = append(b, x)
	}
	io.PrintLn(b...)
}

func (io *Io) PrintStringLn(a []string) {
	b := []interface{}{}
	for _, x := range a {
		b = append(b, x)
	}
	io.PrintLn(b...)
}

func Log(name string, value interface{}) {
	fmt.Fprintf(os.Stderr, "%s=%+v\n", name, value)
}

//________util________//
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

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func pow(p, q int) int {
	var a = 1
	for i := 0; i < q; i++ {
		a = a * p
	}
	return a
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
