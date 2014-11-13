package brato

import "net/http"

var DefaultURL = "https://metrics-api.librato.com/v1"

type Client struct {
	*http.Client
}

func (c *Client) NewRequest(method, path string) (*http.Request, error) {
	return http.NewRequest(method, path, nil)
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.client().Do(req)
}

func (c *Client) client() *http.Client {
	if c.Client == nil {
		return http.DefaultClient
	}

	return c.Client
}
