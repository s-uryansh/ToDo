package model

import (
	"errors"
	"log"
)

// Defining Structs
type ToDo struct {
	ID          int
	UserID      int
	Task        string
	Description string
	Status      bool
	User        *User
}

// Validating structs
func (t *ToDo) ValidateToDo() error {
	if t.Status != true && t.Status != false {
		return errors.New("done must be a boolean value")
	}

	if t.UserID == 0 {
		return errors.New("userID is required ")
	}

	if t.Task == "" {
		log.Println(t.Task)
		return errors.New("title is required or should be under 255 chars ")
	}

	if t.Description == "" {
		return errors.New("description is required")
	}
	return nil
}
