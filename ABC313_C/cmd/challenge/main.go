package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func scan(scanner *bufio.Scanner) {
	if !scanner.Scan() {
		log.Fatal(scanner.Err())
	}
}

func max(arr []int) int {
	var max int
	for i, val := range arr {
		if i == 0 {
			max = val
			continue
		}

		if val > max {
			max = val
		}
	}
	return max
}

func min(arr []int) int {
	var min int
	for i, val := range arr {
		if i == 0 {
			min = val
			continue
		}

		if val < min {
			min = val
		}
	}
	return min
}

func getMaxIndex(arr []int) int {
	var max, index int

	for i, val := range arr {
		if i == 0 {
			max, index = val, i
			continue
		}

		if val > max {
			max, index = val, i
		}
	}
	return index
}

func getMinIndex(arr []int) int {
	var min, index int

	for i, val := range arr {
		if i == 0 {
			min, index = val, i
			continue
		}

		if val < min {
			min, index = val, i
		}
	}
	return index
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// input1 int value N
	scan(scanner)
	_, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalf("input value must be integer. err: %s\n", err.Error())
	}

	// input2 sequece of int An
	scan(scanner)
	strArr := strings.Split(scanner.Text(), " ")
	intArr := make([]int, len(strArr))
	for i, s := range strArr {
		ii, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("input value must be sequence of integers. err: %s\n", err.Error())
		}
		intArr[i] = ii
	}

	var count int
	for {
		fmt.Printf("intArr: %v\n", intArr)
		fmt.Println(count)
		if max(intArr)-min(intArr) <= 1 {
			break
		}

		// 最大値を1減らす
		maxi := getMaxIndex(intArr)
		intArr[maxi] -= 1

		// 最小値を1増やす
		mini := getMinIndex(intArr)
		intArr[mini] += 1

		count++
	}

	fmt.Println(count)
}
