package oauthService

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
)

var (
	GoogleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost" + os.Getenv("SERVER_PORT") + "/token/google",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	FacebookOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost" + os.Getenv("SERVER_PORT") + "/token/facebook",
		ClientID:     os.Getenv("FACEBOOK_CLIENT_ID"),
		ClientSecret: os.Getenv("FACEBOOK_CLIENT_SECRET"),
		Scopes:       []string{"email", "user_friends"},
		Endpoint:     facebook.Endpoint,
	}
	State = "rando"
)

func GenerateToken(Email string) (t string) {
	// Create token
	JWTToken := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := JWTToken.Claims.(jwt.MapClaims)
	claims["name"] = Email
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token
	t, err := JWTToken.SignedString([]byte("secret"))
	if err != nil {
		return
	}
	return
}
