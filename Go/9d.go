package main

import (
	"fmt"
)

//lint:ignore U1000 //

// /*
func main() {
	nined()
}

// */

func nined() {
	var a, b, q int
	var order, str string
	fmt.Scanf("%s", &str)
	fmt.Scanf("%d", &q)
	for i := 0; i < q; i++ {
		fmt.Scan(&order, &a, &b)
		var stmp string
		if order == "print" {
			fmt.Println(str[a : b+1])
		} else if order == "replace" {
			fmt.Scanf("%s", &stmp)
			str = str[:a] + stmp + str[b+1:]
			/*
				var sub string
				fmt.Scanf("%s", &sub)
				str = strings.Replace(str, str[a:b+1], sub, 1) //int is number of replacements
				//fmt.Println(str)//debug
			*/
		} else {
			stmp = ""
			for j := b; j >= a; j-- {
				stmp += string(str[j])
			}
			str = str[:a] + stmp + str[b+1:]
			/*
				var tmp []byte
				for j := a; j <= b; j++ {
					tmp = append(tmp, str[j])
				}
				for j := a; j <= b; j++ {
					str[j] = tmp[len(tmp)-(j-a)-1]
				}
			*/
			//fmt.Println(str)//debug
		}
	}
	//fmt.Println(str)
}
