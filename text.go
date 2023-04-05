package goption

// MarshalText marshals the underlying option data
func (o Option[T]) MarshalText() ([]byte, error) {
	return o.MarshalJSON()
}

// UnmarshalText unmarshals the underlying
func (o *Option[T]) UnmarshalText(data []byte) error {
	return o.UnmarshalJSON(data)
}
