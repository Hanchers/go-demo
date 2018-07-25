package main

import (
    "fmt"
    "net/http"
    "strings"
)

type Auth func(string, string) bool

//验证权限
func testAuth(r *http.Request, auth Auth) bool {
    //获取Authorization的header值
    header := r.Header.Get("Authorization")
    fmt.Println(header)
    //按空格分组
    s := strings.SplitN(header, " ", 2)
    if len(s) != 2 || s[0] != "Basic" {
        return false
    }

    //正常密码应该是通过base64加密的，此处没用
    //b, err := base64.StdEncoding.DecodeString(s[1])
    //if err != nil {
    //    return false
    //}
    //pair := strings.SplitN(string(b), ":", 2)

    //通过：分割，获取明文密码
    pair := strings.SplitN(s[1], ":", 2)
    if len(pair) != 2 {
        return false
    }
    return auth(pair[0], pair[1])
}

//验证失败，返回401
func requireAuth(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("WWW-Authenticate",
        "Basic realm=\"private\"")
    w.WriteHeader(401)
    w.Write([]byte("401 Unauthorized\n"))
}

//权限验证：通过=hello world,失败=401
func wrapAuth(h http.HandlerFunc, a Auth) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if testAuth(r, a) {
            h(w, r)
        } else {
            requireAuth(w, r)
        }
    }
}

//权限通过返回值
func hello68(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello secret world!")
}

//url=localhost:5000
//header={"Authorization":"Basic pwd:supersecret"}
func main() {
    //验证密码是否匹配
    checkPassword := func(_, password string) bool {
        return password == "supersecret"
    }

    handler1 := http.HandlerFunc(hello68)
    handler2 := wrapAuth(handler1, checkPassword)
    http.ListenAndServe(":5000", handler2)
}
