package handler

import (
	"CRUD-SQL/model"
	"CRUD-SQL/service"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ToDoHandler struct {
	toDoService service.ToDoService
}

func NewToDoHandler(toDoService service.ToDoService) *ToDoHandler {
	return &ToDoHandler{toDoService: toDoService}
}

func (h *ToDoHandler) CreateToDo(c *gin.Context) {
	var json model.ToDo
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

	existingTask, err := h.toDoService.GetToDos(json.UserID)
	if err != nil || existingTask != nil {
		fmt.Println("error creating task ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}

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
	user_id_str := c.Request.URL.Query().Get("ID")
	user_id, err := strconv.Atoi(user_id_str)
	if err != nil {
		log.Println("can not convert useridstr to userid", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}
	todos, err := h.toDoService.GetToDos(user_id)
	if err != nil {
		fmt.Println("error getting todos", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Description": todos[0].Description,

		"Title": todos[0].Task,
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
