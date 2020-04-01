package iokit

import (
	"io"
	"os"
)

func File(path string) LuckyInputOutput {
	return LuckyInputOutput{file(path)}
}

type file string

func (f file) Open() (io.ReadCloser, error) {
	return os.Open(string(f))
}

func (f file) Create() (Whole, error) {
	x, err := os.Create(string(f))
	if err != nil {
		return nil, err
	}
	return &whole{regular{x}}, nil
}

type regular struct {
	*os.File
}

func (f regular) Reset() error {
	_, err := f.File.Seek(0, 0)
	return err
}

func (f regular) Size() int64 {
	st, _ := f.File.Stat()
	return st.Size()
}

func (f regular) Fail() {
	fname := f.File.Name()
	_ = f.File.Truncate(0)
	_ = f.File.Close()
	_ = os.Remove(fname)
}
