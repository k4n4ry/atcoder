package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type xyz struct {
	id  int
	x   int
	y   int
	z   int
	flg bool
}

type xyzQueue []xyz

func (queue *xyzQueue) enqueue(xyz xyz) {
	*queue = append(*queue, xyz)
}

func (queue *xyzQueue) dequeue() (xyz xyz) {
	result := (*queue)[0]
	*queue = (*queue)[1:]
	return result
}

var tmpCost int
var ans = 0

func main() {
	sc := newScanner()
	var n = sc.readInt()
	var xyzs []xyz
	for i := 0; i < n; i++ {
		var x = sc.readInt()
		var y = sc.readInt()
		var z = sc.readInt()
		var xyz = xyz{
			id: i,
			x:  x,
			y:  y,
			z:  z,
		}
		xyzs = append(xyzs, xyz)
	}
	var xyzq xyzQueue
	xyzs[0].flg = true
	xyzq.enqueue(xyzs[0])
	for len(xyzq) > 0 {
		var now = xyzq.dequeue()
		for i := range xyzs {
			// if now.id == 0 && xyzs[i].id == 0 {
			// 	continue
			// }
			if xyzs[i].flg {
				continue
			}
			tmpCost += kyori(now, xyzs[i])
			xyzs[i].flg = true
			xyzq.enqueue(xyzs[i])
		}
		tmpCost += kyori(now, xyzs[0])
		fmt.Println(tmpCost)
		if ans == 0 {
			ans = tmpCost
		} else if ans > tmpCost {
			ans = tmpCost
		}
	}
	fmt.Println(ans)

}

func kyori(f, d xyz) int {
	var ans int
	ans += abs(f.x - d.x)
	ans += abs(f.y - d.y)
	var max = 0
	if f.z-d.z > 0 {
		max = f.z - d.z
	}
	ans += max
	return ans
}

/*
snipet--------------------------------------
*/
func getNums(sc *scanner, len int) (nums []int) {
	var a = make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = sc.readInt()
	}
	return a
}

func getNums64(sc *scanner, len int) (nums []int64) {
	var a = make([]int64, len)
	for i := 0; i < len; i++ {
		a[i] = sc.readInt64()
	}
	return a
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func pow(p, q int) int {
	return int(math.Pow(float64(p), float64(q)))
}

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

type modint struct {
	x      int64
	modVal int64
	inv    int64
}

func new(x, modVal int64) *modint {
	// 逆元を求める
	return &modint{
		x:      x,
		modVal: modVal,
	}
}

func (m *modint) add(n int64) int64 {
	return (m.x + n) % m.modVal
}

func (m *modint) sub(n int64) int64 {
	if m.x-n < 0 {
		return (m.x - n) + m.modVal
	}
	return m.x - n
}

func (m *modint) mul(n int64) int64 {
	return (m.x * n) % m.modVal
}

func (m *modint) div(n int64) int64 {
	return (m.inv * n) % m.modVal
}

/*
io------------------------------------------
*/

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

func (sc *scanner) readUint64() uint64 {
	bufSc := sc.bufScanner
	bufSc.Scan()
	text := bufSc.Text()

	num, err := strconv.ParseUint(text, 10, 64)
	if err != nil {
		panic(err)
	}
	return num
}
