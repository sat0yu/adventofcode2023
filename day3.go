package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

var sum int
var m [][]byte

func main() {
    sc := bufio.NewScanner(os.Stdin)
    for i := 0; sc.Scan(); i++ {
        m = append(m, []byte(sc.Text()))
        // if i > 0 {
        //     fmt.Printf("row %03d: %s\n", i-1, m[i-1])
        // }
        fmt.Printf("row %03d: %s\n", i, m[i])
        for j := 0; i > 0 && j < len(m[i-1]); {
            c := m[i-1][j]
            // fmt.Printf("i=%d, j=%d, c=%c\n", i, j, c)
            if !is_number(c) {
                j++
                continue
            }
            flag, d := search(i-1, j, false)
            n := m[i-1][j:j+d]
            fmt.Printf("n=%s, i=%d, j=%d, d=%d, flag=%v\n", n, i, j, d, flag)
            j += d
            if flag {
                atoi, err := strconv.Atoi(string(n))
                if err != nil {
                    panic(err)
                }
                sum += atoi
                fmt.Printf("sum=%d, j=%d\n", sum, j)
            }
        }
        // if i > 0 {
        //     fmt.Printf("row %03d: %s\n", i-1, m[i-1])
        // }
    }
    last_row_idx := len(m) - 1
    m = append(m, m[last_row_idx-1])
    for j := 0; j < len(m[last_row_idx]); {
        c := m[last_row_idx][j]
        if !is_number(c) {
            j++
            continue
        }
        flag, d := search(last_row_idx, j, false)
        n := m[last_row_idx][j:j+d]
        fmt.Printf("n=%s, i=%d, j=%d, d=%d, flag=%v\n", n, last_row_idx, j, d, flag)
        j += d
        if flag {
            atoi, err := strconv.Atoi(string(n))
            if err != nil {
                panic(err)
            }
            sum += atoi
            fmt.Printf("sum=%d, j=%d\n", sum, j)
        }
    }
    fmt.Printf("sum=%d\n", sum)
}

func search(y, x int, flag bool) (bool, int) {
    if !is_number(m[y][x]) {
        panic(fmt.Sprintf("invalid arguments. x=%d y=%d", x, y))
    }
    if !flag {
        if x > 0 {
            if y > 0 && is_valid_symbol(m[y-1][x-1]) {
                flag = true
            }
            if is_valid_symbol(m[y][x-1]) {
                flag = true
            }
            if is_valid_symbol(m[y+1][x-1]) {
                flag = true
            }
        }
        if y > 0 && is_valid_symbol(m[y-1][x]) {
            flag = true
        }
        if is_valid_symbol(m[y+1][x]) {
            flag = true
        }
        if x < len(m[y]) - 1 {
            if y > 0 && is_valid_symbol(m[y-1][x+1]) {
                flag = true
            }
            if is_valid_symbol(m[y+1][x+1]) {
                flag = true
            }
            if is_valid_symbol(m[y][x+1]) {
                flag = true
            }
        }
    }
    d := 1
    if x < len(m[y]) - 1 && is_number(m[y][x+1]) {
        ff, dd := search(y, x+1, flag)
        d += dd
        flag = flag || ff
    }
    return flag, d
}

func is_valid_symbol(c byte) bool {
    return !is_number(c) && c != '.'
}

func is_number(c byte) bool {
    return '0' <= c && c <= '9'
}
