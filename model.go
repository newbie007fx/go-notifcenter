package notif

type NotifBody struct {
	Id          string       `json:"id,omitempty"`
	From        string       `json:"from"`
	To          []string     `json:"to"`
	Cc          []string     `json:"cc,omitempty"`
	Subject     string       `json:"subject"`
	Body        string       `json:"body"`
	Channel     string       `json:"channel"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

type NotifResponse struct {
	Status  *bool      `json:"status"`
	Message *NotifBody `json:"message"`
	Channel *string    `json:"channel"`
}

type Attachment struct {
	Filename string `json:"filename"`
	Url      string `json:"url"`
}
