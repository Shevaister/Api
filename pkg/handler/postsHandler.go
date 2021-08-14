package handler

import (
	"Api/pkg/models/posts"
	"Api/pkg/models/users"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// @Summary List posts
// @Description Get all posts
// @Tags posts
// @Accept json
// @Produce json
// @Produce xml
// @Success 200 {array} posts.Posts
// @Router /posts [get]
func ReturnAllPosts(c echo.Context) error {
	result := posts.GetAllPosts()
	Accept := c.Request().Header.Get("Accept")
	if Accept == "" || Accept == "application/json" {
		return c.JSON(http.StatusOK, result)
	}
	return c.XML(http.StatusOK, result)
}

// @Summary Get post
// @Description Get post by id
// @Tags posts
// @Accept json
// @Produce json
// @Produce xml
// @Param id path int true "ID of post to return"
// @Success 200 {object} posts.Posts
// @Failure 400 "Post_not_found"
// @Router /posts/{id} [get]
func ReturnPost(c echo.Context) error {
	result, err := posts.GetPost(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	Accept := c.Request().Header.Get("Accept")
	if Accept == "" || Accept == "application/json" {
		return c.JSON(http.StatusOK, result)
	}
	return c.XML(http.StatusOK, result)
}

// @Summary Create post
// @Description Create post
// @Tags posts
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 "OK"
// @Failure 400 "Bad_request"
// @Param comment body posts.Posts true "Data for post to create"
// @Router /restricted/posts [post]
func CreatePost(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["name"].(string)
	request := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	request["UserID"] = (users.GetUser(email)).ID
	request["id"] = 0
	posts.CreatePost(request)
	return c.JSON(http.StatusOK, nil)
}

// @Summary Update post
// @Description Update post by id
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "ID of post to update"
// @Param comment body posts.Posts true "Data for post to update"
// @Success 200 "OK"
// @Failure 400 "Comment_not_found"
// @Router /posts/{id} [put]
func UpdatePost(c echo.Context) error {
	request := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	request["id"], err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	err = posts.UpdatePost(request, c.Param("id"))
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
func DeletePost(c echo.Context) error {
	posts.DeletePost(c.Param("id"))
	return c.JSON(http.StatusOK, nil)
}
