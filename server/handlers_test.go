package main

import (
    "bytes"
    "testing"
    "net/http"
    "io/ioutil"
    "net/http/httptest"
    "github.com/golang/protobuf/proto"
    "github.com/wojciechw/encrypt-service/models"
)

func TestEncryptDecryptEndpoints(t *testing.T) {
    text := "TestTest"
    r := NewRouters()
    // Encrypt Data
    eReq := &models.EncryptReq{Id:[]byte("1"), Data:[]byte(text)}
    data, err := proto.Marshal(eReq)
    if err != nil {
        t.Error(err)
    }
    req, err := http.NewRequest("POST", "/encrypt", bytes.NewBuffer(data))
    if err != nil {
        t.Error(err)
    }

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    if w.Code != 200 {
        t.Errorf("HTTP Status Encrypt excepted: 200, got: %d", w.Code)
    }

    data, err = ioutil.ReadAll(w.Body)
    if err != nil {
        t.Error(err)
    }

    var eRes models.EncryptRes
    err = proto.Unmarshal(data, &eRes)
    if err != nil {
        t.Error(err)
    }

    // Decrypt Data
    dReq := &models.DecryptReq{Id:[]byte("1"), Key:eRes.GetKey()}
    data, err = proto.Marshal(dReq)
    if err != nil {
        t.Error(err)
    }

    req, err = http.NewRequest("POST", "/decrypt", bytes.NewBuffer(data))
    if err != nil {
        t.Error(err)
    }

    r.ServeHTTP(w, req)
    if w.Code != 200 {
        t.Errorf("HTTP Status Decrypt excepted: 200, got: %d", w.Code)
    }

    if w.Code != 200 {
        t.Errorf("HTTP Status excepted: 200, got: %d", w.Code)
    }

    data, err = ioutil.ReadAll(w.Body)
    if err != nil {
        t.Error(err)
    }

    var dRes models.DecryptRes
    err = proto.Unmarshal(data, &dRes)
    if err != nil {
        t.Error(err)
    }

    if text != string(dRes.GetData()) {
        t.Errorf("%s != %s", text, string(dRes.GetData()))
    }
}


