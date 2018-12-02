package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	l "github.com/tigrouxyz/email-subscribe-manager/log"
)

func main() {
	startServer()
}

func startServer() {
	setUpServer()
}

func setUpServer() {
	http.HandleFunc("/subscribe", suscribe)
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
	email := strings.Split(string(body), "=")
	l.Info(email[1])

	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("subscribe.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(email[1] + "\n")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	l.Info("Sending back response to confirm subscription")

	var jsonStr = []byte("<html><head></head><body><h1>Votre inscription a bien été prise en compte.</h1></body></html>")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	//Write json response back to response
	w.Write(jsonStr)
}
