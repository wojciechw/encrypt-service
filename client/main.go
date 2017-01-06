package main

import (
    "log"
    "fmt"
    "flag"
    "net/url"
    "encoding/hex"
)

const (
    URL = "http://localhost:12345"
)

func Encrypt(c *HttpClient, id, data string) {
    key, err := c.Store([]byte(id), []byte(data))
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(hex.EncodeToString(key))
}

func Decrypt(c *HttpClient, id, key string) {
    b, err := hex.DecodeString(key)
    if err != nil {
        log.Fatal(err)
    }
    data, err := c.Retrieve([]byte(id), b)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(data))
}

func main() {
    var e bool
    var d bool
    var id string
    var data string
    var key string
    var addr string
    flag.BoolVar(&e, "e", false, "encrypt")
    flag.BoolVar(&d, "d", false, "decrypt")
    flag.StringVar(&id, "id", "", "user id")
    flag.StringVar(&data, "data", "", "data")
    flag.StringVar(&key, "key", "", "encryption key")
    flag.StringVar(&addr, "addr", URL, "server address")
    flag.Parse()

    if flag.NFlag() < 3 {
        flag.PrintDefaults()
        return
    }

    if e && d {
        log.Fatal("Can't use -d -e, choose one option!")
    }

    if e && (len(id) == 0 || len(data) == 0) {
        log.Fatal("Invalid id or data length!")
    }

    if d && (len(id) == 0 || len(key) == 0) {
        log.Fatal("Invalid id or key length!")
    }

    if _, err := url.Parse(addr); err != nil {
        log.Fatal("Invalid server url ", addr)
    }

    c := NewHttpClient(addr)
    if e {
        Encrypt(c, id, data)
    }

    if d {
        Decrypt(c, id, key)
    }
}
