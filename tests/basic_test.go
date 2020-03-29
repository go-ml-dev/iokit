package tests

import (
	"github.com/sudachen/go-iokit/iokit"
	"gotest.tools/assert"
	"io/ioutil"
	"os"
	"testing"
)

func Test_Open(t *testing.T) {
	S := `test`
	w,err := os.Create(iokit.CacheFile("go-iokit/tests/create.txt"))
	assert.NilError(t,err)
	_,err = w.WriteString(S)
	assert.NilError(t,err)
	err = w.Close()
	assert.NilError(t,err)
	f := iokit.File(iokit.CacheFile("go-iokit/tests/create.txt"))
	r,err := f.Open()
	assert.NilError(t,err)
	x,err := ioutil.ReadAll(r)
	assert.NilError(t,err)
	assert.Assert(t, string(x) == S)
}

func Test_CreateOpen(t *testing.T) {
	S := `test`
	file := iokit.File(iokit.CacheFile("go-iokit/tests/createopen.txt"))
	w,err := file.Create()
	assert.NilError(t,err)
	defer w.End()
	_,err = w.Write([]byte(S))
	assert.NilError(t,err)
	err = w.Commit()
	assert.NilError(t,err)
	r,err := file.Open()
	assert.NilError(t,err)
	defer r.Close()
	x,err := ioutil.ReadAll(r)
	assert.NilError(t,err)
	assert.Assert(t, string(x) == S)
}
