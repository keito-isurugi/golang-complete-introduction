package demo

import "fmt"

func BasicSyntax() {
	fmt.Println("----1. 基本構文----")
	const  (
		x = 100
		y = 200
		w
		z
	)

	// n := x + y
	// m := n + 100

	msg := "Hello World!"

	if y == 100 {
		msg = "100です"
	} else {
			msg = "100じゃないです"
	}

	var msg2 string
	switch x {
	case 100, 300:
		msg2 = "x is 100 or 300"
	default:
		msg2 = "defalut" 
	}

	// for i := 0; i <= 3; i++ {
	// 	fmt.Println(i)
	// }

	ary := []int{100, 200, 300}
	for i, v := range ary {
		if v == 200 {break}
		fmt.Println(i)
		fmt.Println(v)
	}

	// fmt.Println(x, y, w, z)
	// fmt.Println(m)
	fmt.Println(msg)
	fmt.Println(msg2)
}