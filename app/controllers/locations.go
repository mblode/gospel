package controllers

import (
	"fmt"
	"net/http"

	"github.com/mblode/gospel/app/models"
	"github.com/mblode/gospel/db"

	"github.com/labstack/echo"
)

func GetAllLocations(c echo.Context) error {
	db := db.GetDb()
	locations := []models.Location{}
	db.Find(&locations)
	// return c.JSON(http.StatusOK, locations)
	return c.Render(http.StatusOK, "locations/index.html", "Beans")
}

func GetLocation(c echo.Context) error {
	db := db.GetDb()
	id := c.Param("id")
	location := []models.Location{}
	db.Where("id=?", id).First(&location)
	// return c.JSON(http.StatusOK, location)
	return c.Render(http.StatusOK, "locations/show.html", "Wow")
}

func NewLocation(c echo.Context) error {
	db := db.GetDb()
	slug := c.QueryParam("slug")
	title := c.QueryParam("title")
	description := c.QueryParam("description")
	body := c.QueryParam("body")

	location := models.Location{Slug: slug, Title: title, Description: description, Body: body}
	db.Create(&location)
	return c.String(http.StatusOK, title+" location successfully created")
}

func DeleteLocation(c echo.Context) error {
	db := db.GetDb()
	location := []models.Location{}

	id := c.Param("id")
	db.Where("id = ?", id).Delete(&location)

	return c.String(http.StatusOK, id+" location successfully deleted")
}

func UpdateLocation(c echo.Context) error {
	db := db.GetDb()
	location := []models.Location{}

	id := c.Param("id")
	email := c.QueryParam("email")
	name := c.QueryParam("name")
	locationname := c.QueryParam("locationname")
	password := c.QueryParam("password")

	db.Model(&location).Where("id=?", id).Updates(map[string]interface{}{"email": email, "name": name, "locationname": locationname, "password": password})

	return c.String(http.StatusOK, id+" location successfully updated")
}

func GetLocationDistance(c echo.Context) error {
	db := db.GetDb()
	locations := []models.Location{}

	lat := c.Param("lat")
	lng := c.QueryParam("lng")
	dst := c.QueryParam("dst")

	sql := fmt.Sprintf("SELECT id, (6371 * acos(cos(radians(%v)) * cos(radians(lat)) * cos(radians(lng) - radians(%v)) + sin(radians(%v)) * sin(radians(lat)))) AS distance FROM shops WHERE active=1 HAVING distance < %v ORDER BY distance LIMIT 0, 20;", lat, lng, lat, dst)

	rows, err := db.Raw(sql).Rows()

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&locations)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"locations": &locations,
	})
}
