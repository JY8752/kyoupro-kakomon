package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	sc       *bufio.Scanner = bufio.NewScanner(os.Stdin)
	dataPath                = "../tests/sample-1.in"
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

	N := scanInt(sc)

	var list []int
	for i := 1; i <= 9; i++ {
		if N%i == 0 {
			list = append(list, i)
		}
	}

	var sb strings.Builder
	sb.Grow(N + 1)

	for i := 0; i < N+1; i++ {
		var found bool
		for _, j := range list {
			if i%(N/j) == 0 {
				sb.WriteString(strconv.Itoa(j))
				found = true
				break
			}
		}
		if !found {
			sb.WriteString("-")
		}
	}
	fmt.Println(sb.String())
}
