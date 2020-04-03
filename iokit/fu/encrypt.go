package fu

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"io"
)

func Encrypt(password string, data []byte) ([]byte, error) {
	key := sha256.Sum256([]byte(password))
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}
	bf := bytes.Buffer{}
	bf.Write(iv)
	wr := &cipher.StreamWriter{S: cipher.NewOFB(block, iv), W: &bf}
	if err := binary.Write(wr, binary.LittleEndian, uint32(len(data))); err != nil {
		return nil, err
	}
	if _, err := io.Copy(wr, bytes.NewReader(data)); err != nil {
		return nil, err
	}
	_ = wr.Close()
	return bf.Bytes(), nil
}
