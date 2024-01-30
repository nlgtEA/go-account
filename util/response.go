package util

import (
	"log"
	"net/http"
)

func ResponseWithError(w http.ResponseWriter, code int, msg string){
    w.WriteHeader(code)
    if code > 299 {
        log.Println(msg)
    }
    w.Write([]byte(msg))
    return
}
