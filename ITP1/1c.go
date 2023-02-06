package main

import "fmt"

//lint:ignore U1000 //
/*
	func main() {
		onec()
	}
//*/
func onec() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(a*b, 2*a+2*b)
	//Printlnはオペラント","の間にスペースが入る
}
