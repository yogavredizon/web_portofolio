package models

type Project struct {
	Id                                      int
	Title, LinkDoc, LinkGithub, Description string
	Stack                                   []TechStack
}
