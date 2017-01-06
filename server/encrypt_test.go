package main

import (
    "testing"
)

func TestEncryptDecrypt(t *testing.T) {
    msg := []byte("Message to encrypt")
    key, err := GenerateKey()

    if err != nil {
        t.Error(err)
    }

    enc, err := Encrypt(key, msg)
    if err != nil {
        t.Error(err)
    }

    dec, err := Decrypt(key, enc)
    if err != nil {
        t.Error(err)
    }

    sEnc := string(msg)
    sDec := string(dec)

    if sEnc != sDec {
        t.Errorf("Msg[%s] != Decrypt(msg)[%s]", sEnc, sDec)
    }
}
