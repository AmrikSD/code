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

func processImages(in chan string, c *gosseract.Client){
    for path := range in {
        c.SetImage(path)
        _, err := c.Text()
        if err != nil {
            continue //TODO DLQ or something?
        }
    }
}

func processDirectory(startingDir string, fileChan chan File, imageChan chan string){
    filepath.Walk(startingDir, func(path string, info os.FileInfo, err error) error {
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
                imageChan <- path
            }

            f := File{
                info.Name(),
                info.Size(),
                info.ModTime(),
                strings.TrimPrefix(path, fmt.Sprintf("%s/",startingDir)),
                text,
            }
            fileChan <- f

        }
        return nil
    })
}

func main(){
    client := gosseract.NewClient()
    defer client.Close()

    imageChan := make(chan string)
    go processImages(imageChan, client)
    defer close(imageChan)

    fileChan := make(chan File)
    filesList := make([]File, 0)
    go func(){
        for file := range fileChan {
            filesList = append(filesList, file)
        }
    }()
    defer close(fileChan)

    pathToStuff := os.Args[1]

    go processDirectory(pathToStuff, fileChan, imageChan)

    e := echo.New()
    e.Static("/static", "public/static")

    e.Static("/original", pathToStuff) //TODO Make this configurable

    e.Renderer = &Template{
        templates: template.Must(template.ParseGlob("public/views/*.html")),
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
