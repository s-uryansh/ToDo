package main

import (
	"CRUD-SQL/handler"
	"CRUD-SQL/internal/config"
	"CRUD-SQL/pkg/di"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	tmplPath := "../../template/registration_email.html"
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, nil)
		if err != nil {
			log.Fatalf("Error executing template: %v", err)
		}
	})
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("error loading config")
		return
	}
	// log.Println(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
	// 	cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName))

	db, errs := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName))
	if errs != nil {
		fmt.Println("error connecting db", errs)
		return
	}
	defer db.Close()

	Container := di.NewContainer(db, cfg)
	userHandler := handler.NewUserHandler(Container.UserService)
	todoHandler := handler.NewToDoHandler(Container.ToDoService)
	// userHandler := handler.NewUserHandler(userContainer.UserService)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)
	r.POST("/todo", todoHandler.CreateToDo)
	r.GET("/todo", todoHandler.GetToDos)
	r.Run()
}
