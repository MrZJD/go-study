package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?#*&()'-=@~"

	sEnc := base64.StdEncoding.EncodeToString([]byte(data)) // 标准编码

	fmt.Println(sEnc)

	sDec, _ := base64.StdEncoding.DecodeString(sEnc) // 标准解码

	fmt.Println(string(sDec))

	// 兼容URL编解码

	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)

	uDesc, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDesc))
}
