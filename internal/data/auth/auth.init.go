package auth

import (
	"skeleton/pkg/httpclient"
)

// Data ...
type Data struct {
	client  *httpclient.Client
	baseURL string
}

// New ...
func New(client *httpclient.Client, baseURL string) Data {
	d := Data{
		client:  client,
		baseURL: baseURL,
	}
	return d
}
