


package main

import "fmt"





func main() {

    var a []int
    printSlice("a", a)    

    a = append(a, 0)
    printSlice("a", a)    
    
    a = append(a, 2, 3, 4)
    printSlice("a", a)    


    var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

    for i, v := range pow {
        fmt.Printf("2**%d = %d\n", i, v)
    }

    pov := make([]int, 10)
    for i := range pov {
        pov[i] = 1 << uint(i)
    }

    for _, value := range pov {
        fmt.Printf("%d\n", value)
    }

}

func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
