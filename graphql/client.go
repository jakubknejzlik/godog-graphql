package graphql

import (
	"context"

	"github.com/machinebox/graphql"
)

// Client ...
type Client struct {
	c *graphql.Client
}

var client *Client

// NewClient ...
func NewClient(URL string) (*Client, error) {
	if client == nil {

		c := graphql.NewClient(URL)
		// c.Log = func(s string) {
		// 	log.Println(s)
		// }
		client = &Client{c}
	}
	return client, nil
}

// SendQuery ...
func (c *Client) SendQuery(ctx context.Context, query string, variables map[string]interface{}, headers map[string]string, data interface{}) error {
	req := graphql.NewRequest(query)
	for key, value := range variables {
		req.Var(key, value)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return c.c.Run(ctx, req, data)
}
