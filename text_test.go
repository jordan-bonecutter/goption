package goption

import (
	"testing"
)

func TestTextMarshal(t *testing.T) {
	foo := Foo{
		Things: []int{1, 2, 3},
	}

	optFoo := Some(foo)

	encoded, err := optFoo.MarshalText()
	if err != nil {
		t.Fatalf("Failed marshalling json: %s", err)
	}

	if string(encoded) != `{"Stuff":null,"Things":[1,2,3]}` {
		t.Errorf("Unexpected encoded data: %s", string(encoded))
	}

	foo.Stuff = Some(Bar{Baz: "hey!"})
	optFoo = Some(foo)
	encoded, err = optFoo.MarshalText()
	if err != nil {
		t.Fatalf("Failed marshalling json: %s", err)
	}

	if string(encoded) != `{"Stuff":{"Baz":"hey!"},"Things":[1,2,3]}` {
		t.Errorf("Unexpected encoded data: %s", string(encoded))
	}
}

func TestTextUnmarshal(t *testing.T) {
	var foo Option[Foo]
	if err := foo.UnmarshalText([]byte(`{"Stuff":null,"Things":[1,2,3]}`)); err != nil {
		t.Errorf("Failed unmarshalling into foo: %s", err)
	} else if len(foo.UnwrapOrDefault().Things) != 3 ||
		foo.UnwrapOrDefault().Things[0] != 1 ||
		foo.UnwrapOrDefault().Things[1] != 2 ||
		foo.UnwrapOrDefault().Things[2] != 3 {
		t.Errorf("Failed unmarshalling foo.Things: %v", foo.UnwrapOrDefault().Things)
	} else if foo.UnwrapOrDefault().Stuff.Ok() {
		t.Errorf("Expected optional value to be empty.")
	}

	if err := foo.UnmarshalText([]byte(`{"Stuff":{"Baz":"hey!"},"Things":[1,2,3]}`)); err != nil {
		t.Errorf("Failed unmarshalling into foo: %s", err)
	} else if len(foo.UnwrapOrDefault().Things) != 3 ||
		foo.UnwrapOrDefault().Things[0] != 1 ||
		foo.UnwrapOrDefault().Things[1] != 2 ||
		foo.UnwrapOrDefault().Things[2] != 3 {
		t.Errorf("Failed unmarshalling foo.Things: %v", foo.UnwrapOrDefault().Things)
	} else if !foo.UnwrapOrDefault().Stuff.Ok() || foo.UnwrapOrDefault().Stuff.Unwrap().Baz != "hey!" {
		t.Errorf("Expected optional value to be present.")
	}
}
