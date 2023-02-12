package main

import (
	"fmt"
	"os"
)

//lint:ignore U1000 //

/*
func main() {
	eightd()
}

//*/

func eightd() {
	var s, p string
	fmt.Scanf("%s", &s)
	fmt.Scanf("%s", &p)
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if s[(i+j)%len(s)] != p[j] {
				break
			} else if j == len(p)-1 {
				fmt.Println("Yes")
				os.Exit(0)
			}
		}
	}
	fmt.Println("No")
}
