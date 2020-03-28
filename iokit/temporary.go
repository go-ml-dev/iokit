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

func Tempfile(pattern string) (_ io.ReadWriteCloser, err error) {
	var f *os.File
	if f, err = ioutil.TempFile("", pattern); err != nil {
		return
	}
	return &temporary{regular{f}, false}, nil
}

