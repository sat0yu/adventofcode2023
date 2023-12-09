package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type player struct {
	hand [5]int
	bid  int
}

func (p player) calc_type() int {
	m := make(map[int]int)
	for _, h := range p.hand {
		m[h] += 1
	}
	var keys []int
	var max_count int
	for k := range m {
		keys = append(keys, k)
		if max_count < m[k] {
			max_count = m[k]
		}
	}
	if max_count == 5 {
		return 6
	} else if max_count == 4 {
		return 5
	} else if max_count == 3 {
		if len(keys) == 2 {
			return 4
		} else {
			return 3
		}
	} else if max_count == 2 {
		if len(keys) == 3 {
			return 2
		} else {
			return 1
		}
	}
	return 0
}

func (p player) calc_score() int {
	var s int
	for i, h := range p.hand {
		s += h * int(math.Pow(14.0, float64(4-i)))
	}
	return s
}

type game []player

func (g game) sort() {
	sort.Slice(g, func(i, j int) bool { return g[i].calc_score() < g[j].calc_score() })
}

func main() {
	var input game
	var h [5]byte
	var n [5]int
	var b int
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		row := sc.Text()
		fmt.Sscanf(row, "%c%c%c%c%c %d", &h[0], &h[1], &h[2], &h[3], &h[4], &b)
		for i, c := range h {
			if c == 'A' {
				n[i] = 14
			} else if c == 'K' {
				n[i] = 13
			} else if c == 'Q' {
				n[i] = 12
			} else if c == 'J' {
				n[i] = 11
			} else if c == 'T' {
				n[i] = 10
			} else {
				n[i] = int(c - '0')
			}
		}
		p := player{n, b}
		input = append(input, p)
	}
	var games [7]game
	for _, p := range input {
		t := p.calc_type()
		// fmt.Printf("player: %v \ttype: %d\n", p, t)
		games[t] = append(games[t], p)
	}
	var sum int
	var i int
	for idx, g := range games {
		fmt.Printf("===== type: %d (%d-players) =====\n", idx, len(g))
		g.sort()
		for j, p := range g {
			sum += (i + j + 1) * p.bid
			fmt.Println(i+j+1, p.hand, p.bid, sum)
		}
		i += len(g)
	}
	fmt.Println(sum)
}
