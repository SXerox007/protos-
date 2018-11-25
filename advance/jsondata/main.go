package main

import (
	"io/ioutil"
	"log"
	"protos-/advance/readwritefileusingproto/proto"

	"github.com/golang/protobuf/jsonpb"
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
	err := writeFile("readwrite.bin", data)
	if err == nil {
		err, fileData := readFile("readwrite.bin")
		if err == nil {
			str, err := toJson(fileData)
			if err == nil {
				log.Println("Json Data:-", str)
			}
		}
	}
}

func toJson(pb proto.Message) (string, error) {
	marshal := jsonpb.Marshaler{}
	return marshal.MarshalToString(pb)
}

func fromJson(jsonString string, pb proto.Message) error {
	return jsonpb.UnmarshalString(jsonString, pb)
}

func writeFile(filename string, pb proto.Message) error {
	arr, _ := proto.Marshal(pb)
	err := ioutil.WriteFile(filename, arr, 0644)
	if err == nil {
		log.Println("Write Successfully")
	}
	return err
}

func readFile(filename string) (error, *userdata.UserData) {
	pb := &userdata.UserData{}
	arr, err := ioutil.ReadFile(filename)
	err = proto.Unmarshal(arr, pb)
	return err, pb

}
