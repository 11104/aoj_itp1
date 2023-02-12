package main

import "fmt"

//lint:ignore U1000 //

/*
func main() {
	sevend()
}

//*/

func sevend() {
	var n, m, l int
	a := [101][101]int{}
	b := [101][101]int{}
	c := [101][101]int{}
	fmt.Scanf("%d %d %d", &n, &m, &l)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &a[i][j])
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < l; j++ {
			fmt.Scanf("%d", &b[i][j])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < l; j++ {
			for k := 0; k < m; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < l; j++ {
			if j == l-1 {
				fmt.Printf("%d\n", c[i][j])
			} else {
				fmt.Printf("%d ", c[i][j])
			}
		}
	}
}
