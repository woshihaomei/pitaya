package services

import (
	"context"

	"github.com/woshihaomei/pitaya/component"
	"github.com/woshihaomei/pitaya/examples/demo/worker/protos"
	"github.com/woshihaomei/pitaya/logger"
)

// Metagame server
type Metagame struct {
	component.Base
}

// LogRemote logs argument when called
func (m *Metagame) LogRemote(ctx context.Context, arg *protos.Arg) (*protos.Response, error) {
	logger.Log.Infof("argument %+v\n", arg)
	return &protos.Response{Code: 200, Msg: "ok"}, nil
}
