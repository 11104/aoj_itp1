package main

import "fmt"

//lint:ignore U1000 //

/*
func main() {
	threeb()
}

//*/

func threeb() {
	a := []int{}
	//スライス
	count := 0
	var tmp int
	for i := 0; i < 10000; i++ {
		//Golangにwhileは無い
		fmt.Scanf("%d", &tmp)
		if tmp == 0 {
			break
		}
		count++
		a = append(a, tmp)
	}
	for i := 0; i < count; i++ {
		fmt.Printf("Case %d: %d\n", i+1, a[i])
	}

}
