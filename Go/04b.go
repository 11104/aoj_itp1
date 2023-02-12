package main

import (
	"fmt"
	"math" //円周率に必要
)

//lint:ignore U1000 //

/*
func main() {
	fourb()
}

//*/

func fourb() {
	var r float64
	fmt.Scanf("%f", &r)
	p := float64(math.Pi)
	fmt.Printf("%.8f %.8f", p*r*r, 2*p*r)
}
