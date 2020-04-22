package models

import (
	"time"
)

type WorkRecord struct {
	ID          uint64     `gorm:"primary_key"`
	UserId      uint64     `json:"-"`
	Date        string     `json:"date"`
	StartOfWork string     `json:"start_of_work"`
	EndOfWork   string     `json:"end_of_work"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`

	User User `json:"-"`
}
