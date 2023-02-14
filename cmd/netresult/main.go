package main

import (
	"Skillbox-diploma/pkg/netresult"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", netresult.HandleConnection)
	r.HandleFunc("/data", netresult.PickDataConnection)
	http.ListenAndServe("127.0.0.1:8282", r)
}
