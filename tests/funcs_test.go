package tests

// in this file I want to continually test the functions in the funcs package and afterwards possibly also the agent file

import (
	"fmt"
	"github.com/KennyZ69/gobootai/funcs"
	"os"
	"testing"
)

func Test_Get_files_info(t *testing.T) {
	// Set up a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "gobootai_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up after test

	// Create some test files and directories
	testFiles := []string{"file1.txt", "file2.txt", "subdir"}
	for _, file := range testFiles {
		path := fmt.Sprintf("%s/%s", tempDir, file)
		if file == "subdir" {
			if err := os.Mkdir(path, 0755); err != nil {
				t.Fatalf("Failed to create subdirectory: %v", err)
			}
		} else {
			if _, err := os.Create(path); err != nil {
				t.Fatalf("Failed to create file %s: %v", file, err)
			}
		}
	}

	// Call the function to test
	resp, err := funcs.Get_files_info(tempDir, ".")
	if err != nil {
		t.Fatalf("Get_files_info failed: %v", err)
	}

	expectedOutput := "File: file1.txt, Size: 0 bytes, IsDir: false\n" +
		"File: file2.txt, Size: 0 bytes, IsDir: false\n" +
		"File: subdir, Size: 4096 bytes, IsDir: true\n"

	if resp != expectedOutput {
		t.Errorf("Expected:\n%s\nGot:\n%s", expectedOutput, resp)
	}
}

func Test_Get_file_content(t *testing.T) {
	// Set up a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "gobootai_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up after test

	// Create a test file with some content
	testFile := fmt.Sprintf("%s/testfile.txt", tempDir)
	content := "This is a test file."
	if err := os.WriteFile(testFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Call the function to test
	resp, err := funcs.Get_file_content(tempDir, "testfile.txt")
	if err != nil {
		t.Fatalf("Get_file_content failed: %v", err)
	}

	if resp != content {
		t.Errorf("Expected content:\n%s\nGot:\n%s", content, resp)
	}
}

func Test_Write_file(t *testing.T) {
	// Set up a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "gobootai_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up after test

	// Define the file path and content
	// filePath := fmt.Sprintf("%s/testfile.txt", tempDir)
	content := "This is a test file."

	// Call the function to test
	bytesWritten, err := funcs.Write_file(tempDir, "testfile.txt", content)
	if err != nil {
		t.Fatalf("Write_file failed: %v", err)
	}

	if bytesWritten != len(content) {
		t.Errorf("Expected to write %d bytes, but wrote %d bytes", len(content), bytesWritten)
	}

	// Verify that the file was created with the correct content
	readContent, err := funcs.Get_file_content(tempDir, "testfile.txt")
	if err != nil {
		t.Fatalf("Failed to read back the file: %v", err)
	}

	if readContent != content {
		t.Errorf("Expected content:\n%s\nGot:\n%s", content, readContent)
	}
}
