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
    max_r := 12
    max_g := 13
    max_b := 14
    GAME:
    for sc.Scan() {
        s := sc.Text()
        t1 := strings.Split(s, ": ")
        t2 := strings.Split(t1[0], " ")
        id, _ := strconv.Atoi(t2[1])
        for _, t := range strings.Split(t1[1], "; ") {
            var r,g,b int
            for _, tt := range strings.Split(t, ", ") {
                ttt := strings.Split(tt, " ")
                if ttt[1] == "red" {
                    r, _ = strconv.Atoi(ttt[0])
                } else if ttt[1] == "green" {
                    g, _ = strconv.Atoi(ttt[0])
                } else if ttt[1] == "blue" {
                    b, _ = strconv.Atoi(ttt[0])
                } else {
                    fmt.Printf("unexpected input\n", tt)
                }
            }
            if max_r < r {
                fmt.Printf("ID: %d, too many red cubes: %d\n", id, r)
                continue GAME
            } else if max_g < g {
                fmt.Printf("ID: %d, too many green cubes: %d\n", id, g)
                continue GAME
            } else if max_b < b {
                fmt.Printf("ID: %d, too many blue cubes: %d\n", id, b)
                continue GAME
            }
        }
        fmt.Printf("%d\n", id)
        sum += id
    }
    fmt.Printf("sum: %d\n", sum)
}
