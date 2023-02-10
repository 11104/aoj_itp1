package main

import "fmt"

//lint:ignore U1000 //

/*
func main() {
	sevenb()
}

//*/

func sevenb() {
	var n, x, count int
	for i := 0; i < 10001; i++ {
		fmt.Scanf("%d %d", &n, &x)
		if n == 0 && x == 0 {
			break
		}
		count = 0
		for j := 1; j <= n-2; j++ {
			for k := j + 1; k <= n-1; k++ {
				for l := k + 1; l <= n; l++ {
					if j+k+l == x {
						count++
					}
				}
			}
		}
		fmt.Println(count)
	}
}
