package iokit

import (
	"io"
	"io/ioutil"
	"os"
)

type temporary struct {
	regular
	deleted bool
}

func (tf temporary) Close() error {
	_ = tf.File.Close()
	if !tf.deleted {
		_ = os.Remove(tf.File.Name())
		tf.deleted = true
	}
	return nil
}

func (tf temporary) End()          { _ = tf.Close() }
func (tf temporary) Commit() error { return nil }

type TemporaryFile interface {
	Whole
	io.Reader
	io.Closer
}

func Tempfile(pattern string) (_ TemporaryFile, err error) {
	var f *os.File
	if f, err = ioutil.TempFile("", pattern); err != nil {
		return
	}
	return &temporary{regular{f}, false}, nil
}
