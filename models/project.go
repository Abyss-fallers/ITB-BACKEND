package models

import (
	"time"

	"gorm.io/gorm"
)

type projectStatus int

const (
	ProjectOpen projectStatus = iota
	ProjectInProgress
	ProjectCompleted
	ProjectCanceled
)

type Project struct {
	gorm.Model
	Title       string
	Description string
	Budget      uint32
	ExecutorID  uint32
	ClientID    uint32
	Deadline    time.Time
	Status      projectStatus
}

type Review struct {
	gorm.Model
	ProjectID  uint64
	Project    Project
	ReviewerID uint32
	Reviewer   User
	RevieweeID uint32
	Reviewee   User
	Rating     byte
	Comment    string
	CreatedAt  time.Time
}
