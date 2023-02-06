package main

import "fmt"

//lint:ignore U1000 //

// /*
func main() {
	twod()
}

//*/

func twod() {
	var x, y, w, h, r int
	fmt.Scanf("%d%d%d%d%d", &w, &h, &x, &y, &r)
	if 0 <= (x-r) && (x+r) <= h && 0 <= (y-r) && (y+r) <= h {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
