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

type Link struct {
	Description       map[string]string    `protobuf:"bytes,1,rep,name=description" json:"description,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Group             map[string]string    `protobuf:"bytes,2,rep,name=group" json:"group,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	FallbackJobStatus Job_Status           `protobuf:"varint,3,opt,name=fallbackJobStatus,enum=pb.Job_Status" json:"fallbackJobStatus,omitempty"`
	FallbackLinkId    string               `protobuf:"bytes,4,opt,name=fallbackLinkId" json:"fallbackLinkId,omitempty"`
	ExitCodes         []*Link_LinkExitCode `protobuf:"bytes,5,rep,name=exitCodes" json:"exitCodes,omitempty"`
	Config            *Link_LinkConfig     `protobuf:"bytes,6,opt,name=config" json:"config,omitempty"`
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
	FilterFileStart    string `protobuf:"bytes,4,opt,name=filterFileStart" json:"filterFileStart,omitempty"`
	FilterFileEnd      string `protobuf:"bytes,5,opt,name=filterFileEnd" json:"filterFileEnd,omitempty"`
	StdoutFile         string `protobuf:"bytes,6,opt,name=stdoutFile" json:"stdoutFile,omitempty"`
	StderrFile         string `protobuf:"bytes,7,opt,name=stderrFile" json:"stderrFile,omitempty"`
	RequiresOutputLock bool   `protobuf:"varint,8,opt,name=requiresOutputLock" json:"requiresOutputLock,omitempty"`
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

func (m *Link_LinkConfig_Standard) GetFilterFileStart() string {
	if m != nil {
		return m.FilterFileStart
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
	// 1088 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x72, 0xdb, 0x44,
	0x14, 0xb6, 0xe5, 0x38, 0xb1, 0x8f, 0x52, 0xd7, 0x6c, 0x43, 0x51, 0x45, 0x27, 0xed, 0x78, 0x3a,
	0x4c, 0x3a, 0x74, 0x3c, 0x83, 0xcb, 0x45, 0x29, 0x05, 0x4a, 0xe3, 0x34, 0x49, 0x09, 0x49, 0x67,
	0xdd, 0xb4, 0xc3, 0xe5, 0x5a, 0xda, 0xd8, 0x4b, 0x14, 0x49, 0x5d, 0xad, 0x9a, 0xfa, 0x61, 0x18,
	0xee, 0x78, 0x00, 0xae, 0xb9, 0xe7, 0x82, 0x97, 0xe0, 0x51, 0x98, 0xfd, 0xd1, 0x9f, 0x23, 0x08,
	0x19, 0xb8, 0xf1, 0xe8, 0x9c, 0xfd, 0xbe, 0xcf, 0xab, 0xf3, 0x1d, 0x9d, 0x5d, 0xe8, 0x9d, 0x47,
	0xfc, 0xf4, 0x24, 0x88, 0xce, 0x87, 0x31, 0x8f, 0x44, 0x84, 0xac, 0x78, 0xea, 0xf6, 0x63, 0x1e,
	0x79, 0x34, 0x49, 0x58, 0x38, 0xd3, 0xd9, 0xc1, 0x3d, 0x40, 0x6f, 0x0c, 0x6e, 0x97, 0x0a, 0x4c,
	0xdf, 0xa6, 0x34, 0x11, 0xa8, 0x07, 0x16, 0xf3, 0x9d, 0xe6, 0xdd, 0xe6, 0x56, 0x17, 0x5b, 0xcc,
	0x1f, 0xfc, 0x00, 0x37, 0x2a, 0xa8, 0x24, 0x8e, 0xc2, 0x84, 0xa2, 0x07, 0xd0, 0xc9, 0xfe, 0x44,
	0x81, 0xed, 0x51, 0x7f, 0x18, 0x4f, 0x87, 0x19, 0x74, 0x4c, 0x04, 0xc1, 0x39, 0x02, 0x6d, 0x40,
	0x9b, 0x72, 0x1e, 0x71, 0xc7, 0x52, 0xba, 0x3a, 0x18, 0xfc, 0x6e, 0xc1, 0x7a, 0x99, 0x80, 0xc6,
	0x80, 0xce, 0x89, 0xf0, 0xe6, 0xd4, 0x1f, 0x33, 0x4e, 0x3d, 0x11, 0x71, 0x46, 0x13, 0xa7, 0x79,
	0xb7, 0xb5, 0x65, 0x8f, 0x36, 0x94, 0x7c, 0x75, 0x75, 0x81, 0x6b, 0xf0, 0xe8, 0x73, 0x58, 0xf5,
	0xe6, 0x84, 0x85, 0x89, 0x63, 0x29, 0xe6, 0xed, 0xe5, 0x8d, 0x0d, 0xb7, 0xd5, 0xf2, 0x4e, 0x28,
	0xf8, 0x02, 0x1b, 0x2c, 0xfa, 0x0c, 0xda, 0x01, 0x0b, 0x4f, 0x13, 0xa7, 0xa5, 0x48, 0x1f, 0x5f,
	0x20, 0x1d, 0xc8, 0x55, 0xcd, 0xd1, 0x48, 0x77, 0x0c, 0x76, 0x49, 0x09, 0xf5, 0xa1, 0x75, 0x4a,
	0x17, 0xa6, 0x74, 0xf2, 0x11, 0xdd, 0x81, 0xf6, 0x3b, 0x12, 0xa4, 0x54, 0xbd, 0xb6, 0x3d, 0xea,
	0x4a, 0x4d, 0xc5, 0xc0, 0x3a, 0xff, 0xd8, 0x7a, 0xd4, 0x74, 0x9f, 0x01, 0x14, 0xd2, 0x35, 0x22,
	0x9b, 0x55, 0x91, 0x8e, 0x14, 0x91, 0x84, 0x92, 0xc6, 0xe0, 0xcf, 0x26, 0xf4, 0x97, 0x6b, 0x83,
	0xbe, 0x82, 0x15, 0xb1, 0x88, 0xa9, 0xd2, 0xea, 0x8d, 0xee, 0xd7, 0xd5, 0xef, 0x42, 0xe2, 0xd5,
	0x22, 0xa6, 0x58, 0xd1, 0x10, 0x82, 0x95, 0x98, 0x88, 0xb9, 0xb1, 0x4c, 0x3d, 0x23, 0x17, 0x3a,
	0x51, 0x18, 0x2c, 0xc6, 0x8c, 0xcb, 0x3a, 0x35, 0xb7, 0x3a, 0x38, 0x8f, 0x91, 0x03, 0x6b, 0xaa,
	0x94, 0xfb, 0xbe, 0xb3, 0xa2, 0x28, 0x59, 0x38, 0xd8, 0x86, 0x8d, 0xba, 0xff, 0x41, 0x36, 0xac,
	0x1d, 0x1f, 0x7e, 0x77, 0x78, 0xf4, 0xe6, 0xb0, 0xdf, 0x40, 0xeb, 0xd0, 0x79, 0x85, 0xbf, 0x3d,
	0x9c, 0x3c, 0xdf, 0xc1, 0xfd, 0x26, 0x5a, 0x83, 0xd6, 0x64, 0xff, 0x65, 0xdf, 0x92, 0x0f, 0xe3,
	0xfd, 0x97, 0xfd, 0xd6, 0xe0, 0xa7, 0x26, 0xb4, 0x55, 0xed, 0xd0, 0x13, 0xb0, 0x7d, 0x9a, 0x78,
	0x9c, 0xc5, 0x82, 0x45, 0xa1, 0x69, 0x0f, 0x37, 0xaf, 0xed, 0x70, 0x5c, 0x2c, 0x6a, 0xbb, 0xca,
	0x70, 0x74, 0x13, 0x56, 0xa5, 0x7b, 0xfb, 0xbe, 0x79, 0x31, 0x13, 0xb9, 0x5f, 0x43, 0x7f, 0x99,
	0x58, 0x63, 0xc6, 0x46, 0xd9, 0x8c, 0x6e, 0xd9, 0x82, 0xdf, 0xae, 0xc3, 0x8a, 0xb4, 0x05, 0x7d,
	0x59, 0xb7, 0xbd, 0x5b, 0x99, 0x6b, 0x97, 0xec, 0xee, 0x3e, 0xb4, 0x67, 0x3c, 0x4a, 0x63, 0xd3,
	0xba, 0x37, 0x72, 0xda, 0xae, 0xcc, 0x9a, 0xee, 0x53, 0x08, 0xf4, 0x04, 0x3e, 0x38, 0x21, 0x41,
	0x30, 0x25, 0xde, 0xe9, 0x8b, 0x68, 0x3a, 0x11, 0x44, 0xa4, 0xda, 0x94, 0xde, 0xa8, 0x27, 0x69,
	0x2f, 0xa2, 0xe9, 0x50, 0x67, 0xf1, 0x45, 0x20, 0xfa, 0x04, 0x7a, 0x59, 0xf2, 0x40, 0x97, 0x43,
	0x9b, 0xb6, 0x94, 0x45, 0x0f, 0xa1, 0x4b, 0xdf, 0x33, 0xb1, 0x1d, 0xf9, 0x34, 0x71, 0xda, 0x6a,
	0x53, 0x1f, 0xe6, 0x9b, 0x92, 0x3f, 0x3b, 0x66, 0x15, 0x17, 0x38, 0xf4, 0x29, 0xac, 0x7a, 0x51,
	0x78, 0xc2, 0x66, 0xce, 0xaa, 0xea, 0xd9, 0x1b, 0x15, 0xc6, 0xb6, 0x5a, 0xc2, 0x06, 0xe2, 0xce,
	0x61, 0xbd, 0xac, 0x23, 0xfb, 0xce, 0x8b, 0x7c, 0xdd, 0xb6, 0xd7, 0xb0, 0x7a, 0x46, 0x0f, 0xa0,
	0xfb, 0x63, 0xfe, 0x8e, 0x56, 0xed, 0x3b, 0x16, 0x80, 0x92, 0xc5, 0xad, 0x8a, 0xc5, 0x3f, 0xdb,
	0xfa, 0x53, 0xd3, 0x1b, 0x90, 0x0d, 0x7b, 0x46, 0x42, 0x32, 0xa3, 0xdc, 0x38, 0x9c, 0x85, 0xd2,
	0xe5, 0xb3, 0xc8, 0xa7, 0x41, 0xe6, 0xb2, 0x0a, 0xd0, 0x63, 0xe8, 0x24, 0x82, 0x84, 0x3e, 0xe1,
	0x5a, 0xd8, 0x4c, 0x96, 0xa5, 0xf7, 0x92, 0xfb, 0x51, 0x98, 0xbd, 0x06, 0xce, 0xf1, 0xe8, 0x29,
	0x40, 0x9a, 0x50, 0xbe, 0x3d, 0x8f, 0x98, 0x47, 0x55, 0xa9, 0xed, 0xd1, 0x66, 0x1d, 0xfb, 0x38,
	0x47, 0xed, 0x35, 0x70, 0x89, 0x83, 0x0e, 0xa0, 0x57, 0x44, 0x63, 0xe6, 0x09, 0xa7, 0xad, 0x54,
	0x06, 0xff, 0xac, 0x22, 0x91, 0x7b, 0x0d, 0xbc, 0xc4, 0x95, 0xfb, 0x49, 0xa8, 0x38, 0x0e, 0x99,
	0x78, 0x4d, 0xb8, 0x71, 0xa9, 0x76, 0x3f, 0x93, 0x1c, 0x25, 0xf7, 0x53, 0x70, 0xa4, 0xc2, 0xac,
	0x50, 0x58, 0xfb, 0x7b, 0x85, 0xdd, 0x8a, 0x42, 0xc1, 0x91, 0x13, 0xf7, 0x8c, 0xcc, 0x98, 0xe7,
	0x74, 0x14, 0xf9, 0x56, 0x1d, 0xf9, 0x7b, 0x09, 0xd8, 0x6b, 0x60, 0x8d, 0x74, 0x7f, 0xb1, 0xa0,
	0x93, 0xd5, 0x57, 0xfa, 0x47, 0xdf, 0x53, 0x2f, 0x15, 0x34, 0xf3, 0xcf, 0x84, 0xe8, 0x36, 0x74,
	0x09, 0x9f, 0xa5, 0x67, 0x34, 0x14, 0x89, 0xf1, 0xb0, 0x48, 0xa0, 0x01, 0xac, 0x9f, 0xb0, 0x40,
	0x50, 0x3e, 0x49, 0xa7, 0x3e, 0xe3, 0xa6, 0x49, 0x2a, 0x39, 0xb4, 0x05, 0xd7, 0x75, 0xfc, 0x9c,
	0x05, 0x74, 0x22, 0x08, 0x17, 0xe6, 0xfb, 0x58, 0x4e, 0xa3, 0x7b, 0x70, 0xad, 0x48, 0xed, 0x84,
	0xbe, 0xb2, 0xa5, 0x8b, 0xab, 0x49, 0xb4, 0x09, 0x90, 0x08, 0x3f, 0x4a, 0x85, 0x4c, 0xa8, 0x7a,
	0x77, 0x71, 0x29, 0x63, 0xd6, 0x29, 0x57, 0x04, 0x55, 0x4d, 0xbd, 0x6e, 0x32, 0x68, 0x08, 0x88,
	0xd3, 0xb7, 0x29, 0xe3, 0x34, 0x39, 0x4a, 0x45, 0x9c, 0x8a, 0x83, 0xc8, 0x3b, 0x55, 0x85, 0xeb,
	0xe0, 0x9a, 0x15, 0x77, 0x0b, 0xa0, 0xe8, 0x01, 0x39, 0xb6, 0xcd, 0x2c, 0xd6, 0xa7, 0x69, 0x17,
	0xe7, 0xb1, 0xfb, 0x6b, 0x0b, 0x7a, 0xd5, 0x76, 0x41, 0x18, 0xd6, 0x39, 0x8d, 0x03, 0xe2, 0x51,
	0x5d, 0x41, 0x3d, 0xc2, 0x86, 0x97, 0x37, 0xda, 0x10, 0x17, 0x34, 0x5c, 0xd1, 0x70, 0xff, 0xb0,
	0xc0, 0x2e, 0xad, 0x2e, 0x5f, 0x33, 0x10, 0xa9, 0x4e, 0x4d, 0x3d, 0xfe, 0xbe, 0xb9, 0xda, 0x5f,
	0x5e, 0x32, 0x5b, 0x8f, 0xa0, 0xcd, 0x04, 0x3d, 0xcb, 0x4e, 0xf8, 0x2f, 0xae, 0x28, 0xbe, 0x2f,
	0xb9, 0x66, 0x02, 0x2b, 0x9d, 0xff, 0x7a, 0x64, 0xb8, 0x8f, 0x00, 0x0a, 0xd1, 0x2b, 0x31, 0xe7,
	0x00, 0xc5, 0x87, 0x29, 0xed, 0x7d, 0x47, 0x38, 0x23, 0xd3, 0x20, 0xfb, 0x12, 0xf2, 0x58, 0xb6,
	0x67, 0xf6, 0xfc, 0xba, 0xa4, 0x55, 0x4d, 0x96, 0xcf, 0xee, 0x56, 0xe5, 0xec, 0x96, 0xb7, 0x93,
	0xdd, 0x7f, 0xf7, 0x4f, 0x25, 0x0d, 0xab, 0xaa, 0x71, 0x07, 0xda, 0xea, 0x3b, 0x2e, 0x0d, 0xe6,
	0x66, 0x79, 0x30, 0x3f, 0x5b, 0x83, 0x36, 0x11, 0x82, 0xff, 0x2f, 0x15, 0x2d, 0x0e, 0xca, 0xab,
	0x30, 0x47, 0x07, 0xd0, 0xc9, 0x6e, 0x7b, 0xe8, 0x29, 0xd8, 0xa5, 0x2b, 0x2f, 0xba, 0x59, 0xbe,
	0x0a, 0x16, 0x37, 0x65, 0xf7, 0xa3, 0x0b, 0x79, 0x7d, 0x37, 0x1e, 0x34, 0xa6, 0xab, 0xea, 0x86,
	0xfd, 0xf0, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x02, 0xf5, 0xea, 0x89, 0x0b, 0x00, 0x00,
}
