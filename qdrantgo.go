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

	restClient := restclientgo.New(endpoint).WithDecodeOnError(true)

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

func (c *Client) CollectionUpdate(
	ctx context.Context,
	req *request.CollectionUpdate,
	res *response.CollectionUpdate,
) error {
	return c.restClient.Patch(ctx, req, res)
}

func (c *Client) IndexCreate(
	ctx context.Context,
	req *request.IndexCreate,
	res *response.IndexCreate,
) error {
	return c.restClient.Put(ctx, req, res)
}

func (c *Client) IndexDelete(
	ctx context.Context,
	req *request.IndexDelete,
	res *response.IndexDelete,
) error {
	return c.restClient.Delete(ctx, req, res)
}

func (c *Client) PointsUpsert(
	ctx context.Context,
	req *request.PointsUpsert,
	res *response.PointsUpsert,
) error {
	return c.restClient.Put(ctx, req, res)
}

func (c *Client) PointsSearch(
	ctx context.Context,
	req *request.PointsSearch,
	res *response.PointsSearch,
) error {
	return c.restClient.Post(ctx, req, res)
}

func (c *Client) PointsDelete(
	ctx context.Context,
	req *request.PointsDelete,
	res *response.PointsDelete,
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

func (c *Client) PointsGet(
	ctx context.Context,
	req *request.PointsGet,
	res *response.PointsGet,
) error {
	return c.restClient.Post(ctx, req, res)
}

func (c *Client) VectorsUpdate(
	ctx context.Context,
	req *request.VectorsUpdate,
	res *response.VectorsUpdate,
) error {
	return c.restClient.Put(ctx, req, res)
}

func (c *Client) VectorsDelete(
	ctx context.Context,
	req *request.VectorsDelete,
	res *response.VectorsDelete,
) error {
	return c.restClient.Post(ctx, req, res)
}
