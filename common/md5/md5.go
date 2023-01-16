package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// MdFiveString 生成md5字符串
func MdFiveString(password string) (text string, err error) {
	m5 := md5.New()
	_, err = m5.Write([]byte(password))
	if err != nil {
		return "", err
	}
	text = hex.EncodeToString(m5.Sum(nil))
	return text, nil
}
