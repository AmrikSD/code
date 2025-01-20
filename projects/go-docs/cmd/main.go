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
	"time"

	"github.com/labstack/echo/v4"
	"github.com/otiai10/gosseract/v2"
)


type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type File struct {
    Name string
    Size int64
    Edited time.Time
    Path string
    Text string
}

// image formats and magic numbers
var magicTable = map[string]string{
    "\xff\xd8\xff":      "image/jpeg",
    "\x89PNG\r\n\x1a\n": "image/png",
    "GIF87a":            "image/gif",
    "GIF89a":            "image/gif",
}

func isImage(data []byte) bool {
    dataStr := string(data)
    for magic := range magicTable {
        if strings.HasPrefix(dataStr, magic) {
            return true
        }
    }
    return false
}


func main(){
    client := gosseract.NewClient()
    defer client.Close()


    e := echo.New()
    e.Static("/static", "public/static")

    e.Static("/original", "input") //TODO Make this configurable

    e.Renderer = &Template{
        templates: template.Must(template.ParseGlob("public/views/*.html")),
    }

    filesList := make([]File, 0)
    err := filepath.Walk("./input", func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if info.Mode().IsRegular() {
            text := ""

            file, err := os.ReadFile(path)
            if err != nil {
                panic(err)
            }

            if isImage(file){
                client.SetImage(path)
                text, _ = client.Text()
            } else {
                fmt.Printf("skipping doing OCR for %s\n", path)
            }


            filesList = append(filesList, File{
                info.Name(),
                info.Size(),
                info.ModTime(),
                strings.TrimPrefix(path, "input/"),
                text,
            })

        }
        return nil
    })

    if err != nil {
        panic(err)
    }

    e.GET("/", func(c echo.Context) error {
        return c.Render(http.StatusOK, "index", nil)
    })


    e.GET("/files", func(c echo.Context) error {
        return c.Render(http.StatusOK, "files", filesList)
    })

    e.GET("/modal/:file", func(c echo.Context) error {
        f := c.Param("file")
        return c.Render(http.StatusOK, "modal", f)
    })


    log.Fatal(e.Start(":8080"))
}
