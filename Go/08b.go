package main

import "fmt"

//lint:ignore U1000 //

/*
func main() {
	eightb()
}

//*/

func eightb() {
	var sum int
	var tmp string
	for i := 0; i < 10001; i++ {
		fmt.Scanf("%s", &tmp)
		if tmp == "0" {
			break
		}
		sum = 0
		for j := 0; j < len(tmp); j++ {
			sum += int(tmp[j]) - 48
			//ASCII "1" is 49
		}
		fmt.Println(sum)
	}
}
