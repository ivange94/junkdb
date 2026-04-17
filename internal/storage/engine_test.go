package storage

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestEngineGetReturnsLatestValueFromAppendOnlyLog(t *testing.T) {
	tmpHome := t.TempDir()
	previousHome := os.Getenv("HOME")
	t.Setenv("HOME", tmpHome)
	t.Cleanup(func() {
		if previousHome != "" {
			_ = os.Setenv("HOME", previousHome)
		}
	})

	engine, err := NewEngine()
	if err != nil {
		t.Fatalf("NewEngine() error = %v", err)
	}

	if err := engine.Put("name", "first"); err != nil {
		t.Fatalf("Put() first error = %v", err)
	}
	if err := engine.Put("other", "value"); err != nil {
		t.Fatalf("Put() other error = %v", err)
	}
	if err := engine.Put("name", "second"); err != nil {
		t.Fatalf("Put() second error = %v", err)
	}

	got, err := engine.Get("name")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if got != "second" {
		t.Fatalf("Get() = %q, want %q", got, "second")
	}

	data, err := os.ReadFile(filepath.Join(tmpHome, ".junkdb", "data"))
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}
	wantLog := "name,first\nother,value\nname,second\n"
	if string(data) != wantLog {
		t.Fatalf("log contents = %q, want %q", string(data), wantLog)
	}
}

func TestEngineGetSupportsValuesContainingCommas(t *testing.T) {
	t.Setenv("HOME", t.TempDir())

	engine, err := NewEngine()
	if err != nil {
		t.Fatalf("NewEngine() error = %v", err)
	}
	if err := engine.Put("payload", "a,b,c"); err != nil {
		t.Fatalf("Put() error = %v", err)
	}

	got, err := engine.Get("payload")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if got != "a,b,c" {
		t.Fatalf("Get() = %q, want %q", got, "a,b,c")
	}
}

func TestEngineGetMissingKey(t *testing.T) {
	t.Setenv("HOME", t.TempDir())

	engine, err := NewEngine()
	if err != nil {
		t.Fatalf("NewEngine() error = %v", err)
	}

	_, err = engine.Get("missing")
	if !errors.Is(err, ErrKeyNotFound) {
		t.Fatalf("Get() error = %v, want ErrKeyNotFound", err)
	}
}
