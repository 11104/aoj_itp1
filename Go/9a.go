package main

import (
	"fmt"
	"strings"
)

//lint:ignore U1000 //

/*
func main() {
	ninea()
}

//*/

func ninea() {
	var w, s string
	count := 0
	fmt.Scan(&w)
	for {
		fmt.Scan(&s) //NO Scanf
		if s == "END_OF_TEXT" {
			break
		}
		//strings.ToLower(s)==w OK
		if strings.EqualFold(s, w) {
			count++
			//fmt.Println(count)//debug
		}
	}
	fmt.Println(count)
}
