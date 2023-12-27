package repository

import "github.com/yogavredizon/web_portofolio/models"

type ProjectRepository interface {
	FindAll() ([]models.Project, error)
	// FindByTitle(title string) ([]models.Project, error)
}
