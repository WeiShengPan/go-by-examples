package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"testing"
)

// SHA1散列
// 生成二进制文件或者文本块的短标识
func TestSha1Hashes(t *testing.T) {

	s := "sha1 this string"

	fmt.Println(s)

	// SHA1 值经常以 16 进制输出，例如在 git commit 中。使用`%x` 来将散列结果格式化为 16 进制字符串。
	fmt.Printf("%x\n", hashBySha1(s))

	fmt.Printf("%x\n", hashByMd5(s))

}

func hashBySha1(s string) []byte {

	h := sha1.New()

	// 写入要处理的字节，如果是字符串，需要用[]byte()强转
	h.Write([]byte(s))

	// 用于得到最终的散列值的字符切片。
	bs := h.Sum(nil)

	return bs
}

func hashByMd5(s string) []byte {

	h := md5.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	return bs
}
