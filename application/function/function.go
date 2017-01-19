package function

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
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

//MakeRowkey 反转字符串
func MakeRowkey(num int64) string {
	seed := [36]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	numstr := strconv.FormatInt(num, 10)
	reverseNum, _ := strconv.ParseInt(ReverseString(numstr), 10, 0)
	seedK1 := num % 36
	seedK2 := reverseNum % 36
	seedK3 := (seedK1 + seedK2) % 36
	return seed[seedK3] + seed[seedK1] + seed[seedK2] + "_" + numstr
}
