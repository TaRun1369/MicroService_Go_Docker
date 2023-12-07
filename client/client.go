package client

import "context"

type Client struct {
	endpoint string
}

func NewClient(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}

}

func (c *Client) FetchPrice(ctx context.Context,ticker string)
