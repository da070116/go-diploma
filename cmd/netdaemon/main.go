package main

import (
	"github.com/gorilla/mux"
	"go-diploma/pkg/netdaemon"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", netdaemon.HandleConnection)
	r.HandleFunc("/data", netdaemon.PickDataConnection)
	http.ListenAndServe("127.0.0.1:8484", r)
}
