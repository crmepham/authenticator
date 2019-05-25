package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	s := &server{}
	s.init()
	log.Fatal(http.ListenAndServe(":8002", s.router))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fail(w http.ResponseWriter, e string, c int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("WWW-Authenticate", "Basic realm=\""+serverName+"\"")
	w.WriteHeader(c)
	u, _ := json.Marshal(&Error{e})
	_, err := w.Write([]byte(string(u)))
	check(err)
}
