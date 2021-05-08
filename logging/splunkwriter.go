package logging

import (
	"io"
	"strings"

	"github.com/ZachtimusPrime/Go-Splunk-HTTP/splunk"
	"go.uber.org/zap/zapcore"
)

//SplunkWriter type to support writing to splunk via the io.Writer interface
type SplunkWriter struct {
	Client  splunk.Client
	Writer  io.Writer
	entries chan []byte
}

//Sync functionality required by the zapcore.WriteSyncer interface
func (w *SplunkWriter) Sync() error {
	return nil
}

//Write implementation of the io.Writer interface
func (w *SplunkWriter) Write(b []byte) (int, error) {
	err := w.Client.Log(string(b))
	if err != nil {
		return -1, err
	}
	return len(b), nil
}

//NewSplunkWriter Instanciates a new splunk client wrapped in a zapcore.Lock for thread safety
func NewSplunkWriter(collectorEndpoint string, hecToken string, source string, sourceType string, index string) zapcore.WriteSyncer {
	splunkClient := splunk.NewClient(
		nil,
		collectorEndpoint,
		hecToken,
		source,
		sourceType,
		index,
	)

	writer := &SplunkWriter{}
	writer.Client = *splunkClient
	lockWriter := zapcore.Lock(writer)
	return lockWriter
}

func GetLogLevel(level string) zapcore.Level {
	level = strings.ToLower(level)
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.WarnLevel
	}
}

func GetDefaultJSONEncoder() zapcore.Encoder {
	ecfg := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	enc := zapcore.NewJSONEncoder(ecfg)
	return enc
}
