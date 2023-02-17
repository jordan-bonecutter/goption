package goption

import (
	"database/sql"
	"database/sql/driver"
	"reflect"
	"time"
)

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
	if srcVal.CanConvert(tType) {
		reflect.ValueOf(&o.t).Elem().Set(srcVal.Convert(tType))
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

func (o Option[T]) Value() (driver.Value, error) {
	if !o.ok {
		return nil, nil
	}

	var maybeValuer any = o.t
	if valuer, isValuer := maybeValuer.(driver.Valuer); isValuer {
		return valuer.Value()
	}

	tVal := reflect.ValueOf(o.t)
	int64Type := reflect.TypeOf(int64(0))
	if tVal.CanConvert(int64Type) {
		return tVal.Convert(int64Type).Interface(), nil
	}
	f64Type := reflect.TypeOf(float64(0))
	if tVal.CanConvert(f64Type) {
		return tVal.Convert(f64Type).Interface(), nil
	}
	boolType := reflect.TypeOf(false)
	if tVal.CanConvert(boolType) {
		return tVal.Convert(boolType).Interface(), nil
	}
	bytesType := reflect.TypeOf([]byte(nil))
	if tVal.CanConvert(bytesType) {
		return tVal.Convert(bytesType).Interface(), nil
	}
	stringType := reflect.TypeOf("")
	if tVal.CanConvert(stringType) {
		return tVal.Convert(stringType).Interface(), nil
	}
	timeType := reflect.TypeOf(time.Time{})
	if tVal.CanConvert(timeType) {
		return tVal.Convert(timeType).Interface(), nil
	}

	return o.t, nil
}
