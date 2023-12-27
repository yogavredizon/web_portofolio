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
		for _, r := range res.Stack {
			fmt.Println(r.Alias)
		}
	}
}
