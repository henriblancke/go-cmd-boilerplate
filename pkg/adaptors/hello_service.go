package adaptors

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	d "github.com/henriblancke/go-cmd-boilerplate/pkg/domain"
)

// HelloService implements the user service port
type HelloService struct{}

// Create Implements the user service method
func (hs *HelloService) Create(msg string) (d.User, error) {
	return d.User{
		ID:        uuid.NewString(),
		Message:   msg,
		CreatedAt: time.Now(),
	}, nil
}

// SayMessage implements the user service method
func (hs *HelloService) SayMessage(user d.User) (string, error) {
	if user.Message == "hello" {
		return "", errors.New("hello is lame")
	}

	return fmt.Sprintf("User with id %s, says hello with message: %s", user.ID, user.Message), nil
}
