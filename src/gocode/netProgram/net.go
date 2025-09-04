package netProgram

import "time"

const (
	tcp = "tcp"
)

// 自定义的消息结构类型
type Message struct {
	ID      uint      `json:"id,omitempty"`
	Code    string    `json:"code,omitempty"`
	Content string    `json:"content,omitempty"`
	Time    time.Time `json:"time,omitempty"`
}
