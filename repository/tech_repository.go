package repository

import "github.com/yogavredizon/web_portofolio/models"

func FindById(id int) ([]models.TechStack, error) {
	techStacks := []models.TechStack{}
	stmt, err := NewConn().Prepare(" SELECT t.id, t.title, t.description, t.alias from tech_used join tech_stacks as t on (t.id= tech_used.id_tech) where tech_used.id_experience =?")

	if err != nil {
		return techStacks, err
	}

	rowsStack, _ := stmt.Query(id)

	for rowsStack.Next() {
		var techStack models.TechStack
		rowsStack.Scan(&techStack.Id, &techStack.Title, &techStack.Description, &techStack.Alias)
		techStacks = append(techStacks, techStack)
	}

	return techStacks, nil
}
