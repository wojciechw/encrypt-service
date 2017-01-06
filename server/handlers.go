package main

import (
    "log"
    "net/http"
    "io/ioutil"
    "encoding/hex"
    "github.com/gorilla/mux"
    "github.com/golang/protobuf/proto"
    "github.com/wojciechw/encrypt-service/models"
    "github.com/wojciechw/encrypt-service/storage"
)

func NewRouters() *mux.Router {
    r := mux.NewRouter().StrictSlash(false)
    r.HandleFunc("/encrypt", EncryptHandler).Methods("POST")
    r.HandleFunc("/decrypt", DecryptHandler).Methods("POST")

    return r
}

func EncryptHandler(w http.ResponseWriter, r *http.Request) {
    data, err := ioutil.ReadAll(r.Body)
    if err != nil {
        writeError(w, "Server internal error!", err)
        return
    }

    var eData models.EncryptReq
    err = proto.Unmarshal(data, &eData)
    if err != nil {
        writeError(w, "Invalid request data!", err)
        return
    }

    log.Println(eData.String())

    // Encrypt
    key, _ := GenerateKey()
    encrypted, err := Encrypt(key, eData.GetData())
    if err != nil {
        writeError(w, "Encrypt failed!", err)
        return
    }

    // Store encrypted data
    storage.Store(string(eData.GetId()), encrypted)

    res := &models.EncryptRes{Key:key[:KeySize]}
    data, err = proto.Marshal(res)
    if err != nil {
        writeError(w, "Server internal error!", err)
        return
    }

    log.Println("key:", hex.EncodeToString(res.GetKey()))

    w.Header().Set("Content-Type", "application/x-protobuf")
    w.WriteHeader(http.StatusOK)
    w.Write(data)
}

func DecryptHandler(w http.ResponseWriter, r *http.Request) {
    data, err := ioutil.ReadAll(r.Body)
    if err != nil {
        writeError(w, "Server internal error!", err)
        return
    }

    var dData models.DecryptReq
    err = proto.Unmarshal(data, &dData)
    if err != nil {
        writeError(w, "Invalid request data!", err)
        return
    }

    if len(dData.Id) == 0 ||  len(dData.Key) < KeySize {
        writeError(w, "Invalid id or key length!", nil)
        return
    }

    log.Println(dData.String())

    // Retrieve encrypted data
    encrypted, err := storage.Retrieve(string(dData.GetId()))
    if err != nil {
        writeError(w, "Server internal error!", err)
        return
    }

    // Decrypt
    var key [KeySize]byte
    copy(key[:], dData.Key[:KeySize])
    decrypted, err := Decrypt(&key, encrypted)
    if err != nil {
        writeError(w, "Decrypt failed!", err)
        return
    }

    res := &models.DecryptRes{Data:decrypted}
    data, err = proto.Marshal(res)
    if err != nil {
        writeError(w, "Server internal error!", err)
        return
    }

    log.Println(res.String())

    w.Header().Set("Content-Type", "application/x-protobuf")
    w.WriteHeader(http.StatusOK)
    w.Write(data)
}

func writeError(w http.ResponseWriter, text string, err error) {
    w.WriteHeader(http.StatusInternalServerError)
    log.Println(text, err)
}

