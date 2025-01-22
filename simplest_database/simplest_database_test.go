package main

import (
	"os"
	"testing"
)

func TestSet(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "data.db")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	set("name", "John", tmpfile)

	content, err := os.ReadFile(tmpfile.Name())
	tmpfile.Close()

	if err != nil {
		t.Error("Error opening file:", err)
		return
	}

	if string(content) != "name,John\n" {
		t.Error("Unexpected content in file:", content)
	}
}

func TestGet(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "data.db")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	set("name", "John", tmpfile)
	set("name", "Jane", tmpfile)

	tmpfile.Seek(0, 0)
	value, _ := get("name", tmpfile)
	tmpfile.Close()

	if value != "Jane" {
		t.Error("Unexpected value:", value)
	}
}
