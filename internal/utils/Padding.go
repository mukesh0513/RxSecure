package utils

import (
	"bytes"
	"encoding/hex"
	"strings"
)

//PKCS5Padding pads the ciphertext to make it suitable for encryption. Padding depends on length of cipherText
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//PKCS5UnPadding removes the padding from the cipherText to give back the original text.
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//PadKeyAndIvForAES... For AES Encryption, key length can be 16, 24 or 32. If the key length is more, we trim the key. Else we pad it with null
//bytes. Initialization vector length should be equal to block size, which is 16 for AES. We do the same operations on IV
//to make its length equal to 16.
func PadKeyAndIvForAES(key string, iv string, keyLength int) (string, string) {

	switch true {
	case keyLength <= 16:
		keyLength = 16
		break
	case keyLength <= 24:
		keyLength = 24
		break
	case keyLength <= 32:
		keyLength = 32
		break
	default:
		keyLength = 32
	}

	if len(key) > keyLength {
		key = key[0:keyLength]
	} else {
		for {
			if len(key) == keyLength {
				break
			}

			key += "\000"
		}
	}

	if len(iv) > 16 {
		iv = iv[0:16]
	} else {
		for {
			if len(iv) == 16 {
				break
			}

			iv += "\000"
		}
	}

	return key, iv
}

func PadKeyWithKeyForAes(key string, keyLength int) string {

	pad := key

	switch true {
	case keyLength <= 16:
		keyLength = 16
	case keyLength <= 24:
		keyLength = 24
	case keyLength <= 32:
		keyLength = 32
	default:
		keyLength = 32
	}

	if len(key) < keyLength {
		for len(key) < keyLength {
			key += pad
		}
	}

	return key[0:keyLength]
}

//PadKeyForTripleDes... For Triple DES Encryption, Key length must be 24. This function pads different key lengths to the valid key length required
//for triple des encryption.
func PadKeyForTripleDes(key string) string {
	keyLength := len(key)

	switch true {
	case keyLength <= 16:
		for {
			if len(key) == 24 {
				break
			}
			key += string(rune(0))
		}
		key = key[0:16] + key[0:8]
		break
	case keyLength <= 24:
		for {
			if len(key) == 24 {
				break
			}
			key += string(rune(0))
		}
		break
	default:
		key = key[0:24]
	}

	return key
}

//PadSbiFssEncData converts the plaintext so a specific format
func PadSbiFssEncData(plaintext string) string {
	var str string
	plaintext += strings.Repeat("^", len(plaintext)%8)

	for i := 0; i < len(plaintext); i++ {
		hexval := hex.EncodeToString([]byte(string(plaintext[i])))
		str += hexval
	}
	str = strings.ToUpper(str)
	return str
}

//PadSbiFssDecData converts the plaintext so a specific format
func PadSbiFssDecData(plaintext string) string {
	var str string

	for i := 0; i < len(plaintext); i += 2 {
		val, err := hex.DecodeString(plaintext[i : i+2])
		if err != nil {
			return "error"
		}
		str += string(val)
	}
	str = strings.ReplaceAll(str, "^", "")
	return str
}

func PadMessageWithIv(plaintext string, iv []byte) string {
	textToEncrypt := make([]byte, len(plaintext)+16)
	bytePlaintext := []byte(plaintext)

	copy(iv[:], textToEncrypt[:])
	copy(textToEncrypt[16:], bytePlaintext)

	return string(textToEncrypt)
}

func UnpadMessageWithIv(plaintext string, iv []byte) string {
	textToDecrypt := make([]byte, len(plaintext)-16)
	bytePlaintext := []byte(plaintext)

	copy(textToDecrypt[:], bytePlaintext[16:])

	return string(textToDecrypt)
}
