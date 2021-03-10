package plugins

import "github.com/Mrs4s/MiraiGo/message"

// NewMessageRequsetFromPrivateMessage NewMessageRequsetFromPrivateMessage
func NewMessageRequsetFromPrivateMessage(msg *message.PrivateMessage) *MessageRequest {
	return &MessageRequest{
		MessageType: PrivateMessage,
		ID:          msg.Id,
		InternalID:  msg.InternalId,
		Sender:      msg.Sender,
		Time:        msg.Time,
		Elements:    msg.Elements,
	}
}

// NewMessageRequsetFromGroupMessage NewMessageRequsetFromGroupMessage
func NewMessageRequsetFromGroupMessage(msg *message.GroupMessage) *MessageRequest {
	return &MessageRequest{
		MessageType:    GroupMessage,
		ID:             msg.Id,
		InternalID:     msg.InternalId,
		GroupCode:      msg.GroupCode,
		GroupName:      msg.GroupName,
		Sender:         msg.Sender,
		Time:           msg.Time,
		Elements:       msg.Elements,
		OriginalObject: msg.OriginalObject,
	}
}
