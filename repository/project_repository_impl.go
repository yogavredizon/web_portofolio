package repository

import (
	"database/sql"

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
	techStacks := []models.TechStack{}

	stmt, err := NewConn().Prepare("SELECT project_experience.title, project_experience.link_doc, project_experience.link_github, project_experience.summary, tech_used.created_at, tech_stacks.alias FROM project_experience join tech_used on (tech_used.id_experience = project_experience.id) join tech_stacks on (tech_stacks.id=tech_used.id_tech)")

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
		var techStack models.TechStack

		rows.Scan(&project.Title, &project.LinkDoc, &project.LinkGithub, &project.Description, &techStack.Date, &techStack.Alias)

		techStacks = append(techStacks, techStack)
		project.Stack = techStacks
		projects = append(projects, project)

	}

	if err != nil {
		return projects, err
	}

	return projects, nil
}

// func (p *ProjectRepositoryImpl) FindByTitle(title string) ([]models.Project, error) {

// }
