package main

import "fmt"

//lint:ignore U1000 //

/*
func main() {
	nineb()
}

//*/

func nineb() {
	var s string
	var m, h int
	for {
		fmt.Scan(&s)
		if s == "-" {
			break
		}
		fmt.Scan(&m)
		for i := 0; i < m; i++ {
			fmt.Scanf("%d", &h)
			s = s[h:] + s[:h]
		}
		fmt.Println(s)
	}
}
