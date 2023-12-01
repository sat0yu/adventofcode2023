package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
    sc := bufio.NewScanner(os.Stdin)
    var sum int
    for sc.Scan() {
        s := sc.Text()
        fmt.Printf("%s, len: %d\n", s, len(s))
        var l, r, c int
        i, length := 0, len(s)
        for i < length {
           if strings.HasPrefix(s[i:], "one") {
               c = 1
           } else if strings.HasPrefix(s[i:], "two") {
               c = 2
           } else if strings.HasPrefix(s[i:], "three") {
               c = 3
           } else if strings.HasPrefix(s[i:], "four") {
               c = 4
           } else if strings.HasPrefix(s[i:], "five") {
               c = 5
           } else if strings.HasPrefix(s[i:], "six") {
               c = 6
           } else if strings.HasPrefix(s[i:], "seven") {
               c = 7
           } else if strings.HasPrefix(s[i:], "eight") {
               c = 8
           } else if strings.HasPrefix(s[i:], "nine") {
               c = 9
           } else {
               c = int(s[i] - '0')
           }
           i += 1
           if 0 < c && c <= 9 {
               if l == 0 {
                   l = c
               }
               r = c
           }
        }
        fmt.Printf("%d, %d\n", 10 * l,r)
        sum += 10 * l + r
    }
    fmt.Printf("sum: %d\n", sum)
}
