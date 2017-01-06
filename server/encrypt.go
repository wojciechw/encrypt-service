package main

import (
    "io"
    "errors"
    "crypto/rand"
    "golang.org/x/crypto/nacl/secretbox"
)

// https://leanpub.com/gocrypto/read#leanpub-auto-chapter-3-symmetric-security

const (
    KeySize = 32
    NonceSize = 24
)
var (
    ErrEncrypt = errors.New("encryption failed")
    ErrDecrypt = errors.New("decryption failed")
)

func GenerateKey() (*[KeySize]byte, error) {
    key := new([KeySize]byte)
    _, err := io.ReadFull(rand.Reader, key[:])
    if err != nil {
        return nil, err
    }
    return key, nil
}

func GenerateNonce() (*[NonceSize]byte, error) {
    nonce := new([NonceSize]byte)
    _, err := io.ReadFull(rand.Reader, nonce[:])
    if err != nil {
        return nil, err
    }
    return nonce, nil
}

func Encrypt(key *[KeySize]byte, message []byte) ([]byte, error) {
    nonce, err := GenerateNonce()
    if err != nil {
        return nil, ErrEncrypt
    }

    out := make([]byte, len(nonce))
    copy(out, nonce[:])
    out = secretbox.Seal(out, message, nonce, key)
    return out, nil
}

func Decrypt(key *[KeySize]byte, message []byte) ([]byte, error) {
    if len(message) < (NonceSize + secretbox.Overhead) {
        return nil, ErrDecrypt
    }

    var nonce [NonceSize]byte
    copy(nonce[:], message[:NonceSize])
    out, ok := secretbox.Open(nil, message[NonceSize:], &nonce, key)
    if !ok {
        return nil, ErrDecrypt
    }

    return out, nil
}

