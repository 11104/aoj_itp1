package main

import "fmt"

//lint:ignore U1000 //

/*
	func main() {
		twoa()
	}

//*/

func twoa() {
	var a, b int
	fmt.Scanf("%d%d", &a, &b)
	if a > b {
		fmt.Println("a > b")
	} else if a < b {
		fmt.Println("a < b")
	} else {
		fmt.Println("a == b")
	}
}
