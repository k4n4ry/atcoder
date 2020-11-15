package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := newScanner()
	n := sc.readInt()
	a := getNums64(sc, n)
	var ans = a[0]
	for i := range a {
		ans = gcd(ans, a[i])
	}
	fmt.Println(ans)

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

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

type modint struct {
	num    int64
	modulo int64
	inv    int64
}

func new(num, modulo int64) *modint {
	// 逆元を求める
	return &modint{
		num:    num,
		modulo: modulo,
	}
}

func (m *modint) add(n int64) int64 {
	return (m.num + n) % m.modulo
}

func (m *modint) sub(n int64) int64 {
	if m.num-n < 0 {
		return (m.num - n) + m.modulo
	}
	return m.num - n
}

// func gcd(m, n int64) int64 {
// 	x := new(big.Int)
// 	y := new(big.Int)
// 	z := new(big.Int)
// 	a := new(big.Int).SetInt64(m)
// 	b := new(big.Int).SetInt64(n)
// 	return z.GCD(x, y, a, b).Int64()
// }

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
