package slack

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/co0p/go-tls-watch/pkg/usecases"
	"github.com/nlopes/slack"
)

type Handler struct {
	ValidateUsecase *usecases.ValidateUsecase
	Client          *slack.Client
}

func (h *Handler) respond(rtm *slack.RTM, msg *slack.MessageEvent, prefix string) {

	text := msg.Text
	text = strings.TrimPrefix(text, prefix)
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	channel := msg.Channel

	// slack wraps urls with <...>
	re := regexp.MustCompile("<(.*?)>")
	url := re.FindString(text)
	if len(url) == 0 {
		rtm.SendMessage(rtm.NewOutgoingMessage("¯\\_(ツ)_/¯ no url found", channel))
		return
	}

	cert, err := h.ValidateUsecase.Validate(url[1 : len(url)-1])

	if err != nil {
		rtm.SendMessage(rtm.NewOutgoingMessage("¯\\_(ツ)_/¯ "+err.Error(), channel))
		return
	}

	var responseText string
	if cert.IsValid() {
		responseText += "Yeah, cert is valid"
	} else {
		responseText += "Nope, cert is invalid"
	}

	rtm.SendMessage(rtm.NewOutgoingMessage(responseText, channel))
}

func (h *Handler) Handle() {
	rtm := h.Client.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {

		case *slack.MessageEvent:
			info := rtm.GetInfo()
			prefix := fmt.Sprintf("<@%s> ", info.User.ID)

			if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {
				h.respond(rtm, ev, prefix)
			}

		default:

			// Ignore other events..
			// fmt.Printf("Unexpected: %v\n", msg.Data)
		}
	}
}
