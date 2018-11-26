protoc -I expert/ --go_out=expert/ expert/oneof-and-maps/proto/contact.proto

//for runprotobasic file 
go run expert/oneof-and-maps/main.go

Note:- 
For changing the proto buffer file name use:
option go_package = "basic_newpb";
package basic;