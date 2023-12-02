package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func main() {
    sc := bufio.NewScanner(os.Stdin)
    var sum int
    for sc.Scan() {
        s := sc.Text()
        t1 := strings.Split(s, ": ")
        t2 := strings.Split(t1[0], " ")
        id, _ := strconv.Atoi(t2[1])
        var max_r, max_g, max_b int
        for _, t := range strings.Split(t1[1], "; ") {
            for _, tt := range strings.Split(t, ", ") {
                ttt := strings.Split(tt, " ")
                if ttt[1] == "red" {
                    r, _ := strconv.Atoi(ttt[0])
                    if r > max_r {
                        max_r = r
                    }
                } else if ttt[1] == "green" {
                    g, _ := strconv.Atoi(ttt[0])
                    if g > max_g {
                        max_g = g
                    }
                } else if ttt[1] == "blue" {
                    b, _ := strconv.Atoi(ttt[0])
                    if b > max_b {
                        max_b = b
                    }
                } else {
                    fmt.Printf("unexpected input\n", tt)
                }
            }
        }
        score := max_r * max_g * max_b
        fmt.Printf("GAME %d: %d (r: %d, g: %d, b: %d)\n", id, score, max_r, max_g, max_b)
        sum += score
    }
    fmt.Printf("sum: %d\n", sum)
}
