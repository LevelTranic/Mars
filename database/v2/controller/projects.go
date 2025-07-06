package controller

import (
	"Mars/database/shared"
	"Mars/database/v2/schemas"
)

func FindAllProjects() []schemas.Project {
	var projects []schemas.Project
	shared.DB.Find(&projects)
	return projects
}
