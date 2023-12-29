package repository

import (
	"database/sql"
	"errors"
	"regexp"
	"strings"

	"github.com/yogavredizon/web_portofolio/app"
	"github.com/yogavredizon/web_portofolio/models"
)

type ProjectRepositoryImpl struct{}

func NewProjectRepositoryImpl() ProjectRepository {
	return &ProjectRepositoryImpl{}
}

func NewConn() *sql.DB {
	return app.NewDB()
}

func (p *ProjectRepositoryImpl) FindAll() ([]models.Project, error) {
	projects := []models.Project{}

	stmt, err := NewConn().Prepare("SELECT p.id, p.title, p.link_doc, p.link_github, p.summary, p.created_at FROM project_experience as p")

	if err != nil {
		return projects, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return projects, err
	}

	for rows.Next() {
		var project models.Project

		rows.Scan(&project.Id, &project.Title, &project.LinkDoc, &project.LinkGithub, &project.Description, &project.Created_at)
		projects = append(projects, project)
	}

	for i := 0; i < len(projects); i++ {
		s, _ := FindById(projects[i].Id)
		projects[i].Stack = s
	}

	return projects, nil
}

func (p *ProjectRepositoryImpl) Search(data string) ([]models.Project, error) {
	projects, err := p.FindAll()

	if err != nil {
		return projects, err
	}
	result := []models.Project{}
	data = strings.ToLower(data)

	for _, project := range projects {

		searchByTitle, _ := regexp.Match(`^.*`+data+`.*`, []byte(strings.ToLower(project.Title)))
		for i := 0; i < len(project.Stack); i++ {
			searchByTechStack, _ := regexp.Match(`^.*`+data+`.*`, []byte(strings.ToLower(project.Stack[i].Title)))

			if searchByTitle || searchByTechStack {
				result = append(result, project)
				break
			}
		}
	}

	if len(result) == 0 {
		return result, errors.New(`we don't have what you looking for`)
	}

	return result, nil
}
