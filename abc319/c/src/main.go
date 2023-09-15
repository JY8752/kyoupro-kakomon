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

func main() {
	if close != nil {
		defer close()
	}

	var (
		cells [9]int
	)

	for i := 0; i < 9; i++ {
		cells[i] = scanInt(sc)
	}

	rows := [8][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	order := [9]int{
		0, 1, 2, 3, 4, 5, 6, 7, 8,
	}

	var (
		notDisappointedCount int
		all                  int
	)

	for {
		all++
		var isDisappoint bool
		for _, row := range rows {
			a, b, c := row[0], row[1], row[2]
			switch {
			case cells[a] == cells[b] && order[a] < order[c] && order[b] < order[c]:
				isDisappoint = true
			case cells[a] == cells[c] && order[a] < order[b] && order[c] < order[b]:
				isDisappoint = true
			case cells[b] == cells[c] && order[b] < order[a] && order[c] < order[a]:
				isDisappoint = true
			}
		}

		if !isDisappoint {
			notDisappointedCount++
		}

		if !nextPermutation(order[:]) {
			break
		}
	}

	fmt.Printf("%.10f\n", float64(notDisappointedCount)/float64(all))
}

func nextPermutation(nums []int) bool {
	n := len(nums)

	if n < 2 {
		return false
	}

	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i == -1 {
		return false
	}

	j := n - 1
	for nums[j] <= nums[i] {
		j--
	}

	// 要素を入れ替える
	nums[i], nums[j] = nums[j], nums[i]

	// 逆順
	reverse(nums[i+1:])

	return true
}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
