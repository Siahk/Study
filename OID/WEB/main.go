package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Print("ready")
	http.HandleFunc("/login", loginTask)

	err := http.ListenAndServe("localhost:8081", nil)

	if err != nil {
		fmt.Println("ListanAndServer error: ", err.Error())
	}

}

func loginTask(w http.ResponseWriter, req *http.Request) {
	fmt.Println("loginTash is running...")

	time.Sleep(time.Second * 2)

	req.ParseForm()
	param_userName, found1 := req.Form["username"]
	param_password, found2 := req.Form["password"]

	if !(found1 && found2) {
		fmt.Fprint(w, "请勿非法访问")
		return
	}
	result := newBaseJsonBean()
	userName := param_userName[0]
	password := param_password[0]

	s := "userName:" + userName + ", password:" + password
	fmt.Println(s)

	if userName == "zhangsan" && password == "123456" {
		result.Code = 100
		result.Message = "登录成功!"
		result.Date = struct {
		}{}
	} else {
		result.Code = 101
		result.Message = "密码错误"
		result.Date = struct {
		}{}
	}

	bytes, _ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))
}

type BaseJsonBean struct {
	Code    int         `json:"code"`
	Date    interface{} `json:"data"`
	Message string      `json:"message"`
}

func newBaseJsonBean() *BaseJsonBean {
	return &BaseJsonBean{}
}
