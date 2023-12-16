package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	total := 0
	for {
		substr := make([]string, 2)
		i, err := fmt.Scanf("%s %s\n", &substr[0], &substr[1])
		if i != 2 || err != nil {
			break
		}

		var groups []int
		for _, c := range strings.Split(substr[1], ",") {
			if i, err := strconv.Atoi(c); err != nil {
				panic(err)
			} else {
				groups = append(groups, i)
			}
		}
		fmt.Printf("groups: %v\n", groups)

		wildcard_count := 0
		for _, c := range substr[0] {
			if c == '?' {
				wildcard_count++
			}
		}
		l := len(substr[0])
		subtotal := 0
		for i := uint(0); i < (1 << wildcard_count); i++ {
			s := uint64(0)
			for j, k := 0, 0; j < l; j++ {
				c := substr[0][j]
				if c == '#' {
					s |= 1 << j
				} else if c == '?' {
					if i&(1<<k) > 0 {
						s |= 1 << j
					}
					k++
				}
			}
			var g []int
			k := 0
			for j := 0; j < l; j++ {
				if s&(1<<j) > 0 {
					k++
				} else {
					if k > 0 {
						g = append(g, k)
						k = 0
					}
				}
			}
			if k > 0 {
				g = append(g, k)
			}
			// fmt.Printf("%08d: %032b, \t%v\n", i, s, g)
			if len(groups) == len(g) {
				match := true
				for j := 0; j < len(groups); j++ {
					if groups[j] != g[j] {
						match = false
						break
					}
				}
				if match {
					fmt.Printf("%08d: %032b, \t%v\n", i, s, g)
					subtotal++
				}
			}
		}
		fmt.Printf("subtotal: %d\n", subtotal)
		total += subtotal
	}
	fmt.Printf("total: %d\n", total)
}
