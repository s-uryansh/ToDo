// This code is not actually optimized much , hence this part of project is not that good as other
package controllers

import (
	"CRUD-SQL/jobs"
	"CRUD-SQL/model"
	"CRUD-SQL/repository"
	"CRUD-SQL/service"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func Callback(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	state := r.FormValue("state")
	code := r.FormValue("code")

	data, err := getUserData(state, code)
	if err != nil {
		log.Println("can not fetch user", err)
		return
	}

	h := NewUserHandler(service.NewUserService(repository.NewUserRepository(db)))
	// log.Println(data, "Userr Data")

	_, errs := h.registerUserWithSSO(data)
	if errs != nil {
		log.Println("error User SSO")
		return
	}

}

func (h *UserHandler) registerUserWithSSO(data []byte) (string, error) {

	var userData struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	err := json.Unmarshal(data, &userData)
	if err != nil {
		log.Println("error unmarshaling user data")
		return "", err
	}

	parts := strings.Split(userData.Email, "@")
	if len(parts) > 1 {
		userData.Name = parts[0]
	}

	pass := generateRandomPassword(12)

	regUser := &model.UserRegister{} // getting user data

	regUser.Email = userData.Email
	regUser.Name = userData.Name
	regUser.Password = pass

	if errs := h.userService.RegisterUserWithSSO(regUser, pass); errs != nil {
		log.Println("can not register user")
		return "", err
	}

	if errs := jobs.SendEmailRegistrationSSO(regUser.Email, regUser.Name, regUser.Password); errs != nil {
		return "", err
	}

	return regUser.Name, nil
}

func getUserData(state, code string) ([]byte, error) {
	if state != RandomStr {
		log.Println("error state")
		return nil, errors.New("invalid state")
	}

	token, err := sso_golang.Exchange(context.Background(), code)
	if err != nil {
		log.Println("error token")
		return nil, err
	}

	req, err := http.NewRequest("GET", "https://openidconnect.googleapis.com/v1/userinfo", nil)
	if err != nil {
		log.Println("error req")
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error response")
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("error body", resp.StatusCode)
		return nil, errors.New("error fetching user data")
	}

	//Structs
	var userData struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	err = json.NewDecoder(resp.Body).Decode(&userData)
	if err != nil {
		log.Println("error decode")
		return nil, err
	}
	//

	return json.Marshal(userData)
}

func generateRandomPassword(length int) string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}
