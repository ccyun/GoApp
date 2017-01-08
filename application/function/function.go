package function

import (
	"crypto/md5"
	"fmt"
	"io"
)

//Md5 字符串的MD5值
func Md5(str string, size int) string {
	if size == 0 || size > 32 {
		size = 32
	}
	size--
	t := md5.New()
	io.WriteString(t, str)
	s := fmt.Sprintf("%x", t.Sum(nil))
	start := (32 - size) / 2
	end := 32 - start
	s = s[start:end]
	return s
}

//ReverseString 反转字符串
func ReverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
