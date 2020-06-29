package app

import (
	"../libs/websocket"
)

// CommandHandler ..
type CommandHandler interface {
	Handle(s *Session, r *websocket.Request) (*websocket.Response, error)
	Validate(r *websocket.Request) error
}
