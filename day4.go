package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    "sort"
)

var sum int

func main() {
    sc := bufio.NewScanner(os.Stdin)
    for sc.Scan() {
        p := 0
        lsc := bufio.NewScanner(strings.NewReader(sc.Text()))
        lsc.Split(bufio.ScanWords)
        lsc.Scan() // Card
        lsc.Scan() // ${index}:
        var win_nums [10]int
        var given_nums [25]int
        read_and_sort_nums(lsc, win_nums[:])
        lsc.Scan() // |
        read_and_sort_nums(lsc, given_nums[:])
        fmt.Printf("%d: %v\n", win_nums)
        fmt.Printf("%d: %v\n", given_nums)
        for i,j := 0, 0; i < len(win_nums) && j < len(given_nums); {
            fmt.Printf("given_nums[j] = %d, win_nums[i] = %d\n", given_nums[j], win_nums[i])
            if given_nums[j] == win_nums[i] {
                fmt.Printf("matched!\n")
                if p == 0 {
                    p = 1
                } else {
                    p *= 2
                }
                i++
                j++
            } else if given_nums[j] < win_nums[i] {
                j++
            } else { // given_nums[j] > win_nums[i]
                i++
            }
        }
        fmt.Printf("p: %d\n", p)
        sum += p
    }
    fmt.Printf("sum: %d\n", sum)
}

func read_and_sort_nums(sc *bufio.Scanner, arr []int) error {
    for i := 0; i < len(arr); i++ {
        sc.Scan()
        if n, err := strconv.Atoi(sc.Text()); err != nil {
            return err
        } else {
            arr[i] = n
        }
    }
    sort.Ints(arr)
    return nil
}
