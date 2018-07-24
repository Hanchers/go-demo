//一个简单的web服务
package main

import (
	"net/http"
)
//请求相应
func hello66(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Content-Type", "text/plain")
    res.Write([]byte("Hello world\n"))
}

func main() {
	//将hello66负责相应/路径的请求
    http.HandleFunc("/", hello66)
    //绑定端口5000
    http.ListenAndServe(":5000", nil)
}
