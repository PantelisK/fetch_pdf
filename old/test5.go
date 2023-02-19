package main

import (
    "fmt"
    "net/http"
    "io"
)

var client = http.Client{}

func cutterHandler(res http.ResponseWriter, req *http.Request) {
    reqImg, err := client.Get("http://95.217.49.162:3000/dummy.png")
    if err != nil {
        fmt.Fprintf(res, "Error %d", err)
        return
    }
    res.Header().Set("Content-Length", fmt.Sprint(reqImg.ContentLength))
    res.Header().Set("Content-Type", reqImg.Header.Get("Content-Type"))
    if _, err = io.Copy(res, reqImg.Body); err != nil {
        // handle error
    }
    reqImg.Body.Close()
}

func main() {
    http.HandleFunc("/cut", cutterHandler)
    http.ListenAndServe(":8080", nil) /* TODO Configurable */
}