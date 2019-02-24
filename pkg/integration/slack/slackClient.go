package slack

import (
	nlopes "github.com/nlopes/slack"
)

func NewSlackClient(token string) *nlopes.Client {
	return nlopes.New(token)
}
