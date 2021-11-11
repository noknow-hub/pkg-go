//////////////////////////////////////////////////////////////////////
// crypto.go
//////////////////////////////////////////////////////////////////////
package crypto

import (
    "encoding/json"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
    "io"
    "strings"
    "time"
)

type Token struct {
    Data string
    Expires int64
}


//////////////////////////////////////////////////
// Encrypt using CBC mode.
//////////////////////////////////////////////////
func EncryptCBC(plainText, secretKey string) (string, error) {
    p := []byte(padLeft16Times(plainText))
    s := []byte(padLeft16Times(secretKey))
    block, err := aes.NewCipher(s)
    if err != nil {
        return "", err
    }
    cipherText := make([]byte, aes.BlockSize + len(p))
    iv := cipherText[:aes.BlockSize]
    _, err = io.ReadFull(rand.Reader, iv)
    if err != nil {
        return "", err
    }
    cbc := cipher.NewCBCEncrypter(block, iv)
    cbc.CryptBlocks(cipherText[aes.BlockSize:], p)
    return hex.EncodeToString(cipherText), nil
}


//////////////////////////////////////////////////
// Decrypt using CBC mode.
//////////////////////////////////////////////////
func DecryptCBC(cipherText, secretKey string) (string, error) {
    decoded, err := hex.DecodeString(cipherText)
    if err != nil {
        return "", err
    }
    c := []byte(decoded)
    s := []byte(padLeft16Times(secretKey))
    block, err := aes.NewCipher(s)
    if err != nil {
        return "", err
    }
    iv := c[:aes.BlockSize]
    decrypted := make([]byte, len(c[aes.BlockSize:]))

    cbc := cipher.NewCBCDecrypter(block, iv)
    cbc.CryptBlocks(decrypted, c[aes.BlockSize:])
    return strings.TrimLeft(string(decrypted), "0"), nil
}


//////////////////////////////////////////////////
// Verify using CBC mode.
//////////////////////////////////////////////////
func VerifyCBC(plainText, secretKey, cipherText string) bool {
    decrypted, err := DecryptCBC(cipherText, secretKey)
    if err != nil {
        return false
    }
    if plainText == decrypted {
        return true
    } else {
        return false
    }
}


//////////////////////////////////////////////////
// Encrypt using CTR mode.
//////////////////////////////////////////////////
func EncryptCTR(plainText, secretKey string) (string, error) {
    p := []byte(plainText)
    s := []byte(padLeft16Times(secretKey))
    block, err := aes.NewCipher(s)
    if err != nil {
        return "", err
    }
    cipherText := make([]byte, aes.BlockSize + len(p))
    iv := cipherText[:aes.BlockSize]
    _, err = io.ReadFull(rand.Reader, iv)
    if err != nil {
        return "", err
    }
    stream := cipher.NewCTR(block, iv)
    stream.XORKeyStream(cipherText[aes.BlockSize:], p)
    result := hex.EncodeToString(cipherText)
    return result, nil
}


//////////////////////////////////////////////////
// Decrypt using CTR mode.
//////////////////////////////////////////////////
func DecryptCTR(cipherText, secretKey string) (string, error) {
    decoded, err := hex.DecodeString(cipherText)
    if err != nil {
        return "", err
    }
    c := []byte(decoded)
    s := []byte(padLeft16Times(secretKey))
    block, err := aes.NewCipher(s)
    if err != nil {
        return "", err
    }
    decrypted := make([]byte, len(c[aes.BlockSize:]))
    stream := cipher.NewCTR(block, c[:aes.BlockSize])
    stream.XORKeyStream(decrypted, c[aes.BlockSize:])
    return string(decrypted), nil
}


//////////////////////////////////////////////////
// Verify using CTR mode.
//////////////////////////////////////////////////
func VerifyCTR(plainText, secretKey, cipherText string) bool {
    decrypted, err := DecryptCTR(cipherText, secretKey)
    if err != nil {
        return false
    }
    if plainText == decrypted {
        return true
    } else {
        return false
    }
}


//////////////////////////////////////////////////
// Encrypt using GCM mode.
//////////////////////////////////////////////////
func EncryptGCM(plainText, secretKey, authData string) (string, error) {
    p := []byte(plainText)
    s := []byte(pad32Bytes(secretKey))
    a := []byte(authData)
    block, err := aes.NewCipher(s)
    if err != nil {
        return "", err
    }
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }
    iv := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }
    cipherText := gcm.Seal(nil, iv, p, a)
    return hex.EncodeToString(append(cipherText, iv...)), nil
}


//////////////////////////////////////////////////
// Decrypt using GCM mode.
//////////////////////////////////////////////////
func DecryptGCM(cipherText, secretKey, authData string) (string, error) {
    decoded, err := hex.DecodeString(cipherText)
    if err != nil {
        return "", err
    }
    c := []byte(decoded)
    s := []byte(pad32Bytes(secretKey))
    a := []byte(authData)
    block, err := aes.NewCipher(s)
    if err != nil {
        return "", err
    }
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }
    plainText, err := gcm.Open(nil, c[len(c)-gcm.NonceSize():], c[:len(c)-gcm.NonceSize()], a)
    if err != nil {
        return "", err
    }
    return string(plainText), nil
}


//////////////////////////////////////////////////
// Verify using GCM mode.
//////////////////////////////////////////////////
func VerifyGCM(plainText, secretKey, cipherText, authData string) bool {
    decrypted, err := DecryptGCM(cipherText, secretKey, authData)
    if err != nil {
        return false
    }
    if plainText == decrypted {
        return true
    } else {
        return false
    }
}


//////////////////////////////////////////////////
// Generate a token with time limit using CBC mode.
//////////////////////////////////////////////////
func GenTokenCBC(plainText string, secondLimit int, secretKey string) (string, error) {
    now := time.Now()
    limit := now.Add(time.Second * time.Duration(secondLimit))
    token := Token{
        Data: plainText,
        Expires: limit.Unix(),
    }
    b, err := json.Marshal(token)
    if err != nil {
        return "", err
    }
    return EncryptCBC(string(b), secretKey)
}


//////////////////////////////////////////////////
// Verify a token with time limit using CBC mode.
//////////////////////////////////////////////////
func VerifyTokenCBC(plainText, cipherText, secretKey string) bool {
    decrypted, err := DecryptCBC(cipherText, secretKey)
    if err != nil {
        return false
    }
    var token Token
    if err = json.Unmarshal([]byte(decrypted), &token); err != nil {
        return false
    }
    if plainText != token.Data {
        return false
    }
    now := time.Now().Unix()
    return token.Expires > now
}


//////////////////////////////////////////////////
// Generate a token with time limit using CTR mode.
//////////////////////////////////////////////////
func GenTokenCTR(plainText string, secondLimit int, secretKey string) (string, error) {
    now := time.Now()
    limit := now.Add(time.Second * time.Duration(secondLimit))
    token := Token{
        Data: plainText,
        Expires: limit.Unix(),
    }
    b, err := json.Marshal(token)
    if err != nil {
        return "", err
    }
    return EncryptCTR(string(b), secretKey)
}


//////////////////////////////////////////////////
// Verify a token with time limit using CTR mode.
//////////////////////////////////////////////////
func VerifyTokenCTR(plainText, cipherText, secretKey string) bool {
    decrypted, err := DecryptCTR(cipherText, secretKey)
    if err != nil {
        return false
    }
    var token Token
    if err = json.Unmarshal([]byte(decrypted), &token); err != nil {
        return false
    }
    if plainText != token.Data {
        return false
    }
    now := time.Now().Unix()
    return token.Expires > now
}


//////////////////////////////////////////////////
// 0 Padding with 16 times from left side.
// @param text: [string] The text to pad with 0.
//////////////////////////////////////////////////
func padLeft16Times(text string) string {
    padCnt := aes.BlockSize - len(text) % aes.BlockSize
    if padCnt % aes.BlockSize == 0 {
        return text
    } else {
        return strings.Repeat("0", padCnt) + text
    }
}


//////////////////////////////////////////////////
// 0 Padding if text is less than 32 bytes.
// Cut out if text exceeds 32 bytes.
//////////////////////////////////////////////////
func pad32Bytes(text string) string {
    if len(text) == 32 {
        return text
    } else {
        if len(text) > 32 {
            return text[:32]
        } else {
            return strings.Repeat("0", 32 - len(text)) + text
        }
    }
}
