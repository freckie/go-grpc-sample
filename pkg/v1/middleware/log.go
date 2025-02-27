package middleware

import (
	"context"
	"time"

	igrpc "frec.kr/tdoo/internal/grpc"
	"frec.kr/tdoo/trace"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func UnaryLogger(l *trace.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		// logger injection
		ctx = trace.WithLogger(ctx, l)

		tid, ok := trace.TraceIdFrom(ctx)
		if !ok {
			tid = trace.NewTraceId()
		}
		ctx = trace.WithTraceId(ctx, tid)

		// incoming
		in := l.WithField("tid", tid)
		in = in.WithField("mtd", info.FullMethod)
		if l.GetLevel() >= logrus.DebugLevel {
			in.WithField("req", req).Debug("IN")
		} else {
			in.Info("IN")
		}

		// handle
		t := time.Now()
		resp, err := handler(ctx, req)
		dt := time.Since(t)

		// outgoing
		out := l.WithField("tid", tid)
		out = out.WithField("dt", dt)
		if err != nil {
			if s, _ := status.FromError(err); s != nil {
				out.WithFields(logrus.Fields{
					"err_code": int(s.Code()),
					"err_type": s.Code().String(),
					"err_msg":  s.Message(),
				}).Error("FAIL")
			}
		} else {
			if l.GetLevel() >= logrus.DebugLevel {
				out.WithField("resp", resp).Debug("OUT")
			} else {
				out.Info("OUT")
			}
		}

		return resp, err
	}
}

func StreamLogger(l *trace.Logger) grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		// logger injection
		ctx := trace.WithLogger(ss.Context(), l)

		tid, ok := trace.TraceIdFrom(ctx)
		if !ok {
			tid = trace.NewTraceId()
		}
		ctx = trace.WithTraceId(ctx, tid)

		// incoming
		in := l.WithField("tid", tid)
		in = in.WithField("mtd", info.FullMethod)
		in.Info("STREAM_OPENED")

		// handle
		ss_ := igrpc.ServerStream{ServerStream: ss}
		ss_.Into(ctx)
		t := time.Now()
		err := handler(srv, &ss_)
		dt := time.Since(t)

		// outgoing (error)
		out := l.WithField("tid", tid)
		out = out.WithField("dt", dt)
		if err != nil {
			if s, _ := status.FromError(err); s != nil {
				out.WithFields(logrus.Fields{
					"err_code": int(s.Code()),
					"err_type": s.Code().String(),
					"err_msg":  s.Message(),
				}).Error("STREAM_FAIL")
			}
		}
		out.Info("STREAM_CLOSED")
		return err
	}
}
