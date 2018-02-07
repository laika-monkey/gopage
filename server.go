package main

import (
    "io"
    "net/http"
    "golang.org/x/oauth2/fitbit"
)

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello world!")
}

func main() {
    http.HandleFunc("/", hello)

    tmp := fitbit.Endpoint
    io.WriteString(http.ResponseWriter, tmp)
    http.ListenAndServe(":8000", nil)
}
