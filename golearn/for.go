


package main

import (
    "fmt"
    "math"
    "runtime"
)

func sqrt(x float64) string {
    if x < 0 {
        return sqrt(-x) +  "i"
    }

    return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n);  v < lim {
        return v;
    }
    return lim
}

func pow2(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    } else {
        fmt.Printf("%g >= %g\n", v, lim)
    }
    return lim
}

func Sqrt(x float64) float64 {
    z0 := float64(2)
    z1 := float64(1)
    v := math.Abs(z1 -z0)
    for v > 0.01 {
        z0 = z1
        z1 = z0 - ((z0 * z0) - x)/(2*z0);
        fmt.Println("z1 ", z1, " z0 ", z0, " z1 - z0 ", math.Abs(z1-z0))
        v = math.Abs(z1 -z0)
    }
    return z1
}

func main() {

    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)

    sum2 := 1
    for sum2 < 1000 {
        sum2 += sum2
    }
    fmt.Println(sum2)

    fmt.Println(sqrt(2), sqrt(-4))

    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
    )

    fmt.Println(
        pow2(3, 2, 10),
        pow2(3, 3, 20),
    )

    Sqrt(2)

    fmt.Print("Go runs on ")
    switch os := runtime.GOOS; os {
        case "darwin":
            fmt.Println("OS X.")
        case "linux":
            fmt.Println("OS X.")
        default:
            fmt.Printf("%s.\n", os)           
    }
    







}
