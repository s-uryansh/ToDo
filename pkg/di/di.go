package di

import (
	"CRUD-SQL/cache"
	"CRUD-SQL/handler"
	"CRUD-SQL/internal/config"
	"CRUD-SQL/repository"
	"CRUD-SQL/service"
	"database/sql"
	"log"
	"time"
)

type Container struct {
	UserService service.UserService
	ToDoService service.ToDoService
	UserHandler *handler.UserHandler
	ToDoHandler *handler.ToDoHandler
}

func NewContainer(db *sql.DB, cfg *config.Config) *Container {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	toDoRepo := repository.NewToDoRepository(db)
	toDoService := service.NewToDoService(toDoRepo)

	redisSessionStore, err := cache.NewRedisSessionStore("localhost:6379", "", 0, "myapp:sessions:", 30*time.Minute)
	if err != nil {
		log.Fatal(err)
	}

	userHandler := handler.NewUserHandler(userService, redisSessionStore)
	toDoHandler := handler.NewToDoHandler(toDoService, "localhost:6379") // <--- Add this line

	return &Container{
		UserService: userService,
		ToDoService: toDoService,
		UserHandler: userHandler,
		ToDoHandler: toDoHandler,
	}
}
