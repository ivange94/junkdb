package junkdb

import (
	"os"
	"path/filepath"
	"testing"
)

func TestJunkDB_Insert(t *testing.T) {
	db := setup(t)

	err := db.Insert("foo", "bar")
	if err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(db.uri)
	if err != nil {
		t.Fatal(err)
	}
	want := "foo,bar\n"
	got := string(data)
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestJunkDB_Get(t *testing.T) {
	db := setup(t)

	err := os.WriteFile(db.uri, []byte("foo,bar\n"), 0600)
	if err != nil {
		t.Fatal(err)
	}

	value, err := db.Get("foo")
	if err != nil {
		t.Fatal(err)
	}
	if value != "bar" {
		t.Errorf("want %q, got %q", "bar", value)
	}
}

func TestJunkDB_Get_Most_Recent_Entry(t *testing.T) {
	db := setup(t)

	err := os.WriteFile(db.uri, []byte("foo,bar\n"), 0600)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile(db.uri, []byte("foo,bars\n"), 0600)
	if err != nil {
		t.Fatal(err)
	}

	value, err := db.Get("foo")
	if err != nil {
		t.Fatal(err)
	}

	if value != "bars" {
		t.Errorf("want %q, got %q", "bars", value)
	}
}

func TestJunkDB_Get_Missing_Key(t *testing.T) {
	db := setup(t)

	value, err := db.Get("foo")
	if err != nil {
		t.Fatal(err)
	}
	if value != "" {
		t.Errorf("want %q, got %q", "", value)
	}
}

func setup(t *testing.T) *JunkDB {
	dir := t.TempDir()
	file := filepath.Join(dir, "test.db")
	db, err := Connect(file)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		err = db.Close()
		if err != nil {
			t.Fatal(err)
		}
	})
	return db
}
