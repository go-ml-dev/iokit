[![CircleCI](https://circleci.com/gh/sudachen/go-iokit.svg?style=svg)](https://circleci.com/gh/sudachen/go-iokit)
[![Maintainability](https://api.codeclimate.com/v1/badges/9f73e3387f39f92b5169/maintainability)](https://codeclimate.com/github/sudachen/go-iokit/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/9f73e3387f39f92b5169/test_coverage)](https://codeclimate.com/github/sudachen/go-iokit/test_coverage)
[![Go Report Card](https://goreportcard.com/badge/github.com/sudachen/go-iokit)](https://goreportcard.com/report/github.com/sudachen/go-iokit)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)


```golang
 // example uses S3_TEST_URL envronment variable
 // export S3_TEST_URL=s3://id:secret@nyc3.digitaloceanspaces.com/bucket/TEST

// Write and read back S3 object
func Test_Example1(t *testing.T) {
	url := "s3://$test/test_example1.txt"
	S := fmt.Sprintf(`Hello world! %d`, rand.Int())

	wh := iokit.Url(url).LuckyCreate()
	defer wh.End()
	wh.LuckyWrite([]byte(S))
	wh.LuckyCommit()

	rd := iokit.Url(url).LuckyOpen()
	defer rd.LuckyClose()
	q := rd.LuckyReadAll()
	assert.Assert(t,string(q)==S)
}

// Write and read back S3 object
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
```
