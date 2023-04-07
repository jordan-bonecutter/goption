package goption

import (
	"encoding"
)

// MarshalText marshals the underlying option data
func (o Option[T]) MarshalText() ([]byte, error) {
	var maybeMarshaler any = o.t
	if valuer, isMarshaler := maybeMarshaler.(encoding.TextMarshaler); isMarshaler {
		return valuer.MarshalText()
	}

	return o.MarshalJSON()
}

// UnmarshalText unmarshals the underlying
func (o *Option[T]) UnmarshalText(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		o.ok = false
		return nil
	}

	// Try unmarshal
	var maybeUnmarshaler any = &o.t
	if unmarshaler, ok := maybeUnmarshaler.(encoding.TextUnmarshaler); ok {
		o.ok = true
		return unmarshaler.UnmarshalText(data)
	}

	return o.UnmarshalJSON(data)
}
