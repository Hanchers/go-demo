package main

import "net/http"
import "strings"
import "fmt"

type handler http.HandlerFunc

func hello70(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    fmt.Fprintln(w, "Hello canonical world")
}

// 包裹主域名
func wrapCanonicalHost(f handler, chost string) handler {
    return func(w http.ResponseWriter, r *http.Request) {
        // 请求原端口
        hostPort := strings.Split(r.Host, ":")
        // 请求原域名
        host := hostPort[0]
        fmt.Println("origin host : ",host)
        //如果域名与指定不一致，则重定向到指定地址
        if host != chost {
            fmt.Println("redirect to", chost)
            hostPort[0] = chost
            url := "http://" +
                strings.Join(hostPort, ":") +
                r.URL.String()
            http.Redirect(w, r, url, 301)
            return
        }
        f(w, r)
    }
}

// 测试重定向方法
func main() {
    handler := wrapCanonicalHost(hello70, "localhost")
    http.HandleFunc("/", handler)
    http.ListenAndServe(":5000", nil)
}
