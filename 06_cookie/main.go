package main

import "net/http"

func SetCookie(w http.ResponseWriter, r *http.Request) {
	//创建cookie
	cookie1 := http.Cookie{
		Name: "admin",
		Value: "123456",
		HttpOnly: true,
		MaxAge: 5,
	}

	cookie2 := http.Cookie{
		Name: "aaa",
		Value: "888888",
		HttpOnly: true,
	}
	//将cookie发送给浏览器
	//方式1：
	//w.Header().Set("Set-Cookie",cookie1.String())
	//添加第二个cookie
	//w.Header().Add("Set-Cookie",cookie2.String())

	//方式2：
	http.SetCookie(w,&cookie1)
	http.SetCookie(w,&cookie2)
}

func main() {
	http.HandleFunc("/setcookie",SetCookie)
	http.ListenAndServe(":8080",nil)
}
