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

	n, m := io.NextInt(), io.NextInt()

	var x = make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = io.NextInt()
	}

	var cy = make(map[int]int)
	for i := 0; i < m; i++ {
		c, y := io.NextInt(), io.NextInt()
		cy[c] = y
	}
	var dp = make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		tmp := make([]int, n+1)
		for j := 0; j < len(tmp); j++ {
			tmp[j] = -1
		}
		dp[i] = tmp
	}

	dp[0][0] = 0
	dp[1][0] = 0
	dp[1][1] = dp[0][0] + x[0] + cy[1]
	var maxes = make([]int, len(dp))
	maxes[0] = 0
	maxes[1] = dp[1][1]
	for i := 2; i < n+1; i++ {
		var max = -1
		for j := i; j >= 0; j-- {
			if i == j {
				dp[i][j] = dp[i-1][j-1] + x[i-1] + cy[j]
				if dp[i][j] > max {
					max = dp[i][j]
				}
			} else if j == 0 {
				dp[i][j] = maxes[i-1]
			} else {
				if dp[i-1][j-1] == -1 {
					continue
				}
				dp[i][j] = dp[i-1][j-1] + x[i-1] + cy[j]
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
		maxes[i] = max
	}
	fmt.Println(max(dp[n]...))
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
