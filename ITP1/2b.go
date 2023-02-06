package main

import "fmt"

//lint:ignore U1000 //

/*
	func main() {
		twob()
	}

//
*/
func twob() {
	var a, b, c int
	fmt.Scanf("%d%d%d", &a, &b, &c)
	if a < b && b < c {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
