package handler

import (
	"articlefe/client"
	"articlefe/entity"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EditHandler struct {
	articleClient *client.ArticleClient
}

func NewEditHandler(
	articleClient *client.ArticleClient,
) *EditHandler {
	return &EditHandler{
		articleClient: articleClient,
	}
}

func (a *EditHandler) GetEdit(c echo.Context) error {
	v, _ := strconv.Atoi(c.Param("id"))
	post, _ := a.articleClient.GetByID(c.Request().Context(), v)
	return c.Render(http.StatusOK, "edit.html", map[string]interface{}{
		"post": post,
	})
}

func (h *EditHandler) Edit(c echo.Context) error {
	v, _ := strconv.Atoi(c.FormValue("id"))
	title := c.FormValue("title")
	content := c.FormValue("content")
	category := c.FormValue("category")
	status := c.FormValue("status") // "Publish" atau "Draft"

	post := entity.Post{
		ID:       uint(v),
		Title:    title,
		Content:  content,
		Category: category,
		Status:   status,
	}

	err := h.articleClient.Edit(c.Request().Context(), post)
	if err != nil {
		return c.Render(http.StatusOK, "edit.html", map[string]interface{}{
			"error": err.Error(),
			"post":  post,
		})
	}

	return c.Render(http.StatusOK, "edit.html", map[string]interface{}{
		"success": true,
		"post":    post,
	})
}
