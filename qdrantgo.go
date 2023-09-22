package qdrantgo

import (
	"context"
	"net/http"

	"github.com/henomis/restclientgo"

	"github.com/henomis/qdrant-go/request"
	"github.com/henomis/qdrant-go/response"
)

type Client struct {
	restClient *restclientgo.RestClient
	endpoint   string
	apiKey     string
}

func New(endpoint, apiKey string) *Client {

	restClient := restclientgo.New(endpoint)

	if len(apiKey) > 0 {
		restClient.SetRequestModifier(func(req *http.Request) *http.Request {
			req.Header.Set("api-key", apiKey)
			return req
		})
	}
	return &Client{
		endpoint:   endpoint,
		restClient: restClient,
		apiKey:     apiKey,
	}
}

func (c *Client) CollectionList(
	ctx context.Context,
	req *request.CollectionList,
	res *response.CollectionList,
) error {
	return c.restClient.Get(ctx, req, res)
}

func (c *Client) CollectionCreate(
	ctx context.Context,
	req *request.CollectionCreate,
	res *response.CollectionCreate,
) error {
	return c.restClient.Put(ctx, req, res)
}

func (c *Client) CollectionCollectInfo(
	ctx context.Context,
	req *request.CollectionCollectInfo,
	res *response.CollectionCollectInfo,
) error {
	return c.restClient.Get(ctx, req, res)
}

func (c *Client) CollectionDelete(
	ctx context.Context,
	req *request.CollectionDelete,
	res *response.CollectionDelete,
) error {
	return c.restClient.Delete(ctx, req, res)
}

func (c *Client) PointUpsert(
	ctx context.Context,
	req *request.PointUpsert,
	res *response.PointUpsert,
) error {
	return c.restClient.Put(ctx, req, res)
}

func (c *Client) PointSearch(
	ctx context.Context,
	req *request.PointSearch,
	res *response.PointSearch,
) error {
	return c.restClient.Post(ctx, req, res)
}

func (c *Client) PointDelete(
	ctx context.Context,
	req *request.PointDelete,
	res *response.PointDelete,
) error {
	return c.restClient.Post(ctx, req, res)
}

func (c *Client) PointGet(
	ctx context.Context,
	req *request.PointGet,
	res *response.PointGet,
) error {
	return c.restClient.Get(ctx, req, res)
}
