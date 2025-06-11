package main

import (
	"os"
	"testing"
)

func TestParseFile(t *testing.T) {
	content := "#...\n#...\n#...\n#...\n\n....\n....\n..##\n..##\n"
	f, err := os.CreateTemp(t.TempDir(), "sample-*.txt")
	if err != nil {
		t.Fatalf("temp file: %v", err)
	}
	defer os.Remove(f.Name())
	if _, err := f.WriteString(content); err != nil {
		t.Fatalf("write: %v", err)
	}
	f.Close()
	pieces, err := ParseFile(f.Name())
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if len(pieces) != 2 {
		t.Fatalf("expected 2 pieces got %d", len(pieces))
	}
}
