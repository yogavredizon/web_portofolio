package repository_test

import (
	"fmt"
	"testing"

	"github.com/yogavredizon/web_portofolio/repository"
)

func TestRepository(t *testing.T) {
	repo := repository.ProjectRepositoryImpl{}

	result, err := repo.FindAll()
	if err != nil {
		t.Fatal(err)
	}

	for _, res := range result {

		fmt.Println(res)
	}
}

func TestSearch(t *testing.T) {

	type data struct {
		name, value string
	}

	value := data{
		name:  "Test Lower",
		value: "go",
		// {
		// },
		// {
		// 	name:  "Test Lower in Middle",
		// 	value: "q",
		// },
		// {
		// 	name:  "Test Upper",
		// 	value: "GO",
		// },
		// {
		// 	name:  "Test random",
		// 	value: "MySql",
		// },
	}

	t.Run(value.name, func(t *testing.T) {
		r, err := repository.NewProjectRepositoryImpl().Search(value.value)

		if err != nil {
			t.Fatal(err.Error())
		}
		fmt.Println(r)
	})
	// for _, v := range value {
	// }

}
