package model

import "time"

type Message struct {
	ID        int
	Sender    int
	Receiver  int
	Content   string
	Timestamp time.Time
}
