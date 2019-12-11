package main

import "fmt"

func main() {
	var n int = 100
	fmt.Println(n)  //100
	fmt.Println(&n) //0xc000018058

	var p *int = &n
	fmt.Println(p)  //0xc000018058
	fmt.Println(*p) //100

}
