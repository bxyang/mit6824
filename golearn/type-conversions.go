


package main 

import (
    "fmt"
    "math"
)


const (
    Big = 1 << 100
    Small = Big >> 99
)

func needInt(x int) int { return x * 10 + 1}
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
    var x, y int = 3, 4
    var f float64 = math.Sqrt(float64(x*x + y*y))
    var z uint  = uint(f)
    fmt.Println(x, y, z)

    v := 42
    fmt.Printf("v is of type %T\n", v)


    fmt.Println(needInt(Small))
    fmt.Println(needFloat(Small))
    fmt.Println(needFloat(Big))
}
