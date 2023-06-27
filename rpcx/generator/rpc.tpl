syntax = "proto3";

option go_package = ".;msg";
package msg;

message LoginReq {
  string name = 1;
  string password = 2;
  string role = 3;
}

message LoginResp {
  string code = 1;
  string message = 2;
}

message ChatReq {
  string name = 1;
  string password = 2;
  string role = 3;
}

message ChatResp {
  string code = 1;
  string message = 2;
}

service {{.serviceName}} {
  // service
  rpc LoginService (LoginReq) returns (LoginResp);
  rpc ChatService (ChatReq) returns (ChatResp);
}