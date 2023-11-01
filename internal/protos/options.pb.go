// Code generated by protoc-gen-go. DO NOT EDIT.
// source: options.proto

package protos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NumericalUnits int32

const (
	// Should error if UNKNOWN
	NumericalUnits_UNKNOWN NumericalUnits = 0
	// Number has no units (e.g. ID).
	NumericalUnits_NO_UNITS NumericalUnits = 1
	// ====== DISTANCE METRICS =======
	// Meters
	// Display name: "m"
	NumericalUnits_M NumericalUnits = 2
	// Meters / second
	// Display name: "m/s"
	NumericalUnits_MPS NumericalUnits = 3
	// Meters / second ^ 2
	// Display name: "m/s^2"
	NumericalUnits_MPS2 NumericalUnits = 4
	// Meters / second ^ 3
	// Display name: "m/s^3"
	NumericalUnits_MPS3 NumericalUnits = 14
	// Millimeters
	// Display name: "mm"
	NumericalUnits_MM NumericalUnits = 15
	// Micrometers
	// Display name: "µm"
	NumericalUnits_UM NumericalUnits = 16
	// ==== TIME METRICS ======
	// Seconds
	// Dislay name: "s"
	NumericalUnits_S NumericalUnits = 5
	// === ANGLE METRICS ===
	// Radians
	// Display name: "rad"
	NumericalUnits_RAD NumericalUnits = 6
	// Radians / second
	// Display name: "rad/s"
	NumericalUnits_RADPS NumericalUnits = 10
	// Radians / second ^ 2
	// Display name: "rad/s^2"
	NumericalUnits_RADPS2 NumericalUnits = 11
	// Radians / second ^ 3
	// Display name: "rad/s^3"
	NumericalUnits_RADPS3 NumericalUnits = 12
	// Degrees
	// Display name: "deg"
	NumericalUnits_DEG NumericalUnits = 22
	// == MASS METRICS =========
	// Grams
	// Display name: "g"
	NumericalUnits_G NumericalUnits = 7
	// Kilograms
	// Display name: "kg"
	NumericalUnits_KG NumericalUnits = 8
	// === FREQUENCY METRICS ===
	// Hertz (1 / seconds)
	// Display name: "Hz"
	NumericalUnits_HERTZ NumericalUnits = 9
	// ==== SIZE METRICS =====
	// Pixels
	// Display name: "px"
	NumericalUnits_PIXELS NumericalUnits = 17
	// ==== POWER METRICS =====
	// Decibel Milliwatts
	// Display name: "dBm"
	NumericalUnits_DBM NumericalUnits = 18
	// Decibels / Meter ^ 2
	// Display name: "dBsm"
	NumericalUnits_DBSM NumericalUnits = 19
	// Decibel Watts
	// Display name: "dBW"
	NumericalUnits_DBW NumericalUnits = 20
	// Decibels
	// Display name: "dB"
	NumericalUnits_DB NumericalUnits = 21
	// ==== INERTIA METRICS ====
	// Kilogram * Meter ^ 2
	// Display name: "kg*m^2"
	NumericalUnits_KGM2 NumericalUnits = 23
	// ==== FORCE METRICS =====
	// Newtons / Meter
	// Display name: "N/m"
	NumericalUnits_NPM NumericalUnits = 24
	// Newton Seconds / Meter
	// Display name: "N*s/m"
	NumericalUnits_NSPM NumericalUnits = 25
	// Newtons
	// Display name: "N"
	NumericalUnits_N NumericalUnits = 26
	// Newton Meters
	// Display name: "Nm"
	NumericalUnits_NM NumericalUnits = 27
	// Newtons Meters / Second
	// Display name: "Nm/s"
	NumericalUnits_NMPS NumericalUnits = 28
)

var NumericalUnits_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "NO_UNITS",
	2:  "M",
	3:  "MPS",
	4:  "MPS2",
	14: "MPS3",
	15: "MM",
	16: "UM",
	5:  "S",
	6:  "RAD",
	10: "RADPS",
	11: "RADPS2",
	12: "RADPS3",
	22: "DEG",
	7:  "G",
	8:  "KG",
	9:  "HERTZ",
	17: "PIXELS",
	18: "DBM",
	19: "DBSM",
	20: "DBW",
	21: "DB",
	23: "KGM2",
	24: "NPM",
	25: "NSPM",
	26: "N",
	27: "NM",
	28: "NMPS",
}

var NumericalUnits_value = map[string]int32{
	"UNKNOWN":  0,
	"NO_UNITS": 1,
	"M":        2,
	"MPS":      3,
	"MPS2":     4,
	"MPS3":     14,
	"MM":       15,
	"UM":       16,
	"S":        5,
	"RAD":      6,
	"RADPS":    10,
	"RADPS2":   11,
	"RADPS3":   12,
	"DEG":      22,
	"G":        7,
	"KG":       8,
	"HERTZ":    9,
	"PIXELS":   17,
	"DBM":      18,
	"DBSM":     19,
	"DBW":      20,
	"DB":       21,
	"KGM2":     23,
	"NPM":      24,
	"NSPM":     25,
	"N":        26,
	"NM":       27,
	"NMPS":     28,
}

func (x NumericalUnits) String() string {
	return proto.EnumName(NumericalUnits_name, int32(x))
}

func (NumericalUnits) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{0}
}

// Custom FieldOptions
type FieldOptions struct {
	// Fields tagged with this will be omitted from generated schemas
	Ignore bool `protobuf:"varint,1,opt,name=ignore,proto3" json:"ignore,omitempty"`
	// Fields tagged with this will be marked as "required" in generated schemas
	Required bool `protobuf:"varint,2,opt,name=required,proto3" json:"required,omitempty"`
	// Fields tagged with this will constrain strings using the "minLength" keyword in generated schemas
	MinLength int32 `protobuf:"varint,3,opt,name=min_length,json=minLength,proto3" json:"min_length,omitempty"`
	// Fields tagged with this will constrain strings using the "maxLength" keyword in generated schemas
	MaxLength int32 `protobuf:"varint,4,opt,name=max_length,json=maxLength,proto3" json:"max_length,omitempty"`
	// Fields tagged with this will constrain strings using the "pattern" keyword in generated schemas
	Pattern              string   `protobuf:"bytes,5,opt,name=pattern,proto3" json:"pattern,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldOptions) Reset()         { *m = FieldOptions{} }
func (m *FieldOptions) String() string { return proto.CompactTextString(m) }
func (*FieldOptions) ProtoMessage()    {}
func (*FieldOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{0}
}

func (m *FieldOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldOptions.Unmarshal(m, b)
}
func (m *FieldOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldOptions.Marshal(b, m, deterministic)
}
func (m *FieldOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldOptions.Merge(m, src)
}
func (m *FieldOptions) XXX_Size() int {
	return xxx_messageInfo_FieldOptions.Size(m)
}
func (m *FieldOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldOptions.DiscardUnknown(m)
}

var xxx_messageInfo_FieldOptions proto.InternalMessageInfo

func (m *FieldOptions) GetIgnore() bool {
	if m != nil {
		return m.Ignore
	}
	return false
}

func (m *FieldOptions) GetRequired() bool {
	if m != nil {
		return m.Required
	}
	return false
}

func (m *FieldOptions) GetMinLength() int32 {
	if m != nil {
		return m.MinLength
	}
	return 0
}

func (m *FieldOptions) GetMaxLength() int32 {
	if m != nil {
		return m.MaxLength
	}
	return 0
}

func (m *FieldOptions) GetPattern() string {
	if m != nil {
		return m.Pattern
	}
	return ""
}

// Custom FileOptions
type FileOptions struct {
	// Files tagged with this will not be processed
	Ignore bool `protobuf:"varint,1,opt,name=ignore,proto3" json:"ignore,omitempty"`
	// Override the default file extension for schemas generated from this file
	Extension            string   `protobuf:"bytes,2,opt,name=extension,proto3" json:"extension,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileOptions) Reset()         { *m = FileOptions{} }
func (m *FileOptions) String() string { return proto.CompactTextString(m) }
func (*FileOptions) ProtoMessage()    {}
func (*FileOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{1}
}

func (m *FileOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileOptions.Unmarshal(m, b)
}
func (m *FileOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileOptions.Marshal(b, m, deterministic)
}
func (m *FileOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileOptions.Merge(m, src)
}
func (m *FileOptions) XXX_Size() int {
	return xxx_messageInfo_FileOptions.Size(m)
}
func (m *FileOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_FileOptions.DiscardUnknown(m)
}

var xxx_messageInfo_FileOptions proto.InternalMessageInfo

func (m *FileOptions) GetIgnore() bool {
	if m != nil {
		return m.Ignore
	}
	return false
}

func (m *FileOptions) GetExtension() string {
	if m != nil {
		return m.Extension
	}
	return ""
}

// Custom MessageOptions
type MessageOptions struct {
	// Messages tagged with this will not be processed
	Ignore bool `protobuf:"varint,1,opt,name=ignore,proto3" json:"ignore,omitempty"`
	// Messages tagged with this will have all fields marked as "required":
	AllFieldsRequired bool `protobuf:"varint,2,opt,name=all_fields_required,json=allFieldsRequired,proto3" json:"all_fields_required,omitempty"`
	// Messages tagged with this will additionally accept null values for all properties:
	AllowNullValues bool `protobuf:"varint,3,opt,name=allow_null_values,json=allowNullValues,proto3" json:"allow_null_values,omitempty"`
	// Messages tagged with this will have all fields marked as not allowing additional properties:
	DisallowAdditionalProperties bool `protobuf:"varint,4,opt,name=disallow_additional_properties,json=disallowAdditionalProperties,proto3" json:"disallow_additional_properties,omitempty"`
	// Messages tagged with this will have all nested enums encoded to use constants instead of simple types (supports value annotations):
	EnumsAsConstants     bool     `protobuf:"varint,5,opt,name=enums_as_constants,json=enumsAsConstants,proto3" json:"enums_as_constants,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageOptions) Reset()         { *m = MessageOptions{} }
func (m *MessageOptions) String() string { return proto.CompactTextString(m) }
func (*MessageOptions) ProtoMessage()    {}
func (*MessageOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{2}
}

func (m *MessageOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageOptions.Unmarshal(m, b)
}
func (m *MessageOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageOptions.Marshal(b, m, deterministic)
}
func (m *MessageOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageOptions.Merge(m, src)
}
func (m *MessageOptions) XXX_Size() int {
	return xxx_messageInfo_MessageOptions.Size(m)
}
func (m *MessageOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageOptions.DiscardUnknown(m)
}

var xxx_messageInfo_MessageOptions proto.InternalMessageInfo

func (m *MessageOptions) GetIgnore() bool {
	if m != nil {
		return m.Ignore
	}
	return false
}

func (m *MessageOptions) GetAllFieldsRequired() bool {
	if m != nil {
		return m.AllFieldsRequired
	}
	return false
}

func (m *MessageOptions) GetAllowNullValues() bool {
	if m != nil {
		return m.AllowNullValues
	}
	return false
}

func (m *MessageOptions) GetDisallowAdditionalProperties() bool {
	if m != nil {
		return m.DisallowAdditionalProperties
	}
	return false
}

func (m *MessageOptions) GetEnumsAsConstants() bool {
	if m != nil {
		return m.EnumsAsConstants
	}
	return false
}

// Custom EnumOptions
type EnumOptions struct {
	// Enums tagged with this will have be encoded to use constants instead of simple types (supports value annotations):
	EnumsAsConstants bool `protobuf:"varint,1,opt,name=enums_as_constants,json=enumsAsConstants,proto3" json:"enums_as_constants,omitempty"`
	// Enums tagged with this will only provide string values as options (not their numerical equivalents):
	EnumsAsStringsOnly bool `protobuf:"varint,2,opt,name=enums_as_strings_only,json=enumsAsStringsOnly,proto3" json:"enums_as_strings_only,omitempty"`
	// Enums tagged with this will have enum name prefix removed from values:
	EnumsTrimPrefix bool `protobuf:"varint,3,opt,name=enums_trim_prefix,json=enumsTrimPrefix,proto3" json:"enums_trim_prefix,omitempty"`
	// Enums tagged with this will not be processed
	Ignore               bool     `protobuf:"varint,4,opt,name=ignore,proto3" json:"ignore,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnumOptions) Reset()         { *m = EnumOptions{} }
func (m *EnumOptions) String() string { return proto.CompactTextString(m) }
func (*EnumOptions) ProtoMessage()    {}
func (*EnumOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_110d40819f1994f9, []int{3}
}

func (m *EnumOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumOptions.Unmarshal(m, b)
}
func (m *EnumOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumOptions.Marshal(b, m, deterministic)
}
func (m *EnumOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumOptions.Merge(m, src)
}
func (m *EnumOptions) XXX_Size() int {
	return xxx_messageInfo_EnumOptions.Size(m)
}
func (m *EnumOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumOptions.DiscardUnknown(m)
}

var xxx_messageInfo_EnumOptions proto.InternalMessageInfo

func (m *EnumOptions) GetEnumsAsConstants() bool {
	if m != nil {
		return m.EnumsAsConstants
	}
	return false
}

func (m *EnumOptions) GetEnumsAsStringsOnly() bool {
	if m != nil {
		return m.EnumsAsStringsOnly
	}
	return false
}

func (m *EnumOptions) GetEnumsTrimPrefix() bool {
	if m != nil {
		return m.EnumsTrimPrefix
	}
	return false
}

func (m *EnumOptions) GetIgnore() bool {
	if m != nil {
		return m.Ignore
	}
	return false
}

var E_FieldOptions = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*FieldOptions)(nil),
	Field:         1125,
	Name:          "protoc.gen.jsonschema.field_options",
	Tag:           "bytes,1125,opt,name=field_options",
	Filename:      "options.proto",
}

var E_ManualLink = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50015,
	Name:          "protoc.gen.jsonschema.manual_link",
	Tag:           "bytes,50015,opt,name=manual_link",
	Filename:      "options.proto",
}

var E_IgnoreInAutocomplete = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         50000,
	Name:          "protoc.gen.jsonschema.ignore_in_autocomplete",
	Tag:           "varint,50000,opt,name=ignore_in_autocomplete",
	Filename:      "options.proto",
}

var E_Units = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*NumericalUnits)(nil),
	Field:         50010,
	Name:          "protoc.gen.jsonschema.units",
	Tag:           "varint,50010,opt,name=units,enum=protoc.gen.jsonschema.NumericalUnits",
	Filename:      "options.proto",
}

var E_FileOptions = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*FileOptions)(nil),
	Field:         1126,
	Name:          "protoc.gen.jsonschema.file_options",
	Tag:           "bytes,1126,opt,name=file_options",
	Filename:      "options.proto",
}

var E_MessageOptions = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*MessageOptions)(nil),
	Field:         1127,
	Name:          "protoc.gen.jsonschema.message_options",
	Tag:           "bytes,1127,opt,name=message_options",
	Filename:      "options.proto",
}

var E_MessageManualLink = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50000,
	Name:          "protoc.gen.jsonschema.message_manual_link",
	Tag:           "bytes,50000,opt,name=message_manual_link",
	Filename:      "options.proto",
}

var E_EnumOptions = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*EnumOptions)(nil),
	Field:         1128,
	Name:          "protoc.gen.jsonschema.enum_options",
	Tag:           "bytes,1128,opt,name=enum_options",
	Filename:      "options.proto",
}

func init() {
	proto.RegisterEnum("protoc.gen.jsonschema.NumericalUnits", NumericalUnits_name, NumericalUnits_value)
	proto.RegisterType((*FieldOptions)(nil), "protoc.gen.jsonschema.FieldOptions")
	proto.RegisterType((*FileOptions)(nil), "protoc.gen.jsonschema.FileOptions")
	proto.RegisterType((*MessageOptions)(nil), "protoc.gen.jsonschema.MessageOptions")
	proto.RegisterType((*EnumOptions)(nil), "protoc.gen.jsonschema.EnumOptions")
	proto.RegisterExtension(E_FieldOptions)
	proto.RegisterExtension(E_ManualLink)
	proto.RegisterExtension(E_IgnoreInAutocomplete)
	proto.RegisterExtension(E_Units)
	proto.RegisterExtension(E_FileOptions)
	proto.RegisterExtension(E_MessageOptions)
	proto.RegisterExtension(E_MessageManualLink)
	proto.RegisterExtension(E_EnumOptions)
}

func init() { proto.RegisterFile("options.proto", fileDescriptor_110d40819f1994f9) }

var fileDescriptor_110d40819f1994f9 = []byte{
	// 848 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0x9e, 0x92, 0xd8, 0x91, 0x8f, 0x5d, 0x87, 0x61, 0x9a, 0x4c, 0xcb, 0xd2, 0xcd, 0xf0, 0x30,
	0x20, 0x28, 0x06, 0x15, 0x73, 0xee, 0x7c, 0x35, 0xa7, 0x4e, 0xb3, 0x20, 0x91, 0xe2, 0x49, 0xf1,
	0x3a, 0x14, 0x03, 0x04, 0xd5, 0xa6, 0x1d, 0xae, 0x14, 0xa5, 0x89, 0xd2, 0x96, 0xbe, 0x40, 0x1f,
	0x62, 0x7b, 0x8b, 0x3d, 0x49, 0xaf, 0x77, 0xb3, 0xab, 0xfd, 0x3c, 0xc2, 0x2e, 0x07, 0x92, 0x92,
	0x62, 0xaf, 0x49, 0x7d, 0x25, 0x9e, 0xf3, 0x7d, 0xdf, 0x39, 0x3c, 0x87, 0x3c, 0x14, 0x3c, 0x88,
	0x93, 0x8c, 0xc6, 0x5c, 0xd8, 0x49, 0x1a, 0x67, 0x31, 0xde, 0x55, 0x9f, 0x89, 0x3d, 0x27, 0xdc,
	0xfe, 0x41, 0xc4, 0x5c, 0x4c, 0xae, 0x49, 0x14, 0xee, 0x77, 0xe6, 0x71, 0x3c, 0x67, 0xe4, 0x89,
	0x42, 0x5f, 0xe6, 0xb3, 0x27, 0x53, 0x22, 0x26, 0x29, 0x4d, 0xb2, 0x38, 0xd5, 0xc2, 0xee, 0x2f,
	0x06, 0xb4, 0x9e, 0x51, 0xc2, 0xa6, 0x97, 0x3a, 0x1e, 0xde, 0x83, 0x3a, 0x9d, 0xf3, 0x38, 0x25,
	0x96, 0xd1, 0x31, 0x0e, 0x4d, 0xaf, 0xb0, 0xf0, 0x3e, 0x98, 0x29, 0xf9, 0x31, 0xa7, 0x29, 0x99,
	0x5a, 0x6b, 0x0a, 0xa9, 0x6c, 0xfc, 0x08, 0x20, 0xa2, 0x3c, 0x60, 0x84, 0xcf, 0xb3, 0x6b, 0x6b,
	0xbd, 0x63, 0x1c, 0xd6, 0xbc, 0x46, 0x44, 0xf9, 0x85, 0x72, 0x28, 0x38, 0xbc, 0x29, 0xe1, 0x8d,
	0x02, 0x0e, 0x6f, 0x0a, 0xd8, 0x82, 0xcd, 0x24, 0xcc, 0x32, 0x92, 0x72, 0xab, 0xd6, 0x31, 0x0e,
	0x1b, 0x5e, 0x69, 0x76, 0x9f, 0x42, 0xf3, 0x19, 0x65, 0x64, 0xd5, 0xd6, 0x0e, 0xa0, 0x41, 0x6e,
	0x32, 0xc2, 0x05, 0x8d, 0xb9, 0xda, 0x5b, 0xc3, 0xbb, 0x75, 0x74, 0xff, 0x35, 0xa0, 0xed, 0x10,
	0x21, 0xc2, 0xf9, 0xca, 0x40, 0x36, 0xec, 0x84, 0x8c, 0x05, 0x33, 0xd9, 0x0f, 0x11, 0xfc, 0xaf,
	0xdc, 0xed, 0x90, 0x31, 0xd5, 0x29, 0xe1, 0x95, 0x75, 0x3f, 0x06, 0xe9, 0x8c, 0x7f, 0x0e, 0x78,
	0xce, 0x58, 0xf0, 0x53, 0xc8, 0x72, 0x22, 0x54, 0xf9, 0xa6, 0xb7, 0xa5, 0x00, 0x37, 0x67, 0xec,
	0x5b, 0xe5, 0xc6, 0x43, 0xf8, 0x64, 0x4a, 0x85, 0xa6, 0x87, 0xd3, 0x29, 0x95, 0x3b, 0x09, 0x59,
	0x90, 0xa4, 0x71, 0x42, 0xd2, 0x8c, 0x12, 0xa1, 0x1a, 0x63, 0x7a, 0x07, 0x25, 0x6b, 0x50, 0x91,
	0x46, 0x15, 0x07, 0x7f, 0x01, 0x98, 0xf0, 0x3c, 0x12, 0x41, 0x28, 0x82, 0x49, 0xcc, 0x45, 0x16,
	0xf2, 0x4c, 0xa8, 0xb6, 0x99, 0x1e, 0x52, 0xc8, 0x40, 0x3c, 0x2d, 0xfd, 0xdd, 0xdf, 0x0c, 0x68,
	0x9e, 0xf0, 0x3c, 0x2a, 0xeb, 0xbe, 0x5b, 0x6d, 0xdc, 0xad, 0xc6, 0x5f, 0xc2, 0x6e, 0xc5, 0x16,
	0x59, 0x4a, 0xf9, 0x5c, 0x04, 0x31, 0x67, 0xaf, 0x8b, 0x7e, 0xe0, 0x42, 0xe0, 0x6b, 0xe8, 0x92,
	0xb3, 0xd7, 0xb2, 0x21, 0x5a, 0x92, 0xa5, 0x34, 0x0a, 0x92, 0x94, 0xcc, 0xe8, 0x4d, 0xd9, 0x10,
	0x05, 0x5c, 0xa5, 0x34, 0x1a, 0x29, 0xf7, 0xc2, 0x21, 0x6c, 0x2c, 0x1e, 0xc2, 0xe3, 0x5f, 0xd7,
	0xa0, 0xed, 0xe6, 0x11, 0x49, 0xe9, 0x24, 0x64, 0x63, 0x4e, 0x33, 0x81, 0x9b, 0xb0, 0x39, 0x76,
	0xcf, 0xdd, 0xcb, 0xe7, 0x2e, 0xfa, 0x00, 0xb7, 0xc0, 0x74, 0x2f, 0x83, 0xb1, 0x7b, 0x76, 0xe5,
	0x23, 0x03, 0xd7, 0xc0, 0x70, 0xd0, 0x1a, 0xde, 0x84, 0x75, 0x67, 0xe4, 0xa3, 0x75, 0x6c, 0xc2,
	0x86, 0x33, 0xf2, 0x7b, 0x68, 0xa3, 0x58, 0x1d, 0xa1, 0x36, 0xae, 0xc3, 0x9a, 0xe3, 0xa0, 0x2d,
	0xf9, 0x1d, 0x3b, 0x08, 0x49, 0x8d, 0x8f, 0x6a, 0x52, 0xe3, 0x0d, 0x86, 0xa8, 0x8e, 0x1b, 0x50,
	0xf3, 0x06, 0xc3, 0x91, 0x8f, 0x00, 0x03, 0xd4, 0xd5, 0xb2, 0x87, 0x9a, 0xd5, 0xfa, 0x08, 0xb5,
	0x24, 0x77, 0x78, 0x72, 0x8a, 0xf6, 0xa4, 0xf6, 0x14, 0x6d, 0xca, 0x50, 0xe7, 0xa7, 0xc8, 0x94,
	0xd2, 0xaf, 0x4f, 0xbc, 0xab, 0x17, 0xa8, 0x21, 0xe9, 0xa3, 0xb3, 0xef, 0x4e, 0x2e, 0x7c, 0xb4,
	0xad, 0xe8, 0xc7, 0x0e, 0xc2, 0x72, 0x13, 0xc3, 0x63, 0xdf, 0x41, 0x3b, 0xda, 0xf5, 0x1c, 0x3d,
	0x94, 0xd2, 0xe1, 0x31, 0xda, 0x95, 0xd0, 0xf9, 0xa9, 0xd3, 0x43, 0x1f, 0x4a, 0xc8, 0x1d, 0x39,
	0xc8, 0x92, 0x2e, 0xd7, 0x1f, 0x39, 0xe8, 0x23, 0x99, 0xc6, 0x45, 0xfb, 0x92, 0xeb, 0x3a, 0xe8,
	0x63, 0x05, 0xc8, 0xfa, 0x0e, 0xfa, 0xd7, 0xf0, 0x40, 0x5d, 0xcf, 0xa0, 0x98, 0x7f, 0xfc, 0xc8,
	0xd6, 0x33, 0x6e, 0x97, 0x33, 0x6e, 0x2f, 0x8e, 0xb3, 0xf5, 0xa7, 0xd9, 0x31, 0x0e, 0x9b, 0xbd,
	0xcf, 0xec, 0x3b, 0x1f, 0x88, 0x25, 0xae, 0xd7, 0x9a, 0x2d, 0x58, 0xfd, 0xaf, 0xa0, 0x19, 0x85,
	0x3c, 0x0f, 0x59, 0xc0, 0x28, 0x7f, 0xb5, 0x2a, 0xcf, 0x1f, 0x6f, 0xd6, 0xd5, 0xe0, 0x81, 0xd6,
	0x5c, 0x50, 0xfe, 0xaa, 0x3f, 0x86, 0x3d, 0x7d, 0xa6, 0x01, 0xe5, 0x41, 0x98, 0x67, 0xf1, 0x24,
	0x8e, 0x12, 0x46, 0x32, 0xb2, 0x2a, 0xd8, 0xdb, 0x37, 0xfa, 0xce, 0x3c, 0xd4, 0xf2, 0x33, 0x3e,
	0x58, 0x10, 0xf7, 0xbf, 0x87, 0x5a, 0xae, 0xae, 0xc5, 0x8a, 0x28, 0xbf, 0xab, 0x28, 0xed, 0xde,
	0xe7, 0xf7, 0xd4, 0xbe, 0x7c, 0xc9, 0x3c, 0x1d, 0xb4, 0x3f, 0x85, 0xd6, 0x8c, 0x32, 0x52, 0xf5,
	0xf7, 0xe0, 0x8e, 0x24, 0xd5, 0x93, 0x64, 0xfd, 0xa5, 0xdb, 0xdb, 0xbd, 0xb7, 0xbd, 0x15, 0xd5,
	0x6b, 0xce, 0x6e, 0x8d, 0x7e, 0x02, 0x5b, 0x91, 0x7e, 0x93, 0xaa, 0x44, 0x9f, 0xbe, 0x93, 0x68,
	0xf9, 0xd5, 0xb2, 0xfe, 0xd6, 0xb9, 0xee, 0x2b, 0x67, 0x99, 0xed, 0xb5, 0xa3, 0x25, 0xbb, 0xff,
	0x0d, 0xec, 0x94, 0x19, 0x17, 0x8f, 0x75, 0x65, 0xd6, 0xb7, 0xc5, 0xc1, 0x6e, 0x17, 0x6a, 0xe7,
	0xf6, 0x7c, 0xa7, 0xd0, 0x92, 0x43, 0xfd, 0x9e, 0x56, 0x2d, 0x3c, 0x3e, 0xd6, 0x3f, 0xef, 0x6f,
	0xd5, 0x02, 0xd5, 0x6b, 0x92, 0x5b, 0xe3, 0x18, 0x5e, 0x98, 0xb6, 0xfe, 0x81, 0x89, 0x97, 0x75,
	0xf5, 0x3d, 0xfa, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x87, 0x0e, 0x96, 0x46, 0xfe, 0x06, 0x00, 0x00,
}
