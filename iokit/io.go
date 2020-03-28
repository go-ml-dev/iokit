package iokit

import (
	"io"
)

type Input interface {
	Open() (io.ReadCloser, error)
}

type Output interface {
	Create() (Whole, error)
}

type Inout interface {
	Input
	Output
}
