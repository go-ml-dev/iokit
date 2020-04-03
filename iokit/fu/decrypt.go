package fu

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/binary"
	"io"
)

func Decrypt(password string, data []byte) ([]byte, error) {
	key := sha256.Sum256([]byte(password))
	ds := bytes.NewReader(data)
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(ds, iv); err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}
	bf := bytes.Buffer{}
	rd := &cipher.StreamReader{S: cipher.NewOFB(block, iv), R: ds}
	var ln uint32
	if err := binary.Read(rd, binary.LittleEndian, &ln); err != nil {
		return nil, err
	}
	if _, err := io.Copy(&bf, rd); err != nil {
		return nil, err
	}
	return bf.Bytes()[:int(ln)], nil
}
