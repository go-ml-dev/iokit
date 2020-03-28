package iokit

import (
	"github.com/sudachen/go-foo/fu"
	"github.com/sudachen/go-iokit/iokit/internal"
	"os"
	"os/user"
	"path"
	"path/filepath"
)

const cacheGoFp = ".cache/gofper"

var FullCacheDir string

func init() {
	homedir, _ := os.LookupEnv("HOME")
	usr, err := user.Current()
	if err == nil {
		homedir = usr.HomeDir
	}
	if homedir == "" {
		homedir = "/tmp"
	}
	FullCacheDir, _ = filepath.Abs(filepath.Join(homedir, cacheGoFp))
}

func CacheDir(d string) string {
	r := fu.Ifes(filepath.IsAbs(d), d, path.Join(FullCacheDir, d))
	_ = os.MkdirAll(r, 0777)
	return r
}

func CacheFile(f string) string {
	r := fu.Ifes(filepath.IsAbs(f), f, path.Join(FullCacheDir, f))
	_ = os.MkdirAll(path.Dir(r), 0777)
	return r
}
