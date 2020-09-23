package logic

import "time"

type User struct {
	UID      int       `json:"uid"`
	NickName string    `json:"nick_name"`
	EnterAt  time.Time `json:"enter_at"`
	Addr     string    `json:"addr"`
}
