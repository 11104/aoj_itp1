package main

import "fmt"

//lint:ignore U1000 //

/*
func main() {
	fiveb()
}

//*/

func fiveb() {
	var i, j, k, w, h int
	for k = 0; k < 10001; k++ {
		fmt.Scanf("%d%d", &w, &h)
		if w == 0 && h == 0 {
			break
		}
		for i = 0; i < w; i++ {
			for j = 0; j < h; j++ {
				if i == 0 || i == w-1 || j == 0 || j == h-1 {
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
