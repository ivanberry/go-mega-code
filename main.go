package main

// deps
import (
	"fmt"
	"html/template"
	"net/http"
)

// User
type User struct {
	Username string
}

type IndexViewModel struct {
	Title string
	User  User
}

// main
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := User{Username: "谢尔婷"}
		v := IndexViewModel{Title: "主页", User: user}
		fmt.Printf("v's addrss is %p\n", &v)
		tpl, _ := template.ParseFiles("templates/index.html")
		tpl.Execute(w, &v)
	})
	http.ListenAndServe(":8888", nil)
}
