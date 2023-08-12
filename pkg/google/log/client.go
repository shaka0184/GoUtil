package log

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/oauth2/google"
)

type Client struct {
	ProjectID string
	LogName   string
}

func NewClient(ctx context.Context) (*Client, error) {
	cred, err := google.FindDefaultCredentials(ctx, "https://www.googleapis.com/oauth2/v4/token")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Client{ProjectID: cred.ProjectID}, nil
}

func (c *Client) Info(ctx context.Context, m string) {
	Info(ctx, c.ProjectID, c.LogName, m)
}

func (c *Client) Error(ctx context.Context, err error) {
	Error(ctx, c.ProjectID, c.LogName, err)
}
