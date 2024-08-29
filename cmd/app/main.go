package main

import (
	"CRUD-SQL/auth"
	"CRUD-SQL/cache"
	"CRUD-SQL/controllers"
	"CRUD-SQL/handler"
	"CRUD-SQL/internal/config"
	"CRUD-SQL/pkg/di"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.New()

	tmplPath := "../../template/registration_email.html"
	_, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	r.Static("/static", "./public")

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("error loading config")
		return
	}

	db, errs := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName))
	if errs != nil {
		fmt.Println("error connecting db", errs)
		return
	}
	defer db.Close()
	Container := di.NewContainer(db, cfg)

	redisSessionStore, err := cache.NewRedisSessionStore("localhost:6379", "", 0, "myapp:sessions:", 30*time.Minute)
	if err != nil {
		log.Fatal(err)
	}

	userHandler := handler.NewUserHandler(Container.UserService, redisSessionStore)
	todoHandler := handler.NewToDoHandler(Container.ToDoService, "localhost:6379")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/sso", func(c *gin.Context) {
		controllers.SSO_signin(c.Writer, c.Request)
	})

	r.GET("/callback", func(c *gin.Context) {
		controllers.Callback(c.Writer, c.Request, db)
		c.Redirect(http.StatusFound, "/")
		// Add a home page (redirct after registration is done)
		// log.Println("user registered")

	})
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", auth.SessionMiddleware, userHandler.Logout)
	r.POST("/todo", auth.SessionMiddleware, todoHandler.CreateToDo)
	r.GET("/todo", auth.SessionMiddleware, todoHandler.GetToDos)
	r.Run(":8080")
}
