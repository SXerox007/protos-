protoc -I advance/ --go_out=advance/ advance/readwritefileusingproto/proto/userdata.proto

//for runprotobasic file 
go run advance/readwritefileusingproto/main.go

Note:- 
For changing the proto buffer file name use:
option go_package = "basic_newpb";
package basic;