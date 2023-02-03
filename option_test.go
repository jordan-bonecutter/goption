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
  None[struct{}]().UnwrapRef()
}

// TestUnwrapFail tests that unwrapping a ref on some succeeds.
func TestUnwrapRefSuccess(t *testing.T) {
  val := Some(3).UnwrapRef()
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
  val := Some(3).ExpectRef("not my custom message")
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
  val := None[struct{FooBar int}]().UnwrapOrDefault()
  if val.FooBar != 0 {
    t.Errorf("Expected default value of FooBar, got %v", val)
  }

  val = Some[struct{FooBar int}](struct{FooBar int}{FooBar: 11}).UnwrapOrDefault()
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
