package httpUtil

import (
	"github.com/pkg/errors"
	"io"
	"net/http"
)

type Header struct {
	Key   string
	Value string
}

func Request(method, url string, reqParam io.Reader, header []Header) (io.Reader, error) {
	req, err := http.NewRequest(method, url, reqParam)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for _, v := range header {
		req.Header.Add(v.Key, v.Value)
	}

	var (
		client = &http.Client{}
	)

	res, err := client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer res.Body.Close()

	return res.Body, nil
}
