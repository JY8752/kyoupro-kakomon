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

func main() {
	if close != nil {
		defer close()
	}

	var (
		n int
		g [100][100]bool
	)

	n = scanInt(sc)
	for i := 0; i < n; i++ {
		aa, bb, cc, dd := scanInt(sc), scanInt(sc), scanInt(sc), scanInt(sc)
		for j := aa; j < bb; j++ {
			for k := cc; k < dd; k++ {
				g[j][k] = true
			}
		}
	}

	var counter int
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if g[i][j] {
				counter++
			}
		}
	}

	fmt.Println(counter)
}
