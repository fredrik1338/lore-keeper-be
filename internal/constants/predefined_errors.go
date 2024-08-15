package constants

import "errors"

var (
	ErrNilCharacter = errors.New("character is nil")
	ErrNilCity      = errors.New("city is nil")
	ErrNilWorld     = errors.New("world is nil")
	ErrNilFaction   = errors.New("faction is nil")
	ErrNilUniverse  = errors.New("universe is nil")
)
