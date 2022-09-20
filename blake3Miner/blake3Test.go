package main

import (
	"encoding/hex"
	"fmt"
	"lukechampine.com/blake3"
)

func main() {

	// subkey为子秘钥
	// ctx为上下文
	// srckey为
	var subKey []byte
	var ctx string
	var srcKey []byte
	subKey = []byte(`nmsl`)
	ctx = `nmsl`
	srcKey = []byte(`nmsl`)
	fmt.Println(subKey)
	fmt.Println(ctx)
	fmt.Println(srcKey)
	// 产生一个秘钥subKey
	blake3.DeriveKey(subKey, ctx, srcKey)
	fmt.Println(subKey)
	fmt.Println(ctx)
	fmt.Println(srcKey)
	fmt.Println(string(subKey))
	fmt.Println(string(srcKey))

	//进行256位的blake3加密，256位的加密就是512位加密的前半段
	string256 := `string256`
	byte256 := []byte(string256)
	hash256 := blake3.Sum256(byte256)
	fmt.Println(hash256)
	fmt.Println(hex.EncodeToString(hash256[:]))

	//进行256位的blake3加密
	string512 := `string256`
	byte512 := []byte(string512)
	hash512 := blake3.Sum512(byte512)
	fmt.Println(hash512)
	fmt.Println(hex.EncodeToString(hash512[:]))

	// 新建一个blake3加密通道, size为32即Sum256
	hasher := blake3.New(23, nil)
	selfHash := hasher.Sum([]byte(string256))
	fmt.Println(selfHash)
	fmt.Println(hex.EncodeToString(selfHash))

}
