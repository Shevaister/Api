package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	_ "myapp4/cmd/main/docs"
	"myapp4/pkg/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"golang.org/x/oauth2"

	"strconv"

	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8000/token/google",
		ClientID:     "824369757673-fncequl3scdmgf3i77ocrv8ssg0jf159.apps.googleusercontent.com",
		ClientSecret: "PtctplhvPhI0nlfn6Gi6gvqJ",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	facebookOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8000/token/facebook",
		ClientID:     "142940704631263",
		ClientSecret: "9e4faec575849aa5ed0eabd7f9bef712",
		Scopes:       []string{"email", "user_friends"},
		Endpoint:     facebook.Endpoint,
	}
	state = "rando"
)

// @title Echo Swagger API
// @version 1.0
// @description This is an echo swagger API server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /
// @schemes http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	handleRequests()
}

func homePage(c echo.Context) error {
	return c.String(http.StatusOK, "HomePage: localhost:8000/swagger/index.html")
}

func handleRequests() {
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", homePage)
	e.GET("/comments", returnAllComments)
	e.GET("/comments/:id", returnComment)
	e.DELETE("/comments/:id", deleteComment)
	e.PUT("/comments/:id", updateComment)
	e.GET("/posts", returnAllPosts)
	e.GET("/posts/:id", returnPost)
	e.DELETE("/posts/:id", deletePost)
	e.PUT("/posts/:id", updatePost)
	e.GET("/signin/google", googleSignIn)
	e.GET("/signin/facebook", facebookSignIn)
	e.GET("/token/:source", getAuthToken)
	e.POST("/signup", signUp)
	e.POST("/signin", signIn)

	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.POST("/comments", createComment)
	r.POST("/posts", createPost)
	e.Logger.Fatal(e.Start(":8000"))
}

// @Summary List comments
// @Description Get all comments
// @Tags comments
// @Accept json
// @Produce json
// @Produce xml
// @Success 200 {array} models.Comments
// @Router /comments [get]
func returnAllComments(c echo.Context) error {
	result := models.GetAllComments()
	Accept := c.Request().Header.Get("Accept")
	if Accept == "" || Accept == "application/json" {
		return c.JSONPretty(http.StatusOK, result, "    ")
	}
	return c.XMLPretty(http.StatusOK, result, "    ")
}

// @Summary Comment by id
// @Description Get comment by id
// @Tags comments
// @Accept json
// @Produce json
// @Produce xml
// @Param id path int true "ID of comment to return"
// @Success 200 {object} models.Comments
// @Failure 400 "Record_not_found"
// @Router /comments/{id} [get]
func returnComment(c echo.Context) error {
	result, err := models.GetComment(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	Accept := c.Request().Header.Get("Accept")
	if Accept == "" || Accept == "application/json" {
		return c.JSONPretty(http.StatusOK, result, "    ")
	}
	return c.XMLPretty(http.StatusOK, result, "    ")
}

// @Summary Create comment
// @Description Create comment
// @Tags comments
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param comment body models.Comments true "Data for comment to create"
// @Success 200 "OK"
// @Failure 400 "Bad_request"
// @Router /restricted/comments [post]
func createComment(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["name"].(string)
	request := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	request["email"] = email
	request["id"] = 0
	models.CreateComment(request)
	return c.JSON(http.StatusOK, nil)
}

// @Summary Update comment
// @Description Update comment by id
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "ID of comment to update"
// @Param comment body models.Comments true "Data for comment to update"
// @Success 200 "OK"
// @Failure 400 "Comment_not_found"
// @Router /comments/{id} [put]
func updateComment(c echo.Context) error {
	request := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	request["id"], err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	err = models.UpdateComment(request, c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, nil)
}

// @Summary Delete comment
// @Description Delete comment by id
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "ID of comment to delete"
// @Success 200 "OK"
// @Router /comments/{id} [delete]
func deleteComment(c echo.Context) error {
	models.DeleteComment(c.Param("id"))
	return c.JSON(http.StatusOK, nil)
}

// @Summary List posts
// @Description Get all posts
// @Tags posts
// @Accept json
// @Produce json
// @Produce xml
// @Success 200 {array} models.Posts
// @Router /posts [get]
func returnAllPosts(c echo.Context) error {
	result := models.GetAllPosts()
	Accept := c.Request().Header.Get("Accept")
	if Accept == "" || Accept == "application/json" {
		return c.JSONPretty(http.StatusOK, result, "    ")
	}
	return c.XMLPretty(http.StatusOK, result, "    ")
}

// @Summary Get post
// @Description Get post by id
// @Tags posts
// @Accept json
// @Produce json
// @Produce xml
// @Param id path int true "ID of post to return"
// @Success 200 {object} models.Posts
// @Failure 400 "Post_not_found"
// @Router /posts/{id} [get]
func returnPost(c echo.Context) error {
	result, err := models.GetPost(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	Accept := c.Request().Header.Get("Accept")
	if Accept == "" || Accept == "application/json" {
		return c.JSONPretty(http.StatusOK, result, "    ")
	}
	return c.XMLPretty(http.StatusOK, result, "    ")
}

// @Summary Create post
// @Description Create post
// @Tags posts
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 "OK"
// @Failure 400 "Bad_request"
// @Param comment body models.Posts true "Data for post to create"
// @Router /restricted/posts [post]
func createPost(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["name"].(string)
	request := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	request["UserID"] = (models.GetUser(email)).ID
	request["id"] = 0
	models.CreatePost(request)
	return c.JSON(http.StatusOK, nil)
}

// @Summary Update post
// @Description Update post by id
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "ID of post to update"
// @Param comment body models.Posts true "Data for post to update"
// @Success 200 "OK"
// @Failure 400 "Comment_not_found"
// @Router /posts/{id} [put]
func updatePost(c echo.Context) error {
	request := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	request["id"], err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	err = models.UpdatePost(request, c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, nil)
}

// @Summary Delete post
// @Description Delete post by id
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "ID of post to delete"
// @Success 200 "OK"
// @Router /posts/{id} [delete]
func deletePost(c echo.Context) error {
	models.DeletePost(c.Param("id"))
	return c.JSON(http.StatusOK, nil)
}

// @Summary Register
// @Description Register an account with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.Users true "Data for user to create"
// @Success 200 "OK"
// @Failure 400 "Account_with_this_email_is_already_registred"
// @Router /signup [post]
func signUp(c echo.Context) error {
	request := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	success := models.Register(request["email"].(string), request["password"].(string))
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
// @Param user body models.Users true "Data for user to login"
// @Success 200 "OK"
// @Failure 400 "Wrong_login_info"
// @Router /signin [post]
func signIn(c echo.Context) error {
	request := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	err = models.Login(request["email"].(string), request["password"].(string))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	t := generateToken(request["email"].(string))
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
func googleSignIn(c echo.Context) error {
	result := googleOauthConfig.AuthCodeURL(state)
	return c.JSON(http.StatusOK, result)
}

// @Summary Sign in throгgh facebook
// @Description Get facebook authorization link
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /signin/facebook [get]
func facebookSignIn(c echo.Context) error {
	result := facebookOauthConfig.AuthCodeURL(state)
	return c.JSON(http.StatusOK, result)
}

func getAuthToken(c echo.Context) error {
	if c.FormValue("state") != state {
		return c.String(http.StatusOK, "state is not valid")
	}

	var (
		response *http.Response
		token    *oauth2.Token
		err      error
	)
	source := c.Param("source")

	if source == "google" {
		token, err = googleOauthConfig.Exchange(context.Background(), c.FormValue("code"))
	} else {
		token, err = facebookOauthConfig.Exchange(context.Background(), c.FormValue("code"))
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
	user := models.GetUser((data["email"]).(string))

	t := generateToken(user.Email)
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func generateToken(Email string) (t string) {
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
