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
	n, m := io.NextInt(), io.NextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = io.NextInt()
	}
	b := make([]int, m)
	for i := 0; i < m; i++ {
		b[i] = io.NextInt()
	}
	sort.Ints(a)
	sort.Ints(b)
	var aidx, bidx int
	var shu bool
	if a[0] < b[0] {
		shu = true
	}
	var ans = 1000000010
	for {
		if len(a) == aidx || len(b) == bidx {
			break
		}
		if shu {
			aidx++
			if len(a) == aidx {
				if abs(a[aidx-1]-b[bidx]) < ans {
					ans = abs(a[aidx-1] - b[bidx])
				}
				break
			}
			if a[aidx] >= b[bidx] {
				dif := compare(a[aidx-1], b[bidx], a[aidx])
				if ans > dif {
					ans = dif
				}
				shu = false
			} else {
				continue
			}
		} else {
			bidx++
			if len(b) == bidx {
				if abs(b[bidx-1]-a[aidx]) < ans {
					ans = abs(b[bidx-1] - a[aidx])
				}
				break
			}
			if b[bidx] >= a[aidx] {
				dif := compare(b[bidx-1], a[aidx], b[bidx])
				if ans > dif {
					ans = dif
				}
				shu = true
			} else {
				continue
			}

		}

	}
	fmt.Println(ans)

}

func compare(a, b, c int) int {
	if c-b < b-a {
		return c - b
	} else {
		return b - a
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
