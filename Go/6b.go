package main

import "fmt"

//lint:ignore U1000 //

// /*
func main() {
	sixb()
}

//*/

func sixb() {
	var n, tmp int
	var t string
	d := [4][13]int{}
	/* Don't need to initialize when using Golang
	for i := 0; i < 4; i++ {
		for j = 0; j < 13; j++ {
			d[i][j] = 0
		}
	}
	*/
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s %d", &t, &tmp)
		if t == "S" {
			d[0][tmp-1] = 1
		} else if t == "H" {
			d[1][tmp-1] = 1
		} else if t == "C" {
			d[2][tmp-1] = 1
		} else {
			d[3][tmp-1] = 1
		}

	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			if d[i][j] == 0 {
				if i == 0 {
					fmt.Printf("S %d\n", j+1)
				} else if i == 1 {
					fmt.Printf("H %d\n", j+1)
				} else if i == 2 {
					fmt.Printf("C %d\n", j+1)
				} else {
					fmt.Printf("D %d\n", j+1)
				}
			}
		}
	}
	fmt.Println()
}
