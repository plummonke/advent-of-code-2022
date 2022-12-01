package utility

import (
    "bytes"
    "os"
)

func ReadFile(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()
    
    var t []byte
    text := bytes.NewBuffer(t)
    _, err = text.ReadFrom(f)
    if err != nil {
        return "", err
    }

    return string(text.Bytes()), nil
}

