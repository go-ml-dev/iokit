package tests

import (
	"fmt"
	"github.com/sudachen/go-iokit/iokit"
	"gotest.tools/assert"
	"math/rand"
	"testing"
)

func Test_Lzma2(t *testing.T) {
	S := fmt.Sprintf("test string %v", rand.Int())
	func() {
		w := iokit.Lzma2(iokit.Cache("test.lzma2").File()).LuckyCreate()
		defer w.End()
		w.LuckyWrite([]byte(S))
		w.LuckyCommit()
	}()
	s := func() string {
		r := iokit.Compressed(iokit.Cache("test.lzma2").File()).LuckyOpen()
		defer r.Close()
		return string(r.LuckyReadAll())
	}()
	assert.Assert(t, s == S)
}
