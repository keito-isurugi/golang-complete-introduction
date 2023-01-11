package demo

import (
	"bufio"
	"flag"
	"fmt"
	// "text/scanner"
	"os"
	// "strings"
)

var msg = flag.String("msg", "デフォルト値", "説明")
var n int
func init() {
	flag.IntVar(&n, "n", 4, "回数")
}

func CliTool() {
	fmt.Println("----4. コマンドラインツール----")
	// fmt.Println(os.Args)
	
	flag.Parse()
	// fmt.Println(strings.Repeat(*msg, n)) // go run main.go -msg=こんにちは -n=2
	// fmt.Println(flag.Args()) // go run main.go -msg=こんにちは -n=2 やぁ
	// fmt.Fprintln(os.Stderr, "エラー")  // 標準エラー出力に出力
	// os.Exit(1) // プログラムの終了
	// fmt.Fprintln(os.Stderr, "Hello")  // 標準出力に出力

	// 関数の遅延実行 defer
	// msg := "!!!"
	// defer fmt.Println(msg)
	// msg = "world"
	// defer fmt.Println(msg)
	// fmt.Println("hello")
	fmt.Print("continue? (Y/n) >")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		in := scanner.Text()
		fmt.Println("in: ", in)
		switch in {
		case "", "Y":
			fmt.Println("デフォルトの処理をします")
			goto L
		case "n":
			fmt.Println("何もしない")
			goto L
		default:
			fmt.Println("コマンドが不正なのでもう一度入力を促す")
			continue
		}
	}
	L:
	


}