package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_DoMagic(t *testing.T) {
	inStr := "Some test Data..."
	m := myStruct{
		data: []byte(inStr),
	}

	exp := strings.ToUpper(inStr)
	m.doMagic()

	if exp != string(m.data) {
		t.Errorf("Expected '%s', got '%s'\n", exp, m.data)
	}
}

func Test_NewFromFile(t *testing.T) {
	exp := "Some dummy data for test purposes only\n\n"

	m, err := newFromFile("./testdata/test1.txt")

	if err != nil {
		t.Fatalf("There should be no error. Got %s\n", err)
	}

	if exp != string(m.data) {
		t.Errorf("Expected '%s', got '%s'\n", exp, m.data)
	}
}

func Test_NewFromFile_TempFile(t *testing.T) {
	tempFile, err := ioutil.TempFile(os.TempDir(), "test1_")
	if err != nil {
		t.Fatalf("Cannot create temp file: %s\n", err)
	}
	defer os.Remove(tempFile.Name())

	exp := []byte("Some test data.\n")
	if _, err := tempFile.Write(exp); err != nil {
		t.Fatalf("Cannot write to temp file: %s\n", err)
	}

	if err := tempFile.Close(); err != nil {
		t.Fatalf("Cannot close temp file: %s\n", err)
	}

	m, err := newFromFile(tempFile.Name())

	if err != nil {
		t.Fatalf("There should be no error. Got %s\n", err)
	}

	if bytes.Compare(exp, m.data) != 0 {
		t.Errorf("Expected '%s', got '%s'\n", exp, m.data)
	}
}

// START HELPER OMIT
func createTempFile(t *testing.T, data []byte) (filename string, cleanFunc func()) {
	t.Helper()
	tempFile, err := ioutil.TempFile(os.TempDir(), "test2_")
	if err != nil {
		t.Fatalf("Cannot create temp file: %s\n", err)
	}

	if _, err := tempFile.Write(data); err != nil {
		t.Fatalf("Cannot write to temp file: %s\n", err)
	}

	if err := tempFile.Close(); err != nil {
		t.Fatalf("Cannot close temp file: %s\n", err)
	}

	return tempFile.Name(), func() { os.Remove(tempFile.Name()) }
}

func Test_NewFromFile_TempFile_Helper(t *testing.T) {
	exp := []byte("Some test data.\n")
	tempFile, cleanFunc := createTempFile(t, exp) // HL
	defer cleanFunc()                             // HL

	m, err := newFromFile(tempFile)

	if err != nil {
		t.Fatalf("There should be no error. Got %s\n", err)
	}

	if bytes.Compare(exp, m.data) != 0 {
		t.Errorf("Expected '%s', got '%s'\n", exp, m.data)
	}
}

// END HELPER OMIT
