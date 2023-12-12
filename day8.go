package main

import (
	"bufio"
	"fmt"
	"os"
)

const A = 26
const AA = A * A
const AAA = AA * A

var graph [A * A * A]node

var guide []int

type node [2]int

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	row := sc.Text() // RLRLRLRL....
	for _, ch := range row {
		if ch == 'L' {
			guide = append(guide, 0)
		} else if ch == 'R' {
			guide = append(guide, 1)
		} else {
			panic("unexpected error.")
		}
	}
	// fmt.Println(guide)
	sc.Scan() // empty line
	var t [3]string
	for sc.Scan() {
		row := sc.Text()
		fmt.Sscanf(row, "%3s = (%3s, %3s)", &t[0], &t[1], &t[2])
		n := int(t[0][0]-'A')*AA + int(t[0][1]-'A')*A + int(t[0][2]-'A')
		l := int(t[1][0]-'A')*AA + int(t[1][1]-'A')*A + int(t[1][2]-'A')
		r := int(t[2][0]-'A')*AA + int(t[2][1]-'A')*A + int(t[2][2]-'A')
		graph[n] = node{l, r}
		fmt.Println(t[0], t[1], t[2], n, graph[n])
	}
	var steps, guide_len uint
	guide_len = uint(len(guide))
	for i := 0; i != AAA-1; steps++ {
		lf := guide[steps%guide_len]
		i = graph[i][lf]
		// fmt.Println(steps, i)
	}
	fmt.Println(steps)
}
