package main

import "fmt"

//lint:ignore U1000 //

/*
func main() {
	fivec()
}

//*/

func fivec() {
	var i, j, k, w, h int
	for k = 0; k < 10001; k++ {
		fmt.Scanf("%d%d", &w, &h)
		if w == 0 && h == 0 {
			break
		}
		for i = 1; i <= w; i++ {
			for j = 1; j <= h; j++ {
				if i%2 != 0 && j%2 != 0 || i%2 == 0 && j%2 == 0 {
					fmt.Printf("#")
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Printf("\n")
		}
		fmt.Println()
	}
}
