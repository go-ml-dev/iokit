package iokit

import (
	"bytes"
	"io"
)

type StringIO string

func (s StringIO) Open() (io.ReadCloser, error) {
	return reader{bytes.NewBufferString(string(s)), nil},
		nil
}

