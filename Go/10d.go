package main

import (
	"fmt"
	"math"
)

//lint:ignore U1000 //

// /*
func main() {
	tend()
}

//*/

func tend() {
	var n int
	x := [1001]float64{}
	y := [1001]float64{}
	fmt.Scanf("%d", &n)
	var d1, d2, d3, d0 float64
	for i := 0; i < n; i++ {
		fmt.Scanf("%f", &x[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scanf("%f", &y[i])
	}
	for i := 0; i < n; i++ {
		d1 += math.Abs(x[i] - y[i])
		d2 += math.Abs(x[i]-y[i]) * math.Abs(x[i]-y[i])
		d3 += math.Abs(x[i]-y[i]) * math.Abs(x[i]-y[i]) * math.Abs(x[i]-y[i])
		if d0 < math.Abs(x[i]-y[i]) {
			d0 = math.Abs(x[i] - y[i])
		}
	}
	d2 = math.Pow(d2, 1.0/2.0)
	d3 = math.Pow(d3, 1.0/3.0)
	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Println(d0)
}
