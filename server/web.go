package server

import (
    `fmt`
    `log`
    `net/http`
)

type WebServer struct {
}

var WS = WebServer{}

func (ws *WebServer) Run(dir string, port int) string{
    fs := http.FileServer(http.Dir(dir))
    addr := fmt.Sprintf("http://127.0.0.1:%d", port)
    addr2 := fmt.Sprintf(":%d", port)
    go func() {
        err := http.ListenAndServe(addr2, fs)
        if err != nil {
            log.Fatal(err)
        }
    }()
    return addr
}