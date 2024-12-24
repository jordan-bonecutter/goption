package goption

import "iter"

// Option represents a value whose presence is optional.
type Option[T any] struct {
	t  T
	ok bool
}

// Unwrap forcefully unwraps the Optional value.
// If the optional is not ok this function will panic.
func (o Option[T]) Unwrap() T {
	return o.Expect("Unwrapped empty optional")
}

// UnwrapRef returns a reference to the underlying T.
func (o *Option[T]) UnwrapRef() *T {
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
func (o *Option[T]) ExpectRef(msg string) *T {
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

// UnwrapRefOr returns a reference if optional is present, otherwise it returns default.
func (o *Option[T]) UnwrapRefOr(def *T) *T {
	if !o.ok {
		return def
	}

	return &o.t
}

// UnwrapRefOrNil returns a reference if optional is present, otherwise it returns nil.
func (o *Option[T]) UnwrapRefOrNil() *T {
	var def *T
	return o.UnwrapRefOr(def)
}

// Ok returns if the optional is present.
func (o Option[T]) Ok() bool {
	return o.ok
}

// Get returns the underlying value and a boolean indicating if it's present.
func (o Option[T]) Get() (T, bool) {
	return o.t, o.ok
}

// Some returns an Option whose underlying value is present.
func Some[T any](t T) Option[T] {
	return Option[T]{
		t:  t,
		ok: true,
	}
}

// None returns an empty optional value.
func None[T any]() Option[T] {
	return Option[T]{
		ok: false,
	}
}

// FromRef returns an Option whose underlying value is present if t is not nil.
// Otherwise, it returns an empty optional value.
func FromRef[T any](t *T) Option[T] {
	if t == nil {
		return None[T]()
	}
	return Some(*t)
}

// Apply f to the optional value.
func Apply[In, Out any](in Option[In], f func(In) Out) Option[Out] {
	if !in.ok {
		return None[Out]()
	}

	return Some(f(in.t))
}

// Do runs the function f which may panic.
// If f does not panic Some(f()) is returned.
// Otherwise none is returned.
func Do[T any](f func() T) (o Option[T]) {
	defer func() {
		if r := recover(); r != nil {
			o = None[T]()
		}
	}()

	o = Some(f())
	return
}

// With allows you to use the optional value in it's own scope.
// For example:
//
// opt := Some(3)
// for value := opt.With() {
//   fmt.Println(value)
// }
//
// This is a hack on the iter.Seq interface.
func (o *Option[T]) With() iter.Seq[T] {
	return func(yield func(T) bool) {
		if o.ok {
			yield(o.t)
		}
	}
}

// IsZero returns true if o is not present.
// This is for use with encoding/json in go1.24+.
func (o *Option[T]) IsZero() bool {
	return !o.ok
}
