package demo

import (
	// "fmt"
	"testing"
	// "mypkg"
)

// type Hex1 int
// func (h Hex) String1() string {
// 	return fmt.Sprintf("%x", int(h))
// }

// func TestHex_String(t *testing.T) {
// 	want := "a"
// 	got := mypkg.Hex(10).String1()
// 	if got != want {
// 		t.Errorf("want %q, got %q", want, got)
// 	}
// }

func add1(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal",
			args: args{a: 1, b: 2},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add1(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestAndTestabiriti() {
// 	fmt.Println("----7. テストとテスタビリティ----")
// }