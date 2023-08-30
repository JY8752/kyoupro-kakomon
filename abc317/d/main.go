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
	dataPath                = "./data/data4.txt"
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

func sum(ints []int) int {
	var m int
	for _, i := range ints {
		m += i
	}
	return m
}

func chmin(a *int, b int) {
	if *a > b {
		*a = b
	}
}

const inf = int(1e18)

func main() {
	if close != nil {
		defer close()
	}

	var (
		n int
		x []int
		y []int
		z []int
	)

	n = scanInt(sc)

	x = make([]int, n)
	y = make([]int, n)
	z = make([]int, n)

	for i := 0; i < n; i++ {
		x[i] = scanInt(sc)
		y[i] = scanInt(sc)
		z[i] = scanInt(sc)
	}

	// zの合計(最低限必要な議席数)
	z_sum := sum(z)

	// dpテーブル
	dp := make([]int, z_sum+1)
	for i := 0; i < z_sum+1; i++ {
		dp[i] = inf
	}
	dp[0] = 0

	// dpテーブルを埋めていく
	for i := 0; i < n; i++ {
		// w = 鞍替えする人数
		var w int
		if x[i] < y[i] {
			w = (x[i]+y[i])/2 + 1 - x[i]
		}
		for j := z_sum; j >= z[i]; j-- {
			// fmt.Println(dp[j-z[i]] + w)
			chmin(&dp[j], dp[j-z[i]]+w)
		}
	}

	result := inf
	for j := z_sum/2 + 1; j <= z_sum; j++ {
		chmin(&result, dp[j])
	}

	// fmt.Printf("%v", dp)
	fmt.Println(result)
}
