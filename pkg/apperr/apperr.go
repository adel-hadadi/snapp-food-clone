package apperr

import (
	"snapp-food/pkg/logger"

	"go.uber.org/zap"
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
	ConflictClientMsg     = "اطلاعات وارد شده تکراری می‌باشد"
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
		logger.NewLogger().Error(a.sysMsg, zap.String("error", a.err.Error()))
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
