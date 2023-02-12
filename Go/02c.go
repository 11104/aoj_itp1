package main

import "fmt"

//lint:ignore U1000 //

/*
	func main() {
		twoc()
	}

//
*/
func twoc() {
	var a, b, c int
	fmt.Scanf("%d%d%d", &a, &b, &c)
	if a < b {
		if b < c {
			fmt.Println(a, b, c)
		} else if a < c {
			fmt.Println(a, c, b)
		} else {
			fmt.Println(c, a, b)
		}
	} else {
		if c < b {
			fmt.Println(c, b, a)
		} else if a < c {
			fmt.Println(b, a, c)
		} else {
			fmt.Println(b, c, a)
		}
	}
}
