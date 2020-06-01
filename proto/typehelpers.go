package proto

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// StringToStringValue returns a StringValue from a string
func StringToStringValue(value string) *wrappers.StringValue {
	if len(value) <= 0 {
		return nil
	}
	return &wrappers.StringValue{Value: value}
}

func StringValueToString(value *wrappers.StringValue) string {
	if value == nil {
		return ""
	}
	return value.Value
}

// TimeToTimestamp returns a Timestamp from *time.Time
func TimeToTimestamp(value *time.Time) *timestamppb.Timestamp {
	if value == nil {
		return nil
	}
	val, err := ptypes.TimestampProto(*value)
	if err == nil {
		return val
	}
	return nil
}
