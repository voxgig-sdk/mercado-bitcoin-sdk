package utility

import "github.com/voxgig-sdk/mercado-bitcoin-sdk/core"

func makeErrorUtil(ctx *core.Context, err error) (any, error) {
	if ctx == nil {
		ctx = &core.Context{
			Ctrl: &core.Control{},
			Op:   core.NewOperation(map[string]any{}),
		}
	}

	op := ctx.Op
	if op == nil {
		op = core.NewOperation(map[string]any{})
	}
	opname := op.Name
	if opname == "" || opname == "_" {
		opname = "unknown operation"
	}

	result := ctx.Result
	if result == nil {
		result = core.NewResult(map[string]any{})
	}
	result.Ok = false

	if err == nil {
		err = result.Err
	}
	if err == nil {
		err = ctx.MakeError("unknown", "unknown error")
	}

	errmsg := err.Error()
	msg := "MercadoBitcoinSDK: " + opname + ": " + errmsg
	msg = cleanUtil(ctx, msg).(string)

	result.Err = nil

	spec := ctx.Spec

	if ctx.Ctrl.Explain != nil {
		ctx.Ctrl.Explain["err"] = map[string]any{
			"message": msg,
		}
	}

	sdkErr := &core.MercadoBitcoinError{
		IsMercadoBitcoinError: true,
		Sdk:              "MercadoBitcoin",
		Code:             "",
		Msg:              msg,
		Ctx:              ctx,
		Result:           cleanUtil(ctx, result),
		Spec:             cleanUtil(ctx, spec),
	}
	if se, ok := err.(*core.MercadoBitcoinError); ok {
		sdkErr.Code = se.Code
	}

	ctx.Ctrl.Err = sdkErr

	if ctx.Ctrl.Throw != nil && *ctx.Ctrl.Throw == false {
		return result.Resdata, nil
	}

	return nil, sdkErr
}
