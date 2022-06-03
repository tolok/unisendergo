package unisendergo

import (
	"encoding/base64"
	"io"
	"io/ioutil"
	"time"
)

type Message struct {
	FromEmail           string         `json:"from_email"`
	Recipients          []Recipient    `json:"recipients"`
	TemplateId          string         `json:"template_id"`
	TemplateEngine      TemplateEngine `json:"template_engine"`
	GlobalSubstitutions Map            `json:"global_substitutions"`
	GlobalMetadata      Map            `json:"global_metadata"`
	GlobalLanguage      LanguageCode   `json:"global_language"`
	FromName            string         `json:"from_name"`
	ReplyTo             string         `json:"reply_to"`
	TrackLinks          int            `json:"track_links"`
	TrackRead           int            `json:"track_read"`
	Headers             Map            `json:"headers"`
	Body                Body           `json:"body"`
	SkipUnsubscribe     int            `json:"skip_unsubscribe"`
	Attachments         []Attachment   `json:"attachments"`
	InlineAttachments   []Attachment   `json:"inline_attachments"`
	Options             Options        `json:"options"`
}

type Recipient struct {
	Email         string `json:"email"`
	Substitutions Map    `json:"substitutions"`
	Metadata      Map    `json:"metadata"`
}

type Body struct {
	HTML      string `json:"html"`
	PlainText string `json:"plaintext"`
	AMP       string `json:"amp"`
}

type Attachment struct {
	ContentType string `json:"type"`
	Name        string `json:"name"`
	Content     string `json:"content"`
}

type Options struct {
	SendAt          time.Time `json:"send_at"`
	UnsubscribeURL  string    `json:"unsubscribe_url"`
	CustomBackendId int       `json:"custom_backend_id"`
	SMTPPoolId      string    `json:"smtp_pool_id"`
}

func NewMessage(fromEmail string, recipients []Recipient) Message {
	return Message{
		FromEmail:       fromEmail,
		Recipients:      recipients,
		SkipUnsubscribe: 0,
	}
}

func (m Message) WithTemplate(templateId string) Message {
	m.TemplateId = templateId
	if m.TemplateEngine == "" {
		m.TemplateEngine = Simple
	}
	return m
}

func (m Message) WithSubstitutions(substitutions Map) Message {
	m.GlobalSubstitutions = substitutions
	return m
}

func (m Message) WithMetadata(metadata Map) Message {
	m.GlobalMetadata = metadata
	return m
}

func (m Message) WithBody(html io.Reader, plaintext io.Reader, amp io.Reader) Message {
	htmlBody, _ := ioutil.ReadAll(html)
	plaintextBody, _ := ioutil.ReadAll(plaintext)
	ampBody, _ := ioutil.ReadAll(amp)

	m.Body.HTML = string(htmlBody)
	m.Body.PlainText = string(plaintextBody)
	m.Body.AMP = string(ampBody)

	return m
}

func (m Message) WithGlobalLanguage(code LanguageCode) Message {
	m.GlobalLanguage = code
	return m
}

func (m Message) WithFromName(fromName string) Message {
	m.FromName = fromName
	return m
}

func (m Message) WithReplyTo(replyTo string) Message {
	m.ReplyTo = replyTo
	return m
}

func (m Message) WithTracking(trackLinks int, trackRead int) Message {
	m.TrackLinks = trackLinks
	m.TrackRead = trackRead
	return m
}

func (m Message) WithHeader(key string, value interface{}) Message {
	m.Headers.Add(key, value)
	return m
}

func (m Message) WithAttachment(contentType string, name string, content io.Reader) Message {
	attachment := prepareAttachment(contentType, name, content)

	m.Attachments = append(m.Attachments, attachment)

	return m
}

func (m Message) WithInlineAttachment(contentType string, name string, content io.Reader) Message {
	attachment := prepareAttachment(contentType, name, content)

	m.InlineAttachments = append(m.InlineAttachments, attachment)

	return m
}

func prepareAttachment(contentType string, name string, content io.Reader) Attachment {
	b, _ := ioutil.ReadAll(content)
	b64s := base64.StdEncoding.EncodeToString(b)

	return Attachment{
		ContentType: contentType,
		Name:        name,
		Content:     b64s,
	}
}

func (m Message) WithOptions(options Options) Message {
	m.Options = options
	return m
}
