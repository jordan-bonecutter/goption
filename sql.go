package goption

import (
  "reflect"
  "database/sql"
  "unsafe"
)

type exportOption[T any] struct {
  Val T
  Ok bool
}

// Scan implements sql.Scanner for Options
func (o *Option[T]) Scan(src any) error {
  if src == nil {
    *o = None[T]()
    return nil
  }

  // Try scanning
  var maybeScanner any = &o.t
  if scanner, isScanner := maybeScanner.(sql.Scanner); isScanner {
    o.ok = true
    return scanner.Scan(src)
  }

  // Try reflecting
  srcVal := reflect.ValueOf(src)
  tType := reflect.TypeOf(o.t)
  exportPtr := (*exportOption[T])(unsafe.Pointer(o))
  if srcVal.CanConvert(tType) {
    reflect.ValueOf(exportPtr).Elem().Field(0).Set(srcVal.Convert(tType))
    o.ok = true
    return nil
  }

  return ErrNotAScanner
}

type errNotAScanner struct{}
func (errNotAScanner) Error() string {
  return "Not a scanner"
}
var ErrNotAScanner errNotAScanner
