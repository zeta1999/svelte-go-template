package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	GOOGLE_Client_ID     string
	GOOGLE_Client_Secret string
	GOOGLEAuthConfig     = &oauth2.Config{}
	State                string
	BACKEND_URL          string
	FRONTEND_URL         string
	JWT_KEY              []byte
	SQL_DB_USER          string
	SQL_DB_PASS          string
	SQL_DB_DB            string
	SQL_DB_HOST          string
	SQL_DB_PORT          string
	PGSQL_CS             string
	PORT                 string
)

var Log = logrus.New()

func LoadConfig() {

	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.Out = file
	} else {
		Log.Fatal("Failed to log to file, using default stderr")
	}

	err = godotenv.Load()
	if err != nil {
		Log.WithFields(logrus.Fields{
			"fn":  "LoadConfig",
			"err": err.Error(),
		}).Fatal("unable to load configs")
	}

	GOOGLE_Client_ID = os.Getenv("GOOGLE_Client_ID")
	GOOGLE_Client_Secret = os.Getenv("GOOGLE_Client_Secret")
	BACKEND_URL = os.Getenv("BACKEND_URL")
	FRONTEND_URL = os.Getenv("FRONTEND_URL")
	JWT_KEY = []byte(os.Getenv("JWT_KEY"))
	SQL_DB_USER = os.Getenv("SQL_DB_USER")
	SQL_DB_PASS = os.Getenv("SQL_DB_PASS")
	SQL_DB_DB = os.Getenv("SQL_DB_DB")
	SQL_DB_HOST = os.Getenv("SQL_DB_HOST")
	SQL_DB_PORT = os.Getenv("SQL_DB_PORT")
	State = os.Getenv("State")
	PORT = ":" + os.Getenv("PORT")

	GOOGLEAuthConfig.RedirectURL = fmt.Sprintf("%s/callback", BACKEND_URL)
	GOOGLEAuthConfig.ClientID = GOOGLE_Client_ID
	GOOGLEAuthConfig.ClientSecret = GOOGLE_Client_Secret
	GOOGLEAuthConfig.Scopes = []string{"https://www.googleapis.com/auth/userinfo.profile"}
	GOOGLEAuthConfig.Endpoint = google.Endpoint
	PGSQL_CS = fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", SQL_DB_USER, SQL_DB_DB, SQL_DB_PASS, SQL_DB_HOST, SQL_DB_PORT)
}
