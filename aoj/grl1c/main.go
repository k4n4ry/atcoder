package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var inf = math.MaxInt32

func main() {
	sc := newScanner()
	v := sc.readInt()
	e := sc.readInt()
	dp := make([][]int, v) //dp[i][j]はi-j間の最短路長を管理する
	for i := 0; i < len(dp); i++ {
		tmp := make([]int, v)
		for j := 0; j < len(tmp); j++ {
			if i == j {
				tmp[j] = 0
			} else {
				tmp[j] = inf
			}
		}
		dp[i] = tmp
	}
	for i := 0; i < e; i++ {
		s, t, d := sc.readInt(), sc.readInt(), sc.readInt()
		dp[s][t] = d
	}
	// Warshall–Floyd
	for tmp := 0; tmp < v; tmp++ {
		for i := 0; i < v; i++ {
			for j := 0; j < v; j++ {
				// inf(到達不可)の場合は更新しない
				if dp[i][tmp] == inf || dp[tmp][j] == inf {
					continue
				}
				// tmpのノードを経由したほうが路長が短くなる場合、更新
				dp[i][j] = min(dp[i][j], dp[i][tmp]+dp[tmp][j])
			}
		}
	}
	// 出力
	for i := 0; i < len(dp); i++ {
		if dp[i][i] < 0 {
			fmt.Println("NEGATIVE CYCLE")
			return
		}
	}
	for i := 0; i < len(dp); i++ {
		var s string
		for j := 0; j < len(dp[i]); j++ {
			if dp[i][j] == inf {
				s += "INF "
			} else {
				s += fmt.Sprintf("%d ", dp[i][j])
			}
		}
		fmt.Println(strings.Trim(s, " "))
	}
}
func min(nums ...int) int {
	ret := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] < ret {
			ret = nums[i]
		}
	}
	return ret
}

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
