package logging

import (
	"io"

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
