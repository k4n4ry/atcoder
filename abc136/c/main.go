package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ac
func main() {
	hash := make(map[string]int64, 0)
	N := getStdinInt64()
	var i int64 = 0
	for i = 0; i < N; i++ {
		s := getStdin()
		hash[sortString(s)]++
	}

	var ans int64 = 0
	for _, val := range hash {
		if val > 1 {
			ans += sumfrom1(val - 1)
		}
	}
	fmt.Println(ans)
}

func sumfrom1(val int64) int64 {
	var a int64 = 0
	var i int64 = 0
	for i = 0; i < val; i++ {
		a += i + 1
	}
	return a
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

var sc = bufio.NewReaderSize(os.Stdin, 1024*1024*10)

func getStdin() string {
	return readLine()
}

func getStdinInt() int {
	str := getStdin()
	rtn, _ := strconv.Atoi(str)
	return rtn
}

func getStdinInt64() int64 {
	str := getStdin()
	rtn, _ := strconv.ParseInt(str, 10, 64)
	return rtn
}
func getStdinIntArr() []int {
	str := getStdin()
	list := strings.Split(str, " ")
	rtn := make([]int, len(list))
	for idx, val := range list {
		rtn[idx], _ = strconv.Atoi(val)
	}
	return rtn
}
func getStdinIntArr64() []int64 {
	str := getStdin()
	list := strings.Split(str, " ")
	rtn := make([]int64, len(list))
	for idx, val := range list {
		rtn[idx], _ = strconv.ParseInt(val, 10, 64)
	}
	return rtn
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max64(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func abs64(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

func readLine() string {
	buf := make([]byte, 0, 0)
	for {
		l, p, _ := sc.ReadLine()
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}
