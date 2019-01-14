package qqbotapi

import (
	"encoding/json"
	"strconv"

	"github.com/siskinc/qq-bot-api/cqcode"
)

// APIResponse is a response from the Coolq HTTP API with the result
// stored raw.
type APIResponse struct {
	Status  string          `json:"status"`
	Data    json.RawMessage `json:"data"`
	RetCode int             `json:"retcode"`
	Echo    interface{}     `json:"echo"`
}

type WebSocketRequest struct {
	Action string                 `json:"action"`
	Params map[string]interface{} `json:"params"`
	Echo   interface{}            `json:"echo"`
}

// Update is an update response, from GetUpdates.
type Update struct {
	Time          int64       `json:"time"`
	PostType      string      `json:"post_type"`
	MessageType   string      `json:"message_type"`
	SubType       string      `json:"sub_type"`
	MessageID     int64       `json:"message_id"`
	GroupID       int64       `json:"group_id"`
	DiscussID     int64       `json:"discuss_id"`
	UserID        int64       `json:"user_id"`
	Font          int         `json:"font"`
	RawMessage    interface{} `json:"message"`
	Anonymous     interface{} `json:"anonymous"`
	AnonymousFlag string      `json:"anonymous_flag"`
	Event         string      `json:"event"`
	NoticeType    string      `json:"notice_type"`
	OperatorID    int64       `json:"operator_id"`
	File          *File       `json:"file"`
	RequestType   string      `json:"request_type"`
	Flag          string      `json:"flag"`
	Text          string      `json:"-"` // Known as "message", in a message or request
	Message       *Message    `json:"-"`
	Sender        *User       `json:"sender"`
}

// UpdatesChannel is the channel for getting updates.
type UpdatesChannel <-chan Update

// File is a file.
type File struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	BusID int64  `json:"busid"`
}

// User is a user on QQ.
type User struct {
	ID       int64  `json:"user_id"`
	NickName string `json:"nickname"`
	Sex      string `json:"sex"` // "male"、"female"、"unknown"
	Age      int    `json:"age"`
	Area     string `json:"area"`
	// Group member
	Card                string `json:"card"`
	CardChangeable      bool   `json:"card_changeable"`
	Title               string `json:"title"`
	TitleExpireTimeUnix int64  `json:"title_expire_time"`
	Level               string `json:"level"`
	Role                string `json:"role"` // "owner"、"admin"、"member"
	Unfriendly          bool   `json:"unfriendly"`
	JoinTimeUnix        int64  `json:"join_time"`
	LastSentTimeUnix    int64  `json:"last_sent_time"`
	AnonymousID         int64  `json:"anonymous_id" anonymous:"id"`
	AnonymousName       string `json:"anonymous_name" anonymous:"name"`
	AnonymousFlag       string `json:"anonymous_flag" anonymous:"flag"`
}

// Group is a group on QQ.
type Group struct {
	ID   int64  `json:"group_id"`
	Name string `json:"group_name"`
}

// String displays a simple text version of a user.
//
// It is normally a user's card, but falls back to a nickname as available.
func (u *User) String() string {
	p := ""
	if u.Title != "" {
		p = "[" + u.Title + "]"
	}
	return p + u.Name()
}

// Name displays a simple text version of a user.
func (u *User) Name() string {
	if u.AnonymousName != "" {
		return u.AnonymousName
	}
	if u.Card != "" {
		return u.Card
	}
	if u.NickName != "" {
		return u.NickName
	}
	return strconv.FormatInt(u.ID, 10)
}

// Chat contains information about the place a message was sent.
type Chat struct {
	ID      int64  `json:"id"`
	Type    string `json:"type"`     // "private"、"group"、"discuss"
	SubType string `json:"sub_type"` // (only when Type is "private") "friend"、"group"、"discuss"、"other"
}

// IsPrivate returns if the Chat is a private conversation.
func (c Chat) IsPrivate() bool {
	return c.Type == "private"
}

// IsGroup returns if the Chat is a group.
func (c Chat) IsGroup() bool {
	return c.Type == "group"
}

// IsDiscuss returns if the Chat is a discuss.
func (c Chat) IsDiscuss() bool {
	return c.Type == "discuss"
}

// Message is returned by almost every request, and contains data about
// almost anything.
type Message struct {
	*cqcode.Message `json:"message"`
	MessageID       int64  `json:"message_id"`
	From            *User  `json:"from"`
	Chat            *Chat  `json:"chat"`
	Text            string `json:"text"`
	SubType         string `json:"sub_type"` // (only when Chat.Type is "group") "normal"、"anonymous"、"notice"
	Font            int    `json:"font"`
}

// IsAnonymous returns if a message is an anonymous message.
func (m Message) IsAnonymous() bool {
	return m.SubType == "anonymous"
}

// IsNotice returns if a message is a notice.
func (m Message) IsNotice() bool {
	return m.SubType == "notice"
}
