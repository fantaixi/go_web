package main

import (
	"fmt"
	"net/http"
	"time"
)

//使用默认的处理器
/*
//创建处理器函数
func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"hello.....",r.URL.Path)
}
func main() {
	http.HandleFunc("/",handler)

	//创建路由
	http.ListenAndServe(":8080",nil)
}
*/

//创建自己的处理器
type MyHandler struct {

}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"通过自己创建的处理器处理请求.......")
}

func main() {
	handler := MyHandler{}

	//http.Handle("/my",&handler)
	//http.ListenAndServe(":8082",nil)

	//创建Serve结构,并详细配置里面的字段
	server := http.Server{
		Addr: ":8083",
		Handler: &handler, 
		ReadTimeout: 2*time.Second,
	}
	server.ListenAndServe()
}