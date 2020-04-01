package iokit

import (
	"io"
	"io/ioutil"
)

type Input interface {
	Open() (io.ReadCloser, error)
}

type Output interface {
	Create() (Whole, error)
}

type InputOutput interface {
	Input
	Output
}

type LuckyInputOutput struct{ InputOutput }

func (iox LuckyInputOutput) LuckyCreate() LuckyWhole {
	return LuckyOutput{iox.InputOutput}.LuckyCreate()
}

func (iox LuckyInputOutput) LuckyOpen() LuckyReader {
	return LuckyInput{iox.InputOutput}.LuckyOpen()
}

type LuckyOutput struct{ Output }

func (iox LuckyOutput) LuckyCreate() LuckyWhole {
	wr, err := iox.Create()
	if err != nil {
		panic(err)
	}
	return LuckyWhole{wr}
}

type LuckyWhole struct{ Whole }

func (lw LuckyWhole) LuckyWrite(b []byte) {
	if _, err := lw.Write(b); err != nil {
		panic(err)
	}
}

func (lr LuckyWhole) LuckyCommit() {
	if err := lr.Commit(); err != nil {
		panic(err)
	}
}

type LuckyInput struct{ Input }

func (iox LuckyInput) LuckyOpen() LuckyReader {
	rd, err := iox.Open()
	if err != nil {
		panic(err)
	}
	return LuckyReader{rd}
}

type LuckyReader struct{ io.ReadCloser }

func (lr LuckyReader) LuckyReadAll() []byte {
	bs, err := ioutil.ReadAll(lr)
	if err != nil {
		panic(err)
	}
	return bs
}
