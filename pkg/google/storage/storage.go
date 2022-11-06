package storage

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/pkg/errors"
	"io"
)

func GetReader(ctx context.Context, bucketName, objectPath string) (*storage.Reader, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	obj := client.Bucket(bucketName).Object(objectPath)
	reader, err := obj.NewReader(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return reader, nil
}

func GetByteSlice(ctx context.Context, bucketName, objectPath string) ([]byte, error) {
	r, err := GetReader(ctx, bucketName, objectPath)
	if err != nil {
		return nil, err
	}

	res, err := io.ReadAll(r)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer r.Close()

	return res, nil
}
