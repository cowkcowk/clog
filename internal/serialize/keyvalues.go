package serialize

import "bytes"

type textWriter interface {
	WriteText(*bytes.Buffer)
}

func WithValues(oldKV, newKV []interface{}) []interface{} {
	if len(newKV) == 0 {
		return oldKV
	}
	newLen := len(oldKV) + len(newKV)
	hasMissingValue := newLen%2 != 0
	if hasMissingValue {
		newLen++
	}
	// The new LogSink must have its own slice.
	kv := make([]interface{}, 0, newLen)
	kv = append(kv, oldKV...)
	kv = append(kv, newKV...)
	if hasMissingValue {
		kv = append(kv, missingValue)
	}
	return kv
}

const missingValue = "(MISSING)"
