package main

import "fmt"

//lint:ignore U1000 //

/*
func main() {
	threed()
}

//*/

func threed() {
	var a, b, c int
	flag := [10001]int{}
	//Goは初期化時に0が代入されている
	fmt.Scanf("%d %d %d", &a, &b, &c)
	for i := 1; i <= c; i++ {
		if c%i == 0 {
			flag[i] = 1
		}
	}
	count := 0
	for i := a; i <= b; i++ {
		if flag[i] == 1 {
			count++
		}
	}
	fmt.Println(count)
}
