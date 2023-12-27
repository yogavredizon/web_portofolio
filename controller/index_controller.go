package controller

import (
	"net/http"

	"github.com/yogavredizon/web_portofolio/repository"
)

func IndexController(w http.ResponseWriter, r *http.Request) map[string]interface{} {

	if r.Method == "GET" {
		repo := repository.ProjectRepositoryImpl{}

		result, _ := repo.FindAll()
		data := map[string]interface{}{
			"title": "My Portofolio",
			"value": result,
		}

		return data
	}

	return map[string]interface{}{}
}
