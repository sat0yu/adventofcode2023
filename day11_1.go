package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type galaxy struct {
	i, j, idx int
}

func main() {
	var acc_expansion_count_i, acc_expansion_count_j []int
	var s [][]byte
	var galaxies []*galaxy
	sc := bufio.NewScanner(os.Stdin)
	for i, idx := 0, 0; sc.Scan(); i++ {
		row := []byte(sc.Text())
		should_expand := true
		for j := 0; j < len(row); j++ {
			if row[j] == '#' {
				g := galaxy{i, j, idx}
				galaxies = append(galaxies, &g)
				fmt.Printf("%d: %v\n", idx, g)
				idx++
				should_expand = false
			}
		}
		prev := 0
		if i > 0 {
			prev = acc_expansion_count_i[i-1]
		}
		if should_expand {
			acc_expansion_count_i = append(acc_expansion_count_i, prev+1)
		} else {
			acc_expansion_count_i = append(acc_expansion_count_i, prev)
		}
		s = append(s, row)
	}
	for j := 0; j < len(s[0]); j++ {
		prev := 0
		if j > 0 {
			prev = acc_expansion_count_j[j-1]
		}
		should_expand := true
		for i := 0; i < len(s); i++ {
			if s[i][j] == '#' {
				should_expand = false
			}
		}
		if should_expand {
			acc_expansion_count_j = append(acc_expansion_count_j, prev+1)
		} else {
			acc_expansion_count_j = append(acc_expansion_count_j, prev)
		}
	}
	var dist int
	for i, g := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			h := galaxies[j]
			diff_i := math.Abs(float64(g.i - h.i))
			diff_j := math.Abs(float64(g.j - h.j))
			expansion_i := math.Abs(float64(acc_expansion_count_i[g.i] - acc_expansion_count_i[h.i]))
			expansion_j := math.Abs(float64(acc_expansion_count_j[g.j] - acc_expansion_count_j[h.j]))
			d := int(diff_i+diff_j) + int(expansion_i) + int(expansion_j)
			fmt.Printf("(%d, %d): %d (diff_i:%d, diff_j:%d, expansion_i:%d, expansion_j:%d)\n", g.idx, h.idx, d, diff_i, diff_j, expansion_i, expansion_j)
			dist += d
		}
	}
	fmt.Println(dist)
}
