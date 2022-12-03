package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"message-board/model"
)

func Encrypt(raw string) (res string, err error) {
	has := md5.New()
	_, err = io.WriteString(has, raw)
	if err != nil {
		log.Printf("Write failed:%v", err)
		return
	}
	tem := has.Sum(model.MyKey)
	res = hex.EncodeToString(tem)
	return
}
