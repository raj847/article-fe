package handler

import (
	"articlefe/client"
	"articlefe/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AddNewHandler struct {
	articleClient *client.ArticleClient
}

func NewAddNewHandler(
	articleClient *client.ArticleClient,
) *AddNewHandler {
	return &AddNewHandler{
		articleClient: articleClient,
	}
}

func (a *AddNewHandler) GetNew(c echo.Context) error {
	return c.Render(http.StatusOK, "new.html", map[string]interface{}{})
}

func (h *AddNewHandler) AddNew(c echo.Context) error {
	title := c.FormValue("title")
	content := c.FormValue("content")
	category := c.FormValue("category")
	status := c.FormValue("status") // "Publish" atau "Draft"

	post := entity.Post{
		Title:    title,
		Content:  content,
		Category: category,
		Status:   status,
	}

	err := h.articleClient.AddNew(c.Request().Context(), post)
	if err != nil {
		return c.Render(http.StatusOK, "new.html", map[string]interface{}{
			"error": err.Error(),
		})
	}

	// Redirect atau berikan respons sukses
	return c.Render(http.StatusOK, "new.html", map[string]interface{}{
		"success": true,
	})
}
