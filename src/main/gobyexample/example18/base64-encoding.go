package main

import (
	"encoding/base64"
	"fmt"
)

// Base64编码
func main() {

	var p = fmt.Println

	data := "abc123!?$*&()'-=@~"

	// 标准编解码
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	p(sEnc)
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	p(string(sDec))
	p()

	// 使用url兼容的base64格式进行呢编解码
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	p(uEnc)
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	p(string(uDec))
}
