package domain

import (
	"time"
)

type User struct {
	ID        string
	Message   string `validate:"required"`
	CreatedAt time.Time
}
