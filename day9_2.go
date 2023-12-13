package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sum := 0
	for sc.Scan() {
		var arr []int
		lsc := bufio.NewScanner(strings.NewReader(sc.Text()))
		lsc.Split(bufio.ScanWords)
		for lsc.Scan() {
			w := lsc.Text()
			i, err := strconv.Atoi(w)
			if err != nil {
				panic(fmt.Sprintf("Failed to convert \"%s\" to int", w))
			}
			arr = append(arr, i)
		}
		diff := predict(arr, 0)
		sum += diff
	}
	fmt.Println(sum)
}

func predict(arr []int, d int) int {
	diffs := make([]int, len(arr)-1)
	flag := true
	for i := 1; i < len(arr); i++ {
		diffs[i-1] = arr[i] - arr[i-1]
		if diffs[i-1] != 0 {
			flag = false
		}
	}
	if flag {
		return arr[0]
	}
	diff := predict(diffs, d+1)
	prev := arr[0] - diff
	fmt.Printf("[%d] %v => diff: %d, prev: %d\n", d, arr, diff, prev)
	return prev
}
