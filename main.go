package main

import (
	"net/http"

	"github.com/rasyidev/dts-h8-assignment-iii/controller"
)

func main() {

}

func RunServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		controller.RenderHTML(w, r)
	})

	server := http.Server{
		Addr:    "localhost:9090",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
