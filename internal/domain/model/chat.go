package model

type Chat struct {
	UserID string `json:"user_id"`
	ChatID string `json:"chat_id"`
}

type ChatRequest struct {
	UserID string `json:"user_id"`
}

type ChatResponse struct {
	UserID string `json:"user_id"`
	ChatID string `json:"chat_id"`
}

type ChatResponseList struct {
	Page    int             `json:"page"`
	HasMore bool            `json:"has_more"`
	Chats   []*ChatResponse `json:"chats"`
}

type ChatResponseListByUser struct {
	Page    int             `json:"page"`
	HasMore bool            `json:"has_more"`
	Chats   []*ChatResponse `json:"chats"`

	UserID string `json:"user_id"`
}
