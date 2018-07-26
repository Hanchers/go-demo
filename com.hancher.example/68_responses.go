package main

import "net/http"
import "fmt"
//测试请求返回值
func hello68(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Content-Type", "text/plain")
    if req.URL.Path == "/protected" {
        res.WriteHeader(401)
        fmt.Fprintln(res, "Not allowed")
    } else {
        fmt.Fprintln(res, "Hello")
    }
}

func main() {
    http.HandleFunc("/", hello68)
    http.ListenAndServe(":5000", nil)
}
