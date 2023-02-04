package goption

import (
  "database/sql"
  "database/sql/driver"
  "testing"
  "time"

  epg "github.com/fergusstrange/embedded-postgres"
  _ "github.com/lib/pq"
)

func TestSQLScanner(t *testing.T) {
  eDB := epg.NewDatabase(epg.DefaultConfig().Username("test").Password("test").Database("test").Port(2345))
  eDB.Start()
  t.Cleanup(func() { eDB.Stop() })

  db, err := sql.Open("postgres", "host=127.0.0.1 port=2345 user=test password=test dbname=test sslmode=disable")
  if err != nil {
    t.Fatalf("Failed connecting to database: %s", err.Error())
  }

  if _, err := db.Exec(`CREATE SCHEMA test;`); err != nil {
    t.Fatalf("Failed creating test schema: %s", err.Error())
  }

  if _, err := db.Exec(`CREATE TABLE test(
    key integer not null,
    maybe_empty integer,
    ts timestamptz
  );`); err != nil {
    t.Fatalf("Failed creating test table: %s", err.Error())
  }

  var testImplementsValuer any = Some[int](0)
  if _, isValuer := testImplementsValuer.(driver.Valuer); !isValuer {
    t.Errorf("Not a valuer")
  }

  now := time.Now()
  if _, err := db.Exec(`INSERT INTO test(
    key,
    maybe_empty,
    ts
  ) VALUES (
    0,
    $1,
    $2
  );`, Some[int](123), Some[time.Time](now)); err != nil {
    t.Fatalf("Failed inserting test data: %s", err.Error())
  }

  if _, err := db.Exec(`INSERT INTO test(
    key,
    maybe_empty,
    ts
  ) VALUES (
    1,
    $1,
    $2
  );`, None[int](), None[time.Time]()); err != nil {
    t.Fatalf("Failed inserting test data: %s", err.Error())
  }

  rows, err := db.Query("SELECT * FROM test;")
  if err != nil {
    t.Fatalf("Failed selecting test data: %s", err.Error())
  }
  defer rows.Close()

  for rows.Next() {
    var i Option[int]
    var ts Option[time.Time]
    var key int
    if err := rows.Scan(&key, &i, &ts); err != nil {
      t.Errorf("Failed scanning row: %s", err.Error())
      continue
    }

    switch key {
    case 0:
      if i.Unwrap() != 123 {
        t.Errorf("Unexpected value: %v", i.Unwrap())
      }
      if tstr := ts.Unwrap().Format("2006-01-02 15:04:05"); tstr != now.Format("2006-01-02 15:04:05") {
        t.Errorf("Bad time staring: %s", tstr)
      }
    case 1:
      if i.Ok() {
        t.Errorf("Expected value to be empty")
      }
      if ts.Ok() {
        t.Errorf("Expected tstr to be empty")
      }
    }
  }
}

