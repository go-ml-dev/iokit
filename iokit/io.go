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

type Inout interface {
	Input
	Output
}

type InoutExt struct { Inout }

func (iox InoutExt) LuckyCreate() LuckyWhole {
	wr, err := iox.Create()
	if err != nil { panic(err) }
	return LuckyWhole{ wr }
}

type LuckyWhole struct { Whole }

func (lw LuckyWhole) LuckyWrite(b []byte) {
	if _, err := lw.Write(b); err != nil { panic(err) }
}

func (lr LuckyWhole) LuckyCommit() {
	if err := lr.Commit(); err != nil { panic(err) }
}

func (iox InoutExt) LuckyOpen() LuckyReader {
	rd, err := iox.Open()
	if err != nil { panic(err) }
	return LuckyReader{rd }
}

type LuckyReader struct { io.ReadCloser }

func (lr LuckyReader) LuckyReadAll() []byte {
	bs, err := ioutil.ReadAll(lr)
	if err != nil { panic(err) }
	return bs
}

func (lr LuckyReader) LuckyClose() {
	if err := lr.Close(); err != nil { panic(err) }
}
