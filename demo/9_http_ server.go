package demo

import (
	// "bytes"
	// "encoding/json"
	"encoding/json"
	"fmt"
	// "text/template"
	// "log"
	"net/http"
)




func handler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// fmt.Fprintln(w, "Hello, GoでHTTPサーバ")
	// fmt.Fprintln(w, "リクエストパラメータから取得", r.FormValue("msg"))
	// contentType := r.Header.Get("Content-Type")
	// fmt.Fprintln(w, "リクエストパラメータから取得", contentType)

	type Person struct {
		Name string `json: "name"`
		Age int `json: "age"`
	}

	// HTTPリクエストを送る
	resp, err := http.Get("https://www.google.com/")
	if err != nil {fmt.Println("エラーです")}
	defer resp.Body.Close()
	var p Person
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&p); err != nil {
		fmt.Println("エラーです")
	}
	fmt.Println(p)

	// ミドルウェア
	// type MiddleWare interface {
	// 	ServeNext(h http.Handler) http.Handler
	// }
	// type MiddleWareFunc func(h http.Handler) http.Handler
	// func (f MiddleWareFunc) ServeNext(h http.Handler) http.Handler {
	// 	return f(h)
	// }
	// func With(h http.Handler, ms ...MiddleWare) http.Handler {
	// 	for _, m := range ms { h = m.ServeNext(h) }
	// 	return h
	// }	

	// リダイレクト
	// http.Redirect(w, r, "/hoge", http.StatusFound)

	//  テンプレートエンジン
	// var tmpl = template.Must(template.New("sign").Parse("<html><body>{{.}}</body></html>"))
	// var result string
	// if r.FormValue("p") == "hoge"{
	// 	result = "大吉です"
	// 	} else {
	// 	result = "凶です"
	// }
	// tmpl.Execute(w, result)

	// jsonエンコード
	// v := struct {
	// 	Msg string `json :"msg"`
	// }{
	// 	Msg: "hello",
	// }
	// if err := json.NewEncoder(w).Encode(v); err != nil {
	// 	log.Println("Error:", err)
	// }


	// type Person struct {
	// 	Name string `json: "name"`
	// 	Age int `json: "age"`
	// }
	// p := &Person{Name: "tenntenn", Age: 32}

	// var buf bytes.Buffer
	// enc := json.NewEncoder(&buf)
	// if err := enc.Encode(p); err != nil {log.Fatal(err)}
	// fmt.Println(buf.String())
	// fmt.Fprint(w, buf.String())
}

func HttpServer() {
	fmt.Println("----9. HTTPサーバー----")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}