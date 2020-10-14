package logic

import (
	"time"
)

const (
	MsgTypeNormal    = iota // 普通 用户消息
	MsgTypeWelcome          // 当前用户欢迎消息
	MsgTypeUserEnter        // 用户进入
	MsgTypeUserLeave        // 用户退出
	MsgTypeError            // 错误消息
)

// 给用户发送的消息
type Message struct {
	// 哪个用户发送的消息
	User    *User     `json:"user"`
	Type    int       `json:"type"`
	Content string    `json:"content"`
	MsgTime time.Time `json:"msg_time"`

	ClientSendTime time.Time `json:"client_send_time"`

	// 消息 @ 了谁
	Ats []string `json:"ats"`

	// 用户列表不通过 WebSocket 下发
	// Users []*User `json:"users"`
}

//func NewMessage(user *User, content, clientTime string) *Message {
//	message := &Message{
//		User:    user,
//		Type:    MsgTypeNormal,
//		Content: content,
//		MsgTime: time.Now(),
//	}
//	if clientTime != "" {
//		message.ClientSendTime = time.Unix(0, cast.ToInt64(clientTime))
//	}
//	return message
//}
func NewMessage(content string) *Message {
	message := &Message{
		//User:    user,
		Type:    MsgTypeNormal,
		Content: content,
		MsgTime: time.Now(),
	}

	return message
}
