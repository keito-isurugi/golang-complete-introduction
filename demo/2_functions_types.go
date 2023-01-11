package demo

import "fmt"

func FunctionsAndTypes() {
	fmt.Println("----2. 変数と型----")

	// var sum float64
	// sum = 5 + 6 + 3
	// avg := sum / 3
	// if avg > 4.5 {
	// 	println("good")
	// }

	// ===== 構造体 =====
	// p := struct {
	// 	name string
	// 	age int
	// }{
	// 	name: "keito",
	// 	age: 30,
	// }
	// p.age ++
	// fmt.Println(p.name, p.age)
	
	// // ===== 配列 =====
	// // ゼロ値で初期化
	// var ns1 [20]int
	// // 配列リテラルで初期化
	// var ns2 = [5]int{10, 20, 30, 40, 50}
	// // 要素数を値から推論
	// ns3 := [...]int{10, 20, 30, 40, 50}
	// // 5番目が50、10番目が100で他が0の要素数11の配列
	// ns4 := [...]int{5: 50, 10: 100}

	// fmt.Println(ns1)
	// fmt.Println(ns2)
	// fmt.Println(ns3[3])
	// fmt.Println(ns4)
	
	// ns := [...]int{10, 20, 30, 40, 50, 10, 20, 30, 40, 50}
	// // 要素にアクセス
	// println(ns[3]) // 添字は変数でもよい
	// // 長さ
	// println(len(ns))
	// // スライス演算
	// fmt.Println(ns[1:3])



	// ===== スライス =====
	// // ゼロ値はnil
	// var ns5 []int
	// // 長さと容量を指定して初期化
	// // 各要素はゼロ値で初期化される
	// ns5 = make([]int, 3, 10)
	// // スライスリテラルで初期化
	// // 要素数は指定しなくてよい
	// // 自動で配列は作られる
	// var ns6 = []int{10, 20, 30, 40, 50}
	// // 5番目が50、10番目が100で他が0の要素数11のスライス
	// ns7 := []int{5: 50, 10: 100}
	// fmt.Println(ns5)
	// fmt.Println(ns6)
	// fmt.Println(ns7)

	// ns := []int{10, 20, 30, 40, 50}
	// // 要素にアクセス
	// println(ns[3])
	// // 長さ
	// println(len(ns))
	// // 容量
	// println(cap(ns))
	// // 要素の追加
	// // 容量が足りない場合は背後の配列が再確保される
	// ns = append(ns, 60, 70, 80, 90 ,100, 110, 120)
	// println(len(ns), cap(ns)) // 長さと容量

	// n, m := 2, 4
	// // n番目以降のスライスを取得する
	// fmt.Println(ns[n:]) // [30 40 50]

	// // 先頭からm-1番目までのスライスを取得する
	// fmt.Println(ns[:m]) // [10 20 30 40]

	// // capを指定する
	// ms := ns[:m:m]
	// fmt.Println(cap(ms)) // 4
	
	// for i, v := range ns {
	// 	fmt.Println(i) 
	// 	fmt.Println(v) 
	// }

	// // n1 := 111
	// // n2 := 222
	// // n3 := 333
	// // n4 := 444

	// ns1110 := [...]int{111, 222, 333, 444}
	// var sum int

	// for _, v := range ns1110 {
	// 	sum += v
	// }

	// println(sum)



	// ===== マップ =====
	// ゼロ値はnil
	// var m map[string]int
	// // makeで初期化
	// m = make(map[string]int)
	// // 容量を指定できる
	// m = make(map[string]int, 10)
	// リテラルで初期化
	// m := map[string]int{"x": 10, "y": 20}
	// 空の場合
	// m := map[string]int{}
	// m := map[string]int{"x": 10, "y": 20}
	// // キーを指定してアクセス
	// println(m["x"])
	// // キーを指定して入力
	// m["z"] = 30
	// // 存在を確認する
	// n, ok := m["z"]
	// println(n, ok)
	// // キーを指定して削除する
	// delete(m, "z")
	// // 削除されていることを確認
	// n, ok = m["z"] // ゼロ値とfalseを返す
	// println(n, ok)
	// fmt.Printf("m[%%v] -> %v\n", m)
	
	// for i, v := range m {
	// 	println(i,"->",v)
	// }

	// ===== 関数 =====
	println("- 関数")
	println(add(1, 2))
	x, y := swap(10, 20)
	// x, _ := swap(10, 20)
	println(x, y)
	
	func() {
		println("これは無名関数です。")
	}()

	// 関数型
	// fs := make([]func() string, 2)
	// fs[0] = func() string { return "hoge" }
	// fs[1] = func() string { return "fuga" }
	// // fs = [hogeを返す関数, hugaを返す関数]
	// for _, f := range fs {
	// 	fmt.Println(f())
	// }

	fs := make([]func(), 3)
	for i := range fs {
		fs[i] = func() { fmt.Println(i) }
	}
	for _, f := range fs { f() }

	// 値のコピー
	p := struct {
		age int
		name string
	}{
		age: 10,
		name: "keito",
	}
	p2 := p
	p2.age = 20
	println(p.age, p.name)
	println(p2.age, p2.name)

	// ポインタ
	// var z int
	// f(&z)
	// println(z)

	// for i := 1; i <= 100; i++ {
	// 	print(i)
	// 	println(odd_or_even(i))
	// }

	// println(int_swap(10, 20))

	// a, b := 10, 20
	// int_swap2(&a, &b)
	// println(a, b)


	// メソッド
	var hex Hex = 100
	fmt.Println(hex.String())

	// レシーバ
	var v T
	(&v).f()
	v.f()

	var n MyInt
	println(n)
	n.Inc()
	println(n)

	// メソッド値
	var hex2 Hex2 = 200
	f2 := Hex2.String2
	fmt.Println(f2(hex2))

}


// 関数
func add(x int, y int) int {
	return x + y
}
func swap (x, y int) (x2 int, y2 int) {
	y2, x2 = x, y
	return
}
func odd_or_even(x int) string {
		if x%2 == 0 {
			return "-偶数"
		} else {
			return "-奇数"
		}
}
func int_swap(x, y int) (x2 int, y2 int) {
	y2, x2 = x, y
	return
}
func int_swap2(x, y *int) (x2 int, y2 int) {
	y2, x2 = *x, *y
	return
}

// ポインタ
func f(xp *int) {
	*xp = 100
}

// メソッド
type Hex int
func (h Hex) String() string {
	return fmt.Sprint(int(h))
}

// レシーバ
type T int
func (t *T) f() { println("hi!") }

type MyInt int
func (n *MyInt) Inc() { *n++ }

// メソッド値
type Hex2 int
func (h Hex2) String2() string {
	return fmt.Sprint(int(h))
}