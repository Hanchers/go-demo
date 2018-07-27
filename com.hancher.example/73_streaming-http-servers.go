package main

import (
    "fmt"
    "net/http"
    "time"
	"strconv"
)

func stream(resp http.ResponseWriter, req *http.Request) {
    respf, ok := resp.(http.Flusher)
    if !ok {
        panic("not flushable")
    }
    fmt.Println("stream begin")
    resp.WriteHeader(200)
    // 循环向前端发送消息
    for i := 0; i < 10; i++ {
    	str := "tick"+strconv.Itoa(i)
        n, err := resp.Write([]byte(str+"\n"))
        respf.Flush()
        fmt.Println(str, n, err)
        time.Sleep(time.Second * 1)
    }
}

// http 流处理
func main() {
    http.HandleFunc("/", stream)
    fmt.Println("serve begin")
    http.ListenAndServe(":5000", nil)
}
