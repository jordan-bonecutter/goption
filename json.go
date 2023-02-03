package goption

import (
  "encoding/json"
)

// MarshalJSON marshals the underlying option data
func (o Option[T]) MarshalJSON() ([]byte, error) {
  if !o.ok {
    return []byte("null"), nil
  }

  return json.Marshal(o.t)
}

// UnmarshalJSON unmarshals the underlying 
func (o *Option[T]) UnmarshalJSON(data []byte) error {
  if string(data) == "null" {
    o.ok = false
    return nil
  }

  o.ok = true
  return json.Unmarshal(data, &o.t)
}
