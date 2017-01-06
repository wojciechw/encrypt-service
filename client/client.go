package main

import (
    "bytes"
    "errors"
    "net/http"
    "io/ioutil"
    "github.com/golang/protobuf/proto"
    "github.com/wojciechw/encrypt-service/models"
)

type HttpClient struct {
    Http http.Client
    Url string
}

func NewHttpClient(url string) *HttpClient {
    return &HttpClient{Url:url}
}

func (c *HttpClient) PerformRequest(url string, data []byte) ([]byte, error) {
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
    if (err != nil) {
        return nil, err
    }

    req.Header.Set("Content-Type", "application/x-protobuf")
    res, err := c.Http.Do(req)
    if err != nil {
        return nil, err
    }

    data, err = ioutil.ReadAll(res.Body)
    if err != nil {
        return nil, err
    }

    if res.StatusCode != http.StatusOK {
        return nil, errors.New("Server internal error!")
    }
    return data, nil
}

func (c *HttpClient) Store(id, payload []byte) (aesKey []byte, err error) {
    eReq := &models.EncryptReq{Id:id, Data:payload}
    data, err := proto.Marshal(eReq)
    if err != nil {
        return nil, err
    }

    data, err = c.PerformRequest(c.Url + "/encrypt", data)
    if err != nil {
        return nil, err
    }

    var eRes models.EncryptRes
    err = proto.Unmarshal(data, &eRes)
    if err != nil {
        return nil, err
    }
    return eRes.GetKey(), nil
}

func (c *HttpClient) Retrieve(id, aesKey []byte) (payload []byte, err error) {
    dReq := &models.DecryptReq{Id:id, Key:aesKey}
    data, err := proto.Marshal(dReq)
    if err != nil {
        return nil, err
    }

    data, err = c.PerformRequest(c.Url + "/decrypt", data)
    if err != nil {
        return nil, err
    }

    var dRes models.DecryptRes
    err = proto.Unmarshal(data, &dRes)
    if err != nil {
        return nil, err
    }
    return dRes.Data, nil
}

