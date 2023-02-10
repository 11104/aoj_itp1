package main

import "fmt"

//lint:ignore U1000 //

/*
func main() {
	sevenc()
}

//*/

func sevenc() {
	var r, c int
	t := [101][101]int{}
	fmt.Scanf("%d %d", &r, &c)
	for i := 0; i < r; i++ {
		for j := 0; j <= c; j++ {
			if j == c {
				for k := 0; k < c; k++ {
					t[i][j] += t[i][k]
				}
			} else {
				fmt.Scanf("%d", &t[i][j])
			}
		}
	}
	for j := 0; j < c; j++ {
		for i := 0; i < r; i++ {
			t[r][j] += t[i][j]
		}
		t[r][c] += t[r][j]
	}
	for i := 0; i <= r; i++ {
		for j := 0; j <= c; j++ {
			if j == c {
				fmt.Printf("%d\n", t[i][j])
			} else {
				fmt.Printf("%d ", t[i][j])
			}
		}
	}
}
