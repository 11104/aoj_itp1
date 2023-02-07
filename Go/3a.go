package main

import "fmt"

//lint:ignore U1000 //

/*
func main() {
	threea()
}

//*/

func threea() {
	for i := 0; i < 1000; i++ {
		//初期値を設定するときはvar,intを使わず:=で宣言できる
		fmt.Println("Hello World")
	}
}
