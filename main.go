package main

import (
	"io/ioutil"
	"net/http"

	l "github.com/tigrouxyz/email-suscribe-manager/log"
)

func main() {
	startServer()
}

func startServer() {
	setUpServer()
}

func setUpServer() {
	http.HandleFunc("/suscribe", suscribe)
	// set router
	err := http.ListenAndServe(":9090", nil)
	// set listen port
	if err != nil {
		l.Err("ListenAndServe: ", err)
	}
}

func suscribe(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	l.Info(string(body))
}
