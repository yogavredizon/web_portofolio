package controller

import (
	"net/http"

	"github.com/yogavredizon/web_portofolio/repository"
)

func IndexController(w http.ResponseWriter, r *http.Request) map[string]interface{} {

	if r.Method == "GET" {
		project := repository.ProjectRepositoryImpl{}

		result, _ := project.FindAll()

		data := map[string]interface{}{
			"title": "My Portofolio",
			"value": result,
		}

		return data
	}

	return map[string]interface{}{}
}
func SearchProjectController(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	search := r.FormValue("search")

	if r.Method == "POST" {

		repo := repository.ProjectRepositoryImpl{}

		result, _ := repo.Search(search)

		data := map[string]interface{}{
			"value": result,
		}

		return data
	}

	return map[string]interface{}{}
}
