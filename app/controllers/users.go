package controllers

import (
	"net/http"

	"github.com/mblode/gospel/app/models"
	"github.com/mblode/gospel/db"

	"github.com/labstack/echo"
)

func GetAllUsers(c echo.Context) error {
	db := db.GetDb()
	users := []models.User{}
	db.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	db := db.GetDb()
	id := c.Param("id")
	user := []models.User{}
	db.Where("id=?", id).First(&user)
	return c.JSON(http.StatusOK, user)
}

func NewUser(c echo.Context) error {
	db := db.GetDb()
	name := c.QueryParam("name")
	email := c.QueryParam("email")
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	user := models.User{Name: name, Email: email, Username: username, Password: password}
	db.Create(&user)
	return c.String(http.StatusOK, name+" user successfully created")
}

func DeleteUser(c echo.Context) error {
	db := db.GetDb()
	user := []models.User{}

	id := c.Param("id")
	db.Where("id = ?", id).Delete(&user)

	return c.String(http.StatusOK, id+" user successfully deleted")
}

func UpdateUser(c echo.Context) error {
	db := db.GetDb()
	user := []models.User{}

	id := c.Param("id")
	email := c.QueryParam("email")
	name := c.QueryParam("name")
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	db.Model(&user).Where("id=?", id).Updates(map[string]interface{}{"email": email, "name": name, "username": username, "password": password})

	return c.String(http.StatusOK, id+" user successfully updated")
}
