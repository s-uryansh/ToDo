package repository

import (
	"CRUD-SQL/model"
	"database/sql"
)

// Defining Structs
type UserRepository interface {
	CreateUser(*model.UserRegister) error
	// GetUser(int) (*model.UserRegister, error)
	GetUserByEmail(string) (*model.UserRegister, error)
	CreateUserBySSO(string, string, string) error
	// UpdateUser(*model.User) error
	// DeleteUser(string) error
}
type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

//Concrete-implementations code
//This will interact with database to perform CRUD

func (r *userRepository) CreateUser(u *model.UserRegister) error {
	query := "INSERT INTO users_register (Name , Email , Password , Role) VALUES (?,?,?,?)"
	_, err := r.db.Exec(query, u.Name, u.Email, u.Password, u.Role)
	// log.Println(err)
	return err
}

func (r *userRepository) GetUserByEmail(emial string) (*model.UserRegister, error) {
	query := "SELECT * FROM users_register WHERE email = ?"
	row := r.db.QueryRow(query, emial)
	user := &model.UserRegister{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) CreateUserBySSO(pass, email, name string) error {
	// Generate a random password
	password := pass
	query := "INSERT INTO users_register (Name , Email , Password , Role) VALUES (?,?,?,?)"
	_, err := r.db.Exec(query, name, email, password, "sso_user")
	return err
}

// func (r *MySQLUserRepository) UpdateUser(u *model.User) error {
// 	query := ""
// }

// func (r *MySQLUserRepository) DeleteUser(id string) error {

// }

//Mock-implementations code
//This will not interact with database will only be used for testing

// type MockUserRepository struct{}

// func (r *MockUserRepository) CreateUser(u *model.User) error {

// }

// func (r *MockUserRepository) GetUser(id string) (*model.User, error) {

// }

// func (r *MockUserRepository) UpdateUser(u *model.User) error {

// }

// func (r *MockUserRepository) DeleteUser(id string) error {

// }
