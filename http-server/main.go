package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", sayHello)

    log.Fatal(http.ListenAndServe(":8080", nil))

}

func sayHello(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html")
        fmt.Fprintf(w, "Hello World! From Golang.")
}
