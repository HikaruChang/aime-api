package types

import "time"

type TelegramOauthChatProps struct {
	ID       int    `json:"id"`
	PhotoURL string `json:"photo_url,omitempty"`
	Type     string `json:"type,omitempty"`
	Title    string `json:"title"`
	Username string `json:"username,omitempty"`
}

type TelegramOauthUserProps struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name,omitempty"`
	Username     string `json:"username,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
	IsBot        bool   `json:"is_bot,omitempty"`
	IsPrenium    bool   `json:"is_prenium,omitempty"`
}

type TelegramOauthProps struct {
	AuthDate     time.Time              `json:"auth_date"`
	CanSendAfter int                    `json:"can_send_after,omitempty"`
	Chat         TelegramOauthChatProps `json:"chat,omitempty"`
	ChatType     string                 `json:"chat_type,omitempty"`
	ChatInstance string                 `json:"chat_instance,omitempty"`
	User         TelegramOauthUserProps `json:"user,omitempty"`
	QueryID      string                 `json:"query_id,omitempty"`
	Receiver     TelegramOauthUserProps `json:"receiver,omitempty"`
	StartParam   string                 `json:"start_param,omitempty"`
	Hash         string                 `json:"hash"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
	Username  string `json:"username,omitempty"`
	PhotoURL  string `json:"photo_url,omitempty"`
}
