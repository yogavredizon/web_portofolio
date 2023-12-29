package repository_test

import (
	"fmt"
	"testing"

	"github.com/yogavredizon/web_portofolio/repository"
)

func TestTechRepository(t *testing.T) {
	// repo := repository.ProjectRepositoryImpl{}

	result, err := repository.FindById(2)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(result)
}
