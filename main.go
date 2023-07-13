package main

import (
	"errors"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"articlefe/client"
	"articlefe/handler"
)

type TemplateRegistry struct {
	templates map[string]*template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	templates := make(map[string]*template.Template)
	templates["home.html"] = template.Must(template.ParseFiles("view/home.html", "view/base.html"))
	templates["all-post.html"] = template.Must(template.ParseFiles("view/all-post.html", "view/base.html"))
	templates["preview.html"] = template.Must(template.ParseFiles("view/preview.html", "view/base.html"))
	templates["new.html"] = template.Must(template.ParseFiles("view/new.html", "view/base.html"))
	templates["edit.html"] = template.Must(template.ParseFiles("view/edit.html", "view/base.html"))
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}
	httpClient := http.Client{Timeout: 5 * time.Second}
	articleClient := client.NewArticleClient(&httpClient, "http://localhost:1323")
	allposthandler := handler.NewAllPostHandler(articleClient)
	previewhandler := handler.NewPreviewHandler(articleClient)
	addNewHandler := handler.NewAddNewHandler(articleClient)
	editHandler := handler.NewEditHandler(articleClient)

	// Route => handler
	e.GET("/", handler.Home)
	e.GET("/all-post", allposthandler.GetAllPost)
	e.GET("/trash/:id", allposthandler.UpdateTrash)
	e.GET("/preview", previewhandler.GetPreview)
	e.GET("/new", addNewHandler.GetNew)
	e.POST("/new", addNewHandler.AddNew)
	e.GET("/edit/:id", editHandler.GetEdit)
	e.POST("/edit", editHandler.Edit)
	// Start the Echo server
	e.Logger.Fatal(e.Start(":80"))
}
