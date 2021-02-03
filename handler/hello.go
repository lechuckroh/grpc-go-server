package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/lechuckroh/grpc-go-server/pb/hellopb"
	"log"
)

type Hello struct {
}

// Call is a single request handler called via client.Call or the generated client code
func (c *Hello) Call(ctx context.Context, req *hellopb.CallRequest) (*hellopb.CallResponse, error) {
	log.Printf("Received Hello.Call request: %+v", req)
	return &hellopb.CallResponse{
		Msg:       "Hello " + req.Name,
		Timestamp: ptypes.TimestampNow(),
	}, nil
}

func (c *Hello) Healthcheck(ctx context.Context, req *hellopb.Empty) (*hellopb.Empty, error) {
	log.Printf("Received Common.Healthcheck request: %+v", req)
	return &hellopb.Empty{}, nil
}
