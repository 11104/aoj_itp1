package main

import "fmt"

//lint:ignore U1000 //

/*
func main() {
	fourc()
}

//*/

func fourc() {
	var a, b int
	var op string
	for i := 0; i < 10001; i++ {
		fmt.Scanf("%d %s %d", &a, &op, &b)
		if op == "?" {
			break
		}
		switch op {
		case "+":
			fmt.Println(a + b)
		case "-":
			fmt.Println(a - b)
		case "*":
			fmt.Println(a * b)
		case "/":
			fmt.Println(a / b)
		}
	}
}
