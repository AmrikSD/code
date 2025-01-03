package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)


type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}


func listAllFiles() {
	// Get the RUNFILES_DIR environment variable, which points to the runfiles directory
	runfilesDir := os.Getenv("RUNFILES_DIR")
	if runfilesDir == "" {
		fmt.Println("RUNFILES_DIR not set. This should be set by Bazel during runtime.")
		return
	}

	// Define the relative path to the directory inside runfiles
	publicDir := filepath.Join(runfilesDir, "./")

	// Walk through the public directory and list all files
	err := filepath.Walk(publicDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only list files, not directories
		if !info.IsDir() {
			// Print the relative file path inside the public directory
			relativePath := strings.TrimPrefix(path, runfilesDir+"/")
			fmt.Println(relativePath)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error listing files: %v\n", err)
	}
}

func main(){

    listAllFiles()


    e := echo.New()
    e.Renderer = &Template{
        templates: template.Must(template.ParseGlob("public/views/*.html")),
    }

    e.GET("/", func(c echo.Context) error {
        return c.Render(http.StatusOK, "index", nil)
    })

    log.Fatal(e.Start(":8080"))
}
