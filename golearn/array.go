




package main

import (
    "fmt"
    "strings"
)


func main() {
    var a[2]string
    a[0] = "hello"
    a[1] = "world"
    fmt.Println(a[0], a[1])
    fmt.Println(a)


    s := []int{2, 3, 5, 7, 11, 13}
    fmt.Println("s == ", s)

    for i := 0; i < len(s); i++ {
        fmt.Printf("s[%d] == %d\n", i, s[i])
    }

    fmt.Println("s ==", s)
    fmt.Println("s[1:4] ==", s[1:4])
    fmt.Println("s[:3] ==", s[:3])
    fmt.Println("s[4:] ==", s[4:])


    game := [][]string {
        []string{"_", "_", "_"},
        []string{"_", "_", "_"},
        []string{"_", "_", "_"},
    }

    game[0][0] = "X"
    game[2][2] = "O"
    game[2][0] = "X"
    game[1][0] = "O"
    game[0][2] = "X"
    
    printBoard(game)

    aa := make([]int, 5)
    printSlice("aa", aa)
    b := make([]int, 0, 5)
    printSlice("b", b)
    c := b[:2]
    printSlice("c", c)
    d := c[2:5]
    printSlice("d", d)

    var z []int
    fmt.Println(z, len(z), cap(z))
    if z == nil {
        fmt.Println("nil!")
    }



}

func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func printBoard(s [][]string) {
    for i := 0; i < len(s); i++ {
        fmt.Printf("%s\n", strings.Join(s[i], " "))
    }
}








