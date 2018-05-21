package main

import (
    "log"
    "net/http"
)

func outputRequestInfo(w http.ResponseWriter, r *http.Request) {
    url := ParseURL(r.URL.RequestURI())

    w.WriteHeader(http.StatusTeapot)
    w.Write([]byte(url.Output()))
}

func main() {
    http.HandleFunc("/", outputRequestInfo)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("server start error: ", err)
    }
}
