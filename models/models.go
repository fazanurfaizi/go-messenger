package models

type Chat struct {
	ID             int64
	Name           string
	AddrUsers      []string
	MessageBlockID int64
	LastSender     string
	LastMessage    string
}

type MessageBlock struct {
	ChatID   int64
	Messages []Message
}

type Message struct {
	ChatID     int64
	AddrAuthor string
	Content    string
	Type       string
}

type UserChatInfo struct {
	ID              int64           `json:"id"`
	Name            string          `json:"name"`
	Type            int             `json:"type"`
	LastSender      string          `json:"last_sender"`
	AdminID         int64           `json:"admin_id"`
	LastMessage     *MessageContent `json:"last_message"`
	LastMessageTime int64           `json:"last_message_time"`
	View            int             `json:"view"`
	Deleted         bool            `json:"deleted"`
	Banned          bool            `json:"banned"`
	Online          int64           `json:"online"`
}

type MessageContent struct {
	Message   string  `json:"content"`
	Documents []int64 `json:"documents"`
	Type      int     `json:"type"`
	Command   int     `json:"command,integer"`
}

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Language string `json:"language"`
	pass     string
}

type NewMessageToUser struct {
	ID         int64                 `json:"id"`
	ChatID     int64                 `json:"chat_id"`
	AuthorID   int64                 `json:"author_id"`
	Type       string                `json:"mtype"`
	Content    *MessageContentToUser `json:"message"`
	AuthorName string                `json:"author_name"`
	Time       int64                 `json:"time"`
}

type MessageContentToUser struct {
	Message   string  `json:"content"`
	Documents *[]File `json:"documents"`
	Type      int     `json:"type"`
	Command   int     `json:"command,integer"`
}

type File struct {
	ID        int64   `json:"id"`
	AuthorID  int64   `json:"author_id"`
	ChatID    int64   `json:"chat_id"`
	Name      string  `json:"name"`
	Path      string  `json:"path"`
	RatioSize float64 `json:"ratio"`
	Size      int64   `json:"size"`
	Duration  int64   `json:"duration"`
}

type CreateDHData struct {
	CommonName   string
	Organization string
	DNSNames1    string
	DNSNames2    string
	Type         string
}

type ForceMessageToUser struct {
	UserID  int64
	Message NewMessageToUser
}

type UserSettings struct {
	Name     string
	Language string
}

type ChatSettings struct {
	Name string `json:"name"`
}

type FolkChatsInfo struct {
	ID         int64  `json:"id"`
	Login      string `json:"login"`
	Name       string `json:"name"`
	DeleteLast int64  `json:"delete_last"`
	Ban        bool   `json:"ban"`
}

type EncryptedMessage struct {
	Type string `json:"mtype"`
	Data string `json:"data"`
	IV   string `json:"iv"`
	Key  string `json:"key"`
}

type MiddlewareMessage struct {
	ID       int64
	chatID   int64
	command  MessageCommand
	authorID int64
	time     int64
}
