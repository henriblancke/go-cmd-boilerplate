package ports

import (
	d "github.com/henriblancke/go-cmd-boilerplate/pkg/domain"
)

type UserService interface {
	Create(msg string) (d.User, error)
	SayMessage(user d.User) (string, error)
}
