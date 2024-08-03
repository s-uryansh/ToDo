package model

import "errors"

//Creating structs
type User struct {
	ID   int
	Name string
	Age  int
	Role string
	// ToDos []ToDo
}

type UserRegister struct {
	ID       int
	Name     string
	Email    string
	Password string
	Role     string
}

type UserLogin struct {
	Email    string
	Password string
}

//Validating structs

func (u *User) ValidateUser() error {
	if u.Name == "" || len(u.Name) < 3 {
		return errors.New("username is required and must be at least 3 characters long")
	}
	if u.Age < 0 || u.Age > 100 {
		return errors.New("invalid age")
	}
	if u.Role != "admin" && u.Role != "user" {
		return errors.New("role must be either 'admin' or 'user'")
	}
	return nil
}

func (u *UserRegister) ValidatingRegistration() error {
	if u.Name == "" || len(u.Name) > 50 {
		return errors.New("name is required (length b/w 0 to 50)")
	}

	if u.Password == "" || len(u.Password) > 255 {
		return errors.New("password is required (length b/w 0 to 255)")
	}

	if u.Email == "" || len(u.Email) > 100 {
		return errors.New("email is required (length b/w 0 to 100)")
	}

	if u.Role == "" {
		u.Role = "user"
	} else if u.Role == "admin_key" {
		u.Role = "admin"
	}

	return nil
}
