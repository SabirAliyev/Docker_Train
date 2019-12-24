package main

import (
        "fmt"
        "net/http"
)

func main() {
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                _, _ = fmt.Fprintf(w, "Wellcome to the main page! If you see this message it means that the page is working. ")
        })

        http.HandleFunc("hello/", func(w http.ResponseWriter, r *http.Request) {
                _, _ = fmt.Fprintf(w, "Hello World")
        })

        fmt.Println("Start server")
        _ = http.ListenAndServe(":8001", nil)
}
