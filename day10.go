package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	i, j, d int
}

func main() {
	var m [][]byte
	var q []*point
	sc := bufio.NewScanner(os.Stdin)
	var st *point
	for i := 0; sc.Scan(); i++ {
		row := []byte(sc.Text())
		for j := 0; j < len(row); j++ {
			if row[j] == 'S' {
				st = &point{i, j, 0}
				q = append(q, &point{i, j, 1})
			}
		}
		fmt.Printf("%d: %s\n", i, row)
		m = append(m, row)
	}
	c := make([][]bool, len(m))
	for i := 0; i < len(c); i++ {
		c[i] = make([]bool, len(m[i]))
	}
	if st == nil {
		panic("No start point")
	}
	var dist int
	for len(q) > 0 {
		h := q[0]
		q = q[1:]
		c[h.i][h.j] = true
		ch := m[h.i][h.j]
		fmt.Printf("(%3d, %3d) [%4d]: %c\n", h.i, h.j, h.d, ch)
		if dist < h.d {
			dist = h.d
		}
		if ch == 'S' {
			if 0 < h.i {
				q = append(q, &point{h.i - 1, h.j, h.d + 1})
			}
			if h.i < len(m) {
				q = append(q, &point{h.i + 1, h.j, h.d + 1})
			}
			if 0 < h.j {
				q = append(q, &point{h.i, h.j - 1, h.d + 1})
			}
			if h.j < len(m[h.i]) {
				q = append(q, &point{h.i, h.j + 1, h.d + 1})
			}
		} else if ch == 'F' {
			if h.i < len(m) && !c[h.i+1][h.j] {
				q = append(q, &point{h.i + 1, h.j, h.d + 1})
			}
			if h.j < len(m[h.i]) && !c[h.i][h.j+1] {
				q = append(q, &point{h.i, h.j + 1, h.d + 1})
			}
		} else if ch == '7' {
			if h.i < len(m) && !c[h.i+1][h.j] {
				q = append(q, &point{h.i + 1, h.j, h.d + 1})
			}
			if 0 < h.j && !c[h.i][h.j-1] {
				q = append(q, &point{h.i, h.j - 1, h.d + 1})
			}
		} else if ch == 'L' {
			if 0 < h.i && !c[h.i-1][h.j] {
				q = append(q, &point{h.i - 1, h.j, h.d + 1})
			}
			if h.j < len(m[h.i]) && !c[h.i][h.j+1] {
				q = append(q, &point{h.i, h.j + 1, h.d + 1})
			}
		} else if ch == 'J' {
			if 0 < h.i && !c[h.i-1][h.j] {
				q = append(q, &point{h.i - 1, h.j, h.d + 1})
			}
			if 0 < h.j && !c[h.i][h.j-1] {
				q = append(q, &point{h.i, h.j - 1, h.d + 1})
			}
		} else if ch == '|' {
			if 0 < h.i && !c[h.i-1][h.j] {
				q = append(q, &point{h.i - 1, h.j, h.d + 1})
			}
			if h.i < len(m) && !c[h.i+1][h.j] {
				q = append(q, &point{h.i + 1, h.j, h.d + 1})
			}
		} else if ch == '-' {
			if 0 < h.j && !c[h.i][h.j-1] {
				q = append(q, &point{h.i, h.j - 1, h.d + 1})
			}
			if h.j < len(m[h.i]) && !c[h.i][h.j+1] {
				q = append(q, &point{h.i, h.j + 1, h.d + 1})
			}
		} else if ch == '.' {
			// do nothing
		} else {
			panic(fmt.Sprintf("Unknown char: %c", ch))
		}
	}
	fmt.Println(dist)
}
