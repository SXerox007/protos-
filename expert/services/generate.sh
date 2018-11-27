protoc -I expert/ --go_out=expert/ expert/services/proto/search.proto

//To generate the grpc code
protoc expert/services/proto/search.proto --go_out=plugins=grpc:.

//for runprotobasic file 
go run expert/services/main.go

Note:- 
For changing the proto buffer file name use:
option go_package = "basic_newpb";
package basic;