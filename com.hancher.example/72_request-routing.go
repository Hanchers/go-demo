package main

import "net/http"
import "fmt"
//获得请求的查询参数
func hello72(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL)
	fmt.Fprintln(w, "hello "+req.URL.Query().Get("name"))
}

//http://localhost:5000/hello?name=hancher
func main() {
	http.HandleFunc("/hello", hello72)
	http.ListenAndServe(":5000", nil)
}
