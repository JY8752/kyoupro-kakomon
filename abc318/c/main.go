package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	sc       *bufio.Scanner = bufio.NewScanner(os.Stdin)
	dataPath                = "./data/data1.txt"
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

func chmin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	if close != nil {
		defer close()
	}

	var (
		n, d, p int
		f, s    []int
	)

	n, d, p = scanInt(sc), scanInt(sc), scanInt(sc)
	f = make([]int, n)
	for i := 0; i < n; i++ {
		f[i] = scanInt(sc)
	}

	sort.Ints(f)

	// kセットで固定したときのN日間の最小値を考える
	// 通常料金をi回払ったときの合計料金をSiとすると
	s = make([]int, n)
	s[0] = f[0]
	for i := 0; i < n-1; i++ {
		s[i+1] = s[i] + f[i+1]
	}

	// k := n / d
	k := (n + d - 1) / d
	ans := p * k

	for i := 0; i < k; i++ {
		ans = chmin(ans, s[n-1-(i*d)]+p*i)
	}

	fmt.Println(ans)
}
