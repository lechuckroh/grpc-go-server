package handler

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/lechuckroh/grpc-go-server/proto/hello"
	"testing"
)

func TestHello_Call(t *testing.T) {
	h := Hello{}
	req := hello.CallRequest{Name: "world"}
	resp, err := h.Call(context.Background(), &req)
	if err != nil {
		t.Error(err)
	}
	if diff := cmp.Diff("Hello world", resp.Msg); diff != "" {
		t.Errorf("TestHello_Call() mismatch (-expected +actual):\n%s", diff)
	}
}
