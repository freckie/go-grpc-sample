package grpc

import (
	"context"

	"google.golang.org/grpc"
)

type ServerStream struct {
	grpc.ServerStream
	Ctx context.Context
}

func (ss *ServerStream) Into(ctx context.Context) {
	ss.Ctx = ctx
}

func (ss *ServerStream) Context() context.Context {
	return ss.Ctx
}
