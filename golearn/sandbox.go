



package main 

import (
    "fmt"
    "time"
    "math/rand"
    "math"
)


func add(x, y int) int {
    return x + y
}

func swap(x, y string) (string, string) {
    return y, x
}

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

var c, python, java bool

var m, n int = 1, 2
func main() {
    
    fmt.Println("Welcom to the playground!")
    fmt.Println("The time is", time.Now())
    fmt.Println("My favorite number is", rand.Intn(10))
    fmt.Println("Now you have %g problems.", math.Sqrt(7))
    fmt.Println(math.Pi)
    fmt.Println(add(42, 13))
    
    a, b := swap("hello", "world")
    fmt.Println(a, b)

    fmt.Println(split(17))

    var i int
    fmt.Println(i, c, python, java)

    var bc, bp, sj = true, false, "no!"
    fmt.Println(m, n, bc, bp, sj)
}
