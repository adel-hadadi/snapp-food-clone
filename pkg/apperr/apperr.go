package apperr

import (
	"log"
)

type Type int

const (
	Forbidden Type = iota + 1
	Unauthorized
	NotFound
	Conflict
	Invalid
	Unexpected

	ForbiddenClientMsg    = "forbidden"
	UnAuthorizedClientMsg = "unauthorized"
	NotFoundClientMsg     = "not found"
	ConflictClientMsg     = "DuplicateEntry"
	UnexpectedClientMsg   = "Unexpected"
	DefaultClientMsg      = "Unexpected"
)

var errTypeToClientMessage = map[Type]string{
	Forbidden:    ForbiddenClientMsg,
	Unauthorized: UnAuthorizedClientMsg,
	NotFound:     NotFoundClientMsg,
	Conflict:     ConflictClientMsg,
	Unexpected:   UnexpectedClientMsg,
}

type AppErr struct {
	err       error
	clientMsg string
	sysMsg    string
	Type      Type
}

func (a *AppErr) Error() string {
	if a.Type == Unexpected {
		log.Printf("[%s]: %v", a.sysMsg, a.err)
		// logger.NewLogger().Errorf("something bad happened, clientMsg: %s, sysMsg: %s, err: %s", a.clientMsg, a.sysMsg, a.err)
	}

	if a.clientMsg == "" {
		a.setClientMsg()
	}

	return a.clientMsg
}

func New(t Type) *AppErr {
	return &AppErr{Type: t}
}

func (a *AppErr) WithErr(err error) *AppErr {
	if a.err != nil {
		if iErr, ok := a.err.(*AppErr); ok {
			a.err = iErr.err
			a.Type = iErr.Type

			return a
		}
	}

	a.err = err

	return a
}

func (a *AppErr) WithSysMsg(msg string) *AppErr {
	a.sysMsg = msg

	return a
}

func (a *AppErr) WithMsg(msg string) *AppErr {
	a.clientMsg = msg

	return a
}

func (a *AppErr) setClientMsg() {
	msg, exists := errTypeToClientMessage[a.Type]
	if !exists {
		msg = DefaultClientMsg
	}

	a.clientMsg = msg
}
