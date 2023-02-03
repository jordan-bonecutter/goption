package goption

import (
  "testing"
  "encoding/json"
)

type Bar struct {
  Baz string
}

type Foo struct {
  Stuff Option[Bar]
  Things []int
}


func TestJSONMarshal(t *testing.T) {
  foo := Foo{
    Things: []int{1, 2, 3},
  }
  encoded, err := json.Marshal(foo)
  if err != nil {
    t.Fatalf("Failed marshalling json: %s", err)
  }

  if string(encoded) != `{"Stuff":null,"Things":[1,2,3]}` {
    t.Errorf("Unexpected encoded data: %s", string(encoded))
  }

  foo.Stuff = Some(Bar{Baz: "hey!"})
  encoded, err = json.Marshal(foo)
  if err != nil {
    t.Fatalf("Failed marshalling json: %s", err)
  }

  if string(encoded) != `{"Stuff":{"Baz":"hey!"},"Things":[1,2,3]}` {
    t.Errorf("Unexpected encoded data: %s", string(encoded))
  }
}

func TestJSONUnmarshal(t *testing.T) {
  var foo Foo
  if err := json.Unmarshal([]byte(`{"Stuff":null,"Things":[1,2,3]}`), &foo); err != nil {
    t.Errorf("Failed unmarshalling into foo: %s", err)
  } else if len(foo.Things) != 3 || foo.Things[0] != 1 || foo.Things[1] != 2 || foo.Things[2] != 3 {
    t.Errorf("Failed unmarshalling foo.Things: %v", foo.Things)
  } else if foo.Stuff.Ok() {
    t.Errorf("Expected optional value to be empty.")
  }

  if err := json.Unmarshal([]byte(`{"Stuff":{"Baz":"hey!"},"Things":[1,2,3]}`), &foo); err != nil {
    t.Errorf("Failed unmarshalling into foo: %s", err)
  } else if len(foo.Things) != 3 || foo.Things[0] != 1 || foo.Things[1] != 2 || foo.Things[2] != 3 {
    t.Errorf("Failed unmarshalling foo.Things: %v", foo.Things)
  } else if !foo.Stuff.Ok() || foo.Stuff.Unwrap().Baz != "hey!" {
    t.Errorf("Expected optional value to be present.")
  }
}
