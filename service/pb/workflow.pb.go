// Code generated by protoc-gen-go.
// source: workflow.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WatchedDirectory_WatchedDirectoryType int32

const (
	WatchedDirectory_UNKNOWN  WatchedDirectory_WatchedDirectoryType = 0
	WatchedDirectory_TRANSFER WatchedDirectory_WatchedDirectoryType = 1
	WatchedDirectory_SIP      WatchedDirectory_WatchedDirectoryType = 2
	WatchedDirectory_DIP      WatchedDirectory_WatchedDirectoryType = 3
)

var WatchedDirectory_WatchedDirectoryType_name = map[int32]string{
	0: "UNKNOWN",
	1: "TRANSFER",
	2: "SIP",
	3: "DIP",
}
var WatchedDirectory_WatchedDirectoryType_value = map[string]int32{
	"UNKNOWN":  0,
	"TRANSFER": 1,
	"SIP":      2,
	"DIP":      3,
}

func (x WatchedDirectory_WatchedDirectoryType) String() string {
	return proto.EnumName(WatchedDirectory_WatchedDirectoryType_name, int32(x))
}
func (WatchedDirectory_WatchedDirectoryType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor2, []int{3, 0}
}

type WorkflowGetRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *WorkflowGetRequest) Reset()                    { *m = WorkflowGetRequest{} }
func (m *WorkflowGetRequest) String() string            { return proto.CompactTextString(m) }
func (*WorkflowGetRequest) ProtoMessage()               {}
func (*WorkflowGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *WorkflowGetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type WorkflowGetResponse struct {
	Workflow *WorkflowData `protobuf:"bytes,1,opt,name=workflow" json:"workflow,omitempty"`
	Error    string        `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *WorkflowGetResponse) Reset()                    { *m = WorkflowGetResponse{} }
func (m *WorkflowGetResponse) String() string            { return proto.CompactTextString(m) }
func (*WorkflowGetResponse) ProtoMessage()               {}
func (*WorkflowGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *WorkflowGetResponse) GetWorkflow() *WorkflowData {
	if m != nil {
		return m.Workflow
	}
	return nil
}

func (m *WorkflowGetResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type WorkflowData struct {
	WatchedDirectories []*WatchedDirectory `protobuf:"bytes,1,rep,name=watchedDirectories" json:"watchedDirectories,omitempty"`
	Chains             map[string]*Chain   `protobuf:"bytes,2,rep,name=chains" json:"chains,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Links              map[string]*Link    `protobuf:"bytes,3,rep,name=links" json:"links,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *WorkflowData) Reset()                    { *m = WorkflowData{} }
func (m *WorkflowData) String() string            { return proto.CompactTextString(m) }
func (*WorkflowData) ProtoMessage()               {}
func (*WorkflowData) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *WorkflowData) GetWatchedDirectories() []*WatchedDirectory {
	if m != nil {
		return m.WatchedDirectories
	}
	return nil
}

func (m *WorkflowData) GetChains() map[string]*Chain {
	if m != nil {
		return m.Chains
	}
	return nil
}

func (m *WorkflowData) GetLinks() map[string]*Link {
	if m != nil {
		return m.Links
	}
	return nil
}

type WatchedDirectory struct {
	Type     WatchedDirectory_WatchedDirectoryType `protobuf:"varint,1,opt,name=type,enum=pb.WatchedDirectory_WatchedDirectoryType" json:"type,omitempty"`
	Path     string                                `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	OnlyDirs bool                                  `protobuf:"varint,3,opt,name=onlyDirs" json:"onlyDirs,omitempty"`
	ChainId  string                                `protobuf:"bytes,4,opt,name=chainId" json:"chainId,omitempty"`
}

func (m *WatchedDirectory) Reset()                    { *m = WatchedDirectory{} }
func (m *WatchedDirectory) String() string            { return proto.CompactTextString(m) }
func (*WatchedDirectory) ProtoMessage()               {}
func (*WatchedDirectory) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *WatchedDirectory) GetType() WatchedDirectory_WatchedDirectoryType {
	if m != nil {
		return m.Type
	}
	return WatchedDirectory_UNKNOWN
}

func (m *WatchedDirectory) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *WatchedDirectory) GetOnlyDirs() bool {
	if m != nil {
		return m.OnlyDirs
	}
	return false
}

func (m *WatchedDirectory) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

type Chain struct {
	Description map[string]string `protobuf:"bytes,1,rep,name=description" json:"description,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	LinkId      string            `protobuf:"bytes,2,opt,name=linkId" json:"linkId,omitempty"`
	Id          string            `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
}

func (m *Chain) Reset()                    { *m = Chain{} }
func (m *Chain) String() string            { return proto.CompactTextString(m) }
func (*Chain) ProtoMessage()               {}
func (*Chain) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *Chain) GetDescription() map[string]string {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *Chain) GetLinkId() string {
	if m != nil {
		return m.LinkId
	}
	return ""
}

func (m *Chain) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Link struct {
	Description       map[string]string    `protobuf:"bytes,1,rep,name=description" json:"description,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Group             map[string]string    `protobuf:"bytes,2,rep,name=group" json:"group,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	FallbackJobStatus Job_Status           `protobuf:"varint,3,opt,name=fallbackJobStatus,enum=pb.Job_Status" json:"fallbackJobStatus,omitempty"`
	FallbackLinkId    string               `protobuf:"bytes,4,opt,name=fallbackLinkId" json:"fallbackLinkId,omitempty"`
	ExitCodes         []*Link_LinkExitCode `protobuf:"bytes,5,rep,name=exitCodes" json:"exitCodes,omitempty"`
	Config            *Link_LinkConfig     `protobuf:"bytes,6,opt,name=config" json:"config,omitempty"`
	Id                string               `protobuf:"bytes,7,opt,name=id" json:"id,omitempty"`
}

func (m *Link) Reset()                    { *m = Link{} }
func (m *Link) String() string            { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()               {}
func (*Link) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *Link) GetDescription() map[string]string {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *Link) GetGroup() map[string]string {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *Link) GetFallbackJobStatus() Job_Status {
	if m != nil {
		return m.FallbackJobStatus
	}
	return Job_UNKNOWN
}

func (m *Link) GetFallbackLinkId() string {
	if m != nil {
		return m.FallbackLinkId
	}
	return ""
}

func (m *Link) GetExitCodes() []*Link_LinkExitCode {
	if m != nil {
		return m.ExitCodes
	}
	return nil
}

func (m *Link) GetConfig() *Link_LinkConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Link) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Link_LinkExitCode struct {
	Code      uint32     `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	JobStatus Job_Status `protobuf:"varint,2,opt,name=jobStatus,enum=pb.Job_Status" json:"jobStatus,omitempty"`
	LinkId    string     `protobuf:"bytes,3,opt,name=linkId" json:"linkId,omitempty"`
}

func (m *Link_LinkExitCode) Reset()                    { *m = Link_LinkExitCode{} }
func (m *Link_LinkExitCode) String() string            { return proto.CompactTextString(m) }
func (*Link_LinkExitCode) ProtoMessage()               {}
func (*Link_LinkExitCode) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5, 0} }

func (m *Link_LinkExitCode) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Link_LinkExitCode) GetJobStatus() Job_Status {
	if m != nil {
		return m.JobStatus
	}
	return Job_UNKNOWN
}

func (m *Link_LinkExitCode) GetLinkId() string {
	if m != nil {
		return m.LinkId
	}
	return ""
}

type Link_LinkConfig struct {
	Manager string `protobuf:"bytes,1,opt,name=manager" json:"manager,omitempty"`
	Model   string `protobuf:"bytes,2,opt,name=model" json:"model,omitempty"`
	// Types that are valid to be assigned to Attrs:
	//	*Link_LinkConfig_Standard_
	//	*Link_LinkConfig_UserChoice_
	//	*Link_LinkConfig_UserChoiceDict_
	//	*Link_LinkConfig_SetUnitVar_
	//	*Link_LinkConfig_GetUnitVar_
	//	*Link_LinkConfig_Magic_
	Attrs isLink_LinkConfig_Attrs `protobuf_oneof:"attrs"`
}

func (m *Link_LinkConfig) Reset()                    { *m = Link_LinkConfig{} }
func (m *Link_LinkConfig) String() string            { return proto.CompactTextString(m) }
func (*Link_LinkConfig) ProtoMessage()               {}
func (*Link_LinkConfig) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5, 1} }

type isLink_LinkConfig_Attrs interface {
	isLink_LinkConfig_Attrs()
}

type Link_LinkConfig_Standard_ struct {
	Standard *Link_LinkConfig_Standard `protobuf:"bytes,3,opt,name=standard,oneof"`
}
type Link_LinkConfig_UserChoice_ struct {
	UserChoice *Link_LinkConfig_UserChoice `protobuf:"bytes,4,opt,name=userChoice,oneof"`
}
type Link_LinkConfig_UserChoiceDict_ struct {
	UserChoiceDict *Link_LinkConfig_UserChoiceDict `protobuf:"bytes,5,opt,name=userChoiceDict,oneof"`
}
type Link_LinkConfig_SetUnitVar_ struct {
	SetUnitVar *Link_LinkConfig_SetUnitVar `protobuf:"bytes,6,opt,name=setUnitVar,oneof"`
}
type Link_LinkConfig_GetUnitVar_ struct {
	GetUnitVar *Link_LinkConfig_GetUnitVar `protobuf:"bytes,7,opt,name=getUnitVar,oneof"`
}
type Link_LinkConfig_Magic_ struct {
	Magic *Link_LinkConfig_Magic `protobuf:"bytes,8,opt,name=magic,oneof"`
}

func (*Link_LinkConfig_Standard_) isLink_LinkConfig_Attrs()       {}
func (*Link_LinkConfig_UserChoice_) isLink_LinkConfig_Attrs()     {}
func (*Link_LinkConfig_UserChoiceDict_) isLink_LinkConfig_Attrs() {}
func (*Link_LinkConfig_SetUnitVar_) isLink_LinkConfig_Attrs()     {}
func (*Link_LinkConfig_GetUnitVar_) isLink_LinkConfig_Attrs()     {}
func (*Link_LinkConfig_Magic_) isLink_LinkConfig_Attrs()          {}

func (m *Link_LinkConfig) GetAttrs() isLink_LinkConfig_Attrs {
	if m != nil {
		return m.Attrs
	}
	return nil
}

func (m *Link_LinkConfig) GetManager() string {
	if m != nil {
		return m.Manager
	}
	return ""
}

func (m *Link_LinkConfig) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *Link_LinkConfig) GetStandard() *Link_LinkConfig_Standard {
	if x, ok := m.GetAttrs().(*Link_LinkConfig_Standard_); ok {
		return x.Standard
	}
	return nil
}

func (m *Link_LinkConfig) GetUserChoice() *Link_LinkConfig_UserChoice {
	if x, ok := m.GetAttrs().(*Link_LinkConfig_UserChoice_); ok {
		return x.UserChoice
	}
	return nil
}

func (m *Link_LinkConfig) GetUserChoiceDict() *Link_LinkConfig_UserChoiceDict {
	if x, ok := m.GetAttrs().(*Link_LinkConfig_UserChoiceDict_); ok {
		return x.UserChoiceDict
	}
	return nil
}

func (m *Link_LinkConfig) GetSetUnitVar() *Link_LinkConfig_SetUnitVar {
	if x, ok := m.GetAttrs().(*Link_LinkConfig_SetUnitVar_); ok {
		return x.SetUnitVar
	}
	return nil
}

func (m *Link_LinkConfig) GetGetUnitVar() *Link_LinkConfig_GetUnitVar {
	if x, ok := m.GetAttrs().(*Link_LinkConfig_GetUnitVar_); ok {
		return x.GetUnitVar
	}
	return nil
}

func (m *Link_LinkConfig) GetMagic() *Link_LinkConfig_Magic {
	if x, ok := m.GetAttrs().(*Link_LinkConfig_Magic_); ok {
		return x.Magic
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Link_LinkConfig) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Link_LinkConfig_OneofMarshaler, _Link_LinkConfig_OneofUnmarshaler, _Link_LinkConfig_OneofSizer, []interface{}{
		(*Link_LinkConfig_Standard_)(nil),
		(*Link_LinkConfig_UserChoice_)(nil),
		(*Link_LinkConfig_UserChoiceDict_)(nil),
		(*Link_LinkConfig_SetUnitVar_)(nil),
		(*Link_LinkConfig_GetUnitVar_)(nil),
		(*Link_LinkConfig_Magic_)(nil),
	}
}

func _Link_LinkConfig_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Link_LinkConfig)
	// attrs
	switch x := m.Attrs.(type) {
	case *Link_LinkConfig_Standard_:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Standard); err != nil {
			return err
		}
	case *Link_LinkConfig_UserChoice_:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UserChoice); err != nil {
			return err
		}
	case *Link_LinkConfig_UserChoiceDict_:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UserChoiceDict); err != nil {
			return err
		}
	case *Link_LinkConfig_SetUnitVar_:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SetUnitVar); err != nil {
			return err
		}
	case *Link_LinkConfig_GetUnitVar_:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GetUnitVar); err != nil {
			return err
		}
	case *Link_LinkConfig_Magic_:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Magic); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Link_LinkConfig.Attrs has unexpected type %T", x)
	}
	return nil
}

func _Link_LinkConfig_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Link_LinkConfig)
	switch tag {
	case 3: // attrs.standard
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Link_LinkConfig_Standard)
		err := b.DecodeMessage(msg)
		m.Attrs = &Link_LinkConfig_Standard_{msg}
		return true, err
	case 4: // attrs.userChoice
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Link_LinkConfig_UserChoice)
		err := b.DecodeMessage(msg)
		m.Attrs = &Link_LinkConfig_UserChoice_{msg}
		return true, err
	case 5: // attrs.userChoiceDict
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Link_LinkConfig_UserChoiceDict)
		err := b.DecodeMessage(msg)
		m.Attrs = &Link_LinkConfig_UserChoiceDict_{msg}
		return true, err
	case 6: // attrs.setUnitVar
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Link_LinkConfig_SetUnitVar)
		err := b.DecodeMessage(msg)
		m.Attrs = &Link_LinkConfig_SetUnitVar_{msg}
		return true, err
	case 7: // attrs.getUnitVar
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Link_LinkConfig_GetUnitVar)
		err := b.DecodeMessage(msg)
		m.Attrs = &Link_LinkConfig_GetUnitVar_{msg}
		return true, err
	case 8: // attrs.magic
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Link_LinkConfig_Magic)
		err := b.DecodeMessage(msg)
		m.Attrs = &Link_LinkConfig_Magic_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Link_LinkConfig_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Link_LinkConfig)
	// attrs
	switch x := m.Attrs.(type) {
	case *Link_LinkConfig_Standard_:
		s := proto.Size(x.Standard)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Link_LinkConfig_UserChoice_:
		s := proto.Size(x.UserChoice)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Link_LinkConfig_UserChoiceDict_:
		s := proto.Size(x.UserChoiceDict)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Link_LinkConfig_SetUnitVar_:
		s := proto.Size(x.SetUnitVar)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Link_LinkConfig_GetUnitVar_:
		s := proto.Size(x.GetUnitVar)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Link_LinkConfig_Magic_:
		s := proto.Size(x.Magic)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Link_LinkConfig_Standard struct {
	Execute            string `protobuf:"bytes,1,opt,name=execute" json:"execute,omitempty"`
	Arguments          string `protobuf:"bytes,2,opt,name=arguments" json:"arguments,omitempty"`
	FilterSubdir       string `protobuf:"bytes,3,opt,name=filterSubdir" json:"filterSubdir,omitempty"`
	FilterFileEnd      string `protobuf:"bytes,4,opt,name=filterFileEnd" json:"filterFileEnd,omitempty"`
	StdoutFile         string `protobuf:"bytes,5,opt,name=stdoutFile" json:"stdoutFile,omitempty"`
	StderrFile         string `protobuf:"bytes,6,opt,name=stderrFile" json:"stderrFile,omitempty"`
	RequiresOutputLock bool   `protobuf:"varint,7,opt,name=requiresOutputLock" json:"requiresOutputLock,omitempty"`
}

func (m *Link_LinkConfig_Standard) Reset()                    { *m = Link_LinkConfig_Standard{} }
func (m *Link_LinkConfig_Standard) String() string            { return proto.CompactTextString(m) }
func (*Link_LinkConfig_Standard) ProtoMessage()               {}
func (*Link_LinkConfig_Standard) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5, 1, 0} }

func (m *Link_LinkConfig_Standard) GetExecute() string {
	if m != nil {
		return m.Execute
	}
	return ""
}

func (m *Link_LinkConfig_Standard) GetArguments() string {
	if m != nil {
		return m.Arguments
	}
	return ""
}

func (m *Link_LinkConfig_Standard) GetFilterSubdir() string {
	if m != nil {
		return m.FilterSubdir
	}
	return ""
}

func (m *Link_LinkConfig_Standard) GetFilterFileEnd() string {
	if m != nil {
		return m.FilterFileEnd
	}
	return ""
}

func (m *Link_LinkConfig_Standard) GetStdoutFile() string {
	if m != nil {
		return m.StdoutFile
	}
	return ""
}

func (m *Link_LinkConfig_Standard) GetStderrFile() string {
	if m != nil {
		return m.StderrFile
	}
	return ""
}

func (m *Link_LinkConfig_Standard) GetRequiresOutputLock() bool {
	if m != nil {
		return m.RequiresOutputLock
	}
	return false
}

type Link_LinkConfig_UserChoice struct {
	ChainIds []string `protobuf:"bytes,1,rep,name=chainIds" json:"chainIds,omitempty"`
}

func (m *Link_LinkConfig_UserChoice) Reset()         { *m = Link_LinkConfig_UserChoice{} }
func (m *Link_LinkConfig_UserChoice) String() string { return proto.CompactTextString(m) }
func (*Link_LinkConfig_UserChoice) ProtoMessage()    {}
func (*Link_LinkConfig_UserChoice) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{5, 1, 1}
}

func (m *Link_LinkConfig_UserChoice) GetChainIds() []string {
	if m != nil {
		return m.ChainIds
	}
	return nil
}

type Link_LinkConfig_UserChoiceDict struct {
	Replacements []*Link_LinkConfig_UserChoiceDict_Replacement `protobuf:"bytes,1,rep,name=replacements" json:"replacements,omitempty"`
}

func (m *Link_LinkConfig_UserChoiceDict) Reset()         { *m = Link_LinkConfig_UserChoiceDict{} }
func (m *Link_LinkConfig_UserChoiceDict) String() string { return proto.CompactTextString(m) }
func (*Link_LinkConfig_UserChoiceDict) ProtoMessage()    {}
func (*Link_LinkConfig_UserChoiceDict) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{5, 1, 2}
}

func (m *Link_LinkConfig_UserChoiceDict) GetReplacements() []*Link_LinkConfig_UserChoiceDict_Replacement {
	if m != nil {
		return m.Replacements
	}
	return nil
}

type Link_LinkConfig_UserChoiceDict_Replacement struct {
	Id          string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Description map[string]string `protobuf:"bytes,2,rep,name=description" json:"description,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Items       map[string]string `protobuf:"bytes,3,rep,name=items" json:"items,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Link_LinkConfig_UserChoiceDict_Replacement) Reset() {
	*m = Link_LinkConfig_UserChoiceDict_Replacement{}
}
func (m *Link_LinkConfig_UserChoiceDict_Replacement) String() string {
	return proto.CompactTextString(m)
}
func (*Link_LinkConfig_UserChoiceDict_Replacement) ProtoMessage() {}
func (*Link_LinkConfig_UserChoiceDict_Replacement) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{5, 1, 2, 0}
}

func (m *Link_LinkConfig_UserChoiceDict_Replacement) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Link_LinkConfig_UserChoiceDict_Replacement) GetDescription() map[string]string {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *Link_LinkConfig_UserChoiceDict_Replacement) GetItems() map[string]string {
	if m != nil {
		return m.Items
	}
	return nil
}

type Link_LinkConfig_SetUnitVar struct {
	Variable      string `protobuf:"bytes,1,opt,name=variable" json:"variable,omitempty"`
	VariableValue string `protobuf:"bytes,2,opt,name=variableValue" json:"variableValue,omitempty"`
	ChainId       string `protobuf:"bytes,3,opt,name=chainId" json:"chainId,omitempty"`
}

func (m *Link_LinkConfig_SetUnitVar) Reset()         { *m = Link_LinkConfig_SetUnitVar{} }
func (m *Link_LinkConfig_SetUnitVar) String() string { return proto.CompactTextString(m) }
func (*Link_LinkConfig_SetUnitVar) ProtoMessage()    {}
func (*Link_LinkConfig_SetUnitVar) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{5, 1, 3}
}

func (m *Link_LinkConfig_SetUnitVar) GetVariable() string {
	if m != nil {
		return m.Variable
	}
	return ""
}

func (m *Link_LinkConfig_SetUnitVar) GetVariableValue() string {
	if m != nil {
		return m.VariableValue
	}
	return ""
}

func (m *Link_LinkConfig_SetUnitVar) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

type Link_LinkConfig_GetUnitVar struct {
	Variable string `protobuf:"bytes,1,opt,name=variable" json:"variable,omitempty"`
	ChainId  string `protobuf:"bytes,2,opt,name=chainId" json:"chainId,omitempty"`
}

func (m *Link_LinkConfig_GetUnitVar) Reset()         { *m = Link_LinkConfig_GetUnitVar{} }
func (m *Link_LinkConfig_GetUnitVar) String() string { return proto.CompactTextString(m) }
func (*Link_LinkConfig_GetUnitVar) ProtoMessage()    {}
func (*Link_LinkConfig_GetUnitVar) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{5, 1, 4}
}

func (m *Link_LinkConfig_GetUnitVar) GetVariable() string {
	if m != nil {
		return m.Variable
	}
	return ""
}

func (m *Link_LinkConfig_GetUnitVar) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

type Link_LinkConfig_Magic struct {
	LinkId string `protobuf:"bytes,1,opt,name=linkId" json:"linkId,omitempty"`
}

func (m *Link_LinkConfig_Magic) Reset()                    { *m = Link_LinkConfig_Magic{} }
func (m *Link_LinkConfig_Magic) String() string            { return proto.CompactTextString(m) }
func (*Link_LinkConfig_Magic) ProtoMessage()               {}
func (*Link_LinkConfig_Magic) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5, 1, 5} }

func (m *Link_LinkConfig_Magic) GetLinkId() string {
	if m != nil {
		return m.LinkId
	}
	return ""
}

func init() {
	proto.RegisterType((*WorkflowGetRequest)(nil), "pb.WorkflowGetRequest")
	proto.RegisterType((*WorkflowGetResponse)(nil), "pb.WorkflowGetResponse")
	proto.RegisterType((*WorkflowData)(nil), "pb.WorkflowData")
	proto.RegisterType((*WatchedDirectory)(nil), "pb.WatchedDirectory")
	proto.RegisterType((*Chain)(nil), "pb.Chain")
	proto.RegisterType((*Link)(nil), "pb.Link")
	proto.RegisterType((*Link_LinkExitCode)(nil), "pb.Link.LinkExitCode")
	proto.RegisterType((*Link_LinkConfig)(nil), "pb.Link.LinkConfig")
	proto.RegisterType((*Link_LinkConfig_Standard)(nil), "pb.Link.LinkConfig.Standard")
	proto.RegisterType((*Link_LinkConfig_UserChoice)(nil), "pb.Link.LinkConfig.UserChoice")
	proto.RegisterType((*Link_LinkConfig_UserChoiceDict)(nil), "pb.Link.LinkConfig.UserChoiceDict")
	proto.RegisterType((*Link_LinkConfig_UserChoiceDict_Replacement)(nil), "pb.Link.LinkConfig.UserChoiceDict.Replacement")
	proto.RegisterType((*Link_LinkConfig_SetUnitVar)(nil), "pb.Link.LinkConfig.SetUnitVar")
	proto.RegisterType((*Link_LinkConfig_GetUnitVar)(nil), "pb.Link.LinkConfig.GetUnitVar")
	proto.RegisterType((*Link_LinkConfig_Magic)(nil), "pb.Link.LinkConfig.Magic")
	proto.RegisterEnum("pb.WatchedDirectory_WatchedDirectoryType", WatchedDirectory_WatchedDirectoryType_name, WatchedDirectory_WatchedDirectoryType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Workflow service

type WorkflowClient interface {
	WorkflowGet(ctx context.Context, in *WorkflowGetRequest, opts ...grpc.CallOption) (*WorkflowGetResponse, error)
}

type workflowClient struct {
	cc *grpc.ClientConn
}

func NewWorkflowClient(cc *grpc.ClientConn) WorkflowClient {
	return &workflowClient{cc}
}

func (c *workflowClient) WorkflowGet(ctx context.Context, in *WorkflowGetRequest, opts ...grpc.CallOption) (*WorkflowGetResponse, error) {
	out := new(WorkflowGetResponse)
	err := grpc.Invoke(ctx, "/pb.Workflow/WorkflowGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Workflow service

type WorkflowServer interface {
	WorkflowGet(context.Context, *WorkflowGetRequest) (*WorkflowGetResponse, error)
}

func RegisterWorkflowServer(s *grpc.Server, srv WorkflowServer) {
	s.RegisterService(&_Workflow_serviceDesc, srv)
}

func _Workflow_WorkflowGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServer).WorkflowGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Workflow/WorkflowGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServer).WorkflowGet(ctx, req.(*WorkflowGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Workflow_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Workflow",
	HandlerType: (*WorkflowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WorkflowGet",
			Handler:    _Workflow_WorkflowGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workflow.proto",
}

func init() { proto.RegisterFile("workflow.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 1085 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x56, 0xdf, 0x72, 0xdb, 0xc4,
	0x17, 0xb6, 0xe5, 0x38, 0x96, 0x8f, 0x53, 0xff, 0xfc, 0xdb, 0x84, 0xa2, 0x8a, 0x4e, 0xda, 0xf1,
	0x74, 0x98, 0x74, 0xe8, 0x78, 0x06, 0x97, 0x8b, 0x52, 0x0a, 0x94, 0xc6, 0x69, 0x92, 0x12, 0x92,
	0xce, 0xa6, 0x69, 0x87, 0xcb, 0xb5, 0xb4, 0xb1, 0x97, 0x28, 0x92, 0xba, 0x5a, 0x35, 0xf5, 0xeb,
	0x70, 0xcf, 0x0d, 0x2f, 0xc1, 0x05, 0x2f, 0xc1, 0x4b, 0x70, 0xc7, 0x05, 0xb3, 0x7f, 0xf4, 0xcf,
	0x11, 0x84, 0x0c, 0xdc, 0x78, 0x74, 0xce, 0x7e, 0xdf, 0xe7, 0xd5, 0xf9, 0x8e, 0xce, 0x2e, 0xf4,
	0x2f, 0x22, 0x7e, 0x76, 0x1a, 0x44, 0x17, 0xa3, 0x98, 0x47, 0x22, 0x42, 0x56, 0x3c, 0x75, 0x07,
	0x31, 0x8f, 0x3c, 0x9a, 0x24, 0x2c, 0x9c, 0xe9, 0xec, 0xf0, 0x1e, 0xa0, 0x37, 0x06, 0xb7, 0x4b,
	0x05, 0xa6, 0x6f, 0x53, 0x9a, 0x08, 0xd4, 0x07, 0x8b, 0xf9, 0x4e, 0xf3, 0x6e, 0x73, 0xab, 0x8b,
	0x2d, 0xe6, 0x0f, 0xbf, 0x87, 0xf5, 0x0a, 0x2a, 0x89, 0xa3, 0x30, 0xa1, 0xe8, 0x01, 0xd8, 0xd9,
	0x9f, 0x28, 0x70, 0x6f, 0x3c, 0x18, 0xc5, 0xd3, 0x51, 0x06, 0x9d, 0x10, 0x41, 0x70, 0x8e, 0x40,
	0x1b, 0xd0, 0xa6, 0x9c, 0x47, 0xdc, 0xb1, 0x94, 0xae, 0x0e, 0x86, 0xbf, 0x58, 0xb0, 0x56, 0x26,
	0xa0, 0x09, 0xa0, 0x0b, 0x22, 0xbc, 0x39, 0xf5, 0x27, 0x8c, 0x53, 0x4f, 0x44, 0x9c, 0xd1, 0xc4,
	0x69, 0xde, 0x6d, 0x6d, 0xf5, 0xc6, 0x1b, 0x4a, 0xbe, 0xba, 0xba, 0xc0, 0x35, 0x78, 0xf4, 0x19,
	0xac, 0x7a, 0x73, 0xc2, 0xc2, 0xc4, 0xb1, 0x14, 0xf3, 0xf6, 0xf2, 0xc6, 0x46, 0xdb, 0x6a, 0x79,
	0x27, 0x14, 0x7c, 0x81, 0x0d, 0x16, 0x7d, 0x0a, 0xed, 0x80, 0x85, 0x67, 0x89, 0xd3, 0x52, 0xa4,
	0x8f, 0x2e, 0x91, 0x0e, 0xe4, 0xaa, 0xe6, 0x68, 0xa4, 0x3b, 0x81, 0x5e, 0x49, 0x09, 0x0d, 0xa0,
	0x75, 0x46, 0x17, 0xa6, 0x74, 0xf2, 0x11, 0xdd, 0x81, 0xf6, 0x3b, 0x12, 0xa4, 0x54, 0xbd, 0x76,
	0x6f, 0xdc, 0x95, 0x9a, 0x8a, 0x81, 0x75, 0xfe, 0xb1, 0xf5, 0xa8, 0xe9, 0x3e, 0x03, 0x28, 0xa4,
	0x6b, 0x44, 0x36, 0xab, 0x22, 0xb6, 0x14, 0x91, 0x84, 0x92, 0xc6, 0xf0, 0xb7, 0x26, 0x0c, 0x96,
	0x6b, 0x83, 0xbe, 0x84, 0x15, 0xb1, 0x88, 0xa9, 0xd2, 0xea, 0x8f, 0xef, 0xd7, 0xd5, 0xef, 0x52,
	0xe2, 0xd5, 0x22, 0xa6, 0x58, 0xd1, 0x10, 0x82, 0x95, 0x98, 0x88, 0xb9, 0xb1, 0x4c, 0x3d, 0x23,
	0x17, 0xec, 0x28, 0x0c, 0x16, 0x13, 0xc6, 0x65, 0x9d, 0x9a, 0x5b, 0x36, 0xce, 0x63, 0xe4, 0x40,
	0x47, 0x95, 0x72, 0xdf, 0x77, 0x56, 0x14, 0x25, 0x0b, 0x87, 0xdb, 0xb0, 0x51, 0xf7, 0x3f, 0xa8,
	0x07, 0x9d, 0x93, 0xc3, 0x6f, 0x0f, 0x8f, 0xde, 0x1c, 0x0e, 0x1a, 0x68, 0x0d, 0xec, 0x57, 0xf8,
	0x9b, 0xc3, 0xe3, 0xe7, 0x3b, 0x78, 0xd0, 0x44, 0x1d, 0x68, 0x1d, 0xef, 0xbf, 0x1c, 0x58, 0xf2,
	0x61, 0xb2, 0xff, 0x72, 0xd0, 0x1a, 0xfe, 0xd4, 0x84, 0xb6, 0xaa, 0x1d, 0x7a, 0x02, 0x3d, 0x9f,
	0x26, 0x1e, 0x67, 0xb1, 0x60, 0x51, 0x68, 0xda, 0xc3, 0xcd, 0x6b, 0x3b, 0x9a, 0x14, 0x8b, 0xda,
	0xae, 0x32, 0x1c, 0xdd, 0x84, 0x55, 0xe9, 0xde, 0xbe, 0x6f, 0x5e, 0xcc, 0x44, 0xa6, 0xef, 0x5b,
	0x59, 0xdf, 0xbb, 0x5f, 0xc1, 0x60, 0x59, 0xa8, 0xc6, 0x9c, 0x8d, 0xb2, 0x39, 0xdd, 0xb2, 0x25,
	0x3f, 0xfe, 0x0f, 0x56, 0xa4, 0x4d, 0xe8, 0x8b, 0xba, 0xed, 0xde, 0xca, 0x5c, 0xbc, 0x62, 0xb7,
	0xf7, 0xa1, 0x3d, 0xe3, 0x51, 0x1a, 0x9b, 0x56, 0x5e, 0xcf, 0x69, 0xbb, 0x32, 0x6b, 0xba, 0x51,
	0x21, 0xd0, 0x13, 0xf8, 0xff, 0x29, 0x09, 0x82, 0x29, 0xf1, 0xce, 0x5e, 0x44, 0xd3, 0x63, 0x41,
	0x44, 0xaa, 0x4d, 0xea, 0x8f, 0xfb, 0x92, 0xf6, 0x22, 0x9a, 0x8e, 0x74, 0x16, 0x5f, 0x06, 0xa2,
	0x8f, 0xa1, 0x9f, 0x25, 0x0f, 0x74, 0x79, 0xb4, 0x89, 0x4b, 0x59, 0xf4, 0x10, 0xba, 0xf4, 0x3d,
	0x13, 0xdb, 0x91, 0x4f, 0x13, 0xa7, 0xad, 0x36, 0xf5, 0x41, 0xbe, 0x29, 0xf9, 0xb3, 0x63, 0x56,
	0x71, 0x81, 0x43, 0x9f, 0xc0, 0xaa, 0x17, 0x85, 0xa7, 0x6c, 0xe6, 0xac, 0xaa, 0x1e, 0x5e, 0xaf,
	0x30, 0xb6, 0xd5, 0x12, 0x36, 0x10, 0x63, 0x44, 0x27, 0x37, 0x62, 0x0e, 0x6b, 0x65, 0x5d, 0xd9,
	0x97, 0x5e, 0xe4, 0xeb, 0xb6, 0xbe, 0x81, 0xd5, 0x33, 0x7a, 0x00, 0xdd, 0x1f, 0xf2, 0x77, 0xb6,
	0x6a, 0xdf, 0xb9, 0x00, 0x94, 0x5a, 0xa0, 0x55, 0x6e, 0x01, 0xf7, 0x77, 0xd0, 0x9f, 0xa2, 0xde,
	0x90, 0x6c, 0xe8, 0x73, 0x12, 0x92, 0x19, 0xe5, 0xc6, 0xf1, 0x2c, 0x94, 0xae, 0x9f, 0x47, 0x3e,
	0x0d, 0x32, 0xd7, 0x55, 0x80, 0x1e, 0x83, 0x9d, 0x08, 0x12, 0xfa, 0x84, 0x6b, 0x61, 0x33, 0x79,
	0x96, 0xde, 0x53, 0xee, 0x47, 0x61, 0xf6, 0x1a, 0x38, 0xc7, 0xa3, 0xa7, 0x00, 0x69, 0x42, 0xf9,
	0xf6, 0x3c, 0x62, 0x1e, 0x55, 0xa5, 0xef, 0x8d, 0x37, 0xeb, 0xd8, 0x27, 0x39, 0x6a, 0xaf, 0x81,
	0x4b, 0x1c, 0x74, 0x00, 0xfd, 0x22, 0x9a, 0x30, 0x4f, 0x38, 0x6d, 0xa5, 0x32, 0xfc, 0x7b, 0x15,
	0x89, 0xdc, 0x6b, 0xe0, 0x25, 0xae, 0xdc, 0x4f, 0x42, 0xc5, 0x49, 0xc8, 0xc4, 0x6b, 0xc2, 0x8d,
	0x6b, 0xb5, 0xfb, 0x39, 0xce, 0x51, 0x72, 0x3f, 0x05, 0x47, 0x2a, 0xcc, 0x0a, 0x85, 0xce, 0x5f,
	0x2b, 0xec, 0x56, 0x14, 0x0a, 0x8e, 0x9c, 0xc8, 0xe7, 0x64, 0xc6, 0x3c, 0xc7, 0x56, 0xe4, 0x5b,
	0x75, 0xe4, 0xef, 0x24, 0x60, 0xaf, 0x81, 0x35, 0xd2, 0xfd, 0xa3, 0x09, 0x76, 0x56, 0x5f, 0xe9,
	0x1f, 0x7d, 0x4f, 0xbd, 0x54, 0xd0, 0xcc, 0x3f, 0x13, 0xa2, 0xdb, 0xd0, 0x25, 0x7c, 0x96, 0x9e,
	0xd3, 0x50, 0x24, 0xc6, 0xc3, 0x22, 0x81, 0x86, 0xb0, 0x76, 0xca, 0x02, 0x41, 0xf9, 0x71, 0x3a,
	0xf5, 0x19, 0x37, 0x4d, 0x52, 0xc9, 0xa1, 0x7b, 0x70, 0x43, 0xc7, 0xcf, 0x59, 0x40, 0x77, 0xc2,
	0xec, 0x6b, 0xa9, 0x26, 0xd1, 0x26, 0x40, 0x22, 0xfc, 0x28, 0x15, 0x32, 0xa1, 0xfc, 0xe8, 0xe2,
	0x52, 0xc6, 0xac, 0x53, 0xae, 0x08, 0xaa, 0xca, 0x7a, 0xdd, 0x64, 0xd0, 0x08, 0x10, 0xa7, 0x6f,
	0x53, 0xc6, 0x69, 0x72, 0x94, 0x8a, 0x38, 0x15, 0x07, 0x91, 0x77, 0xa6, 0x6a, 0x69, 0xe3, 0x9a,
	0x15, 0x77, 0x0b, 0xa0, 0x70, 0x56, 0x0e, 0x6b, 0x33, 0x81, 0xf5, 0x19, 0xda, 0xc5, 0x79, 0xec,
	0xfe, 0xdc, 0x82, 0x7e, 0xb5, 0x09, 0x10, 0x86, 0x35, 0x4e, 0xe3, 0x80, 0x78, 0x54, 0xd7, 0x45,
	0x0f, 0xaa, 0xd1, 0xd5, 0xed, 0x33, 0xc2, 0x05, 0x0d, 0x57, 0x34, 0xdc, 0x5f, 0x2d, 0xe8, 0x95,
	0x56, 0x97, 0x2f, 0x17, 0x88, 0x54, 0x67, 0xa3, 0x1e, 0x72, 0x5f, 0x5f, 0xef, 0x2f, 0xaf, 0x98,
	0xa0, 0x47, 0xd0, 0x66, 0x82, 0x9e, 0x67, 0xe7, 0xfa, 0xe7, 0xd7, 0x14, 0xdf, 0x97, 0x5c, 0x33,
	0x67, 0x95, 0xce, 0xbf, 0x3d, 0x18, 0xdc, 0x47, 0x00, 0x85, 0xe8, 0xb5, 0x98, 0x73, 0x80, 0xe2,
	0x73, 0x93, 0xf6, 0xbe, 0x23, 0x9c, 0x91, 0x69, 0x90, 0xf5, 0x77, 0x1e, 0xcb, 0xf6, 0xcc, 0x9e,
	0x5f, 0x97, 0xb4, 0xaa, 0xc9, 0xf2, 0x89, 0xdd, 0xaa, 0x9c, 0xd8, 0xf2, 0x4e, 0xb2, 0xfb, 0xcf,
	0xfe, 0xa9, 0xa4, 0x61, 0x55, 0x35, 0xee, 0x40, 0x5b, 0x7d, 0x9d, 0xa5, 0x71, 0xdb, 0x2c, 0x8f,
	0xdb, 0x67, 0x1d, 0x68, 0x13, 0x21, 0xf8, 0x7f, 0x52, 0xd1, 0xe2, 0x38, 0xbc, 0x0e, 0x73, 0x7c,
	0x00, 0x76, 0x76, 0xc7, 0x43, 0x4f, 0xa1, 0x57, 0xba, 0xe8, 0xa2, 0x9b, 0xe5, 0x0b, 0x60, 0x71,
	0x3f, 0x76, 0x3f, 0xbc, 0x94, 0xd7, 0x37, 0xe2, 0x61, 0x63, 0xba, 0xaa, 0xee, 0xd5, 0x0f, 0xff,
	0x0c, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x6e, 0x06, 0x67, 0x7f, 0x0b, 0x00, 0x00,
}
