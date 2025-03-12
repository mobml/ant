package utils

import (
	"os"
	"testing"
)

func TestGetPath(t *testing.T) {
	fileName := "test.txt"
	expectedPath := os.Getenv("APPDATA") + "\\ant\\" + fileName
	path, err := GetPath(fileName)

	if err != nil {
		t.Errorf("Error: %e", err)
	}

	if path != expectedPath {
		t.Errorf("Expected: %s, got: %s", expectedPath, path)
	}
}
