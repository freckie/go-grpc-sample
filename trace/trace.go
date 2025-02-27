package trace

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

const metadataKey = "x-trace-id"

func NewTraceId() string {
	return uuid.New().String()
}

func TraceIdFrom(ctx context.Context) (string, bool) {
	vs := metadata.ValueFromIncomingContext(ctx, metadataKey)
	if len(vs) > 0 {
		return vs[0], true
	}
	return "", false
}

func WithTraceId(ctx context.Context, traceId string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, metadataKey, traceId)
}
