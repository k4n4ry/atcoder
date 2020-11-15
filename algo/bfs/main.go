package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type xy struct {
	y int
	x int
	d int
}
type que struct {
	data []xy
}

func (q *que) enq(i xy) {
	q.data = append(q.data, i)
}

func (q *que) deq() *xy {
	if len(q.data) == 0 {
		return nil
	}
	v := q.data[0]
	q.data = q.data[1:]
	return &v
}

// https://atcoder.jp/contests/agc033/tasks/agc033_a
// https://qiita.com/ruiu/items/8edf22cd5fc8f7511687
// https://qiita.com/aja_min/items/fed640d6835eb343eedf
func main() {
	sc := newScanner()
	var h = sc.readInt()
	var w = sc.readInt()
	var masu = make([][]rune, h)
	var q = que{data: []xy{}}
	for i := 0; i < h; i++ {
		masu[i] = make([]rune, w)
		var s = sc.readString()
		var j = 0
		for _, l := range s {
			masu[i][j] = l
			if masu[i][j] == '#' {
				q.enq(xy{i, j, 0})
			}
			j++
		}
	}
	// for i := 0; i < h; i++ {
	// 	for j := 0; j < w; j++ {
	// 		fmt.Printf(string(masu[i][j]))
	// 	}
	// 	fmt.Println("")
	// }

	depth := 0
	for {
		deq := q.deq()
		if deq == nil {
			break
		}

		var x = deq.x
		var y = deq.y
		var d = deq.d
		if d >= depth {
			depth = d
		}
		if x+1 < w && masu[y][x+1] == '.' {
			masu[y][x+1] = '#'
			q.enq(xy{y, x + 1, d + 1})
		}
		if x-1 >= 0 && masu[y][x-1] == '.' {
			masu[y][x-1] = '#'
			q.enq(xy{y, x - 1, d + 1})
		}
		if y+1 < h && masu[y+1][x] == '.' {
			masu[y+1][x] = '#'
			q.enq(xy{y + 1, x, d + 1})
		}
		if y-1 >= 0 && masu[y-1][x] == '.' {
			masu[y-1][x] = '#'
			q.enq(xy{y - 1, x, d + 1})
		}

	}
	fmt.Println(depth)

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
