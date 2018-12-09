package main

import (
	"log"
	"net/http"
	"os"
)

func whoami(w http.ResponseWriter, req *http.Request) {
	version := os.Getenv("SHIPPER_RELEASE")
	podName := os.Getenv("POD_NAME")

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(version + " / " + podName + "\n"))
}

func main() {
	http.HandleFunc("/whoami", whoami)
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("ok\n"))
	})
	log.Fatal(http.ListenAndServe(":3000", nil))
}
