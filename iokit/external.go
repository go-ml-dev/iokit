package iokit

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
)

type Cached string
type Streamed_ bool

const Streamed = Streamed_(true)

func (c Cached) String() string {
	return CacheFile(string(c))
}

func (c Cached) Remove() (err error) {
	s := CacheFile(string(c))
	_, err = os.Stat(s)
	if err == nil {
		return os.Remove(s)
	}
	return nil
}

type External_ struct {
	url      string
	cache    string
	streamed bool
}

func External(url string, opts ...interface{}) External_ {
	return External_{url, option(Cached(""), opts).String(), false}
}

func (e External_) Open() (io.ReadCloser, error) {
	if e.streamed {
		return StreamedDownload(e.url)
	}
	return CachedDownload(e.url, e.cache)
}

func CachedDownload(url string, cached string) (_ io.ReadCloser, err error) {
	var f io.ReadWriteCloser
	if cached != "" {
		cached = CacheFile(cached)
		if _, err = os.Stat(cached); err == nil {
			if f, err = os.Open(cached); err != nil {
				return nil, errors.Wrap(err, "filed to open cached file")
			}
			return f, nil
		}
		if f, err = os.Create(cached); err != nil {
			return
		}
	} else {
		if f, err = Tempfile("external-noncached-*"); err != nil {
			return nil, errors.Wrap(err,"could not create temporal file")
		}
	}
	err = download(url, f.(io.Writer))
	if err != nil {
		_ = f.Close()
		return nil, errors.Wrap(err, fmt.Sprintf("download error: ",err.Error()))
	}
	_, _ = f.(io.Seeker).Seek(0, 0)
	return f, nil
}

func download(url string, writer io.Writer) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(writer, resp.Body)
	return err
}

func StreamedDownload(url string) (_ io.ReadCloser, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	rd := resp.Body
	return Reader(rd, func() error { return rd.Close() }), nil
}
