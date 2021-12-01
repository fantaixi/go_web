package main

import (
	"html/template"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request)  {
	//解析模板文件
	//t, _ := template.ParseFiles("E:\\go_workspace\\src\\go_web\\05_template\\index.html")
	//通过Must函数让go自己处理异常
	t := template.Must(template.ParseFiles("E:\\go_workspace\\src\\go_web\\05_template\\index.html","E:\\go_workspace\\src\\go_web\\05_template\\index2.html"))
	//执行模板
	//t.Execute(w, "Hello World!")
	//将响应数据在index2.html中显示
	t.ExecuteTemplate(w,"E:\\go_workspace\\src\\go_web\\05_template\\index2.html","index2.....")
}
func main() {
	http.HandleFunc("/red",handle)
	http.ListenAndServe(":8080",nil)
}
