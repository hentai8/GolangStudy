package main

import (
	"encoding/hex"
	"fmt"
	"lukechampine.com/blake3"
	"math/big"
	"strconv"
)

func main() {

	target := "00b2f4fc0794908cf232ff78625902416a7530755a66c5788c4a6d5331471111"
	targetByte, _ := hex.DecodeString(target)
	targetInt := new(big.Int).SetBytes(targetByte)
	//targetInt, _ := strconv.ParseInt(target, 10, 64)
	fmt.Println("targetInt:", targetInt)

	header := "0000000000000000cf0e020000000000000000000002aaced825176dd9db0701c995760a03a1f42c69b63b4b7d4090b0ff7f32477b07a0cc3c89d6f6335433def2d95ff91be838ae47212ba43794901bb0ce220200000000f6ee7f75663920ae6d8617379629d5130323e6e20c5e19cb5606c71bb97ed7e668d5130100000000000000000007b87e00ba71e3b4a9a27d79dad30a55297da63550092644b289502c8efe8f82010000000000007736f4a168656e7461693800000000000000000000000000000000000000000000000000"
	randomness := header[:16]
	//randomness := 0x0000000000000000
	headerWithoutCalculate := header[16:]
	i := 0
	for {
		print("???", i)
		//randomnessHex, _ := hex.DecodeString(randomness)
		fmt.Println("...randomness:", randomness)
		randomnessByte, err := hex.DecodeString(randomness)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("...randomnessByte:", randomnessByte)
		randomnessInt := new(big.Int).SetBytes(randomnessByte)
		fmt.Println("...randomnessInt:", randomnessInt)
		//randomnessInt, _ := strconv.ParseInt(randomness, 16, 64)
		randomnessInt.Add(randomnessInt, big.NewInt(1))
		//randomnessHex = randomnessHex + 1
		randomness = randomnessInt.String()
		randomnessInt64, _ := strconv.ParseInt(randomness, 10, 64)
		fmt.Println("!!!randomnessInt64:", randomnessInt64)
		randomness = fmt.Sprintf("%016d", randomnessInt64)
		fmt.Println("!!!randomness:", randomness)

		string256 := randomness + headerWithoutCalculate

		byte256 := []byte(string256)
		hash256 := blake3.Sum256(byte256)
		fmt.Println(hash256)
		answer := hex.EncodeToString(hash256[:])
		fmt.Println(answer)
		answerByte, _ := hex.DecodeString(answer)
		answerInt := new(big.Int).SetBytes(answerByte)
		//answerInt, _ := strconv.ParseInt(answer, 16, 64)

		fmt.Println("answer:", answer)
		fmt.Println("target:", target)
		fmt.Println("answerInt:", answerInt)
		fmt.Println("targetInt:", targetInt)
		if answerInt.Cmp(targetInt) == -1 {
			fmt.Println("find answer:", answerInt)
			break
		}
		i++
	}

	////进行256位的blake3加密，256位的加密就是512位加密的前半段
	//string256 := `string256`
	//byte256 := []byte(string256)
	//hash256 := blake3.Sum256(byte256)
	//fmt.Println(hash256)
	//fmt.Println(hex.EncodeToString(hash256[:]))
	//
	////进行256位的blake3加密
	//string512 := `string256`
	//byte512 := []byte(string512)
	//hash512 := blake3.Sum512(byte512)
	//fmt.Println(hash512)
	//fmt.Println(hex.EncodeToString(hash512[:]))
	//
	//// 新建一个blake3加密通道, size为32即Sum256
	//hasher := blake3.New(23, nil)
	//selfHash := hasher.Sum([]byte(string256))
	//fmt.Println(selfHash)
	//fmt.Println(hex.EncodeToString(selfHash))
}
