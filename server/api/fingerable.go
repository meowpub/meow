package api

import "github.com/meowpub/meow/lib/xrd"

// Fingerable is an object which can be WebFingered
type Fingerable interface {
	// Finger fingers an object
	Finger(req Request) (*xrd.XRD, error)
}
