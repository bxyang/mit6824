


package main

import (
    "fmt"
    "math/cmplx"
)

var (
    ToBe    bool   = false
    MaxInt  uint64 = 1<<64 - 1
    z       complex128  = cmplx.Sqrt(-5 + 12i)
)



func main() {
    var i, j int = 1, 2
    k := 3
    c, python, java := true, false, "no!"

    fmt.Println(i, j, k, c, python, java)

    const f = "%T(%v)\n"

    fmt.Printf(f, ToBe, ToBe)
    fmt.Printf(f, MaxInt, MaxInt)
    fmt.Printf(f, z, z)
}
