package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	sc       *bufio.Scanner = bufio.NewScanner(os.Stdin)
	dataPath                = "./data/data3.txt"
	close    func() error
)

func init() {
	isDebug := flag.Bool("d", false, "debug flag")
	flag.Parse()

	if *isDebug {
		debug(dataPath)
	}

	sc.Split(bufio.ScanWords)
}

func debug(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	sc = bufio.NewScanner(f)
	close = func() error { return f.Close() }
}

func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var (
	ans  int
	used []bool = make([]bool, 100)
)

func dfs(v, sum, n int, graph [][]int) {
	// 使用済み
	used[v] = true

	// 現在の最大値と現在の合計値で大きい方で更新しとく
	ans = max(ans, sum)

	for i := 1; i <= n; i++ {
		if !used[i] && graph[i][v] != 0 {
			dfs(i, sum+graph[i][v], n, graph)
		}
	}

	used[v] = false
}

func main() {
	if close != nil {
		defer close()
	}

	var (
		n, m  int
		graph [][]int
	)

	n = scanInt(sc)
	m = scanInt(sc)

	graph = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		graph[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		a := scanInt(sc)
		b := scanInt(sc)
		c := scanInt(sc)
		graph[a][b] = c
		graph[b][a] = c
	}

	for i := 1; i <= n; i++ {
		dfs(i, 0, n, graph)
	}

	fmt.Println(ans)
}
