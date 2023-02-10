package main

import (
	"bufio"
	"fmt"
	"os"
	"strings" //To use "ToUpper" and "Tolower"
)

//lint:ignore U1000 //

/*
func main() {
	eighta()
}

//*/

func eighta() {
	//fmt.Scan be cut by ","
	var s string
	scanner := bufio.NewScanner(os.Stdin) // Allow to input scan 標準入力を受け付けるスキャナ
	scanner.Scan()                        //scan 1 liner １行分の入力を取得する
	s = scanner.Text()                    //input scanner text to s 読み取ったテキストをstringのsへ代入
	for i := 0; i < len(s); i++ {
		if "a" <= string(s[i]) && string(s[i]) <= "z" {
			fmt.Print(strings.ToUpper(string(s[i])))
		} else {
			fmt.Print(strings.ToLower(string(s[i])))
		}
	}
	fmt.Printf("\n")
}
