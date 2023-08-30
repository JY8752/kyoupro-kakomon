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

func main() {
	if close != nil {
		defer close()
	}

	var (
		n, h, x int
		p       []int
	)

	n = scanInt(sc)
	h = scanInt(sc)
	x = scanInt(sc)

	p = make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = scanInt(sc)
	}

	var result int
	for i := 0; i < n; i++ {
		if p[i]+h >= x {
			result = i + 1
			break
		}
	}

	fmt.Println(result)
}
