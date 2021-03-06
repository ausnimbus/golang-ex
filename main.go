package main

import (
    "fmt"
    "net/http"
)

func main() {
    panic(http.ListenAndServe(":8080", &logServer{
        hdl: http.FileServer(http.Dir("../src/main/views")),
    }))
}

type logServer struct {
    hdl http.Handler
}

func (l *logServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.RemoteAddr, r.URL, r.URL.Path)
    l.hdl.ServeHTTP(w, r)
}
