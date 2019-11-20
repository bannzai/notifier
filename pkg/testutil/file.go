package testutil

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"
)

// CallerDirectoryPath can get caller test file directory path
func CallerDirectoryPath(t *testing.T) string {
	t.Log("Exec CallerFilePath")
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		t.Error("Can not read current file when runtime")
	}
	t.Log("filename is " + filename)
	directoryPath := filepath.Dir(filename)
	t.Log("Current test directoryPath: " + directoryPath)
	return directoryPath
}

func ReadFile(t *testing.T, path string) []byte {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Error(err)
	}
	return bytes
}
