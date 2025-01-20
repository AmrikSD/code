package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)


type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main(){

    e := echo.New()
    e.Static("/static", "public/static")
    e.Renderer = &Template{
        templates: template.Must(template.ParseGlob("public/views/*.html")),
    }

    e.GET("/", func(c echo.Context) error {
        return c.Render(http.StatusOK, "index", nil)
    })

    filesList := []string{"hil", "ho"}

    e.GET("/files", func(c echo.Context) error {
        return c.Render(http.StatusOK, "files", filesList)
    })


    log.Fatal(e.Start(":8080"))
}
