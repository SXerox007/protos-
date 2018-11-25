protoc -I advance/ --go_out=advance/ advance/runprotobasic/proto/basic.proto

//for runprotobasic file 
go run advance/runprotobasic/main.go

Note:- 
For changing the proto buffer file name use:
option go_package = "basic_newpb";
package basic;