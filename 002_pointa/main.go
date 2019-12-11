package main

import "fmt"

func one(x int) {
	x = 1
}

func main() {
	var n int = 100
	one(n)
	// ポインタに入ってる値
	fmt.Println(n)
	// ポインタ
	fmt.Println(&n)
	// ポインタに入ってる値
	fmt.Println(*&n)
	// ポインタ
	fmt.Println(&*&n)

	// そのまま出力
	// fmt.Println(n)

	// fmt.Println(&n)

	// // intのポインタ型を宣言
	// var p *int = &n
	// fmt.Println(p)
	// fmt.Println(*p)
}
