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

func main() {
	io := NewIo()
	defer io.Flush()
	n := io.NextInt()
	type xx struct {
		x int
		n int
	}
	type yy struct {
		y int
		n int
	}
	xs := make([]xx, n)
	ys := make([]yy, n)
	for i := 0; i < n; i++ {
		x := io.NextInt()
		y := io.NextInt()
		xs[i] = xx{x, i}
		ys[i] = yy{y, i}
	}
	sort.SliceStable(xs, func(i, j int) bool { return xs[i].x < xs[j].x })
	sort.SliceStable(ys, func(i, j int) bool { return ys[i].y < ys[j].y })
	type ind struct {
		n    int
		from int
		to   int
	}
	xp := ind{xs[n-1].x - xs[0].x, xs[n-1].n, xs[0].n}
	yp := ind{ys[n-1].y - ys[0].y, ys[n-1].n, ys[0].n}
	xpp := ind{xs[n-2].x - xs[0].x, xs[n-2].n, xs[0].n}
	ypp := ind{ys[n-2].y - ys[0].y, ys[n-2].n, ys[0].n}
	if xs[n-1].x-xs[1].x > xs[n-2].x-xs[0].x {
		xpp = ind{xs[n-1].x - xs[1].x, xs[n-1].n, xs[1].n}
	}
	if ys[n-1].y-ys[1].y > ys[n-2].y-ys[0].y {
		ypp = ind{ys[n-1].y - ys[1].y, ys[n-1].n, ys[1].n}
	}
	a := []ind{xp, yp, xpp, ypp}
	sort.SliceStable(a, func(i, j int) bool {
		return a[i].n < a[j].n
	})
	if xp.from == yp.from && xp.to == yp.to {
		fmt.Println(a[len(a)-3].n)
	} else {
		fmt.Println(a[len(a)-2].n)
	}
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
