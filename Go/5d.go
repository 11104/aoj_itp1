package main

import "fmt"

//lint:ignore U1000 //

/*
func main() {
	fived()
}

//*/

func fived() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%10 == 3 || i/10%10 == 3 || i/100%10 == 3 || i/1000%10 == 3 {
			fmt.Printf(" %d", i)
		}
	}
	fmt.Println()
}
