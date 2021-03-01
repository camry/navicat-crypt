package main

import (
	"fmt"
	"navicat-cli/ncx"
)

func main() {
	e := ncx.Encrypt("navicat123456")
	println(fmt.Sprintf("Navicat 加密:%s", e))

	s := ncx.Decrypt(e)
	println(fmt.Sprintf("Navicat 解密:%s", s))
}
