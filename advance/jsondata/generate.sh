protoc -I advance/ --go_out=advance/ advance/jsondata/proto/userdata.proto

//for runprotobasic file 
go run advance/jsondata/main.go

Note:- 
For changing the proto buffer file name use:
option go_package = "basic_newpb";
package basic;