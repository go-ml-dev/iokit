package s3p

import (
	"github.com/aws/aws-sdk-go/aws"
	"strings"
	"sync"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Location struct {
	Bucket *string
	Key *string
}

type AccessPoint struct {
	Endpoint    string
	Region      string
	Bucket      string
	Prefix 		string
	Credentials *credentials.Credentials
	session     *session.Session
}

var mu sync.Mutex
var registry = map[string]*AccessPoint{}

func Register(ep string,ap AccessPoint) {
	mu.Lock()
	registry[strings.ToLower(ep)] = &ap
	mu.Unlock()
}

func Lookup(ep string) (*AccessPoint,bool) {
	mu.Lock()
	defer mu.Unlock()
	ap,ok := registry[strings.ToLower(ep)]
	return ap,ok
}

func (ap *AccessPoint) Session() (ssn *session.Session, err error) {
	if ap.session != nil { return ap.session, nil }
	return session.NewSession(&aws.Config{
		Endpoint:    &ap.Endpoint,
		Region:      &ap.Region,
		Credentials: ap.Credentials,
	})
}
