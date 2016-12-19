package function

import (
	"crypto/md5"
	"fmt"
	"io"
)

//Md5 字符串的MD5值
func Md5(str string) string {
	t := md5.New()
	io.WriteString(t, str)
	return fmt.Sprintf("%x", t.Sum(nil))
}
