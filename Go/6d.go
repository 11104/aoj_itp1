package main

import "fmt"

//lint:ignore U1000 //

/*
func main() {
	sixd()
}

//*/

func sixd() {
	var n, m, i, j int
	fmt.Scanf("%d %d", &n, &m)
	a := [101][101]int{}
	b := [101]int{}
	for i = 0; i < n; i++ {
		for j = 0; j < m; j++ {
			fmt.Scanf("%d", &a[i][j])
		}
	}
	for j = 0; j < m; j++ {
		fmt.Scanf("%d", &b[j])
	}
	c := [101]int{}
	for i = 0; i < n; i++ {
		for j = 0; j < m; j++ {
			c[i] += a[i][j] * b[j]
		}
		fmt.Println(c[i])
	}
}
