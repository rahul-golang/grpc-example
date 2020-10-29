package log_utils

import (
	"context"
	"os"
	"runtime"

	uuid "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

var logger *log.Logger

const RequestId = "trace_id"

func init() {
	logger = log.New()
	logger.SetLevel(log.DebugLevel)
	logger.Formatter = &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetOutput(os.Stdout)
}

func GetLogger(ctx context.Context) *log.Entry {
	var depth = 1
	var requestId string


	if ctxRqID, ok := ctx.Value(RequestId).(string); ok {
		requestId = ctxRqID
	}
	function, _, line, _ := runtime.Caller(depth)
	functionObject := runtime.FuncForPC(function)
	entry := logger.WithFields(log.Fields{
		"request_id": requestId,
		"function": functionObject.Name(),
		"line":     line,
	})
	logger.SetOutput(os.Stdout)
	return entry

}
func WithRqID(ctx context.Context) context.Context {
	return context.WithValue(ctx, RequestId, generateRequestID())
}

func generateRequestID() string {
	requestID := uuid.New()
	return requestID.String()

}
