package main

import (
	"net/http"

	"github.com/yogavredizon/web_portofolio/view"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("asset"))))
	http.HandleFunc("/", view.Index)
	http.ListenAndServe(":8080", nil)
}
