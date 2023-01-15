package demo

import (
	"fmt"
	// "go/types"

	// "golang.org/x/text/cases"
)

type Func func() string
func (f Func) String() string { return f() }
var _ fmt.Stringer = Func(nil)

// type Stringer interface {
// 	String() string
// }

// type I int 
// func(n I) String() string {
// 	return "I"
// }

// type B bool
// func(b B) String() string {
// 	return "B"
// }

// type S string
// func(b S) String() string {
// 	return "S"
// }

// func F(s Stringer) {
// 	switch v := s.(type) {
// 	case I:
// 		fmt.Println(int(v), "I")
// 	case B:
// 		fmt.Println(bool(v), "B")
// 	case S:
// 		fmt.Println(string(v), "S")
// 	}
// }

// type Hex int 
// func (h Hex) String() string {
// 	return fmt.Sprintf("%x", int(h))
// }
// type Hex2 struct {Hex}

func Abstraction() {
	fmt.Println("----5. 抽象化----")

	// var n I = I(100)
	// F(I(100))
	// F(B(true))
	// F(S("hoge"))

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

	// var v interface{}
	// n, ok := v.(int)
	// fmt.Println(n, ok)
	// ss,ok := v.(string)
	// fmt.Println(ss, ok)

	// 型スイッチ
	var i interface{}
	i = true
	switch v := i.(type) {
	case int:
		fmt.Println(v*2)
	case string:
		fmt.Println(v+"hoge")
	default:
		fmt.Println("default")
	}
	
	type Reade interface {
		Reade(p []byte) (n int, err error)
	}
	type Writer interface {
		Write(p []byte) (n int, err error)
	}

	// 構造体の埋め込み
	type Hoge struct {
		N int
	}
	type Fuga struct {
		Hoge
	}
	f := Fuga{Hoge{N:100}}
	fmt.Println(f.N)
	fmt.Println(f.Hoge.N)

	

}