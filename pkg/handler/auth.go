package handler

import (
	"Api/pkg/models/users"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"Api/pkg/oauthService"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

// @Summary Register
// @Description Register an account with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body users.Users true "Data for user to create"
// @Success 200 "OK"
// @Failure 400 "Account_with_this_email_is_already_registred"
// @Router /signup [post]
func SignUp(c echo.Context) error {
	request := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	success := users.Register(request["email"].(string), request["password"].(string))
	if !success {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, nil)
}

// @Summary Sign in
// @Description Sign in with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body users.Users true "Data for user to login"
// @Success 200 "OK"
// @Failure 400 "Wrong_login_info"
// @Router /signin [post]
func SignIn(c echo.Context) error {
	request := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	err = users.Login(request["email"].(string), request["password"].(string))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	t := oauthService.GenerateToken(request["email"].(string))
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

// @Summary Sign in throгgh google
// @Description Get google authorization link
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /signin/google [get]
func GoogleSignIn(c echo.Context) error {
	result := oauthService.GoogleOauthConfig.AuthCodeURL(oauthService.State)
	return c.JSON(http.StatusOK, result)
}

// @Summary Sign in throгgh facebook
// @Description Get facebook authorization link
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /signin/facebook [get]
func FacebookSignIn(c echo.Context) error {
	result := oauthService.FacebookOauthConfig.AuthCodeURL(oauthService.State)
	return c.JSON(http.StatusOK, result)
}

func GetAuthToken(c echo.Context) error {
	if c.FormValue("state") != oauthService.State {
		return c.String(http.StatusOK, "state is not valid")
	}

	var (
		response *http.Response
		token    *oauth2.Token
		err      error
	)
	source := c.Param("source")

	if source == "google" {
		token, err = oauthService.GoogleOauthConfig.Exchange(context.Background(), c.FormValue("code"))
	} else {
		token, err = oauthService.FacebookOauthConfig.Exchange(context.Background(), c.FormValue("code"))
	}
	if err != nil {
		return c.String(http.StatusOK, "could not get the token")
	}

	if source == "google" {
		response, err = http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	} else {
		response, err = http.Get("https://graph.facebook.com/me?fields=email&access_token=" + token.AccessToken)
	}
	if err != nil {
		return c.String(http.StatusOK, "failed getting user info")
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return c.String(http.StatusOK, "failed reading response body")
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(contents, &data)
	if err != nil {
		return c.String(http.StatusOK, "failed unmarshalng response body")
	}
	user := users.GetUser((data["email"]).(string))

	t := oauthService.GenerateToken(user.Email)
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
