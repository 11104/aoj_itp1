package main

import "fmt"

// /*
//
//lint:ignore U1000 //
func main() {
	oned()
}

// //lint:ignore U1000 //*/
func oned() {
	var d, h, m, s int
	fmt.Scanf("%d", &d)
	h = d / 3600
	d %= 3600
	m = d / 60
	s = d % 60
	fmt.Printf("%d:%d:%d\n", h, m, s)
}
