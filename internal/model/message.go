package model

import "time"

type Message struct {
	Timestamp time.Time
	Username  string
	Value     string
	Receiver  string
}
