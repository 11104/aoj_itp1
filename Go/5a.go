package main

import "fmt"

//lint:ignore U1000 //

/*
func main() {
	fivea()
}

//*/

func fivea() {
	var i, j, w, h int
	for k := 0; k < 10001; k++ {
		fmt.Scanf("%d%d", &w, &h)
		if w == 0 && h == 0 {
			break
		}
		for i = 0; i < w; i++ {
			for j = 0; j < h; j++ {
				fmt.Printf("#")
			}
			fmt.Printf("\n")
		}
		fmt.Println()
	}
}
