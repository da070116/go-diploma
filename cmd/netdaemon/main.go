package main

import (
	"fmt"
	"go-diploma/pkg/netdaemon"
	"go-diploma/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mainServer := utils.GetEnvVariable("MAINSERVERPATH")
	mainServerPort := utils.GetEnvVariable("MAINSERVERPORT")
	r := mux.NewRouter()
	r.HandleFunc("/", netdaemon.HandleConnection)
	r.HandleFunc("/data", netdaemon.PickDataConnection)
	http.ListenAndServe(fmt.Sprintf("%s:%s", mainServer, mainServerPort), r)
}
