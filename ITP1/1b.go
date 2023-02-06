package main

import "fmt"

//lint:ignore U1000 //
/*
	func main() {
		oneb()
	}
*/
func oneb() {
	var a int
	fmt.Scanf("%d", &a)
	a = a * a * a
	fmt.Println(a)
}
