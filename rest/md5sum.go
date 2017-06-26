package main

import (
	"bytes"
	"crypto/md5"
	"fmt"

	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/md5sum", ReqHandler).Methods("POST")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}

func ReqHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%x", md5sum(r.FormValue("what")))
}

// function to calculate the md5sum and return the same
func md5sum(data string) string {

	h := md5.New()
	if _, err := io.Copy(h, bytes.NewBufferString(data)); err != nil {
		log.Fatal(err)
	}
	b := h.Sum(nil)
	fmt.Printf("%x", b)
	return string(b[:])
}
