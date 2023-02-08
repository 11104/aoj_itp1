package main

import "fmt"

//lint:ignore U1000 //

// /*
func main() {
	sixa()
}

//*/

func sixa() {
	var n, tmp int
	fmt.Scanf("%d", &n)
	a := []int{}
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &tmp)
		a = append(a, tmp)
	}
	for i := n - 1; i >= 0; i-- {
		if i == 0 {
			fmt.Printf("%d", a[i])
		} else {
			fmt.Printf("%d ", a[i])
		}
	}
	fmt.Println()
}
