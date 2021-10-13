package aespkcs5

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
)

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	if len(ciphertext) == 0 {
		return []byte("")
	}

	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	if len(encrypt) == 0 {
		return []byte("")
	}

	srcLen := len(encrypt)
	paddingLen := int(encrypt[srcLen-1])
	return encrypt[:srcLen-paddingLen]
}

func Encrypt(k, s string) (encrypted []byte, err error) {
	key := []byte(k)
	src := []byte(s)

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	src = PKCS5Padding(src, block.BlockSize())
	encryptedByte := make([]byte, len(src))

	ecb := NewECBEncrypter(block)
	ecb.CryptBlocks(encryptedByte, src)

	encrypted = []byte(base64.StdEncoding.EncodeToString(encryptedByte))
	//encrypted = []byte(strings.ToUpper(hex.EncodeToString(encryptedByte)))

	return
}

func Decrypt(k, s string) (decrypted []byte, err error) {
	baseDecodeStr, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return
	}

	key := []byte(k)
	//src := []byte(s)
	src := baseDecodeStr

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	decrypted = make([]byte, len(src))

	ecb := NewECBDecrypter(block)
	ecb.CryptBlocks(decrypted, src)
	decrypted = PKCS5Trimming(decrypted)
	return
}
