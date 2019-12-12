package controllers

import (
	"github.com/mblode/gospel/app/models"
	"github.com/mblode/gospel/db"
	"net/http"

	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) error {
	db := db.DbManager()
	users := []models.User{}
	db.Find(&users)
	// spew.Dump(json.Marshal(users))
	// return c.JSON(http.StatusOK, users)
	return c.JSON(http.StatusOK, users)
}
