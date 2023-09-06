package wUtil

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
)

type CryptModel struct {
	Key  []byte `json:"k"`
	Data []byte `json:"v"`
}

func EncryptJWT(preshareKey []byte, data []byte) (value string, err error) {
	// Generate new aes key and encrypt data with that key
	aesKey, err := Generate256AESKey()
	if err != nil {
		return
	}

	encryptedData, err := EncryptAES(aesKey, data)
	if err != nil {
		return
	}

	encryptedKey, err := RSAEncryptOAEP(preshareKey, aesKey)
	if err != nil {
		return
	}
	//---
	c := CryptModel{
		Key:  encryptedKey,
		Data: encryptedData,
	}
	//---
	dataJson, err := json.Marshal(c)
	if err != nil {
		return
	}
	value = base64.StdEncoding.EncodeToString(dataJson)
	//---
	return
}

func DecryptJWT(preshareKey []byte, jwtEncrypt string) (jwtStr *string, err error) {
	//---
	bodyData, _ := base64.StdEncoding.DecodeString(jwtEncrypt)

	var cryptModel = &CryptModel{}
	/* ============= decryptData ========= */
	err = json.Unmarshal(bodyData, cryptModel)
	if err != nil {
		return nil, err
	}

	//DecryptKey1
	aesKey, err := RSADecryptOAEP(cryptModel.Key, preshareKey)
	if err != nil {
		return nil, err
	}
	// Decrypt body with aesKey
	jwtData := DecryptAES(aesKey, cryptModel.Data)
	if err != nil {
		return nil, err
	}
	//---
	jwtoken := string(jwtData[:])
	return &jwtoken, nil
}

// GenerateAES256Key : generate an 256 bits AES key
func Generate256AESKey() (key []byte, err error) {
	b := make([]byte, 32)
	_, err = rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return key, err
	}
	return b, nil
}

// EncryptAES string to base64 crypto using AES
func EncryptAES(keyByte []byte, plaintext []byte) ([]byte, error) {

	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
	//return base64.URLEncoding.EncodeToString(ciphertext)
}

// DecryptAES :decrypt AES
func DecryptAES(keySt []byte, ciphertext []byte) (decrypted []byte) {
	block, err := aes.NewCipher(keySt)
	if err != nil {
		log.Println(err)
		return
		//panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		log.Println("ciphertext too short")
		return
		//panic()
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext
}

// NewAESKey : new random 256 aes key
func NewAESKey() (keyBuff []byte, err error) {
	keyBuff = make([]byte, 32)
	_, err = rand.Read(keyBuff)
	return
}
