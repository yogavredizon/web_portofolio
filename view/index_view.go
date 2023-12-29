package view

import (
	"html/template"
	"net/http"
	"path"

	"github.com/yogavredizon/web_portofolio/controller"
)

func Index(w http.ResponseWriter, r *http.Request) {

	var filePath = path.Join("asset", "index.html")
	var tmpl, _ = template.ParseFiles(filePath)

	data := controller.IndexController(w, r)
	tmpl.Execute(w, data)
}

func IndexSearch(w http.ResponseWriter, r *http.Request) {

	var filePath = path.Join("asset", "index.html")
	var tmpl, _ = template.ParseFiles(filePath)

	data := controller.SearchProjectController(w, r)
	tmpl.Execute(w, data)
}
