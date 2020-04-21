package tests

import (
	"fmt"
	"go-ml.dev/pkg/iokit"
	"gotest.tools/assert"
	"math/rand"
	"testing"
)

func Test_Zip(t *testing.T) {
	S := fmt.Sprintf("test string %v", rand.Int())
	func() {
		w := iokit.Zip("test.txt", iokit.Cache("test.zip").File()).LuckyCreate()
		defer w.End()
		w.LuckyWrite([]byte(S))
		w.LuckyCommit()
	}()
	s := func() string {
		r := iokit.ZipFile("test.txt", iokit.Cache("test.zip").File()).LuckyOpen()
		defer r.Close()
		return string(r.LuckyReadAll())
	}()
	assert.Assert(t, s == S)
}

func Test_Zip_Fail(t *testing.T) {
	S := fmt.Sprintf("test string %v", rand.Int())
	func() {
		w := iokit.Zip("test.txt", iokit.Cache("test.zip").File()).LuckyCreate()
		defer w.End()
		w.LuckyWrite([]byte(S))
		// no commit here
	}()
	_, err := iokit.ZipFile("test.txt", iokit.Cache("test.zip").File()).Open()
	assert.Assert(t, err != nil)
}
