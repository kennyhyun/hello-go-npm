package main

import (
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHelloHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(helloHandler)

    handler.ServeHTTP(rr, req)

    expected := "Hello, World!"
    body, _ := ioutil.ReadAll(rr.Body)
    if string(body) != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            string(body), expected)
    }
}
