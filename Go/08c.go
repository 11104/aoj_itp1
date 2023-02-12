package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//lint:ignore U1000 //

/*
func main() {
	eightc()
}

//*/

func eightc() {
	var s string
	t := [26]int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s = scanner.Text()
		/*//For Debug
		if s == "EnD" {
			break
		}
		*/
		s = strings.ToLower(s)
		for i := 0; i < len(s); i++ {
			if "a" <= string(s[i]) && string(s[i]) <= "z" {
				t[s[i]-'a']++
			}
		}
	}
	for i := 0; i < 26; i++ {
		fmt.Printf("%c : %d\n", i+'a', t[i])
	}
}
