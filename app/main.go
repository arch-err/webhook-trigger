package main

import (
	"crypto/tls"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
)

const VERSION = "0.1.0"
const PORT = 8080

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{templates: template.Must(template.ParseGlob("views/*.html"))}
}

type Page struct {
	Title   string
	Version string
}

func newPage() Page {
	return Page{
		Title:   "Webhook Trigger",
		Version: VERSION,
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = newTemplate()
	e.Static("/css", "static/css")

	page := newPage()

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", page)
	})

	e.GET("/button1", func(c echo.Context) error {
		url := ""

		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		response, err := http.Get(url)

		log.Println(response)
		if err != nil || response.StatusCode != 200 {
			log.Println(err)
			return c.Render(http.StatusOK, "fail", page)
		}
		return c.Render(http.StatusOK, "success", page)
	})

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(PORT)))
}
