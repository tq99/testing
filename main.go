package main

import (
	"io"
	"net/http"
)

func Display(w http.ResponseWriter, r *http.Request) {
	io.WriterTo("hellooo world")

}

func main() {

}
