package main

import (
    "fmt"
    "net/http"
)

type MyMux struct {
}

func (*MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
        sayHelloName(w, r)
        return
    }
    http.NotFound(w, r)
}

func sayHelloName(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprint(w, "Hello MyRouter.")
}

func main() {
    http.ListenAndServe(":9090", &MyMux{})
}
