package main

import (
	"fmt"
	"math"
)

//lint:ignore U1000 //

/*
func main() {
	tena()
}

//*/

func tena() {
	var x1, x2, y1, y2 float64
	fmt.Scan(&x1, &y1, &x2, &y2)
	fmt.Println(math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
}
