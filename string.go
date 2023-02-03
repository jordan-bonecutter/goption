package goption

import (
  "fmt"
)

// String implements fmt.Stringer
func (o Option[T]) String() string {
  if !o.ok {
    return "null"
  }

  return tryStringer(o.t)
}

func tryStringer(v any) string {
  if stringer, isStringer := v.(fmt.Stringer); isStringer {
    return stringer.String()
  }

  return fmt.Sprintf("%v", v)
}

// GoString implements fmt.GoStringer
func (o Option[T]) GoString() string {
  if !o.ok {
    return fmt.Sprintf("Option[%T]{ok: false}", o.t)
  }

  return tryGoStringer(o.t)
}

func tryGoStringer(v any) string {
  if stringer, isStringer := v.(fmt.GoStringer); isStringer {
    return stringer.GoString()
  }

  return fmt.Sprintf("%#v", v)
}
