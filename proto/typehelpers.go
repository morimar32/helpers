package proto

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
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

// TimestampToTime returns a *time.Time from a Timestamp
func TimestampToTime(value *timestamp.Timestamp) *time.Time {
	if value == nil {
		return nil
	}
	v := value.AsTime()
	return &v
}

func Int32ToInt32Value(value int32) *wrappers.Int32Value {
	if value == 0 {
		return nil
	}
	return &wrapperspb.Int32Value{Value: value}
}

func BoolToBoolValue(val bool) *wrappers.BoolValue {
	v := &wrappers.BoolValue{Value: val}
	return v

}

func Int64ToInt64Value(value int64) *wrappers.Int64Value {
	if value == 0 {
		return nil
	}
	return &wrappers.Int64Value{Value: value}
}

func DoubleValueToFloat64(val *wrappers.DoubleValue) float64 {
	if val == nil {
		return 0
	}
	return val.Value
}

func Float64ToDoubleValue(val float64) *wrappers.DoubleValue {
	v := &wrappers.DoubleValue{Value: val}
	return v
}
