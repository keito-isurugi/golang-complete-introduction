package demo

import (
	"fmt"
	"time"

)

func ParallelProcessing() {
	fmt.Println("----8.  並行処理----")

	// 無名関数とゴールーチン
	// go func() {
	// 	fmt.Println("別のゴールーチン")
	// }()
	// fmt.Println("mainのゴールーチン")
	// time.Sleep(50 * time.Millisecond)
		
	// defer fmt.Println("main done")
	// go func() {
	// 	defer fmt.Println("gorotine1 done")
	// 	time.Sleep(1 * time.Millisecond)
	// }()
	// go func() {
	// defer fmt.Println("gorotine2 done")
	// time.Sleep(2 * time.Millisecond)
	// }()
	// time.Sleep(10 * time.Millisecond)

	// 共通の変数を使う
	// var v int
	// go func() { 
	// 	time.Sleep(1 * time.Second)
	// 	v = 100
	// }()
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println(v)
	// }()
	// time.Sleep(3 * time.Second)
	
	// n := 1 
	// go func() {
	// 	for i := 2; i <= 5; i++ {
	// 		fmt.Println(n, "*", i)
	// 		n *= i
	// 		time.Sleep(100)
	// 	}
	// }()
	// time.Sleep(1 * time.Second)

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(n, "+", i)
	// 	n += 1
	// 	time.Sleep(100)
	// }

	//  チャネル
	// ch := make(chan int)
	// hoge := 1000
	// go func() {
	// 	ch <- 100
	// 	hoge = 2000
	// }()
	// go func() {
	// 	v := <- ch
	// 	fmt.Println(hoge)
	// 	fmt.Println(v)
	// }()
	// time.Sleep(1 * time.Second)

	ch1 := make(chan int)
	ch2 := make(chan string)
	go func() {
		ch1 <- 100
	}()
	go func() {
		ch2 <- "hi"
	}()

	select {
	case v1 := <-ch1:
		fmt.Println(v1)
	case v2 := <-ch2:
		fmt.Println(v2)
	}
	time.Sleep(1 * time.Second)

}