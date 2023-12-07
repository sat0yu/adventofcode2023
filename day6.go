package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func main() {
    var times []int
    var distances []int

    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    read_nums(sc, &times)
    sc.Scan()
    read_nums(sc, &distances)
    fmt.Printf("%v\n", times)
    fmt.Printf("%v\n", distances)
    mul := 1
    for i := 0; i < len(times); i++ {
        fmt.Printf("t=%d, d=%d\n", times[i], distances[i])
        mul *= search_patterns(times[i], distances[i])
    }
    fmt.Println(mul)
}

func read_nums(sc *bufio.Scanner, arr *[]int) error {
    lsc := bufio.NewScanner(strings.NewReader(sc.Text()))
    lsc.Split(bufio.ScanWords)
    lsc.Scan() // Times: or Distances:
    for i := 0; lsc.Scan(); i++ {
        if n, err := strconv.Atoi(lsc.Text()); err != nil {
            return err
        } else {
            *arr = append(*arr, n)
        }
    }
    return nil
}

func search_patterns(t int, d int) int {
    // Should be done with binary-search for greater integers
    var sum int
    for i := 1; i < t; i++ {
        j := t - i
        if i * j > d {
            sum += 1
        }
    }
    return sum
}
