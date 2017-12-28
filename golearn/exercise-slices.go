

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    
    var ret [][]uint = make([][]uint8, dx)
    for i:= 0; i < dx; i++ {
        ret[i] = make([]int, dy)
    }

    for i:= 0; i < dx; i++ {
        for j := 0; j < dy; j++ {
            ret[i][j] = x*y
        }
    }
}


func main() {
    pic.Show(Pic)
}
