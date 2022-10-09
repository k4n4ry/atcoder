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
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		tmp := make([]int, n)
		for j := 0; j < len(tmp); j++ {
			tmp[j] = -1
		}
		matrix[i] = tmp
	}

	type xystr struct {
		x int
		y int
	}
	var xys []xystr
	for i := 0; i < 1001; i++ {
		for j := 0; j < 1001; j++ {
			tmp := i*i + j*j
			if tmp == m {
				xys = append(xys, xystr{x: i, y: j})
			}
		}
	}

	var q que
	q.enqueue(hw{0, 0, 0})
	for len(q) > 0 {
		now := q.dequeue()
		h, w, cnt := now.h, now.w, now.cnt

		for _, xy := range xys {
			// migishita
			if h+xy.x <= n-1 && w+xy.y <= n-1 {
				tgt := hw{h: h + xy.x, w: w + xy.y, cnt: cnt + 1}
				if matrix[tgt.h][tgt.w] == -1 || matrix[tgt.h][tgt.w] > tgt.cnt {
					matrix[tgt.h][tgt.w] = tgt.cnt
					q.enqueue(tgt)
				}
			}
			if h+xy.y <= n-1 && w+xy.x <= n-1 {
				tgt := hw{h: h + xy.y, w: w + xy.x, cnt: cnt + 1}
				if matrix[tgt.h][tgt.w] == -1 || matrix[tgt.h][tgt.w] > tgt.cnt {
					matrix[tgt.h][tgt.w] = tgt.cnt
					q.enqueue(tgt)
				}
			}

			// migiue
			if h-xy.x >= 0 && w+xy.y <= n-1 {
				tgt := hw{h: h - xy.x, w: w + xy.y, cnt: cnt + 1}
				if matrix[tgt.h][tgt.w] == -1 || matrix[tgt.h][tgt.w] > tgt.cnt {
					matrix[tgt.h][tgt.w] = tgt.cnt
					q.enqueue(tgt)
				}
			}
			if h-xy.y >= 0 && w+xy.x <= n-1 {
				tgt := hw{h: h - xy.y, w: w + xy.x, cnt: cnt + 1}
				if matrix[tgt.h][tgt.w] == -1 || matrix[tgt.h][tgt.w] > tgt.cnt {
					matrix[tgt.h][tgt.w] = tgt.cnt
					q.enqueue(tgt)
				}
			}

			// hidarishita
			if h+xy.x <= n-1 && w-xy.y >= 0 {
				tgt := hw{h: h + xy.x, w: w - xy.y, cnt: cnt + 1}
				if matrix[tgt.h][tgt.w] == -1 || matrix[tgt.h][tgt.w] > tgt.cnt {
					matrix[tgt.h][tgt.w] = tgt.cnt
					q.enqueue(tgt)
				}
			}
			if h+xy.y <= n-1 && w-xy.x >= 0 {
				tgt := hw{h: h + xy.y, w: w - xy.x, cnt: cnt + 1}
				if matrix[tgt.h][tgt.w] == -1 || matrix[tgt.h][tgt.w] > tgt.cnt {
					matrix[tgt.h][tgt.w] = tgt.cnt
					q.enqueue(tgt)
				}
			}

			// hidariue
			if h-xy.x >= 0 && w-xy.y >= 0 {
				tgt := hw{h: h - xy.x, w: w - xy.y, cnt: cnt + 1}
				if matrix[tgt.h][tgt.w] == -1 || matrix[tgt.h][tgt.w] > tgt.cnt {
					matrix[tgt.h][tgt.w] = tgt.cnt
					q.enqueue(tgt)
				}
			}
			if h-xy.y >= 0 && w-xy.x >= 0 {
				tgt := hw{h: h - xy.y, w: w - xy.x, cnt: cnt + 1}
				if matrix[tgt.h][tgt.w] == -1 || matrix[tgt.h][tgt.w] > tgt.cnt {
					matrix[tgt.h][tgt.w] = tgt.cnt
					q.enqueue(tgt)
				}
			}
		}
	}
	matrix[0][0] = 0
	for i := 0; i < len(matrix); i++ {
		io.PrintIntLn(matrix[i])
	}
}

type hw struct {
	h   int
	w   int
	cnt int
}

type que []hw

func (q *que) enqueue(i hw) {
	*q = append(*q, i)
}

func (q *que) dequeue() hw {
	result := (*q)[0]
	*q = (*q)[1:]
	return result
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
