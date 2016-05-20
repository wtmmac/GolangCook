package main

import (
	"fmt"
)

func main() {
	//opcodeByte := make([]byte, 1)
	opcodeByte := []byte("Fasdfasdfasdf")
	opcodeByte[0] = opcodeByte[0] | 127

	FIN := opcodeByte[0] >> 7
	RSV1 := opcodeByte[0] >> 6 & 1
	RSV2 := opcodeByte[0] >> 5 & 1
	RSV3 := opcodeByte[0] >> 4 & 1
	OPCODE := opcodeByte[0] & 15
	fmt.Printf("FIN:%x\n", FIN)
	fmt.Printf("RSV1:%x\n", RSV1)
	fmt.Printf("RSV2:%x\n", RSV2)
	fmt.Printf("RSV3:%x\n", RSV3)
	fmt.Printf("OPCODE:0x%01x\n", OPCODE)
	fmt.Printf("%02x\n", opcodeByte[0])
}
