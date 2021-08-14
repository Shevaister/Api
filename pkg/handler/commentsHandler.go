package handler

import (
	"Api/pkg/models/comments"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// @Summary List comments
// @Description Get all comments
// @Tags comments
// @Accept json
// @Produce json
// @Produce xml
// @Success 200 {array} comments.Comments
// @Router /comments [get]
func ReturnAllComments(c echo.Context) error {
	result := comments.GetAllComments()
	Accept := c.Request().Header.Get("Accept")
	if Accept == "" || Accept == "application/json" {
		return c.JSON(http.StatusOK, result)
	}
	return c.XML(http.StatusOK, result)
}

// @Summary Comment by id
// @Description Get comment by id
// @Tags comments
// @Accept json
// @Produce json
// @Produce xml
// @Param id path int true "ID of comment to return"
// @Success 200 {object} comments.Comments
// @Failure 400 "Record_not_found"
// @Router /comments/{id} [get]
func ReturnComment(c echo.Context) error {
	result, err := comments.GetComment(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	Accept := c.Request().Header.Get("Accept")
	if Accept == "" || Accept == "application/json" {
		return c.JSON(http.StatusOK, result)
	}
	return c.XML(http.StatusOK, result)
}

// @Summary Create comment
// @Description Create comment
// @Tags comments
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param comment body comments.Comments true "Data for comment to create"
// @Success 200 "OK"
// @Failure 400 "Bad_request"
// @Router /restricted/comments [post]
func CreateComment(c echo.Context) error {
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
	comments.CreateComment(request)
	return c.JSON(http.StatusOK, nil)
}

// @Summary Update comment
// @Description Update comment by id
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "ID of comment to update"
// @Param comment body comments.Comments true "Data for comment to update"
// @Success 200 "OK"
// @Failure 400 "Comment_not_found"
// @Router /comments/{id} [put]
func UpdateComment(c echo.Context) error {
	request := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return err
	}
	request["id"], err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	err = comments.UpdateComment(request, c.Param("id"))
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
func DeleteComment(c echo.Context) error {
	comments.DeleteComment(c.Param("id"))
	return c.JSON(http.StatusOK, nil)
}
