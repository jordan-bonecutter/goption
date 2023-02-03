package goption

// Option represents a value whose presence is optional.
type Option[T any] struct {
  t T
  ok bool
}

// Unwrap forcefull unwraps the Optional value.
// If the optional is not ok this function will panic.
func (o Option[T]) Unwrap() T {
  return o.Expect("Unwrapped empty optional")
}

// UnwrapRef returns a reference to the underlying T.
func (o Option[T]) UnwrapRef() *T {
  return o.ExpectRef("Unwrapped empty optional")
}

// Expect unwraps o and panics with msg if it's empty.
func (o Option[T]) Expect(msg string) T {
  if !o.ok {
    panic(msg)
  }

  return o.t
}

// ExpectRef unwraps o and panics with msg if it's empty.
func (o Option[T]) ExpectRef(msg string) *T {
  if !o.ok {
    panic(msg)
  }

  return &o.t
}

// UnwrapOr unwraps the optional if it's present, otherwise it returns default.
func (o Option[T]) UnwrapOr(def T) T {
  if !o.ok {
    return def
  }

  return o.t
}

// UnwrapOrDefault unwraps T if it's present, otherwise it returns the default 
// value for T.
func (o Option[T]) UnwrapOrDefault() T {
  var def T
  return o.UnwrapOr(def)
}

// Ok returns if the optional is present.
func (o Option[T]) Ok() bool {
  return o.ok
}

// Some returns an Option whose underlying value is present.
func Some[T any](t T) Option[T] {
  return Option[T]{
    t: t,
    ok: true,
  }
}

// None returns an empty optional value.
func None[T any]() Option[T] {
  return Option[T]{
    ok: false,
  }
}

// Apply f to the optional value.
func Apply[In, Out any](in Option[In], f func(In) Out) Option[Out] {
  if !in.ok {
    return None[Out]()
  }

  return Some(f(in.t))
}

