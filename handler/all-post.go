package handler

import (
	"articlefe/client"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AllPostHandler struct {
	articleClient *client.ArticleClient
}

func NewAllPostHandler(
	articleClient *client.ArticleClient,
) *AllPostHandler {
	return &AllPostHandler{
		articleClient: articleClient,
	}
}

func (a *AllPostHandler) GetAllPost(c echo.Context) error {
	status := c.QueryParam("status")
	if status == "" {
		status = "publish"
	}
	posts, _ := a.articleClient.ListPostByStatus(c.Request().Context(), status)
	return c.Render(http.StatusOK, "all-post.html", map[string]interface{}{
		"posts": posts,
	})
}
func (a *AllPostHandler) UpdateTrash(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	_ = a.articleClient.UpdateTrash(c.Request().Context(), uint(id))
	referer := (c.Request().Header.Get("Referer"))
	return c.Redirect(http.StatusFound, referer)
}
