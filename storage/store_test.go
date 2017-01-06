package storage

import (
    "testing"
)

func TestStore(t *testing.T) {
    id := "11111"
    b := []byte("Test byte 1234")
    Store(id, b)

    r, err := Retrieve(id)

    if err != nil {
        t.Error("Store does not find id ", id)
    }

    if len(b) != len(r) {
            t.Error(" len(b) != len(r)")
    }

    for i := range b {
        if (b[i] != r[i]) {
            t.Error("Stored bytes != to retrieved bytes", b, r)
        }
    }
}
