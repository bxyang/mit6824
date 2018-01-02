

package main

import "fmt"

func adder() func(int) int {
    sum := 0
    return func (x int) int {
        sum += x
        return sum
    }
}


func fibonacci() func() int {
	sum1 := 0
	sum2 := 1
	return func () int {
		t := sum1 + sum2
		sum1 = sum2
		sum2 = t
		return t
	}
}


func main() {
    pos, neg := adder(), adder()

    for i := 0; i < 10; i++ {
        fmt.Println(
            pos(i),
            neg(-2*i),
        )        
    }

    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
