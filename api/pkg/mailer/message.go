package mailer

import (
	"github.com/jordan-wright/email"
	"net/mail"
	"net/textproto"
)

// Attachment represents an email attachment.
type Attachment struct {
	Filename    string
	ContentType string
	Header      textproto.MIMEHeader
	Content     []byte
}

// Message represents a smtp message.
type Message struct {
	From        mail.Address
	To          []string
	Cc          []string
	Bcc         []string
	ReplyTo     []string
	ReadReceipt []string
	Subject     string
	Body        string
	ContentType ContentType
	Headers     textproto.MIMEHeader
	Attachments []*Attachment

	email *email.Email
}

type ContentType string

func (t ContentType) IsTextHtml() bool {
	return t == TextHtml
}

const (
	TextHtml  ContentType = "text/html"
	TextPlain ContentType = "text/plain"
)

func NewMessage(subject, body string, contentType ...ContentType) Message {
	defContentType := TextPlain
	if len(contentType) >= 1 {
		defContentType = contentType[0]
	}

	return Message{
		Subject:     subject,
		Body:        body,
		ContentType: defContentType,

		email: new(email.Email),
	}
}

func (m Message) AttachFile(filename string) error {
	_, err := m.email.AttachFile(filename)
	return err
}

func (m Message) Bytes() ([]byte, error) {
	return m.email.Bytes()
}

func (m Message) fillEmail(from mail.Address) *email.Email {
	m.email.ReplyTo = m.ReplyTo
	m.email.To = m.To
	m.email.Bcc = m.Bcc
	m.email.Cc = m.Cc
	m.email.Subject = m.Subject
	m.email.Headers = m.Headers
	m.email.ReadReceipt = m.ReadReceipt

	if m.From == (mail.Address{}) {
		m.email.From = from.String()
	} else {
		m.email.From = m.From.String()
	}

	if m.ContentType.IsTextHtml() {
		m.email.HTML = []byte(m.Body)
	} else {
		m.email.Text = []byte(m.Body)
	}

	return m.email
}
