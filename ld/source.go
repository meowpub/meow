package ld

import (
	"fmt"
)

type Source int

const (
	ClientToServer Source = iota
	ServerToServer
)

func (s Source) String() string {
	switch s {
	case ClientToServer:
		return "ClientToServer"
	case ServerToServer:
		return "ServerToServer"
	default:
		return fmt.Sprintf("[%d]", s)
	}
}
