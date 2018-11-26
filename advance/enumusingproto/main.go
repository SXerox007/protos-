package main

import (
	"log"
	"protos-/advance/enumusingproto/proto"
)

func main() {
	Init()
}

func Init() {
	setDummyData()
}

func setDummyData() {
	data := enumexample.CompileData{
		Check:     enumexample.IsSuccessFull_FALSE,
		ErrorCode: 404,
	}

	log.Println("Show Data using Enum", data)
}
