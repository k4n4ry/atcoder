package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
)

type edge struct {
	to int // エッジ終点
	w  int // コスト
}

// ヒープで管理する要素情報
type vert struct {
	d int // 緩和したときのdistの値
	v int // 緩和対象のノード番号
}

// ヒープを表現するプライオリティキューcontainer/heap内のInterfaceを実装する
type pQue []vert

func (pq pQue) Len() int            { return len(pq) }
func (pq pQue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq pQue) Less(i, j int) bool  { return pq[i].d < pq[j].d }
func (pq *pQue) Push(x interface{}) { *pq = append(*pq, x.(vert)) }
func (pq *pQue) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
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
	// dijkstra
	dist[r] = 0 // スタートのノードの路長を0に更新
	var pq pQue
	heap.Push(&pq, vert{d: dist[r], v: r})
	for len(pq) > 0 {
		mn := heap.Pop(&pq).(vert)
		if mn.d > dist[mn.v] {
			continue
		}
		for _, ed := range g[mn.v] {
			// 路長を緩和
			if dist[ed.to] > dist[mn.v]+ed.w {
				dist[ed.to] = dist[mn.v] + ed.w
				heap.Push(&pq, vert{d: dist[ed.to], v: ed.to})
			}
		}
	}
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
