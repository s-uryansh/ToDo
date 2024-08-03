package di

import (
	"CRUD-SQL/internal/config"
	"CRUD-SQL/repository"
	"CRUD-SQL/service"
	"database/sql"
)

type Container struct {
	UserService service.UserService
	ToDoService service.ToDoService
}

func NewContainer(db *sql.DB, cfg *config.Config) *Container {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	toDoRepo := repository.NewToDoRepository(db)
	toDoService := service.NewToDoService(toDoRepo)

	return &Container{
		UserService: userService,
		ToDoService: toDoService,
	}
}
