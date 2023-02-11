package main

import "fmt"

//lint:ignore U1000 //

/*
func main() {
	ninec()
}

//*/

func ninec() {
	var n, pa, pb, l int
	var a, b string
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s %s", &a, &b)
		if len(a) < len(b) {
			l = len(a)
		} else {
			l = len(b)
		}
		for j := 0; j < l; j++ {
			if string(a[j]) < string(b[j]) {
				pb += 3
				break
			} else if string(a[j]) > string(b[j]) {
				pa += 3
				break
			} else if j == l-1 {
				if len(a) == len(b) {
					pa++
					pb++
					break
				} else if len(a) > len(b) {
					pa += 3
					break
				} else {
					pb += 3
					break
				}
			}
		}
	}
	fmt.Println(pa, pb)
}
