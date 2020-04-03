package tests

import (
	"fmt"
	"github.com/sudachen/go-iokit/iokit"
	"gotest.tools/assert"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
)

func Test_Example(t *testing.T) {

	os.Setenv("FILES",".")

	for _,url := range []string{
		"s3://$do_test/test_example.txt",
		"gs://$enctest/test_example.txt",
		"file://$files/test_example.txt"} {

		S := fmt.Sprintf(`Hello world! %d`, rand.Int())

		wh,err := iokit.Url(url).Create()
		assert.NilError(t,err)
		defer wh.End()
		_,err = wh.Write([]byte(S))
		assert.NilError(t,err)
		err = wh.Commit()
		assert.NilError(t,err)

		rd,err := iokit.Url(url).Open()
		assert.NilError(t,err)
		defer rd.Close()
		q,err := ioutil.ReadAll(rd)
		assert.NilError(t,err)
		assert.Assert(t,string(q)==S)
	}
}
