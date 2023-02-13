package main

import (
	"fmt"
	"math"
)

//lint:ignore U1000 //

/*
func main() {
	tenc()
}

//*/

func tenc() {
	var n int
	var ave, dev, tmp float64
	s := [1001]float64{} //Err by no useing 1001
	for {
		ave = 0
		dev = 0
		fmt.Scanf("%d", &n)
		if n == 0 {
			break
		}
		for i := 0; i < n; i++ {
			fmt.Scanf("%f", &tmp)
			ave += tmp
			s[i] = tmp
		}
		ave /= float64(n)
		for i := 0; i < n; i++ {
			dev += (s[i] - ave) * (s[i] - ave)
		}
		fmt.Println(math.Sqrt(dev / float64(n)))
	}
}
