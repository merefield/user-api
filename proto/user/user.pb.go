// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: user/user.proto

//package example;
//package resonate.api.user;

package user

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_user_user_proto protoreflect.FileDescriptor

var file_user_user_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xf3, 0x0f, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x77, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x49, 0x92, 0x41, 0x2e, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x0a, 0x41, 0x64,
	0x64, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x19, 0x41, 0x64, 0x64, 0x20, 0x61, 0x20,
	0x75, 0x73, 0x65, 0x72, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x98, 0x01, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x64, 0x92, 0x41, 0x45, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x2d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x20,
	0x75, 0x73, 0x65, 0x72, 0x20, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x20, 0x6f, 0x6e, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16,
	0x32, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0xeb, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x12,
	0x21, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0xa2, 0x01, 0x92, 0x41, 0x78, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x1a, 0x4e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x20, 0x6f, 0x6e,
	0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x64,
	0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x21, 0x32, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73,
	0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x99, 0x01, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x61, 0x92,
	0x41, 0x45, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x0a, 0x47, 0x65, 0x74, 0x20, 0x61,
	0x20, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x30, 0x47, 0x65, 0x74, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65,
	0x72, 0x27, 0x73, 0x20, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0xd8, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x94, 0x01, 0x92, 0x41, 0x6d, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x23, 0x47, 0x65, 0x74, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x27, 0x73, 0x20,
	0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x3f, 0x47, 0x65, 0x74, 0x20, 0x75, 0x73, 0x65, 0x72,
	0x20, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x69,
	0x6e, 0x67, 0x20, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65,
	0x64, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x89, 0x01, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x5b, 0x92, 0x41, 0x34, 0x0a,
	0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x75,
	0x73, 0x65, 0x72, 0x1a, 0x1e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x61, 0x20, 0x75, 0x73,
	0x65, 0x72, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x2a, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x65, 0x64, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x7c, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4a, 0x92, 0x41, 0x32, 0x0a, 0x05,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x1a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x98, 0x01, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x5d, 0x92, 0x41, 0x3e, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x10, 0x41, 0x64, 0x64, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x1a, 0x1f, 0x41, 0x64, 0x64, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x3a, 0x01, 0x2a,
	0x12, 0xb7, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x79, 0x92, 0x41, 0x55, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x1a, 0x33, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x61, 0x6e, 0x20,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x20, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x32,
	0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0xa7, 0x01, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x16, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x60, 0x92, 0x41, 0x3f, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x10, 0x47, 0x65, 0x74, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x1a, 0x20, 0x47, 0x65, 0x74, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x9b, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x63, 0x92,
	0x41, 0x42, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x10, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x1a,
	0x23, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x2a, 0x16, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0xc2, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7f, 0x92, 0x41, 0x5d, 0x0a, 0x0a, 0x55, 0x73, 0x65,
	0x72, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x1a,
	0x37, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x20, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x63, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0xa1, 0x02, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x63,
	0x6f, 0x6f, 0x70, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x92, 0x41, 0xf0, 0x01, 0x12, 0x2d, 0x0a, 0x24, 0x52,
	0x65, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x55,
	0x73, 0x65, 0x72, 0x32, 0x05, 0x32, 0x2e, 0x30, 0x2e, 0x32, 0x2a, 0x01, 0x02, 0x5a, 0x59, 0x0a,
	0x57, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x4d, 0x08, 0x02, 0x12, 0x38, 0x41,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x64, 0x20, 0x62, 0x79,
	0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20,
	0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x0a, 0x0a, 0x08, 0x0a, 0x06, 0x62, 0x65,
	0x61, 0x72, 0x65, 0x72, 0x72, 0x55, 0x0a, 0x29, 0x67, 0x52, 0x50, 0x43, 0x2d, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x2d, 0x75, 0x73,
	0x65, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x20, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x63, 0x6f,
	0x6f, 0x70, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_user_user_proto_goTypes = []interface{}{
	(*UserAddRequest)(nil),              // 0: user.UserAddRequest
	(*UserUpdateRequest)(nil),           // 1: user.UserUpdateRequest
	(*UserUpdateRestrictedRequest)(nil), // 2: user.UserUpdateRestrictedRequest
	(*UserRequest)(nil),                 // 3: user.UserRequest
	(*Empty)(nil),                       // 4: user.Empty
	(*UserGroupCreateRequest)(nil),      // 5: user.UserGroupCreateRequest
	(*UserGroupUpdateRequest)(nil),      // 6: user.UserGroupUpdateRequest
	(*UserGroupRequest)(nil),            // 7: user.UserGroupRequest
	(*UserPublicResponse)(nil),          // 8: user.UserPublicResponse
	(*UserPrivateResponse)(nil),         // 9: user.UserPrivateResponse
	(*UserListResponse)(nil),            // 10: user.UserListResponse
	(*UserGroupPublicResponse)(nil),     // 11: user.UserGroupPublicResponse
	(*UserGroupListResponse)(nil),       // 12: user.UserGroupListResponse
}
var file_user_user_proto_depIdxs = []int32{
	0,  // 0: user.ResonateUser.AddUser:input_type -> user.UserAddRequest
	1,  // 1: user.ResonateUser.UpdateUser:input_type -> user.UserUpdateRequest
	2,  // 2: user.ResonateUser.UpdateUserRestricted:input_type -> user.UserUpdateRestrictedRequest
	3,  // 3: user.ResonateUser.GetUser:input_type -> user.UserRequest
	3,  // 4: user.ResonateUser.GetUserRestricted:input_type -> user.UserRequest
	3,  // 5: user.ResonateUser.DeleteUser:input_type -> user.UserRequest
	4,  // 6: user.ResonateUser.ListUsers:input_type -> user.Empty
	5,  // 7: user.ResonateUser.AddUserGroup:input_type -> user.UserGroupCreateRequest
	6,  // 8: user.ResonateUser.UpdateUserGroup:input_type -> user.UserGroupUpdateRequest
	7,  // 9: user.ResonateUser.GetUserGroup:input_type -> user.UserGroupRequest
	7,  // 10: user.ResonateUser.DeleteUserGroup:input_type -> user.UserGroupRequest
	3,  // 11: user.ResonateUser.ListUsersGroups:input_type -> user.UserRequest
	4,  // 12: user.ResonateUser.AddUser:output_type -> user.Empty
	4,  // 13: user.ResonateUser.UpdateUser:output_type -> user.Empty
	4,  // 14: user.ResonateUser.UpdateUserRestricted:output_type -> user.Empty
	8,  // 15: user.ResonateUser.GetUser:output_type -> user.UserPublicResponse
	9,  // 16: user.ResonateUser.GetUserRestricted:output_type -> user.UserPrivateResponse
	4,  // 17: user.ResonateUser.DeleteUser:output_type -> user.Empty
	10, // 18: user.ResonateUser.ListUsers:output_type -> user.UserListResponse
	4,  // 19: user.ResonateUser.AddUserGroup:output_type -> user.Empty
	4,  // 20: user.ResonateUser.UpdateUserGroup:output_type -> user.Empty
	11, // 21: user.ResonateUser.GetUserGroup:output_type -> user.UserGroupPublicResponse
	4,  // 22: user.ResonateUser.DeleteUserGroup:output_type -> user.Empty
	12, // 23: user.ResonateUser.ListUsersGroups:output_type -> user.UserGroupListResponse
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_user_user_proto_init() }
func file_user_user_proto_init() {
	if File_user_user_proto != nil {
		return
	}
	file_user_common_proto_init()
	file_user_user_messages_proto_init()
	file_user_usergroup_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_user_proto_goTypes,
		DependencyIndexes: file_user_user_proto_depIdxs,
	}.Build()
	File_user_user_proto = out.File
	file_user_user_proto_rawDesc = nil
	file_user_user_proto_goTypes = nil
	file_user_user_proto_depIdxs = nil
}
