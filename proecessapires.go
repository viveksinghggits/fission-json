package main

import (
    "net/http"
    "strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    v1 := r.Header.Get("previousres")
    lenStr:=len(string(v1))
    strres:=strconv.Itoa(lenStr)
    w.Write([]byte(strres))
}
