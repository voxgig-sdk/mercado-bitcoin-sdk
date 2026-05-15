package core

type MercadoBitcoinError struct {
	IsMercadoBitcoinError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewMercadoBitcoinError(code string, msg string, ctx *Context) *MercadoBitcoinError {
	return &MercadoBitcoinError{
		IsMercadoBitcoinError: true,
		Sdk:              "MercadoBitcoin",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *MercadoBitcoinError) Error() string {
	return e.Msg
}
