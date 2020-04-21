package tests

import (
	"fmt"
	"go-ml.dev/pkg/iokit"
	"gotest.tools/assert"
	"math/rand"
	"testing"
)

/*
	GS tests use GS_URL environment variables
	GS_ENCTEST_URL = gs://bucket/prefix:password:/abspath/credential.json.enc
*/

func Test_GsPath1(t *testing.T) {
	url := "gs://$enctest/test_gspath1.txt"
	S := fmt.Sprintf(`Hello world! %d`, rand.Int())
	wh := iokit.Url(url).LuckyCreate()
	defer wh.End()
	wh.LuckyWrite([]byte(S))
	wh.LuckyCommit()
	rd := iokit.Url(url).LuckyOpen()
	defer rd.Close()
	q := rd.LuckyReadAll()
	assert.Assert(t, string(q) == S)
}
