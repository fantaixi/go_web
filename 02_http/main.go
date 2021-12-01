package main

import (
	"fmt"
	"net/http"
)

//创建处理器函数
func handle(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"测试。。。。11111")
}
func main() {
	//调用处理器处理请求
	http.HandleFunc("/http",handle)
	//路由
	http.ListenAndServe(":8080",nil)
}
