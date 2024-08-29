package controllers

import (
	"CRUD-SQL/utils"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var sso_golang *oauth2.Config
var RandomStr = "inasdj"

func init() {
	sso_golang = &oauth2.Config{
		RedirectURL:  utils.REDIRECT_URI,
		ClientID:     utils.CLIENT_ID,
		ClientSecret: utils.CLIENT_SECRET,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func SSO_signin(w http.ResponseWriter, r *http.Request) {
	url := sso_golang.AuthCodeURL(RandomStr)
	// fmt.Println(url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
