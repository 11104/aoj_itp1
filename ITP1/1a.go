package main // mainパッケージであることを宣言

import "fmt" // fmtモジュールをインポート

//lint:ignore U1000 //
/*
func main() {
	onea()
}
*/
func onea() { // 最初に実行されるmain()関数を定義
	fmt.Println("Hello World")
}
