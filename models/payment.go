package models

import (
	"time"

	"gorm.io/gorm"
)

type transactionStatus int

const (
	TransactionPending transactionStatus = iota
	TransactionCompleted
	TransactionFailed
)

type paymentMethod int

const (
	PaymentMethodCard paymentMethod = iota
	PaymentMethodSBP
)

type Transaction struct {
	gorm.Model
	ID              uint64
	ProjectID       uint64
	Project         Project
	FromUserID      uint32
	FromUser        User
	ToUserID        uint32
	ToUser          User
	Amount          uint32
	TransactionDate time.Time
	Status          transactionStatus
}

type PaymentInfo struct {
	gorm.Model
	ID            uint64
	UserID        uint32
	User          User
	PaymentMethod paymentMethod
	Details       string
	CreatedAt     time.Time
}
