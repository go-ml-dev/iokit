package tests

import (
	"fmt"
	"github.com/sudachen/go-iokit/iokit"
	"gotest.tools/assert"
	"io/ioutil"
	"math/rand"
	"testing"
)

/*
	S3 tests use S3_AWS_TEST_URL and S3_DO_TEST_URL environment variables
*/

func Test_S3Path1(t *testing.T) {
	for i, url := range []string{
		"s3://$do_test/go-iokit/test_s3path.txt",
		"s3://$aws_test/go-iokit/test_s3path.txt"} {

		S := fmt.Sprintf(`Hello world! %d`,i)
		wh := iokit.Url(url).LuckyCreate()
		defer wh.End()
		wh.LuckyWrite([]byte(S))
		wh.LuckyCommit()
		rd := iokit.Url(url).LuckyOpen()
		defer rd.LuckyClose()
		q := rd.LuckyReadAll()
		assert.Assert(t,string(q)==S)
	}
}

func Test_S3Path2(t *testing.T) {
	for i, url := range []string{
		"s3://$do_test/go-iokit/test_s3path.txt",
		"s3://$aws_test/go-iokit/test_s3path.txt"} {

		S := fmt.Sprintf(`Hello world! %d`,i)
		file := iokit.Url(url,iokit.Cache("go-iokit/test_s3path.txt"))
		wh,err := file.Create()
		assert.NilError(t,err)
		defer wh.End()
		_, err = wh.Write([]byte(S))
		assert.NilError(t,err)
		err = wh.Commit()
		assert.NilError(t,err)
		wh.End()
		rd,err := file.Open()
		assert.NilError(t,err)
		defer rd.Close()
		q,err := ioutil.ReadAll(rd)
		assert.NilError(t,err)
		assert.Assert(t,string(q)==S)
		rd.Close()
	}
	for i, url := range []string{
		"s3://$do_test/go-iokit/test_s3path.txt",
		"s3://$aws_test/go-iokit/test_s3path.txt"} {
		S := fmt.Sprintf(`Hello world! %d`,i)
		rd,err := iokit.Url(url).Open()
		assert.NilError(t,err)
		defer rd.Close()
		q,err := ioutil.ReadAll(rd)
		assert.NilError(t,err)
		assert.Assert(t,string(q)==S)
	}
}


func Test_Example1(t *testing.T) {
	url := "s3://$do_test/go-iokit/test_s3path.txt"
	S := fmt.Sprintf(`Hello world! %d`, rand.Int())
	wh := iokit.Url(url).LuckyCreate()
	defer wh.End()
	wh.LuckyWrite([]byte(S))
	wh.LuckyCommit()
	rd := iokit.Url(url).LuckyOpen()
	defer rd.Close()
	q := rd.LuckyReadAll()
	assert.Assert(t,string(q)==S)
}

func Test_Example2(t *testing.T) {
	url := "s3://$do_test/go-iokit/test_s3path.txt"
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
