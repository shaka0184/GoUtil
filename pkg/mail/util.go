package mail

import (
	"fmt"
	"net/smtp"
	"os"
	"strconv"
	"strings"
)

var (
	hostname string
	port     int
	username string
	password string
	from     string
)

type Mail struct {
	Recipients []string `json:"recipients"`
	Subject    string   `json:"subject"`
	Body       string   `json:"body"`
}

func init() {
	hostname = os.Getenv("HOST_NAME")
	port, _ = strconv.Atoi(os.Getenv("PORT"))
	username = os.Getenv("USER_NAME")
	password = os.Getenv("PASSWORD")
	from = os.Getenv("FROM")
}

func SendMail(m Mail) error {
	auth := smtp.CRAMMD5Auth(username, password)
	msg := []byte(strings.ReplaceAll(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", strings.Join(m.Recipients, ","), m.Subject, m.Body), "\n", "\r\n"))
	if err := smtp.SendMail(fmt.Sprintf("%s:%d", hostname, port), auth, from, m.Recipients, msg); err != nil {
		return err
	}

	return nil
}
