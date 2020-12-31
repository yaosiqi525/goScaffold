package utils

import (
	"crypto/md5"
	"fmt"
)

// 计算Md5
func GetMd5String(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}
