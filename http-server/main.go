package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", SayHello)

    log.Fatal(http.ListenAndServe(":8080", nil))

}

func SayHello(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello World!")
}
