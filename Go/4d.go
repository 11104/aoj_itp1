package main

import "fmt"

//lint:ignore U1000 //

/*
func main() {
	fourd()
}

//*/

func fourd() {
	var n, tmp int
	fmt.Scanf("%d", &n)
	max := -9999999
	min := 9999999
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &tmp)
		sum += tmp
		if max < tmp {
			max = tmp
		}
		if tmp < min {
			min = tmp
		}
	}
	fmt.Println(min, max, sum)
}
