package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"hash"
)
//笔记纯当练习用
func main() {
	res := SHA256DoubleString("123456", true)
	fmt.Println(res)
}


func HASH(text string, hashType string, isHex bool) string {
	var hashInstance hash.Hash
	switch hashType {
	case "md4":
		hashInstance = md4.New()
	case "md5":
		hashInstance = md5.New()
	case "sha1":
		hashInstance = sha1.New()
	case "ripemd160":
		hashInstance = ripemd160.New()
	case "sha256":
		hashInstance = sha256.New()
	case "512":
		hashInstance = sha512.New()
	}
	if isHex {
		arr, _ := hex.DecodeString(text)
		hashInstance.Write(arr)
	} else {
		hashInstance.Write([]byte(text))
	}
	bytes := hashInstance.Sum(nil)
	return fmt.Sprintf("%x", bytes)
}

func SHA256Double(text string, isHex bool)[]byte  {
	hashInstance:=sha256.New()
	if isHex {
		arr,_:=hex.DecodeString(text)
		hashInstance.Write(arr)
	}else {
		hashInstance.Write([]byte(text))
	}
	bytes:=hashInstance.Sum(nil)
	hashInstance.Reset()//双hash
	hashInstance.Write(bytes)
	bytes=hashInstance.Sum(nil)
	return bytes
}

func SHA256DoubleString (text string, isHex bool)string {
      bytes:=SHA256Double(text,isHex)
      return fmt.Sprintf("%x",bytes)
}