package navicat

import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "encoding/hex"
    "strings"
)

var (
    key = []byte("libcckeylibcckey")
    iv  = []byte("libcciv libcciv ")
)

// Encrypt Navicat 加密。
func Encrypt(s string) string {
    block, _ := aes.NewCipher(key)
    blockSize := block.BlockSize()

    cipherText := []byte(s)
    paddedText := PKCS7Padding(cipherText, blockSize)

    cryted := make([]byte, len(paddedText))

    blockMode := cipher.NewCBCEncrypter(block, iv)
    blockMode.CryptBlocks(cryted, paddedText)

    return strings.ToUpper(hex.EncodeToString(cryted))
}

// Decrypt Navicat 解密。
func Decrypt(cryted string) string {
    crytedByte, _ := hex.DecodeString(cryted)
    block, _ := aes.NewCipher(key)

    origData := make([]byte, len(crytedByte))

    blockMode := cipher.NewCBCDecrypter(block, iv)
    blockMode.CryptBlocks(origData, crytedByte)

    return string(PKCS7UnPadding(origData))
}

// PKCS7Padding 补码
// AES加密数据块分组长度必须为128bit(byte[16])，密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个。
func PKCS7Padding(cipherText []byte, blockSize int) []byte {
    padding := blockSize - (len(cipherText) % blockSize)
    padText := bytes.Repeat([]byte{byte(padding)}, padding)

    return append(cipherText, padText...)
}

// PKCS7UnPadding 去码
func PKCS7UnPadding(origData []byte) []byte {
    length := len(origData)
    unPadding := int(origData[length-1])

    return origData[:(length - unPadding)]
}
