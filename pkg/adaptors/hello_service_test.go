package adaptors

import (
	"errors"
	"testing"
	"time"

	d "github.com/henriblancke/go-cmd-boilerplate/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func TestHelloServiceCreate(t *testing.T) {
	var testCases = []struct {
		msg         string
		service     *HelloService
		exp         string
		description string
	}{
		{
			msg:         "hello",
			service:     &HelloService{},
			exp:         "hello",
			description: "Creating a user with message hello totally works",
		},
		{
			msg:         "hello world",
			service:     &HelloService{},
			exp:         "hello world",
			description: "hello world should be a valid message",
		},
	}

	for _, test := range testCases {
		user, _ := test.service.Create(test.msg)

		assert.Equal(t, test.exp, user.Message)
	}
}

func TestHelloServiceSayMessage(t *testing.T) {
	var testCases = []struct {
		input       d.User
		service     *HelloService
		exp         string
		expErr      error
		description string
	}{
		{
			input: d.User{
				ID:        "1",
				Message:   "hello",
				CreatedAt: time.Now(),
			},
			service:     &HelloService{},
			exp:         "",
			expErr:      errors.New("hello is lame"),
			description: "Hello should return an error because it is lame",
		},
		{
			input: d.User{
				ID:        "1",
				Message:   "hello world",
				CreatedAt: time.Now(),
			},
			service:     &HelloService{},
			exp:         "User with id 1, says hello with message: hello world",
			expErr:      nil,
			description: "hello world should be a valid message",
		},
	}

	for _, test := range testCases {
		message, err := test.service.SayMessage(test.input)

		assert.Equal(t, test.expErr, err)
		assert.Equal(t, test.exp, message)
	}
}
