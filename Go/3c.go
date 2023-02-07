package main

import "fmt"

//lint:ignore U1000 //

// /*
func main() {
	threec()
}

//*/

func threec() {
	var a, b int
	for i := 0; i < 10000; i++ {
		fmt.Scanf("%d %d", &a, &b)
		if a == 0 && b == 0 {
			break
		} else if a > b {
			fmt.Println(b, a)
		} else {
			fmt.Println(a, b)
		}
	}
}
