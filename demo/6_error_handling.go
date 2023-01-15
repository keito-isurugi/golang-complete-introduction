package demo

import (
	// "bufio"
	"fmt"
	"os"
)

type Stringer interface {
	String() string
}

func ToStringer(v interface{}) (Stringer, error) {
	if s, ok := v.(Stringer); ok {
		return s, nil
	}
	return nil, MyError("CastError")
}

type MyError string

func (e MyError) Error() string {
	return string(e)
}

type S string
func (s S) String() string {
	return string(s)
}



func ErrorHandling() {
	fmt.Println("----6. エラー処理----")

	type error interface {
		Error() string
	}


	f, err := os.Open("file.txt")
	if err != nil {
		// エラー処理
		println("エラーです")
	} else {
		println(f)
	}


	type PathError struct {
		Op   string
		Path string
		Err  error
	}

	v := S("hoge")
	if s, err := ToStringer(v); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
	} else {
		fmt.Println(s.String())
	}

	// s := bufio.NewScanner(r)
	// for s.Scan() {
	// 	fmt.Println(s.Text())
	// }
	// if err := s.Err(); err != nil {

	// }

}