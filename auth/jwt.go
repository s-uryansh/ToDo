package auth

import (
	"CRUD-SQL/model"
	"CRUD-SQL/utils"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Helper Funcs
type session struct {
	id     string
	userID string
	expiry int64
}

var Sessions = map[string]session{}

func getUniqueID() string {
	id := make([]byte, 32)
	rand.Read(id)
	return base64.URLEncoding.EncodeToString(id)
}

func createSession(userID string) (string, error) {
	sessionID := getUniqueID()
	expiry := time.Now().Add(utils.SessionDuration * time.Second).Unix()

	Sessions[sessionID] = session{
		id: sessionID, userID: userID, expiry: expiry,
	}
	return sessionID, nil
}

func GetSession(sessionID string) (session, bool) {
	sess, exists := Sessions[sessionID]
	if !exists || sess.expiry < time.Now().Unix() {
		return session{}, false
	}
	return sess, true
}

func InvalidSession(sessionID string) {
	delete(Sessions, sessionID)
}

// Token Creation
func JwtTokenCreate(c *gin.Context) error {
	var json model.UserLogin

	// Sign and get the complete encoded token as a string using the secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  json.Email,
		"expt": time.Now().Add(utils.SessionDuration * time.Second).Unix(),
	})
	tokenString, err := token.SignedString([]byte(utils.SECRET_KEY_TOKEN))
	if err != nil {
		fmt.Println("error getting token: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return err
	}
	// Session Auth
	session_id, err := createSession(tokenString)
	if err != nil {
		fmt.Println("can not create session")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		c.Abort()
		return err
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "session_id",
		Value:   session_id,
		Expires: time.Now().Add(utils.SessionDuration * time.Second),
		Path:    "/",
	})

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "tokenString",
		Value:   tokenString,
		Expires: time.Now().Add(utils.SessionDuration * time.Second),
		Path:    "/",
	})

	// c.Set("JWTtoken", token)
	c.Writer.WriteHeader(http.StatusOK)
	return nil

}

//Session Middleware

func SessionMiddleware(c *gin.Context) {
	cookie, err := c.Request.Cookie("session_id")
	if err != nil {
		fmt.Println("can not request cookie", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		c.Abort()
		return
	}

	sessionID := cookie.Value
	sess, valid := GetSession(sessionID)
	if !valid {
		fmt.Println("not Authorized")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		c.Abort()
		return
	}
	c.Request.Header.Set("user_id", sess.userID)
	c.Next()

}
