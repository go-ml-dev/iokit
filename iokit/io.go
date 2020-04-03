package iokit

import (
	"bytes"
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

func (iox LuckyInputOutput) ReadAll() ([]byte, error) {
	return LuckyInput{iox.InputOutput}.ReadAll()
}

func (iox LuckyInputOutput) LuckyReadAll() []byte {
	return LuckyInput{iox.InputOutput}.LuckyReadAll()
}

func (iox LuckyInputOutput) WriteAll(bs []byte) error {
	return LuckyOutput{iox.InputOutput}.WriteAll(bs)
}

func (iox LuckyInputOutput) LuckyWriteAll(bs []byte) {
	LuckyOutput{iox.InputOutput}.LuckyWriteAll(bs)
}

type LuckyOutput struct{ Output }

func (iox LuckyOutput) LuckyCreate() LuckyWhole {
	wr, err := iox.Create()
	if err != nil {
		panic(err)
	}
	return LuckyWhole{wr}
}

func (iox LuckyOutput) WriteAll(bs []byte) error {
	wr, err := iox.Create()
	if err != nil {
		return err
	}
	defer wr.End()
	_, err = io.Copy(wr, bytes.NewReader(bs))
	if err != nil {
		return err
	}
	return wr.Commit()
}

func (iox LuckyOutput) LuckyWriteAll(bs []byte) {
	if err := iox.WriteAll(bs); err != nil {
		panic(err)
	}
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

func (iox LuckyInput) ReadAll() ([]byte, error) {
	rd, err := iox.Open()
	if err != nil {
		return nil, err
	}
	defer rd.Close()
	return ioutil.ReadAll(rd)
}

func (iox LuckyInput) LuckyReadAll() []byte {
	bs, err := iox.ReadAll()
	if err != nil {
		panic(err)
	}
	return bs
}

type LuckyReader struct{ io.ReadCloser }

func (lr LuckyReader) LuckyReadAll() []byte {
	bs, err := ioutil.ReadAll(lr)
	if err != nil {
		panic(err)
	}
	return bs
}
