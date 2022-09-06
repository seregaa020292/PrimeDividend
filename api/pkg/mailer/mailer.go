package mailer

import (
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"time"

	"github.com/jordan-wright/email"
)

type Sender interface {
	Send(msg Message) error
}

type Config struct {
	Host        string
	Username    string
	Password    string
	Port        int
	TLS         bool
	From        mail.Address
	PoolConn    int
	PoolTimeout time.Duration
}

func (c Config) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

type Mailer struct {
	auth   smtp.Auth
	pool   *email.Pool
	config Config
}

type unencryptedAuth struct {
	smtp.Auth
	config Config
}

func (a unencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	s := *server
	s.TLS = a.config.TLS
	return a.Auth.Start(&s)
}

// NewMailer
//
// For example:
//  sender := NewMailer(config Config)
//  m := NewMessage("Test", "Body message.")
//  m.From = mail.Address{Name: "From", Address: "test@example.com"}
//  m.To = []string{"test@example.com"}
//  m.Bcc = []string{"test_bcc@example.com"}
//  m.Cc = []string{"test_cc@example.com"}
//  m.Subject = "Awesome Subject"
//  m.Text = []byte("Text Body is, of course, supported!")
//  m.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
//  m.AttachFile("/path/to/file", false)
//
//  fmt.Println(sender.Send(m))
func NewMailer(config Config) Sender {
	auth := unencryptedAuth{
		Auth:   smtp.PlainAuth("", config.Username, config.Password, config.Host),
		config: config,
	}
	pool, err := email.NewPool(config.Addr(), config.PoolConn, auth)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Start Mailer")

	return Mailer{
		auth:   auth,
		pool:   pool,
		config: config,
	}
}

func (m Mailer) Send(msg Message) error {
	return m.pool.Send(msg.fillEmail(m.config.From), m.config.PoolTimeout)
}
