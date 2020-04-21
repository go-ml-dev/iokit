package iokit

import (
	"bytes"
	"io"
)

type stringIO string

func StringIO(str string) LuckyInput {
	return LuckyInput{stringIO(str)}
}

func (s stringIO) Open() (io.ReadCloser, error) {
	return reader{bytes.NewBufferString(string(s)), nil}, nil
}
