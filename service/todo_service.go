package service

import (
	"CRUD-SQL/model"
	"CRUD-SQL/repository"
)

// Interface
type ToDoService interface {
	CreateToDo(*model.ToDo) error
	GetToDos(int) ([]*model.ToDo, error)
	UpdateToDo(*model.ToDo) error
	DeleteToDo(int) error
}

type toDoService struct {
	toDoRepo repository.ToDoRepository
}

// Concrete-Code
func NewToDoService(toDoRepo repository.ToDoRepository) ToDoService {
	return &toDoService{toDoRepo: toDoRepo}
}

func (s *toDoService) CreateToDo(t *model.ToDo) error {
	return s.toDoRepo.CreateToDo(t)
}

func (s *toDoService) GetToDos(userID int) ([]*model.ToDo, error) {
	return s.toDoRepo.GetToDosByUserID(userID)
}

func (s *toDoService) UpdateToDo(t *model.ToDo) error {
	return s.toDoRepo.UpdateToDo(t)
}

func (s *toDoService) DeleteToDo(id int) error {
	return s.toDoRepo.DeleteToDo(id)
}
