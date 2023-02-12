package main

import (
	"fmt"
	"math"
)

//lint:ignore U1000 //

/*
func main() {
	tenb()
}

//*/

func tenb() {
	var a, b, c, S, L, H float64
	fmt.Scan(&a, &b, &c)
	c = c * math.Pi / 180 //degree->radian
	S = a * b * math.Sin(c) / 2
	L = a + b + math.Sqrt(math.Pow(a, 2)+math.Pow(b, 2)-2*a*b*math.Cos(c))
	H = b * math.Sin(c)
	fmt.Println(S, L, H)
}
