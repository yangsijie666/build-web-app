package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
)

func main() {
    http.HandleFunc("/", sayHelloName)
    http.HandleFunc("/login", login)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal(err)
    }
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Method: ", r.Method)
    if r.Method == "GET" {
        t, err := template.ParseFiles("/Users/yangsijie/go/src/build-web-app/4-1/view/login.html")
        if err != nil {
            log.Println(err)
        }
        log.Println(t.Execute(w, nil))
    } else {
        r.ParseForm()
        //请求的是登录数据，那么执行登录的逻辑判断
        fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
        fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
    }
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println(r.Form)
    fmt.Println("path: ", r.URL.Path)
    fmt.Println("scheme: ", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprint(w, "Hello ysj~")
}
