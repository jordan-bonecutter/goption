# goption
![Test](https://github.com/jordan-bonecutter/goption/workflows/Main/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/jordan-bonecutter/goption.svg)](https://pkg.go.dev/github.com/jordan-bonecutter/goption)

Optional values for Go. Inspired by the Rust implementation.

## Why?
I created this out of frustration for the incoherency of optional handling across the language. In sql, we use `NullType`, in json and encoding friends we use pointers, and elsewhere we use `ok`. This package adds a more coherent optional experience across the language and is well tested.

## Features
Attempts to follow a monadic design where if the wrapped type `T` implements some interface, so should `Option[T]`. The following are implemented:
- `json.Marshaler`
- `json.Unmarshaler`
- `fmt.Stringer`
- `fmt.GoStringer`
- `sql.Scanner`
- `sql.driver.Valuer`

If there are any more interfaces which should be wrapped, please open an issue or a PR. All features must be tested.

## Examples

To declare an optional value do:

```go
myOption := Some[int](3)
```

To check if an optional value is present, do:

```go
opt := functionReturningOption()
if opt.Ok() {
  fmt.Println(opt.Unwrap())
}
```

To move values in and out of a sql database do:

```go
db.Exec(`
  INSERT INTO test (optional_field) VALUES ($1);
`, Some[int](0))

rows, err := db.Query(`
  SELECT * FROM test;
`)
for rows.Next() {
  var opt Option[int]
  rows.Scan(&opt)
  fmt.Println(opt)
}
```
