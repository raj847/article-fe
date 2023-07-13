package client

import (
	"articlefe/entity"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type ArticleClient struct {
	client *http.Client
	host   string
}

func NewArticleClient(
	client *http.Client,
	host string,
) *ArticleClient {
	return &ArticleClient{
		client: client,
		host:   host,
	}
}

func (a *ArticleClient) buildUrl(path string) string {
	return fmt.Sprintf("%s/%s", a.host, path)
}

func (a *ArticleClient) ListPostByStatus(ctx context.Context, status string) ([]entity.Post, error) {
	resp, err := a.client.Get(a.buildUrl("article?status=" + status))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	posts := []entity.Post{}
	err = json.NewDecoder(resp.Body).Decode(&posts)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (a *ArticleClient) UpdateTrash(ctx context.Context, id uint) error {
	req, err := http.NewRequestWithContext(ctx, "PATCH", a.buildUrl("article/"+strconv.Itoa(int(id))+"/trash"), nil)
	if err != nil {
		return err
	}
	_, err = a.client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func (a *ArticleClient) GetPublishPagination(ctx context.Context, page int) ([]entity.Post, error) {
	offset := 4 * (page - 1)
	resp, err := a.client.Get(a.buildUrl(fmt.Sprintf("article/%d/%d?status=publish", 4, offset)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	posts := []entity.Post{}
	err = json.NewDecoder(resp.Body).Decode(&posts)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (a *ArticleClient) AddNew(ctx context.Context, post entity.Post) error {
	jsonData, err := json.Marshal(post)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", a.buildUrl("article"), bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to add new post")
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		errors := ErrorResponse{}
		err = json.NewDecoder(resp.Body).Decode(&errors)
		if err != nil {
			return fmt.Errorf("failed to add new post")
		}
		return &errors
	}

	return nil
}

func (a *ArticleClient) Edit(ctx context.Context, post entity.Post) error {
	jsonData, err := json.Marshal(post)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", a.buildUrl(fmt.Sprintf("article/%d", post.ID)), bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to edit post")

	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		errors := ErrorResponse{}
		err = json.NewDecoder(resp.Body).Decode(&errors)
		if err != nil {
			return fmt.Errorf("failed to edit post")
		}
		return &errors
	}

	return nil
}

func (a *ArticleClient) GetByID(ctx context.Context, id int) (entity.Post, error) {
	resp, err := a.client.Get(a.buildUrl(fmt.Sprintf("article/%d", id)))
	if err != nil {
		return entity.Post{}, err
	}
	defer resp.Body.Close()
	posts := entity.Post{}
	err = json.NewDecoder(resp.Body).Decode(&posts)
	if err != nil {
		return entity.Post{}, err
	}
	return posts, nil
}
