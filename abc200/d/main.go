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
	n := io.NextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = io.NextInt()
	}

	type b struct {
		mod int
		arr []int
	}
	mn := min(n, 8)
	mpp := make(map[int][]b)
	for bit := 0; bit < (1 << mn); bit++ {
		var b = b{mod: 0, arr: []int{}}
		var sum int
		for j := 0; j < mn; j++ {
			if (bit>>j)&1 == 1 {
				b.arr = append(b.arr, j+1)
				sum += a[j]
			}
		}
		b.mod = sum % 200
		mpp[b.mod] = append(mpp[b.mod], b)
	}
	var flg bool
	for k, v := range mpp {
		if k != 0 {
			if len(v) > 1 {
				if isSame(v[0].arr, v[1].arr) {
					continue
				}
				flg = true
				fmt.Println("Yes")
				io.Printf("%d ", len(v[0].arr))
				io.PrintIntLn(v[0].arr)
				io.Printf("%d ", len(v[1].arr))
				io.PrintIntLn(v[1].arr)
				break
			}
		} else {
			if len(v) > 2 {
				if isSame(v[1].arr, v[2].arr) {
					continue
				}
				flg = true
				fmt.Println("Yes")
				io.Printf("%d ", len(v[1].arr))
				io.PrintIntLn(v[1].arr)
				io.Printf("%d ", len(v[2].arr))
				io.PrintIntLn(v[2].arr)
				break
			}
		}
	}
	if !flg {
		fmt.Println("No")
	}
}

func isSame(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
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
	return int(math.Pow(float64(p), float64(q)))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
