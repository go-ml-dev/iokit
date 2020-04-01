package tests

import (
	"fmt"
	"github.com/sudachen/go-iokit/iokit"
	"gotest.tools/assert"
	"math/rand"
	"testing"
)

func Test_Gzip(t *testing.T) {
	S := fmt.Sprintf("test string %v", rand.Int())
	func() {
		w := iokit.Gzip(iokit.Cache("test.gz").File()).LuckyCreate()
		defer w.End()
		w.LuckyWrite([]byte(S))
		w.LuckyCommit()
	}()
	s := func() string {
		r := iokit.Compressed(iokit.Cache("test.gz").File()).LuckyOpen()
		defer r.Close()
		return string(r.LuckyReadAll())
	}()
	assert.Assert(t, s == S)
}

func Test_Gzip_Fail(t *testing.T) {
	S := fmt.Sprintf("test string %v", rand.Int())
	func() {
		w := iokit.Gzip(iokit.Cache("test.gz").File()).LuckyCreate()
		defer w.End()
		w.LuckyWrite([]byte(S))
		// no commit here
	}()
	_, err := iokit.Compressed(iokit.Cache("test.gz").File()).Open()
	assert.Assert(t, err != nil)
}
