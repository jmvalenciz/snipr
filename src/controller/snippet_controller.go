package controller

import (
	"snipr/src/model"
	"snipr/src/repository"
)

type SnippetController struct {
	repository repository.ISnippetRepository
}

func NewSnippetController(repository repository.ISnippetRepository) SnippetController {
	return SnippetController{
		repository,
	}
}

func (sc SnippetController) GetSnippets() ([]model.Snippet, error) {
	return sc.repository.GetAll()
}

func (sc SnippetController) DeleteById(id uint) (bool, error) {
	return sc.repository.Delete(id)
}

func (sc SnippetController) Create(newSnippet model.CreateSnippet) (*model.Snippet, error) {
	return sc.repository.Create(newSnippet)
}
