package storage

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/pkg/errors"
)

func Get(ctx context.Context, bucketName, objectPath string) (*storage.Reader, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	obj := client.Bucket(bucketName).Object(objectPath)
	reader, err := obj.NewReader(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer reader.Close()

	return reader, nil
}
