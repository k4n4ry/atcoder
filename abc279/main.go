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
	io := NewIo()
	defer io.Flush()

	a, b := io.NextFloat(), io.NextFloat()
	calc := func(a, b float64, n int) float64 {
		nn := float64(n)
		return b*nn + (a / math.Sqrt(nn+1))
	}
	zero := calc(a, b, 0)
	one := calc(a, b, 1)
	if zero < one {
		fmt.Println(zero)
		return
	}
	var left = 0
	var right = 1000000
	for {
		var mid = (right - left) / 2
		fmt.Printf("mid: %d, right: %d, left: %d\n", mid, right, left)
		tmpM1 := calc(a, b, mid-1)
		tmp := calc(a, b, mid)
		tmpP1 := calc(a, b, mid+1)
		fmt.Println("mid", mid)
		fmt.Println(tmpM1, tmp, tmpP1)
		if tmp <= tmpP1 && tmp <= tmpM1 {
			fmt.Println(tmp)
			return
		}
		if tmpM1 < tmp && tmp < tmpP1 {
			fmt.Println("aaa")
			if right == mid {
				fmt.Println(tmpM1)
				return
			}
			right = mid
		}
		if tmpM1 > tmp && tmp > tmpP1 {
			fmt.Println("bb")
			if left == mid {
				fmt.Println(tmpP1)
				return
			}
			left = mid
		}
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
