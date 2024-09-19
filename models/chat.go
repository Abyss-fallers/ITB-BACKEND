package models

import (
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	ID        uint64
	User1ID   uint32
	User1     User
	User2ID   uint32
	User2     User
	CreatedAt time.Time
}

type Message struct {
	gorm.Model
	ID         uint64
	SenderID   uint32
	Sender     User
	ReceiverID uint32
	Receiver   User
	Content    string
	CreatedAt  time.Time
}
