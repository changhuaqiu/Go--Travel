package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func tmpl(w http.ResponseWriter, r *http.Request) {
	tl, err := template.ParseFiles("test.html")
	if err != nil {
		panic(err)
	}
	tl.Execute(w, "hello world")
}
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello World")
}
func main() {

	//server := http.Server{
	//	Addr: "121.4.127.242:8080",
	//}
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8000", nil)

}
