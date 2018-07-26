// 记录请求日志
package main

import "net/http"
import "time"
import "fmt"

// 打印日志
func runLogging(logs chan string) {
    for log := range logs {
        fmt.Println(log)
    }
}

// 包装日志记录方法
func wrapLogging(f http.HandlerFunc) http.HandlerFunc {
	// 打开日志通道
    logs := make(chan string, 10000)
    // 开启协程记录日志
    go runLogging(logs)
    return func(w http.ResponseWriter, r *http.Request) {
    	// 请求开始时间
        start := time.Now()
        // 处置请求方法
        f(w, r)
        // 请求方法
        method := r.Method
        // 请求路径
        path := r.URL.Path
        // 请求处理时间
        elapsed := float64(time.Since(start)) / 1000000.0
        //打印日志
        logs <- fmt.Sprintf(
            "method=%s path=%s elapsed=%f",
            method, path, elapsed)
    }
}

// 请求操作方法
func hello71(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    time.Sleep(time.Millisecond * 50)
    // 返回请求路径
    fmt.Fprintln(w, r.URL.Path)
}

func main() {
    handler := wrapLogging(hello71)
    http.HandleFunc("/", handler)
    http.ListenAndServe(":5000", nil)
}
