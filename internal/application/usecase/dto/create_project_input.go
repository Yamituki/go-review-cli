package dto

type CreateProjectInput struct {
	Name        string
	Type        string
	Framework   string
	Module      string
	Description string
	Path        string
}
