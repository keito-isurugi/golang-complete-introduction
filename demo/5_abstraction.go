package demo

import (
	"fmt"
)

type Func func() string
func (f Func) String() string { return f() }
var _ fmt.Stringer = Func(nil)

func Abstraction() {
	fmt.Println("----5. 抽象化----")

	type Stringer interface {
		String() string
	}

	var s Stringer = Hex(100)
	fmt.Println(s.String())
	
	// var v interface{}
	// v = 111
	// v = "hoge"
	// fmt.Println(v)

	// var ss fmt.Stringer = Func(func() string{
	// 	return "hi"
	// })
	// fmt.Println(ss)

	var v interface{}
	n, ok := v.(int)
	fmt.Println(n, ok)
	ss,ok := v.(string)
	fmt.Println(ss, ok)

	var i interface{}
	i = 100
	switch v := i.(type) {
	case int:
		fmt.Println(v*2)
	case string:
		fmt.Println(v+"hoge")
	default:
		fmt.Println("default")
	}
}