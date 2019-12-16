package routes

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/mblode/gospel/app/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func ParseTemplates(rootDir string) (*template.Template, error) {
	cleanRoot := filepath.Clean(rootDir)
	pfx := len(cleanRoot) + 1
	root := template.New("")

	err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".html") {
			if e1 != nil {
				return e1
			}

			b, e2 := ioutil.ReadFile(path)
			if e2 != nil {
				return e2
			}

			name := path[pfx:]
			t := root.New(name)
			t, e2 = t.Parse(string(b))
			if e2 != nil {
				return e2
			}
		}

		return nil
	})

	return root, err
}

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	parsed, err := ParseTemplates("app/views")

	if err != nil {
		fmt.Println(err)
	}

	// Templates
	t := &Template{
		templates: parsed,
	}

	e.Renderer = t

	e.GET("/", controllers.Home)

	e.GET("/users", controllers.GetAllUsers)
	e.GET("/users/:id", controllers.GetUser)
	e.POST("/users", controllers.NewUser)
	e.DELETE("/users/:id", controllers.DeleteUser)
	e.PUT("/users/:id", controllers.UpdateUser)

	e.GET("/locations", controllers.GetAllLocations)
	e.GET("/locations/:id", controllers.GetLocation)
	e.POST("/locations", controllers.NewLocation)
	e.DELETE("/locations/:id", controllers.DeleteLocation)
	e.PUT("/locations/:id", controllers.UpdateUser)

	e.GET("/distance", controllers.GetLocationDistance)

	return e
}
