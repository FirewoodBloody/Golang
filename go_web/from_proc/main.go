package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" { //判断请求的方法
		t, err := template.ParseFiles("../login.html")
		if err != nil {
			fmt.Fprintf(w, "load login.html failed,err:%s\n", err)
			return
		}
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		name := r.FormValue("username")
		pass := r.FormValue("password")
		if name == "admin" && pass == "admin" {
			fmt.Fprintf(w, "hello world")
		} else {
			fmt.Fprintf(w, "Error in username or password you entered!")
		}
	}
}

func main() {
	http.HandleFunc("/login", login)         //定义用户访问服务的页面
	err := http.ListenAndServe(":9000", nil) //监听服务端口，并启动服务
	if err != nil {
		fmt.Printf("listen server failed,err:%s\n", err)
		return
	}
}
