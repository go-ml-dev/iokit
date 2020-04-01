package iokit

import (
	"github.com/sudachen/go-iokit/iokit/fu"
	"io"
	"strings"
)

type IoUrl struct {
	Url      string
	Schema   string
	Streamed LiveStream
	Cache    Cache
	Observer AsyncUpload
	Metadata Metadata
}

type Metadata map[string]string
type AsyncUpload struct{ Notify func(url string, err error) }
type LiveStream bool

const Streamed = LiveStream(true)

func Url(url string, opts ...interface{}) InoutExt {
	lurl := strings.ToLower(url)
	schema := ""
	if j := strings.Index(lurl, "://"); j > 0 {
		schema = lurl[:j]
	}
	return InoutExt{IoUrl{
		url,
		schema,
		fu.Option(LiveStream(false), opts).Interface().(LiveStream),
		fu.Option(Cache(""), opts).Interface().(Cache),
		fu.Option(AsyncUpload{nil}, opts).Interface().(AsyncUpload),
		fu.Option(Metadata(nil), opts).Interface().(Metadata),
	}}
}

func (p IoUrl) Open() (rd io.ReadCloser, err error) {
	if p.Schema != "file" {
		return p.openUrlReader()
	}
	return File(p.Url[7:]).Open()
}

func (p IoUrl) Create() (hw Whole, err error) {
	if p.Schema != "file" {
		return p.createUrlWriter()
	}
	return File(p.Url[7:]).Create()
}
