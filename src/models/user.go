package models

import (
	"time"
)

type User struct {
	ID          uint64     `gorm:"primary_key"`
	UID         string     `json:"userId"`
	MailAddress string     `json:"email"`
	Name        string     `json:"userName"`
	Password    string     `json:"-"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`

	WorkRecords []WorkRecord `json:"-"`
}
