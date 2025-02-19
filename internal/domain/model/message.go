package model

type Message struct {
	ID              string `json:"id"`
	ChatID          string `json:"chat_id"`
	UserID          string `json:"user_id"`           // user_id owner of the message
	UserIDRecipient string `json:"user_id_recipient"` // user_id recipient of the message
	Content         string `json:"content"`
	IsRead          bool   `json:"is_read"`
	IsDelivered     bool   `json:"is_delivered"`
	IsSent          bool   `json:"is_sent"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type Messages struct {
	Page     int        `json:"page"`
	HasMore  bool       `json:"has_more"`
	Messages []*Message `json:"messages"`
}

type MessageRequest struct {
	ChatID          string `json:"chat_id"`
	UserIDRecipient string `json:"user_id_recipient"`
	Content         string `json:"content"`
}

type MessageResponse struct {
	ID              string `json:"id"`
	ChatID          string `json:"chat_id"`
	UserID          string `json:"user_id"`
	UserIDRecipient string `json:"user_id_recipient"`
	Content         string `json:"content"`
	IsRead          bool   `json:"is_read"`
	IsDelivered     bool   `json:"is_delivered"`
	IsSent          bool   `json:"is_sent"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type MessageResponseList struct {
	Page     int                `json:"page"`
	HasMore  bool               `json:"has_more"`
	Messages []*MessageResponse `json:"messages"`
}

type MessageResponseListByChat struct {
	Page     int                `json:"page"`
	HasMore  bool               `json:"has_more"`
	Messages []*MessageResponse `json:"messages"`

	ChatID string `json:"chat_id"`
}

type MessageResponseListByUser struct {
	Page     int                `json:"page"`
	HasMore  bool               `json:"has_more"`
	Messages []*MessageResponse `json:"messages"`

	UserID string `json:"user_id"`
}
