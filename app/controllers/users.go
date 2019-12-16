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

	// return c.JSON(http.StatusOK, users)
	return c.Render(http.StatusOK, "users/index.html", users)
}

func GetUser(c echo.Context) error {
	db := db.GetDb()
	id := c.Param("id")
	user := []models.User{}
	db.Where("id=?", id).First(&user)
	// return c.JSON(http.StatusOK, user)
	return c.Render(http.StatusOK, "users/show.html", map[string]interface{}{
		"User":  user,
		"Title": "User",
	})
}

func NewUser(c echo.Context) error {
	db := db.GetDb()
	name := c.FormValue("name")
	email := c.FormValue("email")
	username := c.FormValue("username")
	password := c.FormValue("password")

	user := models.User{Name: name, Email: email, Username: username, Password: password}
	db.Create(&user)
	return c.String(http.StatusOK, name+" user successfully created")
}

func DeleteUser(c echo.Context) error {
	db := db.GetDb()
	user := []models.User{}

	id := c.Param("id")
	db.Where("id = ?", id).Delete(&user)

	return c.NoContent(http.StatusNoContent)
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

	return c.JSON(http.StatusOK, user)
}
