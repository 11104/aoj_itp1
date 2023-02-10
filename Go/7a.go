package main

import "fmt"

//lint:ignore U1000 //

/*
func main() {
	sevena()
}

//*/

func sevena() {
	var m, f, r int
	for i := 0; i < 51; i++ {
		fmt.Scanf("%d %d %d", &m, &f, &r)
		if m == -1 && f == -1 && r == -1 {
			break
		} else if m == -1 || f == -1 {
			fmt.Println("F")
		} else if 80 <= m+f {
			fmt.Println("A")
		} else if 65 <= m+f {
			fmt.Println("B")
		} else if 50 <= m+f || 50 <= r {
			fmt.Println("C")
		} else if 30 <= m+f {
			fmt.Println("D")
		} else {
			fmt.Println("F")
		}
	}
}
