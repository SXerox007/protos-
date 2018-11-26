package main

import (
	"fmt"
	"protos-/expert/oneof-and-maps/proto"

	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func main() {
	Init()
}

func Init() {
	pair := contact.Dictinory{
		Timestamp: &timestamp.Timestamp{
			Seconds: 1000,
		},
		Duration: &duration.Duration{
			Seconds: 1000,
		},
	}
	fmt.Println("data:-", pair)
}
