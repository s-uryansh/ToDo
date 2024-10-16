package utils

const (
	SessionDuration                   = 1200
	SECRET_KEY_TOKEN                  = //Secret Key
	REDIS_KEY_TOKEN                   = //Redis KEY
	ZipLength                         = 6
	FROM_MAIL                         = //Mail to send registration mail
	PASSWORD_MAIL                     = //Password for mail to forward
	SMTP_ADDRESS                      = "smtp.gmail.com"
	SMTP_DIAL                         = "smtp.gmail.com:587"
	REGISTER_TemplatePath             = "../../template/registration_email.html"
	REGISTER_TemplatePath_SSO         = "../../template/register_sso.html"
	REGISTER_Success_TemplatePath_SSO = "C:\\Users\\surya\\OneDrive\\Desktop\\Files\\Codes\\Go\\Projects\\CRUD-SQL\\template"
	CLIENT_ID                         = //Put Client ID here
	CLIENT_SECRET                     = //Client Secret here
	REDIRECT_URI                      = "http://localhost:8080/callback"
	//Email Subjects and bodies

	//After Registration
	REGISTER_Subject = "Welcome to To-Do list by Suryansh Rohil!"
)

var CUSTOMER_LOGGED string
