package main

import (
	"encoding/json"
	"fmt"
	"go_web/04_req/model"
	"net/http"
)

//创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"你发送的请求地址是：",r.URL.Path)
	fmt.Fprintln(w,"你发送的请求地址后的查询字符串是：",r.URL.RawQuery)
	fmt.Fprintln(w,"请求头中的所有信息有：",r.Header)
	fmt.Fprintln(w,"请求头中的Connection信息有：",r.Header["Connection"])
	fmt.Fprintln(w,"请求头中的Connection的属性值是：",r.Header.Get("Connection"))

	//获取请求体内容的长度
	//len := r.ContentLength
	//创建byte切片
	//body := make([]byte,len)
	//将请求体中的内容读到body中，body只能获取一次
	//r.Body.Read(body)
	//在前端显示
	//fmt.Fprintln(w,"请求体中的内容有：",string(body))

	//获取请求参数
	//方法1：
	//解析表单，在调用r.Form之前必须执行该操作
	//r.ParseForm()
	//获取请求参数
	//如果form表单的action属性的url地址中也有与表单参数相同的请求参数
	//那么参数值都可以得到，并且form表单中的参数值在url的参数值的前面
	//fmt.Fprintln(w,"请求参数有：",r.Form)

	//方法2：
	//直接调用FormValue和PostFormValue
	fmt.Fprintln(w,"URL中的user请求参数有：",r.FormValue("user"))
	fmt.Fprintln(w,"form表单中的username请求参数有：",r.PostFormValue("username"))
}

//响应
func testJsonRes(w http.ResponseWriter, r *http.Request) {
	//设置响应内容的类型
	w.Header().Set("Content-Type","application/json")
	//创建user
	user := model.User{
		ID: 1,
		Username: "admin",
		Password: "123456",
		Email: "aasdasad",
	}
	//将user转成json格式
	json,_ := json.Marshal(user)
	//将json格式的数据响应给客户端
	w.Write(json)
}

//测试重定向
func testRedirect(w http.ResponseWriter, r *http.Request) {
	//设置响应头中的Location属性
	w.Header().Set("Location","https://www.google.com")
	//设置响应状态码
	w.WriteHeader(302)
}
func main() {
	http.HandleFunc("/hello",handler)
	http.HandleFunc("/json",testJsonRes)
	http.HandleFunc("/red",testRedirect)

	http.ListenAndServe(":8080",nil)
}
