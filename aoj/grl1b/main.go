package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type edge struct {
	to int // エッジ終点
	w  int // コスト
}

var inf = math.MaxInt32

func main() {
	sc := newScanner()
	v, e, r := sc.readInt(), sc.readInt(), sc.readInt()
	g := make([][]edge, v)
	for i := 0; i < e; i++ {
		s, t, d := sc.readInt(), sc.readInt(), sc.readInt()
		g[s] = append(g[s], edge{to: t, w: d})
	}
	dist := make([]int, v) // 始点からの路長
	for i := 0; i < v; i++ {
		dist[i] = inf // infで初期化
	}
	// bellmanford
	dist[r] = 0 // スタートのノードの路長を0に更新
	for i := 0; i < v; i++ {
		var updateFlg bool
		for j := 0; j < len(g); j++ {
			for _, edge := range g[j] {
				if dist[j] == inf {
					continue
				}
				// 路長を緩和
				if dist[edge.to] > edge.w+dist[j] {
					updateFlg = true
					dist[edge.to] = min(dist[edge.to], edge.w+dist[j])
				}
			}
		}
		// v回目のループで更新がある場合、負閉路があると判定
		if i == v-1 && updateFlg {
			fmt.Println("NEGATIVE CYCLE")
			return
		}
	}
	// 各ノードへの最短路長を出力、到達不可の場合はINF
	for _, v := range dist {
		if v < inf {
			fmt.Println(v)
		} else {
			fmt.Println("INF")
		}
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
