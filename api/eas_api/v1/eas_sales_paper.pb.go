// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.3
// source: eas_api/v1/eas_sales_paper.proto

package v1

import (
	_ "eas_api/api/common"
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

var File_eas_api_v1_eas_sales_paper_proto protoreflect.FileDescriptor

var file_eas_api_v1_eas_sales_paper_proto_rawDesc = []byte{
	0x0a, 0x20, 0x65, 0x61, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x61, 0x73,
	0x5f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x70, 0x61, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x65, 0x61, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x27, 0x65, 0x61, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x61,
	0x73, 0x5f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9b, 0x15, 0x0a, 0x14, 0x45,
	0x61, 0x73, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x9f, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61,
	0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x65, 0x61, 0x73, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65,
	0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x65, 0x61, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x40, 0x92, 0x41, 0x1c, 0x0a, 0x0c, 0xe8, 0xaf, 0x95, 0xe5, 0x8d, 0xb7,
	0xe6, 0xa8, 0xa1, 0xe5, 0x9d, 0x97, 0x12, 0x0c, 0xe5, 0x88, 0x9b, 0xe5, 0xbb, 0xba, 0xe8, 0xaf,
	0x95, 0xe5, 0x8d, 0xb7, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0xae, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c,
	0x65, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x28, 0x2e, 0x65, 0x61, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x65, 0x61, 0x73, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50,
	0x61, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x40, 0x92, 0x41, 0x1c, 0x0a, 0x0c, 0xe8, 0xaf, 0x95, 0xe5, 0x8d,
	0xb7, 0xe6, 0xa8, 0xa1, 0xe5, 0x9d, 0x97, 0x12, 0x0c, 0xe8, 0xaf, 0x95, 0xe5, 0x8d, 0xb7, 0xe5,
	0x88, 0x97, 0xe8, 0xa1, 0xa8, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0xcd, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x50, 0x61,
	0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x2e, 0x65, 0x61, 0x73, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x61, 0x6c,
	0x65, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x65, 0x61, 0x73, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x61, 0x6c,
	0x65, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4d, 0x92, 0x41, 0x22, 0x0a, 0x0c, 0xe8, 0xaf,
	0x95, 0xe5, 0x8d, 0xb7, 0xe6, 0xa8, 0xa1, 0xe5, 0x9d, 0x97, 0x12, 0x12, 0xe5, 0x8f, 0xaf, 0xe7,
	0x94, 0xa8, 0xe8, 0xaf, 0x95, 0xe5, 0x8d, 0xb7, 0xe5, 0x88, 0x97, 0xe8, 0xa1, 0xa8, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0xa5, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x61,
	0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x26,
	0x2e, 0x65, 0x61, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x65, 0x61, 0x73, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65,
	0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x3d, 0x92, 0x41, 0x1c, 0x0a, 0x0c, 0xe8, 0xaf, 0x95, 0xe5, 0x8d, 0xb7, 0xe6, 0xa8, 0xa1, 0xe5,
	0x9d, 0x97, 0x12, 0x0c, 0xe8, 0xaf, 0x95, 0xe5, 0x8d, 0xb7, 0xe8, 0xaf, 0xa6, 0xe6, 0x83, 0x85,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6c, 0x65,
	0x73, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0xa5,
	0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61,
	0x70, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x65, 0x61, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x61, 0x73, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65,
	0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x46,
	0x92, 0x41, 0x22, 0x0a, 0x0c, 0xe8, 0xaf, 0x95, 0xe5, 0x8d, 0xb7, 0xe6, 0xa8, 0xa1, 0xe5, 0x9d,
	0x97, 0x12, 0x12, 0xe4, 0xbf, 0xae, 0xe6, 0x94, 0xb9, 0xe8, 0xaf, 0x95, 0xe5, 0x8d, 0xb7, 0xe4,
	0xbf, 0xa1, 0xe6, 0x81, 0xaf, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x72, 0x2f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0xb3, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x53, 0x61,
	0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26,
	0x2e, 0x65, 0x61, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x53,
	0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x65, 0x61, 0x73, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x4b, 0x92, 0x41, 0x23, 0x0a, 0x0c, 0xe8, 0xaf, 0x95, 0xe5, 0x8d, 0xb7, 0xe6, 0xa8, 0xa1, 0xe5,
	0x9d, 0x97, 0x12, 0x13, 0xe7, 0xa6, 0x81, 0xe7, 0x94, 0xa8, 0x2f, 0xe5, 0x90, 0xaf, 0xe7, 0x94,
	0xa8, 0xe8, 0xaf, 0x95, 0xe5, 0x8d, 0xb7, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a,
	0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x9f, 0x01, 0x0a,
	0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65,
	0x72, 0x12, 0x23, 0x2e, 0x65, 0x61, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x61, 0x73, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50,
	0x61, 0x70, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x40, 0x92, 0x41,
	0x1c, 0x0a, 0x0c, 0xe8, 0xaf, 0x95, 0xe5, 0x8d, 0xb7, 0xe6, 0xa8, 0xa1, 0xe5, 0x9d, 0x97, 0x12,
	0x0c, 0xe5, 0x88, 0xa0, 0xe9, 0x99, 0xa4, 0xe8, 0xaf, 0x95, 0xe5, 0x8d, 0xb7, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6c, 0x65,
	0x73, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0xc0,
	0x01, 0x0a, 0x15, 0x53, 0x61, 0x76, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65,
	0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x2e, 0x65, 0x61, 0x73, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50,
	0x61, 0x70, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x65, 0x61, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x61, 0x76, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x52, 0x92,
	0x41, 0x28, 0x0a, 0x12, 0xe8, 0xaf, 0x95, 0xe5, 0x8d, 0xb7, 0xe8, 0xaf, 0x84, 0xe8, 0xaf, 0xad,
	0xe6, 0xa8, 0xa1, 0xe5, 0x9d, 0x97, 0x12, 0x12, 0xe5, 0x88, 0x9b, 0xe5, 0xbb, 0xba, 0xe8, 0xaf,
	0x95, 0xe5, 0x8d, 0xb7, 0xe8, 0xaf, 0x84, 0xe8, 0xaf, 0xad, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21,
	0x3a, 0x01, 0x2a, 0x22, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x61, 0x76,
	0x65, 0x12, 0xc6, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61,
	0x70, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b,
	0x2e, 0x65, 0x61, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x65, 0x61,
	0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65,
	0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4f, 0x92, 0x41, 0x28, 0x0a, 0x12,
	0xe8, 0xaf, 0x95, 0xe5, 0x8d, 0xb7, 0xe8, 0xaf, 0x84, 0xe8, 0xaf, 0xad, 0xe6, 0xa8, 0xa1, 0xe5,
	0x9d, 0x97, 0x12, 0x12, 0xe8, 0xaf, 0x95, 0xe5, 0x8d, 0xb7, 0xe8, 0xaf, 0x84, 0xe8, 0xaf, 0xad,
	0xe5, 0x88, 0x97, 0xe8, 0xa1, 0xa8, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0xc4, 0x01, 0x0a, 0x19, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x44,
	0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e, 0x65, 0x61, 0x73, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65,
	0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x65, 0x61, 0x73, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50,
	0x61, 0x70, 0x65, 0x72, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4a, 0x92, 0x41, 0x1c, 0x0a, 0x0c, 0xe7, 0xbb, 0xb4, 0xe5,
	0xba, 0xa6, 0xe6, 0xa8, 0xa1, 0xe5, 0x9d, 0x97, 0x12, 0x0c, 0xe5, 0x88, 0x9b, 0xe5, 0xbb, 0xba,
	0xe7, 0xbb, 0xb4, 0xe5, 0xba, 0xa6, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a, 0x22,
	0x20, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x70, 0x61, 0x70, 0x65, 0x72,
	0x5f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0xdf, 0x01, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61,
	0x70, 0x65, 0x72, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x67, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x2e, 0x65, 0x61, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x44,
	0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x65, 0x61, 0x73, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x70,
	0x65, 0x72, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x67, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x56, 0x92, 0x41, 0x28,
	0x0a, 0x12, 0xe8, 0xaf, 0x95, 0xe5, 0x8d, 0xb7, 0xe7, 0xbb, 0xb4, 0xe5, 0xba, 0xa6, 0xe6, 0xa8,
	0xa1, 0xe5, 0x9d, 0x97, 0x12, 0x12, 0xe8, 0xaf, 0x95, 0xe5, 0x8d, 0xb7, 0xe7, 0xbb, 0xb4, 0xe5,
	0xba, 0xa6, 0xe5, 0x88, 0x97, 0xe8, 0xa1, 0xa8, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x12, 0x23,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5f,
	0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0xd6, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73,
	0x50, 0x61, 0x70, 0x65, 0x72, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x2f, 0x2e, 0x65, 0x61, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x44,
	0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x65, 0x61, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72,
	0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53, 0x92, 0x41, 0x28, 0x0a, 0x12, 0xe8, 0xaf,
	0x95, 0xe5, 0x8d, 0xb7, 0xe7, 0xbb, 0xb4, 0xe5, 0xba, 0xa6, 0xe6, 0xa8, 0xa1, 0xe5, 0x9d, 0x97,
	0x12, 0x12, 0xe8, 0xaf, 0x95, 0xe5, 0x8d, 0xb7, 0xe7, 0xbb, 0xb4, 0xe5, 0xba, 0xa6, 0xe8, 0xaf,
	0xa6, 0xe6, 0x83, 0x85, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x6d, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0xd6, 0x01, 0x0a,
	0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65,
	0x72, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e, 0x65, 0x61, 0x73,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x61,
	0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x65, 0x61, 0x73, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65,
	0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5c, 0x92, 0x41, 0x2e, 0x0a, 0x12, 0xe8, 0xaf,
	0x95, 0xe5, 0x8d, 0xb7, 0xe7, 0xbb, 0xb4, 0xe5, 0xba, 0xa6, 0xe6, 0xa8, 0xa1, 0xe5, 0x9d, 0x97,
	0x12, 0x18, 0xe4, 0xbf, 0xae, 0xe6, 0x94, 0xb9, 0xe8, 0xaf, 0x95, 0xe5, 0x8d, 0xb7, 0xe7, 0xbb,
	0xb4, 0xe5, 0xba, 0xa6, 0xe4, 0xbf, 0xa1, 0xe6, 0x81, 0xaf, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25,
	0x3a, 0x01, 0x2a, 0x22, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x5f, 0x70,
	0x61, 0x70, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0xd0, 0x01, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e, 0x65, 0x61, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65,
	0x72, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2d, 0x2e, 0x65, 0x61, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x50, 0x61, 0x70, 0x65, 0x72, 0x44,
	0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x56, 0x92, 0x41, 0x28, 0x0a, 0x12, 0xe8, 0xaf, 0x95, 0xe5, 0x8d, 0xb7, 0xe7, 0xbb, 0xb4,
	0xe5, 0xba, 0xa6, 0xe6, 0xa8, 0xa1, 0xe5, 0x9d, 0x97, 0x12, 0x12, 0xe5, 0x88, 0xa0, 0xe9, 0x99,
	0xa4, 0xe8, 0xaf, 0x95, 0xe5, 0x8d, 0xb7, 0xe7, 0xbb, 0xb4, 0xe5, 0xba, 0xa6, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a, 0x22, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6c, 0x65,
	0x73, 0x5f, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x8f, 0x01, 0x92, 0x41, 0x79, 0x12, 0x1f,
	0x0a, 0x18, 0xe8, 0x80, 0x83, 0xe8, 0xaf, 0x95, 0xe5, 0x90, 0x8e, 0xe5, 0x8f, 0xb0, 0xe8, 0xaf,
	0x95, 0xe5, 0x8d, 0xb7, 0xe6, 0x8e, 0xa5, 0xe5, 0x8f, 0xa3, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a,
	0x01, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x1d, 0x0a, 0x1b, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x4b,
	0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x0d, 0x08, 0x02, 0x1a, 0x07, 0x78, 0x2d, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x20, 0x02, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65,
	0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x5a, 0x11, 0x65, 0x61, 0x73, 0x5f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_eas_api_v1_eas_sales_paper_proto_goTypes = []interface{}{
	(*CreateSalesPaperRequest)(nil),                // 0: eas_api.v1.CreateSalesPaperRequest
	(*GetSalesPaperPageListRequest)(nil),           // 1: eas_api.v1.GetSalesPaperPageListRequest
	(*GetUsableSalesPaperPageListRequest)(nil),     // 2: eas_api.v1.GetUsableSalesPaperPageListRequest
	(*GetSalesPaperDetailRequest)(nil),             // 3: eas_api.v1.GetSalesPaperDetailRequest
	(*UpdateSalesPaperRequest)(nil),                // 4: eas_api.v1.UpdateSalesPaperRequest
	(*SetSalesPaperStatusRequest)(nil),             // 5: eas_api.v1.SetSalesPaperStatusRequest
	(*DeleteSalesPaperRequest)(nil),                // 6: eas_api.v1.DeleteSalesPaperRequest
	(*SaveSalesPaperCommentRequest)(nil),           // 7: eas_api.v1.SaveSalesPaperCommentRequest
	(*GetSalesPaperCommentListRequest)(nil),        // 8: eas_api.v1.GetSalesPaperCommentListRequest
	(*CreateSalesPaperDimensionRequest)(nil),       // 9: eas_api.v1.CreateSalesPaperDimensionRequest
	(*GetSalesPaperDimensionPageListRequest)(nil),  // 10: eas_api.v1.GetSalesPaperDimensionPageListRequest
	(*GetSalesPaperDimensionDetailRequest)(nil),    // 11: eas_api.v1.GetSalesPaperDimensionDetailRequest
	(*UpdateSalesPaperDimensionRequest)(nil),       // 12: eas_api.v1.UpdateSalesPaperDimensionRequest
	(*DeleteSalesPaperDimensionRequest)(nil),       // 13: eas_api.v1.DeleteSalesPaperDimensionRequest
	(*CreateSalesPaperResponse)(nil),               // 14: eas_api.v1.CreateSalesPaperResponse
	(*GetSalesPaperPageListResponse)(nil),          // 15: eas_api.v1.GetSalesPaperPageListResponse
	(*GetUsableSalesPaperPageListResponse)(nil),    // 16: eas_api.v1.GetUsableSalesPaperPageListResponse
	(*GetSalesPaperDetailResponse)(nil),            // 17: eas_api.v1.GetSalesPaperDetailResponse
	(*UpdateSalesPaperResponse)(nil),               // 18: eas_api.v1.UpdateSalesPaperResponse
	(*SetSalesPaperStatusResponse)(nil),            // 19: eas_api.v1.SetSalesPaperStatusResponse
	(*DeleteSalesPaperResponse)(nil),               // 20: eas_api.v1.DeleteSalesPaperResponse
	(*SaveSalesPaperCommentResponse)(nil),          // 21: eas_api.v1.SaveSalesPaperCommentResponse
	(*GetSalesPaperCommentListResponse)(nil),       // 22: eas_api.v1.GetSalesPaperCommentListResponse
	(*CreateSalesPaperDimensionResponse)(nil),      // 23: eas_api.v1.CreateSalesPaperDimensionResponse
	(*GetSalesPaperDimensionPageListResponse)(nil), // 24: eas_api.v1.GetSalesPaperDimensionPageListResponse
	(*GetSalesPaperDimensionDetailResponse)(nil),   // 25: eas_api.v1.GetSalesPaperDimensionDetailResponse
	(*UpdateSalesPaperDimensionResponse)(nil),      // 26: eas_api.v1.UpdateSalesPaperDimensionResponse
	(*DeleteSalesPaperDimensionResponse)(nil),      // 27: eas_api.v1.DeleteSalesPaperDimensionResponse
}
var file_eas_api_v1_eas_sales_paper_proto_depIdxs = []int32{
	0,  // 0: eas_api.v1.EasSalesPaperService.CreateSalesPaper:input_type -> eas_api.v1.CreateSalesPaperRequest
	1,  // 1: eas_api.v1.EasSalesPaperService.GetSalesPaperPageList:input_type -> eas_api.v1.GetSalesPaperPageListRequest
	2,  // 2: eas_api.v1.EasSalesPaperService.GetUsableSalesPaperPageList:input_type -> eas_api.v1.GetUsableSalesPaperPageListRequest
	3,  // 3: eas_api.v1.EasSalesPaperService.GetSalesPaperDetail:input_type -> eas_api.v1.GetSalesPaperDetailRequest
	4,  // 4: eas_api.v1.EasSalesPaperService.UpdateSalesPaper:input_type -> eas_api.v1.UpdateSalesPaperRequest
	5,  // 5: eas_api.v1.EasSalesPaperService.SetSalesPaperStatus:input_type -> eas_api.v1.SetSalesPaperStatusRequest
	6,  // 6: eas_api.v1.EasSalesPaperService.DeleteSalesPaper:input_type -> eas_api.v1.DeleteSalesPaperRequest
	7,  // 7: eas_api.v1.EasSalesPaperService.SaveSalesPaperComment:input_type -> eas_api.v1.SaveSalesPaperCommentRequest
	8,  // 8: eas_api.v1.EasSalesPaperService.GetSalesPaperCommentList:input_type -> eas_api.v1.GetSalesPaperCommentListRequest
	9,  // 9: eas_api.v1.EasSalesPaperService.CreateSalesPaperDimension:input_type -> eas_api.v1.CreateSalesPaperDimensionRequest
	10, // 10: eas_api.v1.EasSalesPaperService.GetSalesPaperDimensionPageList:input_type -> eas_api.v1.GetSalesPaperDimensionPageListRequest
	11, // 11: eas_api.v1.EasSalesPaperService.GetSalesPaperDimensionDetail:input_type -> eas_api.v1.GetSalesPaperDimensionDetailRequest
	12, // 12: eas_api.v1.EasSalesPaperService.UpdateSalesPaperDimension:input_type -> eas_api.v1.UpdateSalesPaperDimensionRequest
	13, // 13: eas_api.v1.EasSalesPaperService.DeleteSalesPaperDimension:input_type -> eas_api.v1.DeleteSalesPaperDimensionRequest
	14, // 14: eas_api.v1.EasSalesPaperService.CreateSalesPaper:output_type -> eas_api.v1.CreateSalesPaperResponse
	15, // 15: eas_api.v1.EasSalesPaperService.GetSalesPaperPageList:output_type -> eas_api.v1.GetSalesPaperPageListResponse
	16, // 16: eas_api.v1.EasSalesPaperService.GetUsableSalesPaperPageList:output_type -> eas_api.v1.GetUsableSalesPaperPageListResponse
	17, // 17: eas_api.v1.EasSalesPaperService.GetSalesPaperDetail:output_type -> eas_api.v1.GetSalesPaperDetailResponse
	18, // 18: eas_api.v1.EasSalesPaperService.UpdateSalesPaper:output_type -> eas_api.v1.UpdateSalesPaperResponse
	19, // 19: eas_api.v1.EasSalesPaperService.SetSalesPaperStatus:output_type -> eas_api.v1.SetSalesPaperStatusResponse
	20, // 20: eas_api.v1.EasSalesPaperService.DeleteSalesPaper:output_type -> eas_api.v1.DeleteSalesPaperResponse
	21, // 21: eas_api.v1.EasSalesPaperService.SaveSalesPaperComment:output_type -> eas_api.v1.SaveSalesPaperCommentResponse
	22, // 22: eas_api.v1.EasSalesPaperService.GetSalesPaperCommentList:output_type -> eas_api.v1.GetSalesPaperCommentListResponse
	23, // 23: eas_api.v1.EasSalesPaperService.CreateSalesPaperDimension:output_type -> eas_api.v1.CreateSalesPaperDimensionResponse
	24, // 24: eas_api.v1.EasSalesPaperService.GetSalesPaperDimensionPageList:output_type -> eas_api.v1.GetSalesPaperDimensionPageListResponse
	25, // 25: eas_api.v1.EasSalesPaperService.GetSalesPaperDimensionDetail:output_type -> eas_api.v1.GetSalesPaperDimensionDetailResponse
	26, // 26: eas_api.v1.EasSalesPaperService.UpdateSalesPaperDimension:output_type -> eas_api.v1.UpdateSalesPaperDimensionResponse
	27, // 27: eas_api.v1.EasSalesPaperService.DeleteSalesPaperDimension:output_type -> eas_api.v1.DeleteSalesPaperDimensionResponse
	14, // [14:28] is the sub-list for method output_type
	0,  // [0:14] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_eas_api_v1_eas_sales_paper_proto_init() }
func file_eas_api_v1_eas_sales_paper_proto_init() {
	if File_eas_api_v1_eas_sales_paper_proto != nil {
		return
	}
	file_eas_api_v1_eas_sales_paper_models_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eas_api_v1_eas_sales_paper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eas_api_v1_eas_sales_paper_proto_goTypes,
		DependencyIndexes: file_eas_api_v1_eas_sales_paper_proto_depIdxs,
	}.Build()
	File_eas_api_v1_eas_sales_paper_proto = out.File
	file_eas_api_v1_eas_sales_paper_proto_rawDesc = nil
	file_eas_api_v1_eas_sales_paper_proto_goTypes = nil
	file_eas_api_v1_eas_sales_paper_proto_depIdxs = nil
}
