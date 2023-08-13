package mail

import (
	"bytes"
	"context"
	"encoding/base64"
	"github.com/pkg/errors"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"io"
	"os"
	"strings"
)

type SendMailData struct {
	Ctx     context.Context
	To      string
	Subject string
	Body    string
}

func SendMail(m SendMailData) error {
	credentials := "credentials.json"
	client, err := GetClientAtGCS(m.Ctx, credentials, os.Getenv("BUCKET_NAME"))
	if err != nil {
		return errors.WithStack(err)
	}

	srv, err := gmail.NewService(m.Ctx, option.WithHTTPClient(client))
	if err != nil {
		return errors.WithStack(err)
	}

	msgStr := "From: 'me'\r\n" +
		"reply-to: " + m.To + "\r\n" + //送信元
		"To: " + m.To + "\r\n" + //送信先
		"Subject:" + m.Subject + "\r\n" +
		"\r\n" + m.Body
	reader := strings.NewReader(msgStr)
	transformer := japanese.ISO2022JP.NewEncoder()
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, transform.NewReader(reader, transformer))
	if err != nil {
		return errors.WithStack(err)
	}
	msgISO2022JP := buf.Bytes()

	message := gmail.Message{}
	message.Raw = base64.StdEncoding.EncodeToString(msgISO2022JP)
	_, err = srv.Users.Messages.Send("me", &message).Do()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
