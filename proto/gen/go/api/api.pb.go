// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: api.proto

package api

import (
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

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x1a, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x73,
	0x75, 0x73, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xee, 0x02, 0x0a, 0x04, 0x54, 0x6f,
	0x64, 0x6f, 0x12, 0x2b, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x0e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12,
	0x2b, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x12, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x07,
	0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64,
	0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41,
	0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x13, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71,
	0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x61, 0x73,
	0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x37, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_api_proto_goTypes = []any{
	(*SignupReq)(nil),       // 0: api.SignupReq
	(*SigninReq)(nil),       // 1: api.SigninReq
	(*AddTaskReq)(nil),      // 2: api.AddTaskReq
	(*GetTaskReq)(nil),      // 3: api.GetTaskReq
	(*GetAllTasksReq)(nil),  // 4: api.GetAllTasksReq
	(*UpdateTaskReq)(nil),   // 5: api.UpdateTaskReq
	(*DeleteTaskReq)(nil),   // 6: api.DeleteTaskReq
	(*SignupResp)(nil),      // 7: api.SignupResp
	(*SigninResp)(nil),      // 8: api.SigninResp
	(*AddTaskResp)(nil),     // 9: api.AddTaskResp
	(*GetTaskResp)(nil),     // 10: api.GetTaskResp
	(*GetAllTasksResp)(nil), // 11: api.GetAllTasksResp
	(*UpdateTaskResp)(nil),  // 12: api.UpdateTaskResp
	(*DeleteTaskResp)(nil),  // 13: api.DeleteTaskResp
}
var file_api_proto_depIdxs = []int32{
	0,  // 0: api.Todo.Signup:input_type -> api.SignupReq
	1,  // 1: api.Todo.Signin:input_type -> api.SigninReq
	2,  // 2: api.Todo.AddTask:input_type -> api.AddTaskReq
	3,  // 3: api.Todo.GetTask:input_type -> api.GetTaskReq
	4,  // 4: api.Todo.GetAllTasks:input_type -> api.GetAllTasksReq
	5,  // 5: api.Todo.UpdateTask:input_type -> api.UpdateTaskReq
	6,  // 6: api.Todo.DeleteTask:input_type -> api.DeleteTaskReq
	7,  // 7: api.Todo.Signup:output_type -> api.SignupResp
	8,  // 8: api.Todo.Signin:output_type -> api.SigninResp
	9,  // 9: api.Todo.AddTask:output_type -> api.AddTaskResp
	10, // 10: api.Todo.GetTask:output_type -> api.GetTaskResp
	11, // 11: api.Todo.GetAllTasks:output_type -> api.GetAllTasksResp
	12, // 12: api.Todo.UpdateTask:output_type -> api.UpdateTaskResp
	13, // 13: api.Todo.DeleteTask:output_type -> api.DeleteTaskResp
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	file_tasks_proto_init()
	file_susi_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
