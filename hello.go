package main

import (
	"io"
	"net/http"
	 "log"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func RootDir(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Choose command with [/command]!\n")
}

func main() {
	http.HandleFunc("/app/hello", HelloServer)
	http.HandleFunc("/", RootDir)
	log.Fatal(http.ListenAndServe(":12345", nil))
}