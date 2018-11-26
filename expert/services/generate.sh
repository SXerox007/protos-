protoc -I expert/ --go_out=expert/ expert/services/proto/search.proto

//for runprotobasic file 
go run expert/services/main.go

Note:- 
For changing the proto buffer file name use:
option go_package = "basic_newpb";
package basic;