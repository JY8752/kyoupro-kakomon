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

func main() {
	if close != nil {
		defer close()
	}

	var (
		n int
		a []int
	)

	n = scanInt(sc)
	a = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt(sc)
	}

	sort.Ints(a)

	var result int
	for i := 0; i < n; i++ {
		if i < n-1 && a[i+1]-a[i] > 1 {
			result = a[i] + 1
		}
	}

	fmt.Println(result)
}
