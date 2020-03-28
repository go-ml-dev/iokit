package iokit

import "fmt"

const Version VersionType = 10000

type VersionType int

func (v VersionType) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major(), v.Minor(), v.Patch())
}

func (v VersionType) Major() int {
	return int(v / 10000)
}

func (v VersionType) Minor() int {
	return int(v / 100 % 100)
}

func (v VersionType) Patch() int {
	return int(v % 100)
}
