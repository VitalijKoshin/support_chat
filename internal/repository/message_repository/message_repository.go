package messagerepository

import "supportchat/internal/domain/model"

type IMessageRepository interface {
	GetMessageByID(messageID string) model.Message
	GetMessagesByUserID(userID string) []string
	AddMessage(message string)
}
