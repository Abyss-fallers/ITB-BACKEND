package models

import (
	"time"

	"gorm.io/gorm"
)

type notificationStatus int

const (
	NotificationSent notificationStatus = iota + 1
	NotificationRecieved
	NotificationRead
)

type Notification struct {
	gorm.Model
	ID        uint64
	UserID    uint32
	User      User
	Type      notificationStatus //ТУТ ВОПРОСЫ
	Content   string
	IsRead    bool
	CreatedAt time.Time
}
