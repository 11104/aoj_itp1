package main

import "fmt"

//lint:ignore U1000 //

/*
func main() {
	sixc()
}

//*/

func sixc() {
	var n, b, f, r, v int
	fmt.Scanf("%d", &n)
	m := [4][3][10]int{}
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d %d", &b, &f, &r, &v)
		m[b-1][f-1][r-1] += v
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 10; k++ {
				fmt.Printf(" %d", m[i][j][k])
			}
			fmt.Printf("\n")
		}
		if i != 3 {
			for k := 0; k < 20; k++ {
				fmt.Printf("#")
			}
			fmt.Printf("\n")
		}
	}
}
