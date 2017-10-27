// Code generated by protoc-gen-go.
// source: test.proto
// DO NOT EDIT!

/*
Package tests is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	Test
	GetTest
	CreateTest
	SetTest
	UpdateTest
	UpdatePathsTest
	DeleteTest
	SetOption
	FieldPath
*/
package tests

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_firestore_v1beta14 "google.golang.org/genproto/googleapis/firestore/v1beta1"
import google_firestore_v1beta1 "google.golang.org/genproto/googleapis/firestore/v1beta1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A Test describes a single client method call and its expected result.
type Test struct {
	Description string `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	// Types that are valid to be assigned to Test:
	//	*Test_Get
	//	*Test_Create
	//	*Test_Set
	//	*Test_Update
	//	*Test_UpdatePaths
	//	*Test_Delete
	Test isTest_Test `protobuf_oneof:"test"`
}

func (m *Test) Reset()                    { *m = Test{} }
func (m *Test) String() string            { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()               {}
func (*Test) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isTest_Test interface {
	isTest_Test()
}

type Test_Get struct {
	Get *GetTest `protobuf:"bytes,2,opt,name=get,oneof"`
}
type Test_Create struct {
	Create *CreateTest `protobuf:"bytes,3,opt,name=create,oneof"`
}
type Test_Set struct {
	Set *SetTest `protobuf:"bytes,4,opt,name=set,oneof"`
}
type Test_Update struct {
	Update *UpdateTest `protobuf:"bytes,5,opt,name=update,oneof"`
}
type Test_UpdatePaths struct {
	UpdatePaths *UpdatePathsTest `protobuf:"bytes,6,opt,name=update_paths,json=updatePaths,oneof"`
}
type Test_Delete struct {
	Delete *DeleteTest `protobuf:"bytes,7,opt,name=delete,oneof"`
}

func (*Test_Get) isTest_Test()         {}
func (*Test_Create) isTest_Test()      {}
func (*Test_Set) isTest_Test()         {}
func (*Test_Update) isTest_Test()      {}
func (*Test_UpdatePaths) isTest_Test() {}
func (*Test_Delete) isTest_Test()      {}

func (m *Test) GetTest() isTest_Test {
	if m != nil {
		return m.Test
	}
	return nil
}

func (m *Test) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Test) GetGet() *GetTest {
	if x, ok := m.GetTest().(*Test_Get); ok {
		return x.Get
	}
	return nil
}

func (m *Test) GetCreate() *CreateTest {
	if x, ok := m.GetTest().(*Test_Create); ok {
		return x.Create
	}
	return nil
}

func (m *Test) GetSet() *SetTest {
	if x, ok := m.GetTest().(*Test_Set); ok {
		return x.Set
	}
	return nil
}

func (m *Test) GetUpdate() *UpdateTest {
	if x, ok := m.GetTest().(*Test_Update); ok {
		return x.Update
	}
	return nil
}

func (m *Test) GetUpdatePaths() *UpdatePathsTest {
	if x, ok := m.GetTest().(*Test_UpdatePaths); ok {
		return x.UpdatePaths
	}
	return nil
}

func (m *Test) GetDelete() *DeleteTest {
	if x, ok := m.GetTest().(*Test_Delete); ok {
		return x.Delete
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Test) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Test_OneofMarshaler, _Test_OneofUnmarshaler, _Test_OneofSizer, []interface{}{
		(*Test_Get)(nil),
		(*Test_Create)(nil),
		(*Test_Set)(nil),
		(*Test_Update)(nil),
		(*Test_UpdatePaths)(nil),
		(*Test_Delete)(nil),
	}
}

func _Test_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Test)
	// test
	switch x := m.Test.(type) {
	case *Test_Get:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Get); err != nil {
			return err
		}
	case *Test_Create:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *Test_Set:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Set); err != nil {
			return err
		}
	case *Test_Update:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Update); err != nil {
			return err
		}
	case *Test_UpdatePaths:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UpdatePaths); err != nil {
			return err
		}
	case *Test_Delete:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Delete); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Test.Test has unexpected type %T", x)
	}
	return nil
}

func _Test_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Test)
	switch tag {
	case 2: // test.get
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GetTest)
		err := b.DecodeMessage(msg)
		m.Test = &Test_Get{msg}
		return true, err
	case 3: // test.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CreateTest)
		err := b.DecodeMessage(msg)
		m.Test = &Test_Create{msg}
		return true, err
	case 4: // test.set
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SetTest)
		err := b.DecodeMessage(msg)
		m.Test = &Test_Set{msg}
		return true, err
	case 5: // test.update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UpdateTest)
		err := b.DecodeMessage(msg)
		m.Test = &Test_Update{msg}
		return true, err
	case 6: // test.update_paths
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UpdatePathsTest)
		err := b.DecodeMessage(msg)
		m.Test = &Test_UpdatePaths{msg}
		return true, err
	case 7: // test.delete
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DeleteTest)
		err := b.DecodeMessage(msg)
		m.Test = &Test_Delete{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Test_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Test)
	// test
	switch x := m.Test.(type) {
	case *Test_Get:
		s := proto.Size(x.Get)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Test_Create:
		s := proto.Size(x.Create)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Test_Set:
		s := proto.Size(x.Set)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Test_Update:
		s := proto.Size(x.Update)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Test_UpdatePaths:
		s := proto.Size(x.UpdatePaths)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Test_Delete:
		s := proto.Size(x.Delete)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Call to the DocumentRef.Get method.
type GetTest struct {
	// The path of the doc, e.g. "projects/projectID/databases/(default)/documents/C/d"
	DocRefPath string `protobuf:"bytes,1,opt,name=doc_ref_path,json=docRefPath" json:"doc_ref_path,omitempty"`
	// The request that the call should send to the Firestore service.
	Request *google_firestore_v1beta14.GetDocumentRequest `protobuf:"bytes,2,opt,name=request" json:"request,omitempty"`
}

func (m *GetTest) Reset()                    { *m = GetTest{} }
func (m *GetTest) String() string            { return proto.CompactTextString(m) }
func (*GetTest) ProtoMessage()               {}
func (*GetTest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetTest) GetDocRefPath() string {
	if m != nil {
		return m.DocRefPath
	}
	return ""
}

func (m *GetTest) GetRequest() *google_firestore_v1beta14.GetDocumentRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

// Call to DocumentRef.Create.
type CreateTest struct {
	// The path of the doc, e.g. "projects/projectID/databases/(default)/documents/C/d"
	DocRefPath string `protobuf:"bytes,1,opt,name=doc_ref_path,json=docRefPath" json:"doc_ref_path,omitempty"`
	// The data passed to Create, as JSON. The strings "Delete" and "ServerTimestamp"
	// denote the two special sentinel values. Values that could be interpreted as integers
	// (i.e. digit strings) should be treated as integers.
	JsonData string `protobuf:"bytes,2,opt,name=json_data,json=jsonData" json:"json_data,omitempty"`
	// The request that the call should generate.
	Request *google_firestore_v1beta14.CommitRequest `protobuf:"bytes,3,opt,name=request" json:"request,omitempty"`
	// If true, the call should result in an error without generating a request.
	// If this is true, request should not be set.
	IsError bool `protobuf:"varint,4,opt,name=is_error,json=isError" json:"is_error,omitempty"`
}

func (m *CreateTest) Reset()                    { *m = CreateTest{} }
func (m *CreateTest) String() string            { return proto.CompactTextString(m) }
func (*CreateTest) ProtoMessage()               {}
func (*CreateTest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateTest) GetDocRefPath() string {
	if m != nil {
		return m.DocRefPath
	}
	return ""
}

func (m *CreateTest) GetJsonData() string {
	if m != nil {
		return m.JsonData
	}
	return ""
}

func (m *CreateTest) GetRequest() *google_firestore_v1beta14.CommitRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *CreateTest) GetIsError() bool {
	if m != nil {
		return m.IsError
	}
	return false
}

// A call to DocumentRef.Set.
type SetTest struct {
	DocRefPath string                                   `protobuf:"bytes,1,opt,name=doc_ref_path,json=docRefPath" json:"doc_ref_path,omitempty"`
	Options    []*SetOption                             `protobuf:"bytes,2,rep,name=options" json:"options,omitempty"`
	JsonData   string                                   `protobuf:"bytes,3,opt,name=json_data,json=jsonData" json:"json_data,omitempty"`
	Request    *google_firestore_v1beta14.CommitRequest `protobuf:"bytes,4,opt,name=request" json:"request,omitempty"`
	IsError    bool                                     `protobuf:"varint,5,opt,name=is_error,json=isError" json:"is_error,omitempty"`
}

func (m *SetTest) Reset()                    { *m = SetTest{} }
func (m *SetTest) String() string            { return proto.CompactTextString(m) }
func (*SetTest) ProtoMessage()               {}
func (*SetTest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SetTest) GetDocRefPath() string {
	if m != nil {
		return m.DocRefPath
	}
	return ""
}

func (m *SetTest) GetOptions() []*SetOption {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *SetTest) GetJsonData() string {
	if m != nil {
		return m.JsonData
	}
	return ""
}

func (m *SetTest) GetRequest() *google_firestore_v1beta14.CommitRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *SetTest) GetIsError() bool {
	if m != nil {
		return m.IsError
	}
	return false
}

// A call to the form of DocumentRef.Update that represents the data as a map
// or dictionary.
type UpdateTest struct {
	DocRefPath   string                                   `protobuf:"bytes,1,opt,name=doc_ref_path,json=docRefPath" json:"doc_ref_path,omitempty"`
	Precondition *google_firestore_v1beta1.Precondition   `protobuf:"bytes,2,opt,name=precondition" json:"precondition,omitempty"`
	JsonData     string                                   `protobuf:"bytes,3,opt,name=json_data,json=jsonData" json:"json_data,omitempty"`
	Request      *google_firestore_v1beta14.CommitRequest `protobuf:"bytes,4,opt,name=request" json:"request,omitempty"`
	IsError      bool                                     `protobuf:"varint,5,opt,name=is_error,json=isError" json:"is_error,omitempty"`
}

func (m *UpdateTest) Reset()                    { *m = UpdateTest{} }
func (m *UpdateTest) String() string            { return proto.CompactTextString(m) }
func (*UpdateTest) ProtoMessage()               {}
func (*UpdateTest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateTest) GetDocRefPath() string {
	if m != nil {
		return m.DocRefPath
	}
	return ""
}

func (m *UpdateTest) GetPrecondition() *google_firestore_v1beta1.Precondition {
	if m != nil {
		return m.Precondition
	}
	return nil
}

func (m *UpdateTest) GetJsonData() string {
	if m != nil {
		return m.JsonData
	}
	return ""
}

func (m *UpdateTest) GetRequest() *google_firestore_v1beta14.CommitRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *UpdateTest) GetIsError() bool {
	if m != nil {
		return m.IsError
	}
	return false
}

// A call to the form of DocumentRef.Update that represents the data as a list
// of field paths and their values.
type UpdatePathsTest struct {
	DocRefPath   string                                 `protobuf:"bytes,1,opt,name=doc_ref_path,json=docRefPath" json:"doc_ref_path,omitempty"`
	Precondition *google_firestore_v1beta1.Precondition `protobuf:"bytes,2,opt,name=precondition" json:"precondition,omitempty"`
	// parallel sequences: field_paths[i] corresponds to json_values[i]
	FieldPaths []*FieldPath                             `protobuf:"bytes,3,rep,name=field_paths,json=fieldPaths" json:"field_paths,omitempty"`
	JsonValues []string                                 `protobuf:"bytes,4,rep,name=json_values,json=jsonValues" json:"json_values,omitempty"`
	Request    *google_firestore_v1beta14.CommitRequest `protobuf:"bytes,5,opt,name=request" json:"request,omitempty"`
	IsError    bool                                     `protobuf:"varint,6,opt,name=is_error,json=isError" json:"is_error,omitempty"`
}

func (m *UpdatePathsTest) Reset()                    { *m = UpdatePathsTest{} }
func (m *UpdatePathsTest) String() string            { return proto.CompactTextString(m) }
func (*UpdatePathsTest) ProtoMessage()               {}
func (*UpdatePathsTest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdatePathsTest) GetDocRefPath() string {
	if m != nil {
		return m.DocRefPath
	}
	return ""
}

func (m *UpdatePathsTest) GetPrecondition() *google_firestore_v1beta1.Precondition {
	if m != nil {
		return m.Precondition
	}
	return nil
}

func (m *UpdatePathsTest) GetFieldPaths() []*FieldPath {
	if m != nil {
		return m.FieldPaths
	}
	return nil
}

func (m *UpdatePathsTest) GetJsonValues() []string {
	if m != nil {
		return m.JsonValues
	}
	return nil
}

func (m *UpdatePathsTest) GetRequest() *google_firestore_v1beta14.CommitRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *UpdatePathsTest) GetIsError() bool {
	if m != nil {
		return m.IsError
	}
	return false
}

// A call to DocmentRef.Delete
type DeleteTest struct {
	DocRefPath   string                                   `protobuf:"bytes,1,opt,name=doc_ref_path,json=docRefPath" json:"doc_ref_path,omitempty"`
	Precondition *google_firestore_v1beta1.Precondition   `protobuf:"bytes,2,opt,name=precondition" json:"precondition,omitempty"`
	Request      *google_firestore_v1beta14.CommitRequest `protobuf:"bytes,3,opt,name=request" json:"request,omitempty"`
	IsError      bool                                     `protobuf:"varint,4,opt,name=is_error,json=isError" json:"is_error,omitempty"`
}

func (m *DeleteTest) Reset()                    { *m = DeleteTest{} }
func (m *DeleteTest) String() string            { return proto.CompactTextString(m) }
func (*DeleteTest) ProtoMessage()               {}
func (*DeleteTest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteTest) GetDocRefPath() string {
	if m != nil {
		return m.DocRefPath
	}
	return ""
}

func (m *DeleteTest) GetPrecondition() *google_firestore_v1beta1.Precondition {
	if m != nil {
		return m.Precondition
	}
	return nil
}

func (m *DeleteTest) GetRequest() *google_firestore_v1beta14.CommitRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *DeleteTest) GetIsError() bool {
	if m != nil {
		return m.IsError
	}
	return false
}

// An option to the DocumentRef.Set call.
type SetOption struct {
	All    bool         `protobuf:"varint,1,opt,name=all" json:"all,omitempty"`
	Fields []*FieldPath `protobuf:"bytes,2,rep,name=fields" json:"fields,omitempty"`
}

func (m *SetOption) Reset()                    { *m = SetOption{} }
func (m *SetOption) String() string            { return proto.CompactTextString(m) }
func (*SetOption) ProtoMessage()               {}
func (*SetOption) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SetOption) GetAll() bool {
	if m != nil {
		return m.All
	}
	return false
}

func (m *SetOption) GetFields() []*FieldPath {
	if m != nil {
		return m.Fields
	}
	return nil
}

// A field path.
type FieldPath struct {
	Field []string `protobuf:"bytes,1,rep,name=field" json:"field,omitempty"`
}

func (m *FieldPath) Reset()                    { *m = FieldPath{} }
func (m *FieldPath) String() string            { return proto.CompactTextString(m) }
func (*FieldPath) ProtoMessage()               {}
func (*FieldPath) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *FieldPath) GetField() []string {
	if m != nil {
		return m.Field
	}
	return nil
}

func init() {
	proto.RegisterType((*Test)(nil), "tests.Test")
	proto.RegisterType((*GetTest)(nil), "tests.GetTest")
	proto.RegisterType((*CreateTest)(nil), "tests.CreateTest")
	proto.RegisterType((*SetTest)(nil), "tests.SetTest")
	proto.RegisterType((*UpdateTest)(nil), "tests.UpdateTest")
	proto.RegisterType((*UpdatePathsTest)(nil), "tests.UpdatePathsTest")
	proto.RegisterType((*DeleteTest)(nil), "tests.DeleteTest")
	proto.RegisterType((*SetOption)(nil), "tests.SetOption")
	proto.RegisterType((*FieldPath)(nil), "tests.FieldPath")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x71, 0x62, 0x27, 0xe3, 0x0a, 0xca, 0x0a, 0x21, 0x53, 0x0e, 0x18, 0x4b, 0x40, 0x04,
	0x28, 0x55, 0xca, 0x91, 0x13, 0x34, 0xb4, 0x88, 0x0b, 0xd5, 0x16, 0xb8, 0x46, 0xae, 0x3d, 0x09,
	0x46, 0x8e, 0xd7, 0xec, 0xae, 0xfb, 0x4e, 0x1c, 0xb9, 0xf3, 0x14, 0x9c, 0x78, 0x11, 0xee, 0x68,
	0x77, 0xed, 0x6e, 0x5c, 0x14, 0x29, 0x87, 0x52, 0x6e, 0xeb, 0x6f, 0xbe, 0xf9, 0xfb, 0x66, 0x26,
	0x01, 0x90, 0x28, 0xe4, 0xa4, 0xe2, 0x4c, 0x32, 0x32, 0x50, 0x6f, 0xb1, 0x37, 0x5e, 0x32, 0xb6,
	0x2c, 0x70, 0x7f, 0x91, 0x73, 0x14, 0x92, 0x71, 0xdc, 0x3f, 0x9f, 0x9e, 0xa1, 0x4c, 0xa6, 0x16,
	0x31, 0x0e, 0x7b, 0x8f, 0x36, 0x32, 0x53, 0xb6, 0x5a, 0xb1, 0xd2, 0xd0, 0xe2, 0x1f, 0x3d, 0xe8,
	0x7f, 0x40, 0x21, 0x49, 0x04, 0x41, 0x86, 0x22, 0xe5, 0x79, 0x25, 0x73, 0x56, 0x86, 0x4e, 0xe4,
	0x8c, 0x47, 0x74, 0x1d, 0x22, 0x31, 0xb8, 0x4b, 0x94, 0x61, 0x2f, 0x72, 0xc6, 0xc1, 0xc1, 0xcd,
	0x89, 0x2e, 0x68, 0x72, 0x8c, 0x52, 0xb9, 0xbf, 0xbd, 0x41, 0x95, 0x91, 0x3c, 0x03, 0x2f, 0xe5,
	0x98, 0x48, 0x0c, 0x5d, 0x4d, 0xbb, 0xdd, 0xd0, 0x0e, 0x35, 0xd8, 0x30, 0x1b, 0x8a, 0x0a, 0x28,
	0x50, 0x86, 0xfd, 0x4e, 0xc0, 0x53, 0x1b, 0x50, 0x98, 0x80, 0x75, 0x95, 0xa9, 0x80, 0x83, 0x4e,
	0xc0, 0x8f, 0x1a, 0x6c, 0x03, 0x1a, 0x0a, 0x79, 0x09, 0x3b, 0xe6, 0x35, 0xaf, 0x12, 0xf9, 0x59,
	0x84, 0x9e, 0x76, 0xb9, 0xdb, 0x71, 0x39, 0x51, 0x96, 0xc6, 0x2f, 0xa8, 0x2d, 0xa4, 0x32, 0x65,
	0x58, 0xa0, 0xc4, 0xd0, 0xef, 0x64, 0x9a, 0x69, 0xb0, 0xcd, 0x64, 0x28, 0xaf, 0x3d, 0xe8, 0x2b,
	0x6b, 0x2c, 0xc0, 0x6f, 0x14, 0x20, 0x11, 0xec, 0x64, 0x2c, 0x9d, 0x73, 0x5c, 0xe8, 0xec, 0x8d,
	0x82, 0x90, 0xb1, 0x94, 0xe2, 0x42, 0xa5, 0x20, 0x47, 0xe0, 0x73, 0xfc, 0x5a, 0xa3, 0x68, 0x45,
	0x7c, 0x3e, 0x31, 0x43, 0x9a, 0xd8, 0xe1, 0x35, 0x43, 0x52, 0xba, 0xce, 0x58, 0x5a, 0xaf, 0xb0,
	0x94, 0xd4, 0xf8, 0xd0, 0xd6, 0x39, 0xfe, 0xe6, 0x00, 0x58, 0x41, 0xb7, 0x48, 0x7c, 0x1f, 0x46,
	0x5f, 0x04, 0x2b, 0xe7, 0x59, 0x22, 0x13, 0x9d, 0x7a, 0x44, 0x87, 0x0a, 0x98, 0x25, 0x32, 0x21,
	0xaf, 0x6c, 0x55, 0x66, 0x66, 0x4f, 0x36, 0x57, 0x75, 0xc8, 0x56, 0xab, 0xfc, 0xaf, 0x82, 0xc8,
	0x3d, 0x18, 0xe6, 0x62, 0x8e, 0x9c, 0x33, 0xae, 0xa7, 0x39, 0xa4, 0x7e, 0x2e, 0xde, 0xa8, 0xcf,
	0xf8, 0xa7, 0x03, 0xfe, 0xe9, 0xd6, 0x0a, 0x3d, 0x05, 0x9f, 0xe9, 0x65, 0x13, 0x61, 0x2f, 0x72,
	0xc7, 0xc1, 0xc1, 0xae, 0xdd, 0x8a, 0xf7, 0xda, 0x40, 0x5b, 0x42, 0xb7, 0x29, 0x77, 0x73, 0x53,
	0xfd, 0x2b, 0x68, 0x6a, 0xd0, 0x6d, 0xea, 0xb7, 0x03, 0x60, 0x17, 0x70, 0x8b, 0xbe, 0xde, 0xc1,
	0x4e, 0xc5, 0x31, 0x65, 0x65, 0x96, 0xeb, 0xeb, 0x32, 0xe3, 0x7f, 0xbc, 0xb9, 0xa6, 0x93, 0x35,
	0x36, 0xed, 0xf8, 0xfe, 0xcf, 0xbe, 0xbf, 0xf7, 0xe0, 0xd6, 0xa5, 0x2b, 0xba, 0xe6, 0xe6, 0xa7,
	0x10, 0x2c, 0x72, 0x2c, 0xb2, 0xe6, 0xc0, 0xdd, 0xce, 0x92, 0x1c, 0x29, 0x8b, 0x4a, 0x49, 0x61,
	0xd1, 0x3e, 0x05, 0x79, 0x00, 0x81, 0xd6, 0xeb, 0x3c, 0x29, 0x6a, 0x14, 0x61, 0x3f, 0x72, 0x55,
	0x7d, 0x0a, 0xfa, 0xa4, 0x91, 0x75, 0xcd, 0x06, 0x57, 0xa0, 0x99, 0xd7, 0xd5, 0xec, 0x97, 0x03,
	0x60, 0x7f, 0x42, 0xae, 0x59, 0xae, 0x7f, 0x7b, 0xdb, 0xc7, 0x30, 0xba, 0xb8, 0x4b, 0xb2, 0x0b,
	0x6e, 0x52, 0x14, 0xba, 0x9f, 0x21, 0x55, 0x4f, 0x32, 0x06, 0x4f, 0x8f, 0xe1, 0xf2, 0x2d, 0xdb,
	0x31, 0x35, 0xf6, 0xf8, 0x21, 0x8c, 0x2e, 0x40, 0x72, 0x07, 0x06, 0x1a, 0x0e, 0x1d, 0x3d, 0x29,
	0xf3, 0x71, 0xe6, 0xe9, 0xbf, 0xab, 0x17, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x76, 0x0e, 0xcc,
	0x07, 0x14, 0x07, 0x00, 0x00,
}