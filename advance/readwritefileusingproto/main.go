package main

import (
	"io/ioutil"
	"log"
	"protos-/advance/readwritefileusingproto/proto"

	"github.com/golang/protobuf/proto"
)

func main() {
	Init()
}

func dummyData() *userdata.UserData {
	userAddress := []*userdata.UserData_Address{}
	address := &userdata.UserData_Address{
		AddressLine1: "abc",
		AddressLine2: "cde",
		AddressLine3: "tyu",
		City:         "Shimla",
		Country:      "India",
	}
	userAddress = append(userAddress, address)
	user := &userdata.UserData{
		Name:        "Sumit",
		Age:         22,
		PhoneNumber: []string{"123456", "543210"},
		Birthday: &userdata.Date{
			Day:   1,
			Month: 1,
			Year:  2000,
		},
		UserAddress: userAddress,
	}
	return user
}

func Init() {
	data := dummyData()
	log.Println("User Data:-", data)
	writeFile("readwrite.bin", data)
}

func writeFile(filename string, pb proto.Message) error {
	arr, _ := proto.Marshal(pb)
	err := ioutil.WriteFile(filename, arr, 0644)
	if err == nil {
		log.Println("Write Successfully")
	}
	return err
}
