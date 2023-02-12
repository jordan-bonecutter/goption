package goption

import (
	"fmt"
)

// String implements fmt.Stringer
func (o Option[T]) String() string {
	if !o.ok {
		return "null"
	}

	var maybeStringer any = o.t
	if stringer, isStringer := maybeStringer.(fmt.Stringer); isStringer {
		return stringer.String()
	}

	return fmt.Sprintf("%v", o.t)
}

// GoString implements fmt.GoStringer
func (o Option[T]) GoString() string {
	if !o.ok {
		return fmt.Sprintf("Option[%T]{ok: false}", o.t)
	}

	var maybeGoStringer any = o.t
	if stringer, isStringer := maybeGoStringer.(fmt.GoStringer); isStringer {
		return stringer.GoString()
	}

	return fmt.Sprintf("%#v", o.t)
}
