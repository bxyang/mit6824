


package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m  = map[string]Vertex{

    "Google": Vertex {
        37.42202, -122.08408,
    },
}

func main() {
    const key string = "Bell Labs"
    m[key] = Vertex{
        40.68433,   -74.39967,
    }
    fmt.Println(m[key])
    fmt.Println(m)

    mm := make(map[string]int)

    mm["Answer"] = 42
    fmt.Println("The value:", mm["Answer"])

    mm["Answer"] = 48
    fmt.Println("The value:", mm["Answer"])
    
    delete(mm, "Answer")
    fmt.Println("The value:", mm["Answer"])
    
    v, ok := mm["Answer"]
    fmt.Println("The value:", v, "Present?", ok)
}
