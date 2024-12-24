package goption

import (
	"testing"
)

// TestUnwrapFail tests that unwrapping none fails..
func TestUnwrapFail(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected to fail unwrapping empty optional")
		}
	}()
	None[struct{}]().Unwrap()
}

// TestUnwrapSuccess tests that unwrapping some succeeds.
func TestUnwrapSuccess(t *testing.T) {
	val := Some(3).Unwrap()
	if val != 3 {
		t.Errorf("Failed unwrapping value, expected 3 but got %v", val)
	}
}

// TestUnwrapFail tests that unwrapping a ref on none fails.
func TestUnwrapRefFail(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected to fail unwrapping empty optional")
		}
	}()
	opt := None[struct{}]()
	opt.UnwrapRef()
}

// TestUnwrapFail tests that unwrapping a ref on some succeeds.
func TestUnwrapRefSuccess(t *testing.T) {
	opt := Some(3)
	val := opt.UnwrapRef()
	if *val != 3 {
		t.Errorf("Failed unwrapping value, expected 3 but got %v", val)
	}
}

// TestExpectMessage tests that the custom panic message delivers on expect.
func TestExpectMessage(t *testing.T) {
	// This should succeed!
	val := Some(3).Expect("not my custom message")
	if val != 3 {
		t.Errorf("Failed expecting int: %v", val)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected to fail unwrapping empty optional")
		} else if r.(string) != "my custom message" {
			t.Errorf("Failed setting expect string: %v", r)
		}
	}()
	None[struct{}]().Expect("my custom message")
}

// TestExpectMessage tests that the custom panic message delivers on expect ref.
func TestExpectRefMessage(t *testing.T) {
	opt := Some(3)
	val := opt.ExpectRef("not my custom message")
	if *val != 3 {
		t.Errorf("Failed expecting int: %v", *val)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected to fail unwrapping empty optional")
		} else if r.(string) != "my custom message" {
			t.Errorf("Failed setting expect string: %v", r)
		}
	}()
	None[struct{}]().Expect("my custom message")
}

// TestUnwrapOr tests that or values are returned when unwrap or-ing none.
// Otherwise it expects the underlying optional value.
func TestUnwrapOr(t *testing.T) {
	val := None[int]().UnwrapOr(10)
	if val != 10 {
		t.Errorf("Failed unwrapping optional: %v", val)
	}

	val = Some(4).UnwrapOr(10)
	if val != 4 {
		t.Errorf("Failed unwrapping optional: %v", val)
	}
}

// TestUnwrapOrDefault tests that default values are returned when unwrap or-ing none.
// Otherwise it expects the underlying optional value.
func TestUnwrapOrDefault(t *testing.T) {
	val := None[struct{ FooBar int }]().UnwrapOrDefault()
	if val.FooBar != 0 {
		t.Errorf("Expected default value of FooBar, got %v", val)
	}

	val = Some[struct{ FooBar int }](struct{ FooBar int }{FooBar: 11}).UnwrapOrDefault()
	if val.FooBar != 11 {
		t.Errorf("Failed unwrapping: %v", val)
	}
}

// TestOk tests that Ok returns true iff the value is Some.
func TestOk(t *testing.T) {
	if !Some(0).Ok() {
		t.Errorf("Some must always be present")
	}

	if None[struct{}]().Ok() {
		t.Errorf("None must always be empty")
	}
}

func TestGet(t *testing.T) {
	val, ok := Some(3).Get()
	if !ok {
		t.Error("Failed unwrapping value")
	}
	if val != 3 {
		t.Errorf("Failed unwrapping value, expected 3 but got %v", val)
	}

	_, ok = None[int]().Get()
	if ok {
		t.Error("Expected empty optional")
	}
}

// TestApply tests Applying a function on an optional.
func TestApply(t *testing.T) {
	square := func(v float64) float64 {
		return v * v
	}

	val := Apply(Some(3.0), square)
	if val.Unwrap() != 9.0 {
		t.Errorf("Expected 3^2 = 9, got %v", val)
	}

	val = Apply(None[float64](), square)
	if val.Ok() {
		t.Errorf("Expected empty optional")
	}
}

func TestDo(t *testing.T) {
	val := Do(func() int {
		return 1
	})
	if val.Unwrap() != 1 {
		t.Errorf("Expected 1, got %v", val)
	}

	val = Do(func() int {
		var a *int
		return *a
	})
	if val.Ok() {
		t.Errorf("Expected empty optional, got %v", val)
	}
}

func TestRef(t *testing.T) {
	myOption := Some[int](3)
	ref := myOption.UnwrapRef()
	if *ref != 3 {
		t.Errorf("Failed unwrapping option ref, expected 3 but got %d", *ref)
	}

	*ref = 4
	if unwrapped := myOption.Unwrap(); unwrapped != 4 {
		t.Errorf("Failed unwrapping option ref, expected 4 but got %d", unwrapped)
	}
}

func TestFromRef(t *testing.T) {
	empty := FromRef((*int)(nil))
	if empty.Ok() {
		t.Errorf("FromRef must be empty for nil")
	}

	val := 10
	ptr := &val
	myOption := FromRef(ptr)

	if !myOption.Ok() {
		t.Errorf("FromRef must be present for non-nil value")
	}

	if unwrapped := myOption.Unwrap(); unwrapped != 10 {
		t.Errorf("FromRef must contain dereferenced value, expected 10 but got %d", unwrapped)
	}
}

// TestUnwrapRefOr tests that or values are returned when unwrap or-ing none.
// Otherwise it expects the underlying optional value.
func TestUnwrapRefOr(t *testing.T) {
	opt := None[int]()
	def := 10

	val := opt.UnwrapRefOr(&def)
	if *val != 10 {
		t.Errorf("Failed unwrapping optional: %v", val)
	}

	opt = Some(4)
	val = opt.UnwrapRefOr(&def)
	if *val != 4 {
		t.Errorf("Failed unwrapping optional: %v", val)
	}
}

// TestUnwrapRefOrNil tests that default values are returned when unwrap or-ing none.
// Otherwise it expects the underlying optional value.
func TestUnwrapRefOrNil(t *testing.T) {
	opt := None[struct{ FooBar int }]()

	val := opt.UnwrapRefOrNil()
	if val != nil {
		t.Errorf("Expected default value of FooBar, got %v", val)
	}

	opt = Some[struct{ FooBar int }](struct{ FooBar int }{FooBar: 11})

	val = opt.UnwrapRefOrNil()
	if val.FooBar != 11 {
		t.Errorf("Failed unwrapping: %v", val)
	}
}

func TestWithSome(t *testing.T) {
	opt := Some(3)
	ranFor := 0
	for range opt.With() {
		ranFor++
	}
	if ranFor != 1 {
		t.Fatalf("Expected to run the loop once, ran %d times", ranFor)
	}
}

func TestWithNone(t *testing.T) {
	opt := None[int]()
	ranFor := 0
	for range opt.With() {
		ranFor++
	}
	if ranFor != 0 {
		t.Fatalf("Expected not to run the loop, ran %d times", ranFor)
	}
}

func TestIsZeroSome(t *testing.T) {
	opt := Some(1)
	if opt.IsZero() {
		t.Fatalf("Expected to not be zero")
	}
}

func TestIsZeroNone(t *testing.T) {
	opt := None[int]()
	if opt.IsZero() {
		t.Fatalf("Expected to not be zero")
	}
}

type fooZeroer bool

func (f fooZeroer) IsZero() bool {
	return bool(f)
}

func TestIsZeroWrappedZero(t *testing.T) {
	opt := Some(fooZeroer(true))
	if !opt.IsZero() {
		t.Fatalf("Expected to be zero")
	}
}

func TestIsZeroWrappedNotZero(t *testing.T) {
	opt := Some(fooZeroer(false))
	if opt.IsZero() {
		t.Fatalf("Expected not to be zero")
	}
}
