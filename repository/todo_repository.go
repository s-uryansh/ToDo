package repository

import (
	"CRUD-SQL/model"
	"database/sql"
	"log"
)

// Defining Structs
type ToDoRepository interface {
	CreateToDo(*model.ToDo) error
	GetToDosByUserID(int) ([]*model.ToDo, error)
	UpdateToDo(*model.ToDo) error
	DeleteToDo(int) error
}

type toDoRepository struct {
	db *sql.DB
}

//Concrete code
//This will interact with database to perform CRUD

func NewToDoRepository(db *sql.DB) ToDoRepository {
	return &toDoRepository{db: db}
}

func (r *toDoRepository) CreateToDo(t *model.ToDo) error {
	query := "INSERT INTO todos (task, description, status, user_id) VALUES (?,?,?,?)"
	_, err := r.db.Exec(query, t.Task, t.Description, t.Status, t.UserID)
	return err
}

func (r *toDoRepository) GetToDosByUserID(userID int) ([]*model.ToDo, error) {
	query := "SELECT * FROM todos WHERE user_id =?"
	rows, err := r.db.Query(query, userID)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	defer rows.Close()

	var todos []*model.ToDo
	for rows.Next() {
		var t model.ToDo
		err := rows.Scan(&t.ID, &t.UserID, &t.Task, &t.Description, &t.Status)
		if err != nil {
			return nil, err
		}
		todos = append(todos, &t)
	}
	return todos, nil
}

func (r *toDoRepository) UpdateToDo(t *model.ToDo) error {
	_, err := r.db.Exec(`UPDATE todos SET task =?, description =?, status =? WHERE id =?`,
		t.Task, t.Description, t.Status, t.ID)
	return err
}

func (r *toDoRepository) DeleteToDo(id int) error {
	_, err := r.db.Exec(`DELETE FROM todos WHERE id =?`, id)
	return err
}

//Mock-implementations code
//This will not interact with database will only be used for testing

// type MockToDoRepository struct{}

// func (r *MockToDoRepository) CreateToDo(t *model.ToDo) error {

// }

// func (r *MockToDoRepository) GetToDo(id string) (*model.ToDo, error) {

// }

// func (r *MockToDoRepository) UpdateToDo(t *model.ToDo) error {

// }

// func (r *MockToDoRepository) DeleteToDo(id string) error {

// }
