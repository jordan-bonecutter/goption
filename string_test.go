package goption

import (
  "testing"
)

type IsStringer struct {}
func (IsStringer) String() string {
  return "haha"
}
func (IsStringer) GoString() string {
  return `"haha"`
}

func TestStringer(t *testing.T) {
  if str := Some(IsStringer{}).String(); str != "haha" {
    t.Errorf("Failed wrapping stringer: %s", str)
  }

  if str := None[IsStringer]().String(); str != "null" {
    t.Errorf("Failed stringer for empty option: %s", str)
  }

  if str := Some(struct{}{}).String(); str != "{}" {
    t.Errorf("Failed wrapping struct{}: %s", str)
  }

  if str := None[struct{}]().String(); str != "null" {
    t.Errorf("Failed wrapping struct{}: %s", str)
  }
}

func TestGoStringer(t *testing.T) {
  if str := Some(IsStringer{}).GoString(); str != `"haha"` {
    t.Errorf("Failed wrapping go stringer: %s", str)
  }

  if str := None[IsStringer]().GoString(); str != "Option[goption.IsStringer]{ok: false}" {
    t.Errorf("Failed go stringer for empty option: %s", str)
  }

  if str := Some(struct{}{}).GoString(); str != "struct {}{}" {
    t.Errorf("Failed wrapping struct{}: %s", str)
  }

  if str := None[struct{}]().GoString(); str != "Option[struct {}]{ok: false}" {
    t.Errorf("Failed wrapping struct{}: %s", str)
  }
}
