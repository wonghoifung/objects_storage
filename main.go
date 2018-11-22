package main

import (
	"log"
	"net/http"
	//"os"
	"./objects"
)

func main() {
	http.HandleFunc("/objects/", objects.Handler)
	//log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
	log.Fatal(http.ListenAndServe("0.0.0.0:8888", nil))
}
