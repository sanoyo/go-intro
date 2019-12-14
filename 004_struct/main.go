package main

import "fmt"

type Vertex struct {
	X int
	Y int
	S string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	//v.X = 1000
	(*v).X = 1000
}

func main() {

	// この場合だと、値は書き変わらない
	v := Vertex{1, 2, "test"}
	changeVertex(v)
	fmt.Println(v)

	// この場合だと値を書き換えることができる
	v2 := &Vertex{1, 2, "test"}
	changeVertex2(v2)
	fmt.Println(v2)

	/*
		v := Vertex{X: 1, Y: 2}
		fmt.Println(v)        //{1 2}
		fmt.Println(v.X, v.Y) //1 2

		v2 := Vertex{X: 1}
		// 定義したいない場合は、0が入る
		fmt.Println(v2)

		v3 := Vertex{1, 2, "test"}
		fmt.Println(v3)

		// v4とv5は同じ意味
		v4 := Vertex{}
		fmt.Println(v4) //{0 0 }

		var v5 Vertex
		// 同じstructになる
		// nilにはならない
		fmt.Println(v5)

		// v6とv7は同じ意味
		v6 := new(Vertex)
		fmt.Println(v6) //&{0 0 }

		v7 := &Vertex{}
		fmt.Println(v7) //&{0 0 }
	*/
}
