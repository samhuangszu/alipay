package util

import (
	"crypto/sha1"
	"fmt"
)

// Sha1 获取散列
func Sha1(key string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(key)))
}
