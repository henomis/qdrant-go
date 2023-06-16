package qdrantgo

import (
	"context"
	"net/http"

	collectionrequest "github.com/henomis/qdrant-go/request/collection"
	collectionresponse "github.com/henomis/qdrant-go/response/collection"
	"github.com/henomis/restclientgo"
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

func (c *Client) CollectionList(ctx context.Context, req *collectionrequest.List, res *collectionresponse.List) error {
	return c.restClient.Get(ctx, req, res)
}
