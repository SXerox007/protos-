package main

import (
	"log"
	"protos-/advance/runprotobasic/proto"
)

func main() {
	Init()

}
func Init() {
	data := basic.MessageExample{
		MessageString: "HelloWorld",
		Content:       1,
	}

	log.Println("Data Proto:-", data)
	log.Println("GetContent", data.GetContent(), data.GetMessageString())
}
