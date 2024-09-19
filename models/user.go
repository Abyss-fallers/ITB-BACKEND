package models

import (
	"time"

	"gorm.io/gorm"
)

type role int

const (
	RoleMember role = iota
	RoleMod
	RoleAdm
)

type skill string

type User struct {
	gorm.Model
	Email                      string    `json:"email"`
	Password                   string    `json:"-" gorm:"type:varchar(100)"`
	Fullname                   string    `json:"fullname"`
	AvatarURL                  string    `json:"-"`
	Role                       role      `json:"-"`
	LastLoginAt                time.Time `json:"last-login"`
	IsActive                   bool      `json:"-"`
	IsVerified                 bool      `json:"is-verified"`
	ResetPasswordToken         string    `json:"-"`
	ResetPasswordExpiresAt     time.Time `json:"-"`
	VerificationToken          string    `json:"verification-token"`
	VerificationTokenExpiresAt time.Time `json:"verification-token-expires-at"`
	CLientProjects             []Project `json:"client-projects" gorm:"foreignKey:ClientID"`
	ExecutorProjects           []Project `json:"executor-projects" gorm:"foreignKey:ExecutorID"`
}

type Skills struct {
	gorm.Model
	UserID     uint32
	User       User
	SkillID    skill //ТУТ ВОПРОСЫ
	ExtraSkill string
}

type Portfolio struct {
	gorm.Model
	UserID      uint32
	User        User
	Title       string
	Description string
	URL         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
