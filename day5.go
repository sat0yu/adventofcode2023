package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    "math"
)

func main() {
    var gen []int
    sc := bufio.NewScanner(os.Stdin)
    // seeds
    sc.Scan()
    lsc := bufio.NewScanner(strings.NewReader(sc.Text()))
    lsc.Split(bufio.ScanWords)
    lsc.Scan() // Seeds:
    for lsc.Scan() {
        if n, err := strconv.Atoi(lsc.Text()); err != nil {
            panic(err)
        }else{
            gen = append(gen, n)
        }
    }
    fmt.Printf("seeds: %v\n\n", gen)
    sc.Scan() // while line

    // load and apply map one by one
    for sc.Scan() {
        m := load_map(sc)
        fmt.Printf("map: %v\n",m)
        for i, item := range gen {
            gen[i] = apply_map(m, item)
        }
        fmt.Printf("gen: %v\n\n", gen)
    }

    // calc. the result location
    lo := math.MaxInt64
    for _, item := range gen {
        if item < lo {
            lo = item
        }
    }
    fmt.Printf("location: %d\n", lo)
}

func load_map(sc *bufio.Scanner) [][3]int {
    fmt.Println(sc.Text())
    var m [][3]int
    for sc.Scan() {
        row := sc.Text()
        if len(row) == 0 {
            break
        }
        var triplet [3]int
        fmt.Sscanf(row, "%d %d %d", &triplet[0], &triplet[1], &triplet[2])
        // fmt.Printf("triplet: %v\n",triplet)
        m = append(m, triplet)
    }
    return m
}

func apply_map(m [][3]int, v int) int {
    for _, f := range m {
        if f[1] <= v && v < f[1] + f[2] {
            return (v - f[1]) + f[0]
        }
    }
    return v
}
