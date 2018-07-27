package main

import "net/http"

// 共享文件系统
func main() {
    handler := http.FileServer(http.Dir("./"))
    http.ListenAndServe(":5000", handler)
}
