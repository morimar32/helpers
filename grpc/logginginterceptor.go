package grpc

import (
	"context"
	"fmt"
	"os"
	"path"
	"time"

	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/logging"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Adaptation of the code from: github.com/grpc-ecosystem/go-grpc-middleware/logging/zap
// Brought in primarily to customize the zap keys

//CustomLoggingInterceptor Customized logging interceptor adapted from: github.com/grpc-ecosystem/go-grpc-middleware/logging/zap
type CustomLoggingInterceptor struct {
	lvl zapcore.Level
}

//NewCustomLoggingInterceptor Creates a new instance of the logging interceptor
func NewCustomLoggingInterceptor(loggingLevel zapcore.Level) *CustomLoggingInterceptor {
	val := &CustomLoggingInterceptor{
		lvl: loggingLevel,
	}
	return val
}

//UnaryServerInterceptor Customized logging interceptor
func (l *CustomLoggingInterceptor) UnaryServerInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()

		newCtx := newLoggerForCall(ctx, logger, info.FullMethod, startTime)

		resp, err := handler(newCtx, req)
		code := grpc_logging.DefaultErrorToCode(err)
		level := zap.DebugLevel
		if code != codes.OK {
			level = grpc_zap.DefaultCodeToLevel(code)
		}

		if level < l.lvl {
			return resp, err
		}

		// re-extract logger from newCtx, as it may have extra fields that changed in the holder.
		ctxzap.Extract(newCtx).Check(level, fmt.Sprintf("finished unary call with code %s", code.String())).Write(
			zap.Error(err),
			zap.String("ReturnCode", code.String()),
			zap.Float32("ElapsedMilliseconds", durationToMilliseconds(time.Since(startTime))),
		)

		return resp, err
	}
}

func newLoggerForCall(ctx context.Context, logger *zap.Logger, fullMethodString string, start time.Time) context.Context {
	var f []zapcore.Field
	f = append(f, zap.String("RequestStartTime", start.Format(time.RFC3339)))
	if d, ok := ctx.Deadline(); ok {
		f = append(f, zap.String("RequestDeadline", d.Format(time.RFC3339)))
	}
	callLog := logger.With(append(f, serverCallFields(fullMethodString)...)...)
	return ctxzap.ToContext(ctx, callLog)
}

func serverCallFields(fullMethodString string) []zapcore.Field {
	service := path.Dir(fullMethodString)[1:]
	method := path.Base(fullMethodString)
	host, _ := os.Hostname()
	return []zapcore.Field{
		zap.String("Hostname", host),
		zap.String("RequestService", service),
		zap.String("RequestMethod", method),
	}
}

func durationToMilliseconds(duration time.Duration) float32 {
	return float32(duration.Nanoseconds()/1000) / 1000
}
