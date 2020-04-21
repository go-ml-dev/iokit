[![CircleCI](https://circleci.com/gh/go-ml-dev/iokit.svg?style=svg)](https://circleci.com/gh/go-ml-dev/iokit)
[![Maintainability](https://api.codeclimate.com/v1/badges/8697f28f706fe04ac59e/maintainability)](https://codeclimate.com/github/go-ml-dev/iokit/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/8697f28f706fe04ac59e/test_coverage)](https://codeclimate.com/github/go-ml-dev/iokit/test_coverage)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-ml-dev/iokit)](https://goreportcard.com/report/github.com/go-ml-dev/iokit)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)


```golang
 // S3Url uses S3_*_URL environment variables
 // export S3_DOTEST_URL=s3://id:secret@nyc3.digitaloceanspaces.com/bucket/prefix
 // export S3_AWSTEST_URL=s3://id:secret@s3.us-west-2.amazonaws.com/bucket/prefix
 // s3://$cerds/path means to use credentials form environment valiable S3_CREDS_URL
 // s3://$/bucket/path means to use default aws credentials
 //
 // GsUrl uses GS_*_URL environment variables
 // export GS_ENCTEST_URL=json://bucket:prefix:password:/abs/path/to/servicecredentials.json.enc
 // export GS_TEST_URL=json://bucket:prefix::/abs/path/to/servicecredentials.json
 // gs://$creds/path menas to use credentials form enviroment variable GS_CREDS_URL
 // gs://$/bucket/path means to use default google cloud credentials
 //
 // File url can use any environment variable
 // export FILES=.

// Write and read back 
func Test_Example(t *testing.T) {
    for _,url := range []string{
                        "s3://$awstest/test_example.txt",
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
```
