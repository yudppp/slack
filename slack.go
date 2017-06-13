package slack

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

var once sync.Once
var slackClient *SlackClient

type SlackClient struct {
	httpClient      *http.Client
	webhookURL      string
	defaultUsername string
	defaultChannel  string
}

type SlackMessage struct {
	Text        string             `json:"text,omitempty"`
	Channel     string             `json:"channel,omitempty"`
	Username    string             `json:"username,omitempty"`
	IconURL     string             `json:"icon_url,omitempty"`
	IconEmoji   string             `json:"icon_emoji,omitempty"`
	LinkNames   int                `json:"link_names,omitempty"`
	Attachments []*SlackAttachment `json:"attachments,omitempty"`
}

type SlackAttachment struct {
	Title      string        `json:"title,omitempty"`
	TitleLink  string        `json:"title_link,omitempty"`
	Text       string        `json:"text,omitempty"`
	Color      string        `json:"color,omitempty"`
	AuthorName string        `json:"author_name,omitempty"`
	AuthorLink string        `json:"author_link,omitempty"`
	AuthorIcon string        `json:"author_icon,omitempty"`
	ImageURL   string        `json:"image_url,omitempty"`
	ThumbURL   string        `json:"thumb_url,omitempty"`
	Footer     string        `json:"footer,omitempty"`
	FooterIcon string        `json:"footer_icon,omitempty"`
	Fields     []*SlackField `json:"fields,omitempty"`
}

type SlackField struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
}

func NewClient() *SlackClient {
	once.Do(func() {
		slackClient = &SlackClient{
			httpClient: http.DefaultClient,
		}
	})
	return slackClient
}

func (c *SlackClient) Send(msg interface{}) error {
	params, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", c.webhookURL, strings.NewReader(url.Values{"payload": {string(params)}}.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()
	return nil
}

func SetWebhookURL(webhookURL string) {
	client := NewClient()
	client.webhookURL = webhookURL
}

func SetDefaultUsername(defaultUsername string) {
	client := NewClient()
	client.defaultUsername = defaultUsername
}

func SetDefaultChannel(defaultChannel string) {
	client := NewClient()
	client.defaultChannel = defaultChannel
}

func NewMessage() *SlackMessage {
	client := NewClient()
	return &SlackMessage{
		Username: client.defaultUsername,
		Channel:  client.defaultChannel,
	}
}

func NewAttachment() *SlackAttachment {
	return &SlackAttachment{}
}

func NewField() *SlackField {
	return &SlackField{}
}

func (msg *SlackMessage) SetText(text string) *SlackMessage {
	msg.Text = text
	return msg
}

func (msg *SlackMessage) SetChannel(channel string) *SlackMessage {
	msg.Channel = channel
	return msg
}

func (msg *SlackMessage) SetUsername(username string) *SlackMessage {
	msg.Username = username
	return msg
}

func (msg *SlackMessage) SetIconURL(iconURL string) *SlackMessage {
	msg.IconURL = iconURL
	return msg
}

func (msg *SlackMessage) SetIconEmoji(iconEmoji string) *SlackMessage {
	msg.IconEmoji = iconEmoji
	return msg
}

func (msg *SlackMessage) UseLinkNames(available bool) *SlackMessage {
	if available {
		msg.LinkNames = 1
	} else {
		msg.LinkNames = 0
	}
	return msg
}

func (msg *SlackMessage) AddAttachment(attachment *SlackAttachment) *SlackMessage {
	if len(msg.Attachments) == 0 {
		msg.Attachments = []*SlackAttachment{attachment}
		return msg
	}
	msg.Attachments = append(msg.Attachments, attachment)
	return msg
}

func (attachment *SlackAttachment) SetTitle(title string) *SlackAttachment {
	attachment.Title = title
	return attachment
}

func (attachment *SlackAttachment) SetTitleLink(link string) *SlackAttachment {
	attachment.TitleLink = link
	return attachment
}

func (attachment *SlackAttachment) SetText(text string) *SlackAttachment {
	attachment.Text = text
	return attachment
}

func (attachment *SlackAttachment) SetColor(color string) *SlackAttachment {
	attachment.Color = color
	return attachment
}
func (attachment *SlackAttachment) SetAuthorName(authorName string) *SlackAttachment {
	attachment.AuthorName = authorName
	return attachment
}

func (attachment *SlackAttachment) SetAuthorLink(authorLink string) *SlackAttachment {
	attachment.AuthorLink = authorLink
	return attachment
}

func (attachment *SlackAttachment) SetAuthorIcon(authorIcon string) *SlackAttachment {
	attachment.AuthorIcon = authorIcon
	return attachment
}

func (attachment *SlackAttachment) SetImageURL(imageURL string) *SlackAttachment {
	attachment.ImageURL = imageURL
	return attachment
}

func (attachment *SlackAttachment) SetThumbURL(thumbURL string) *SlackAttachment {
	attachment.ThumbURL = thumbURL
	return attachment
}

func (attachment *SlackAttachment) SetFooter(footer string) *SlackAttachment {
	attachment.Footer = footer
	return attachment
}

func (attachment *SlackAttachment) SetFooterIcon(footerIcon string) *SlackAttachment {
	attachment.FooterIcon = footerIcon
	return attachment
}

func (attachment *SlackAttachment) AddField(field *SlackField) *SlackAttachment {
	if len(attachment.Fields) == 0 {
		attachment.Fields = []*SlackField{field}
		return attachment
	}
	attachment.Fields = append(attachment.Fields, field)
	return attachment
}

func (field *SlackField) SetTitle(title string) *SlackField {
	field.Title = title
	return field
}

func (field *SlackField) SetValue(value string) *SlackField {
	field.Value = value
	return field
}

func (field *SlackField) SetShort(short bool) *SlackField {
	field.Short = short
	return field
}

func (msg *SlackMessage) Send() error {
	client := NewClient()
	return client.Send(msg)
}
