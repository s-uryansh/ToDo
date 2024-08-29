package handler

import (
	"CRUD-SQL/cache"
	"CRUD-SQL/model"
	"CRUD-SQL/service"
	"CRUD-SQL/utils"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ToDoHandler struct {
	toDoService service.ToDoService
	redisStore  *cache.RedisSessionStore
}

func NewToDoHandler(toDoService service.ToDoService, redisAddr string) *ToDoHandler {
	ttl := 10 * time.Second
	prefix := "myapp:sessions:"
	redisStore, err := cache.NewRedisSessionStore(redisAddr, "", 0, prefix, ttl)
	if err != nil {
		return nil
	}
	return &ToDoHandler{toDoService: toDoService, redisStore: redisStore}
}
func (h *ToDoHandler) getUserFromRedis(sessionId string) (*model.UserRegister, error) {
	user, err := h.redisStore.GetSession(sessionId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (h *ToDoHandler) CreateToDo(c *gin.Context) {
	user, err := h.getUserFromRedis(utils.REDIS_KEY_TOKEN)
	if err != nil {
		log.Println(err, "REDIS")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve user from Redis",
		})
		return
	}
	var json model.ToDo
	json.UserID = user.ID
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("json", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}

	if err := json.ValidateToDo(); err != nil {
		fmt.Println("error validating json", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}

	// existingTask, err := h.toDoService.GetToDos(json.UserID)
	// if err != nil || json.Task = existingTask[0].Task{
	// 	fmt.Println("error creating task ", err)
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "something went wrong",
	// 	})
	// 	return
	// }

	if err := h.toDoService.CreateToDo(&json); err != nil {
		fmt.Println("error creating todo", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "todo created successfully",
	})
}

func (h *ToDoHandler) GetToDos(c *gin.Context) {
	user, err := h.getUserFromRedis(utils.REDIS_KEY_TOKEN)
	if err != nil {
		log.Println(err, "REDIS")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve user from Redis",
		})
		return
	}
	user_id := user.ID
	todos, err := h.toDoService.GetToDos(user_id)
	if err != nil {
		fmt.Println("error getting todos", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}

	var todoList []gin.H
	for _, todo := range todos {
		todoList = append(todoList, gin.H{
			"Task":        todo.Task,
			"Description": todo.Description,
		})
	}
	// jwtToken, err := c.Request.Cookie("tokenString")
	// if err != nil {
	// 	log.Println("error getting jwt token: ", err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "internal server error",
	// 	})
	// 	return
	// }
	// fmt.Println("JWTTOKEN-String", jwtToken.Value)
	c.JSON(http.StatusOK, gin.H{
		"todos": todoList,
	})
}

func (h *ToDoHandler) UpdateToDo(c *gin.Context) {
	var json model.ToDo
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("json")

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}
	if err := json.ValidateToDo(); err != nil {
		fmt.Println("error validating json", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}

	if err := h.toDoService.UpdateToDo(&json); err != nil {
		fmt.Println("error updating todo", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "todo updated successfully",
	})
}
