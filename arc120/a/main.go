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
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = io.NextInt()
	}
	var anses = make([]int, n)
	var r = make([]int, n)
	var s int
	for i := 0; i < n; i++ {
		s += a[i]
		r[i] = s
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			anses[i] = a[i] + a[i]
			continue
		}
		tmp := a[:i+1]
		sort.Ints(tmp)
		for j := 0; j < len(tmp); j++ {
			if j == len(tmp)-1 {
				anses[i] += tmp[j] * (i + 2)
			} else {
				anses[i] += tmp[j] * (i + 1 - j)
			}
		}
	}
	for i := 0; i < n; i++ {
		io.PrintLn(anses[i])
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
