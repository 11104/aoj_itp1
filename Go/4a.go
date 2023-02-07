package main

import "fmt"

//lint:ignore U1000 //

/*
func main() {
	foura()
}

//*/

func foura() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	d := int(a / b)
	r := int(a % b)
	f := float64(a) / float64(b)
	fmt.Printf("%d %d %.8f\n", d, r, f)
}
