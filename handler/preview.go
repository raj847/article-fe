package handler

import (
	"articlefe/client"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PreviewHandler struct {
	articleClient *client.ArticleClient
}

func NewPreviewHandler(
	articleClient *client.ArticleClient,
) *PreviewHandler {
	return &PreviewHandler{
		articleClient: articleClient,
	}
}

func (a *PreviewHandler) GetPreview(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page == 0 {
		page = 1
	}
	post, _ := a.articleClient.GetPublishPagination(c.Request().Context(), page)
	return c.Render(http.StatusOK, "preview.html", map[string]interface{}{
		"posts": post,
		"next":  page + 1,
		"prev":  page - 1,
	})
}
