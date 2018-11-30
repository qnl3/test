package test

import (
	"bytes"
	"io"
	"os"
)

// CaptureStdout accepts a function that return no value. The function is then
// executed capturing everything set to STDOUT
func CaptureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
