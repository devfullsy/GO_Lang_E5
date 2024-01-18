package dictionary

import (
	"os"
	"strings"
	"testing"
)

func TestLogToFile(t *testing.T) {
	testFilePath := "testlogfile.txt"

	file, err := os.Create(testFilePath)
	if err != nil {
		t.Fatalf("Error creating test log file: %v", err)
	}
	defer file.Close()

	testLogToFile := func(message string) {
		logToFile(message)
	}

	logMessage := "Test log message"
	testLogToFile(logMessage)

	content, err := os.ReadFile(testFilePath)
	if err != nil {
		t.Fatalf("Error reading test log file: %v", err)
	}

	if !strings.Contains(string(content), logMessage) {
		t.Errorf("Expected log message in file, got:\n%s", string(content))
	}

	os.Remove(testFilePath)
}
