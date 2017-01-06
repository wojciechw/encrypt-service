package main

import (
    "log"
    "flag"
    "strconv"
    "net/http"
)

func isPortValid(port string) bool {
    v, err := strconv.Atoi(port)
    if err != nil {
        return false
    }
    return v > 0
}

func main() {
    port := flag.String("port", "12345", "Server port")
    flag.Parse()

    if !isPortValid(*port) {
        log.Fatalf("Invalid port:%s\n", *port)
    }

    r := NewRouters()
    server := &http.Server{Addr: ":" + *port, Handler: r}
    log.Printf("Server starts on port:%s\n", *port)
    log.Println(server.ListenAndServe())
}

