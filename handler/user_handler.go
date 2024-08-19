package handler

import (
	"CRUD-SQL/auth"
	"CRUD-SQL/cache"
	"CRUD-SQL/jobs"
	"CRUD-SQL/model"
	"CRUD-SQL/service"
	"CRUD-SQL/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService  service.UserService
	sessionStore *cache.RedisSessionStore
}

func NewUserHandler(userService service.UserService, sessionStore *cache.RedisSessionStore) *UserHandler {
	return &UserHandler{userService: userService, sessionStore: sessionStore}
}

func (h *UserHandler) Register(c *gin.Context) {
	var json model.UserRegister

	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println("json")

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}
	// log.Println(json.Password)

	if err := json.ValidatingRegistration(); err != nil {
		fmt.Println("error validating json", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}

	if err := h.userService.RegisterUser(&json); err != nil {
		fmt.Println("error registring user", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}
	if err := jobs.SendEmailRegistration(json.Email, json.Name); err != nil {
		fmt.Println("error sending register mail: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		json.Name : " has succesfully created account",
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var json model.UserLogin
	if err := c.ShouldBindJSON(&json); err != nil {
		log.Println("error json", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	user, err := h.userService.LoginUser(json.Email, json.Password)
	if err != nil {
		log.Println("error in userLoginService: ", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	errs := auth.JwtTokenCreate(c)
	if errs != nil {
		log.Println("error creating jwt token: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	err = h.sessionStore.SetSession(utils.REDIS_KEY_TOKEN, user)
	if err != nil {
		log.Println("error setting session: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		user.Name: "successfully logged in",
	})
}

func (h *UserHandler) Logout(c *gin.Context) {
	cookie, err := c.Request.Cookie("session_id")
	if err != nil {
		fmt.Println("can not request cookie")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		c.Abort()
		return
	}
	auth.InvalidSession(cookie.Value)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		MaxAge: -1,
	})

	token := utils.REDIS_KEY_TOKEN
	errs := h.sessionStore.DeleteSession(token)
	if errs != nil {
		log.Println("error deleting session: ", errs)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
	c.JSON(http.StatusOK, gin.H{
		"message": "logged out successfully",
	})
}
