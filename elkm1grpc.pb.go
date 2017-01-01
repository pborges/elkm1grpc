// Code generated by protoc-gen-go.
// source: elkm1grpc.proto
// DO NOT EDIT!

/*
Package elkm1grpc is a generated protocol buffer package.

It is generated from these files:
	elkm1grpc.proto

It has these top-level messages:
	ZoneChangeArgs
	ArmingChangeArgs
	ArmingStatusReport
	ArmingStatusArea
	ArmArgs
	ZoneChangeUpdate
*/
package elkm1grpc

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ArmingStatusArea_ArmingStatus int32

const (
	ArmingStatusArea_DISARMED                   ArmingStatusArea_ArmingStatus = 0
	ArmingStatusArea_ARMED_AWAY                 ArmingStatusArea_ArmingStatus = 1
	ArmingStatusArea_ARMED_STAY                 ArmingStatusArea_ArmingStatus = 2
	ArmingStatusArea_ARMED_STAY_INSTANT         ArmingStatusArea_ArmingStatus = 3
	ArmingStatusArea_ARMED_TO_NIGHT             ArmingStatusArea_ArmingStatus = 4
	ArmingStatusArea_ARMED_TO_NIGHT_INSTANT     ArmingStatusArea_ArmingStatus = 5
	ArmingStatusArea_ARMED_TO_VACATION          ArmingStatusArea_ArmingStatus = 6
	ArmingStatusArea_ARM_STEP_TO_NEXT_AWAY_MODE ArmingStatusArea_ArmingStatus = 7
	ArmingStatusArea_ARM_STEP_TO_NEXT_STAY_MODE ArmingStatusArea_ArmingStatus = 8
	ArmingStatusArea_FORCE_AWAY                 ArmingStatusArea_ArmingStatus = 9
	ArmingStatusArea_FORCE_STAY                 ArmingStatusArea_ArmingStatus = 10
)

var ArmingStatusArea_ArmingStatus_name = map[int32]string{
	0:  "DISARMED",
	1:  "ARMED_AWAY",
	2:  "ARMED_STAY",
	3:  "ARMED_STAY_INSTANT",
	4:  "ARMED_TO_NIGHT",
	5:  "ARMED_TO_NIGHT_INSTANT",
	6:  "ARMED_TO_VACATION",
	7:  "ARM_STEP_TO_NEXT_AWAY_MODE",
	8:  "ARM_STEP_TO_NEXT_STAY_MODE",
	9:  "FORCE_AWAY",
	10: "FORCE_STAY",
}
var ArmingStatusArea_ArmingStatus_value = map[string]int32{
	"DISARMED":                   0,
	"ARMED_AWAY":                 1,
	"ARMED_STAY":                 2,
	"ARMED_STAY_INSTANT":         3,
	"ARMED_TO_NIGHT":             4,
	"ARMED_TO_NIGHT_INSTANT":     5,
	"ARMED_TO_VACATION":          6,
	"ARM_STEP_TO_NEXT_AWAY_MODE": 7,
	"ARM_STEP_TO_NEXT_STAY_MODE": 8,
	"FORCE_AWAY":                 9,
	"FORCE_STAY":                 10,
}

func (x ArmingStatusArea_ArmingStatus) String() string {
	return proto.EnumName(ArmingStatusArea_ArmingStatus_name, int32(x))
}
func (ArmingStatusArea_ArmingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

type ArmingStatusArea_ArmUpState int32

const (
	ArmingStatusArea_NOT_READY        ArmingStatusArea_ArmUpState = 0
	ArmingStatusArea_READY            ArmingStatusArea_ArmUpState = 1
	ArmingStatusArea_READY_FORCE      ArmingStatusArea_ArmUpState = 2
	ArmingStatusArea_ARMED_EXIT_TIMER ArmingStatusArea_ArmUpState = 3
	ArmingStatusArea_ARMED            ArmingStatusArea_ArmUpState = 4
	ArmingStatusArea_FORCE_ARMED      ArmingStatusArea_ArmUpState = 5
	ArmingStatusArea_ARMED_BYPASS     ArmingStatusArea_ArmUpState = 6
)

var ArmingStatusArea_ArmUpState_name = map[int32]string{
	0: "NOT_READY",
	1: "READY",
	2: "READY_FORCE",
	3: "ARMED_EXIT_TIMER",
	4: "ARMED",
	5: "FORCE_ARMED",
	6: "ARMED_BYPASS",
}
var ArmingStatusArea_ArmUpState_value = map[string]int32{
	"NOT_READY":        0,
	"READY":            1,
	"READY_FORCE":      2,
	"ARMED_EXIT_TIMER": 3,
	"ARMED":            4,
	"FORCE_ARMED":      5,
	"ARMED_BYPASS":     6,
}

func (x ArmingStatusArea_ArmUpState) String() string {
	return proto.EnumName(ArmingStatusArea_ArmUpState_name, int32(x))
}
func (ArmingStatusArea_ArmUpState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 1}
}

type ArmingStatusArea_AlarmState int32

const (
	ArmingStatusArea_NO_ALARM               ArmingStatusArea_AlarmState = 0
	ArmingStatusArea_ENTRANCE_DELAY         ArmingStatusArea_AlarmState = 1
	ArmingStatusArea_ABORT_DELAY            ArmingStatusArea_AlarmState = 2
	ArmingStatusArea_ALARM_FIRE             ArmingStatusArea_AlarmState = 3
	ArmingStatusArea_ALARM_MEDICAL          ArmingStatusArea_AlarmState = 4
	ArmingStatusArea_ALARM_POLICE           ArmingStatusArea_AlarmState = 5
	ArmingStatusArea_ALARM_BURGLAR          ArmingStatusArea_AlarmState = 6
	ArmingStatusArea_ALARM_AUX1             ArmingStatusArea_AlarmState = 7
	ArmingStatusArea_ALARM_AUX2             ArmingStatusArea_AlarmState = 8
	ArmingStatusArea_ALARM_AUX3             ArmingStatusArea_AlarmState = 9
	ArmingStatusArea_ALARM_AUX4             ArmingStatusArea_AlarmState = 10
	ArmingStatusArea_ALARM_CARBON_MONOXIDE  ArmingStatusArea_AlarmState = 11
	ArmingStatusArea_ALARM_EMERGENCY        ArmingStatusArea_AlarmState = 12
	ArmingStatusArea_ALARM_FREEZE           ArmingStatusArea_AlarmState = 13
	ArmingStatusArea_ALARM_GAS              ArmingStatusArea_AlarmState = 14
	ArmingStatusArea_ALARM_HEAT             ArmingStatusArea_AlarmState = 15
	ArmingStatusArea_ALARM_WATER            ArmingStatusArea_AlarmState = 16
	ArmingStatusArea_ALARM_FIRE_SUPERVISORY ArmingStatusArea_AlarmState = 17
	ArmingStatusArea_ALARM_FIRE_VERIFY      ArmingStatusArea_AlarmState = 18
)

var ArmingStatusArea_AlarmState_name = map[int32]string{
	0:  "NO_ALARM",
	1:  "ENTRANCE_DELAY",
	2:  "ABORT_DELAY",
	3:  "ALARM_FIRE",
	4:  "ALARM_MEDICAL",
	5:  "ALARM_POLICE",
	6:  "ALARM_BURGLAR",
	7:  "ALARM_AUX1",
	8:  "ALARM_AUX2",
	9:  "ALARM_AUX3",
	10: "ALARM_AUX4",
	11: "ALARM_CARBON_MONOXIDE",
	12: "ALARM_EMERGENCY",
	13: "ALARM_FREEZE",
	14: "ALARM_GAS",
	15: "ALARM_HEAT",
	16: "ALARM_WATER",
	17: "ALARM_FIRE_SUPERVISORY",
	18: "ALARM_FIRE_VERIFY",
}
var ArmingStatusArea_AlarmState_value = map[string]int32{
	"NO_ALARM":               0,
	"ENTRANCE_DELAY":         1,
	"ABORT_DELAY":            2,
	"ALARM_FIRE":             3,
	"ALARM_MEDICAL":          4,
	"ALARM_POLICE":           5,
	"ALARM_BURGLAR":          6,
	"ALARM_AUX1":             7,
	"ALARM_AUX2":             8,
	"ALARM_AUX3":             9,
	"ALARM_AUX4":             10,
	"ALARM_CARBON_MONOXIDE":  11,
	"ALARM_EMERGENCY":        12,
	"ALARM_FREEZE":           13,
	"ALARM_GAS":              14,
	"ALARM_HEAT":             15,
	"ALARM_WATER":            16,
	"ALARM_FIRE_SUPERVISORY": 17,
	"ALARM_FIRE_VERIFY":      18,
}

func (x ArmingStatusArea_AlarmState) String() string {
	return proto.EnumName(ArmingStatusArea_AlarmState_name, int32(x))
}
func (ArmingStatusArea_AlarmState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 2}
}

type ZoneChangeUpdate_ZoneStatus int32

const (
	ZoneChangeUpdate_NORMAL   ZoneChangeUpdate_ZoneStatus = 0
	ZoneChangeUpdate_TROUBLE  ZoneChangeUpdate_ZoneStatus = 1
	ZoneChangeUpdate_VIOLATED ZoneChangeUpdate_ZoneStatus = 2
	ZoneChangeUpdate_BYPASSED ZoneChangeUpdate_ZoneStatus = 3
)

var ZoneChangeUpdate_ZoneStatus_name = map[int32]string{
	0: "NORMAL",
	1: "TROUBLE",
	2: "VIOLATED",
	3: "BYPASSED",
}
var ZoneChangeUpdate_ZoneStatus_value = map[string]int32{
	"NORMAL":   0,
	"TROUBLE":  1,
	"VIOLATED": 2,
	"BYPASSED": 3,
}

func (x ZoneChangeUpdate_ZoneStatus) String() string {
	return proto.EnumName(ZoneChangeUpdate_ZoneStatus_name, int32(x))
}
func (ZoneChangeUpdate_ZoneStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 0}
}

type ZoneChangeUpdate_ZoneSubStatus int32

const (
	ZoneChangeUpdate_UNCONFIGURED ZoneChangeUpdate_ZoneSubStatus = 0
	ZoneChangeUpdate_OPEN         ZoneChangeUpdate_ZoneSubStatus = 1
	ZoneChangeUpdate_EOL          ZoneChangeUpdate_ZoneSubStatus = 2
	ZoneChangeUpdate_SHORT        ZoneChangeUpdate_ZoneSubStatus = 3
)

var ZoneChangeUpdate_ZoneSubStatus_name = map[int32]string{
	0: "UNCONFIGURED",
	1: "OPEN",
	2: "EOL",
	3: "SHORT",
}
var ZoneChangeUpdate_ZoneSubStatus_value = map[string]int32{
	"UNCONFIGURED": 0,
	"OPEN":         1,
	"EOL":          2,
	"SHORT":        3,
}

func (x ZoneChangeUpdate_ZoneSubStatus) String() string {
	return proto.EnumName(ZoneChangeUpdate_ZoneSubStatus_name, int32(x))
}
func (ZoneChangeUpdate_ZoneSubStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 1}
}

type ZoneChangeArgs struct {
}

func (m *ZoneChangeArgs) Reset()                    { *m = ZoneChangeArgs{} }
func (m *ZoneChangeArgs) String() string            { return proto.CompactTextString(m) }
func (*ZoneChangeArgs) ProtoMessage()               {}
func (*ZoneChangeArgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ArmingChangeArgs struct {
}

func (m *ArmingChangeArgs) Reset()                    { *m = ArmingChangeArgs{} }
func (m *ArmingChangeArgs) String() string            { return proto.CompactTextString(m) }
func (*ArmingChangeArgs) ProtoMessage()               {}
func (*ArmingChangeArgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ArmingStatusReport struct {
	Areas []*ArmingStatusArea `protobuf:"bytes,1,rep,name=areas" json:"areas,omitempty"`
}

func (m *ArmingStatusReport) Reset()                    { *m = ArmingStatusReport{} }
func (m *ArmingStatusReport) String() string            { return proto.CompactTextString(m) }
func (*ArmingStatusReport) ProtoMessage()               {}
func (*ArmingStatusReport) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ArmingStatusReport) GetAreas() []*ArmingStatusArea {
	if m != nil {
		return m.Areas
	}
	return nil
}

type ArmingStatusArea struct {
	Area         int32                         `protobuf:"varint,1,opt,name=area" json:"area,omitempty"`
	ArmingStatus ArmingStatusArea_ArmingStatus `protobuf:"varint,2,opt,name=arming_status,json=armingStatus,enum=elkm1grpc.ArmingStatusArea_ArmingStatus" json:"arming_status,omitempty"`
	ArmUpState   ArmingStatusArea_ArmUpState   `protobuf:"varint,3,opt,name=arm_up_state,json=armUpState,enum=elkm1grpc.ArmingStatusArea_ArmUpState" json:"arm_up_state,omitempty"`
	AlarmState   ArmingStatusArea_AlarmState   `protobuf:"varint,4,opt,name=alarm_state,json=alarmState,enum=elkm1grpc.ArmingStatusArea_AlarmState" json:"alarm_state,omitempty"`
}

func (m *ArmingStatusArea) Reset()                    { *m = ArmingStatusArea{} }
func (m *ArmingStatusArea) String() string            { return proto.CompactTextString(m) }
func (*ArmingStatusArea) ProtoMessage()               {}
func (*ArmingStatusArea) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ArmingStatusArea) GetArea() int32 {
	if m != nil {
		return m.Area
	}
	return 0
}

func (m *ArmingStatusArea) GetArmingStatus() ArmingStatusArea_ArmingStatus {
	if m != nil {
		return m.ArmingStatus
	}
	return ArmingStatusArea_DISARMED
}

func (m *ArmingStatusArea) GetArmUpState() ArmingStatusArea_ArmUpState {
	if m != nil {
		return m.ArmUpState
	}
	return ArmingStatusArea_NOT_READY
}

func (m *ArmingStatusArea) GetAlarmState() ArmingStatusArea_AlarmState {
	if m != nil {
		return m.AlarmState
	}
	return ArmingStatusArea_NO_ALARM
}

type ArmArgs struct {
	Area int32 `protobuf:"varint,1,opt,name=area" json:"area,omitempty"`
	Pin  int32 `protobuf:"varint,2,opt,name=pin" json:"pin,omitempty"`
}

func (m *ArmArgs) Reset()                    { *m = ArmArgs{} }
func (m *ArmArgs) String() string            { return proto.CompactTextString(m) }
func (*ArmArgs) ProtoMessage()               {}
func (*ArmArgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ArmArgs) GetArea() int32 {
	if m != nil {
		return m.Area
	}
	return 0
}

func (m *ArmArgs) GetPin() int32 {
	if m != nil {
		return m.Pin
	}
	return 0
}

type ZoneChangeUpdate struct {
	Zone      int32                          `protobuf:"varint,1,opt,name=zone" json:"zone,omitempty"`
	Status    ZoneChangeUpdate_ZoneStatus    `protobuf:"varint,2,opt,name=status,enum=elkm1grpc.ZoneChangeUpdate_ZoneStatus" json:"status,omitempty"`
	SubStatus ZoneChangeUpdate_ZoneSubStatus `protobuf:"varint,3,opt,name=sub_status,json=subStatus,enum=elkm1grpc.ZoneChangeUpdate_ZoneSubStatus" json:"sub_status,omitempty"`
}

func (m *ZoneChangeUpdate) Reset()                    { *m = ZoneChangeUpdate{} }
func (m *ZoneChangeUpdate) String() string            { return proto.CompactTextString(m) }
func (*ZoneChangeUpdate) ProtoMessage()               {}
func (*ZoneChangeUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ZoneChangeUpdate) GetZone() int32 {
	if m != nil {
		return m.Zone
	}
	return 0
}

func (m *ZoneChangeUpdate) GetStatus() ZoneChangeUpdate_ZoneStatus {
	if m != nil {
		return m.Status
	}
	return ZoneChangeUpdate_NORMAL
}

func (m *ZoneChangeUpdate) GetSubStatus() ZoneChangeUpdate_ZoneSubStatus {
	if m != nil {
		return m.SubStatus
	}
	return ZoneChangeUpdate_UNCONFIGURED
}

func init() {
	proto.RegisterType((*ZoneChangeArgs)(nil), "elkm1grpc.ZoneChangeArgs")
	proto.RegisterType((*ArmingChangeArgs)(nil), "elkm1grpc.ArmingChangeArgs")
	proto.RegisterType((*ArmingStatusReport)(nil), "elkm1grpc.ArmingStatusReport")
	proto.RegisterType((*ArmingStatusArea)(nil), "elkm1grpc.ArmingStatusArea")
	proto.RegisterType((*ArmArgs)(nil), "elkm1grpc.ArmArgs")
	proto.RegisterType((*ZoneChangeUpdate)(nil), "elkm1grpc.ZoneChangeUpdate")
	proto.RegisterEnum("elkm1grpc.ArmingStatusArea_ArmingStatus", ArmingStatusArea_ArmingStatus_name, ArmingStatusArea_ArmingStatus_value)
	proto.RegisterEnum("elkm1grpc.ArmingStatusArea_ArmUpState", ArmingStatusArea_ArmUpState_name, ArmingStatusArea_ArmUpState_value)
	proto.RegisterEnum("elkm1grpc.ArmingStatusArea_AlarmState", ArmingStatusArea_AlarmState_name, ArmingStatusArea_AlarmState_value)
	proto.RegisterEnum("elkm1grpc.ZoneChangeUpdate_ZoneStatus", ZoneChangeUpdate_ZoneStatus_name, ZoneChangeUpdate_ZoneStatus_value)
	proto.RegisterEnum("elkm1grpc.ZoneChangeUpdate_ZoneSubStatus", ZoneChangeUpdate_ZoneSubStatus_name, ZoneChangeUpdate_ZoneSubStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ElkGRPC service

type ElkGRPCClient interface {
	ArmStay(ctx context.Context, in *ArmArgs, opts ...grpc.CallOption) (*ArmingStatusArea, error)
	ArmAway(ctx context.Context, in *ArmArgs, opts ...grpc.CallOption) (*ArmingStatusArea, error)
	Disarm(ctx context.Context, in *ArmArgs, opts ...grpc.CallOption) (*ArmingStatusArea, error)
	ArmingStatusChange(ctx context.Context, in *ArmingChangeArgs, opts ...grpc.CallOption) (ElkGRPC_ArmingStatusChangeClient, error)
	ZoneChange(ctx context.Context, in *ZoneChangeArgs, opts ...grpc.CallOption) (ElkGRPC_ZoneChangeClient, error)
}

type elkGRPCClient struct {
	cc *grpc.ClientConn
}

func NewElkGRPCClient(cc *grpc.ClientConn) ElkGRPCClient {
	return &elkGRPCClient{cc}
}

func (c *elkGRPCClient) ArmStay(ctx context.Context, in *ArmArgs, opts ...grpc.CallOption) (*ArmingStatusArea, error) {
	out := new(ArmingStatusArea)
	err := grpc.Invoke(ctx, "/elkm1grpc.ElkGRPC/ArmStay", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elkGRPCClient) ArmAway(ctx context.Context, in *ArmArgs, opts ...grpc.CallOption) (*ArmingStatusArea, error) {
	out := new(ArmingStatusArea)
	err := grpc.Invoke(ctx, "/elkm1grpc.ElkGRPC/ArmAway", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elkGRPCClient) Disarm(ctx context.Context, in *ArmArgs, opts ...grpc.CallOption) (*ArmingStatusArea, error) {
	out := new(ArmingStatusArea)
	err := grpc.Invoke(ctx, "/elkm1grpc.ElkGRPC/Disarm", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elkGRPCClient) ArmingStatusChange(ctx context.Context, in *ArmingChangeArgs, opts ...grpc.CallOption) (ElkGRPC_ArmingStatusChangeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ElkGRPC_serviceDesc.Streams[0], c.cc, "/elkm1grpc.ElkGRPC/ArmingStatusChange", opts...)
	if err != nil {
		return nil, err
	}
	x := &elkGRPCArmingStatusChangeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ElkGRPC_ArmingStatusChangeClient interface {
	Recv() (*ArmingStatusReport, error)
	grpc.ClientStream
}

type elkGRPCArmingStatusChangeClient struct {
	grpc.ClientStream
}

func (x *elkGRPCArmingStatusChangeClient) Recv() (*ArmingStatusReport, error) {
	m := new(ArmingStatusReport)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *elkGRPCClient) ZoneChange(ctx context.Context, in *ZoneChangeArgs, opts ...grpc.CallOption) (ElkGRPC_ZoneChangeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ElkGRPC_serviceDesc.Streams[1], c.cc, "/elkm1grpc.ElkGRPC/ZoneChange", opts...)
	if err != nil {
		return nil, err
	}
	x := &elkGRPCZoneChangeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ElkGRPC_ZoneChangeClient interface {
	Recv() (*ZoneChangeUpdate, error)
	grpc.ClientStream
}

type elkGRPCZoneChangeClient struct {
	grpc.ClientStream
}

func (x *elkGRPCZoneChangeClient) Recv() (*ZoneChangeUpdate, error) {
	m := new(ZoneChangeUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ElkGRPC service

type ElkGRPCServer interface {
	ArmStay(context.Context, *ArmArgs) (*ArmingStatusArea, error)
	ArmAway(context.Context, *ArmArgs) (*ArmingStatusArea, error)
	Disarm(context.Context, *ArmArgs) (*ArmingStatusArea, error)
	ArmingStatusChange(*ArmingChangeArgs, ElkGRPC_ArmingStatusChangeServer) error
	ZoneChange(*ZoneChangeArgs, ElkGRPC_ZoneChangeServer) error
}

func RegisterElkGRPCServer(s *grpc.Server, srv ElkGRPCServer) {
	s.RegisterService(&_ElkGRPC_serviceDesc, srv)
}

func _ElkGRPC_ArmStay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArmArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElkGRPCServer).ArmStay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elkm1grpc.ElkGRPC/ArmStay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElkGRPCServer).ArmStay(ctx, req.(*ArmArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElkGRPC_ArmAway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArmArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElkGRPCServer).ArmAway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elkm1grpc.ElkGRPC/ArmAway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElkGRPCServer).ArmAway(ctx, req.(*ArmArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElkGRPC_Disarm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArmArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElkGRPCServer).Disarm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elkm1grpc.ElkGRPC/Disarm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElkGRPCServer).Disarm(ctx, req.(*ArmArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElkGRPC_ArmingStatusChange_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ArmingChangeArgs)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ElkGRPCServer).ArmingStatusChange(m, &elkGRPCArmingStatusChangeServer{stream})
}

type ElkGRPC_ArmingStatusChangeServer interface {
	Send(*ArmingStatusReport) error
	grpc.ServerStream
}

type elkGRPCArmingStatusChangeServer struct {
	grpc.ServerStream
}

func (x *elkGRPCArmingStatusChangeServer) Send(m *ArmingStatusReport) error {
	return x.ServerStream.SendMsg(m)
}

func _ElkGRPC_ZoneChange_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ZoneChangeArgs)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ElkGRPCServer).ZoneChange(m, &elkGRPCZoneChangeServer{stream})
}

type ElkGRPC_ZoneChangeServer interface {
	Send(*ZoneChangeUpdate) error
	grpc.ServerStream
}

type elkGRPCZoneChangeServer struct {
	grpc.ServerStream
}

func (x *elkGRPCZoneChangeServer) Send(m *ZoneChangeUpdate) error {
	return x.ServerStream.SendMsg(m)
}

var _ElkGRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "elkm1grpc.ElkGRPC",
	HandlerType: (*ElkGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ArmStay",
			Handler:    _ElkGRPC_ArmStay_Handler,
		},
		{
			MethodName: "ArmAway",
			Handler:    _ElkGRPC_ArmAway_Handler,
		},
		{
			MethodName: "Disarm",
			Handler:    _ElkGRPC_Disarm_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ArmingStatusChange",
			Handler:       _ElkGRPC_ArmingStatusChange_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ZoneChange",
			Handler:       _ElkGRPC_ZoneChange_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "elkm1grpc.proto",
}

func init() { proto.RegisterFile("elkm1grpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 848 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x16, 0xf5, 0x6b, 0x8d, 0x7e, 0x3c, 0x9e, 0x36, 0x81, 0xe3, 0xa2, 0x85, 0xc1, 0x43, 0xe1,
	0x5e, 0xd2, 0xd8, 0xe9, 0xad, 0x45, 0x8b, 0x15, 0xb9, 0x92, 0x08, 0x88, 0x5c, 0x61, 0xb9, 0x72,
	0xac, 0x5c, 0x08, 0xba, 0x25, 0x5c, 0x23, 0xb6, 0x24, 0x50, 0x32, 0x8a, 0xb6, 0xaf, 0xd5, 0xe7,
	0xe9, 0x13, 0xf4, 0xd4, 0x6b, 0x2f, 0xc5, 0xee, 0x4a, 0xa2, 0x94, 0xb8, 0x49, 0x91, 0xdb, 0x7c,
	0x33, 0xf3, 0x7d, 0x3b, 0x9c, 0x59, 0xce, 0xc2, 0x61, 0x76, 0xf7, 0xe6, 0xfe, 0xfc, 0x26, 0x5f,
	0xfc, 0xf8, 0x7c, 0x91, 0xcf, 0x57, 0x73, 0x6a, 0x6e, 0x1d, 0x2e, 0x42, 0xf7, 0xf5, 0x7c, 0x96,
	0x79, 0x3f, 0xa7, 0xb3, 0x9b, 0x8c, 0xe5, 0x37, 0x4b, 0x97, 0x00, 0x59, 0x7e, 0x7f, 0x3b, 0xbb,
	0xd9, 0xf1, 0x0d, 0x80, 0xac, 0x2f, 0x5e, 0xa5, 0xab, 0x87, 0xa5, 0xcc, 0x16, 0xf3, 0x7c, 0x45,
	0xe7, 0x50, 0x4b, 0xf3, 0x2c, 0x5d, 0x1e, 0x3b, 0xa7, 0x95, 0xb3, 0xd6, 0xc5, 0x67, 0xcf, 0x8b,
	0x73, 0x76, 0xb3, 0x59, 0x9e, 0xa5, 0xd2, 0x66, 0xba, 0x7f, 0x37, 0x36, 0xea, 0x45, 0x8c, 0x08,
	0xaa, 0x3a, 0x7a, 0xec, 0x9c, 0x3a, 0x67, 0x35, 0x69, 0x6c, 0x0a, 0xa1, 0x93, 0x9a, 0xbc, 0x64,
	0x69, 0x12, 0x8f, 0xcb, 0xa7, 0xce, 0x59, 0xf7, 0xe2, 0xec, 0x3d, 0x67, 0xec, 0x39, 0x64, 0x3b,
	0xdd, 0x41, 0x34, 0x04, 0x8d, 0x93, 0x87, 0x85, 0x91, 0xcb, 0x8e, 0x2b, 0x46, 0xed, 0xcb, 0x0f,
	0xa8, 0x4d, 0x16, 0x1a, 0x67, 0x12, 0xd2, 0xad, 0x4d, 0x03, 0x68, 0xa5, 0x77, 0x5a, 0xcb, 0x0a,
	0x55, 0xff, 0x87, 0x90, 0x4e, 0xdf, 0x08, 0x6d, 0x6d, 0xf7, 0x1f, 0x07, 0xda, 0xbb, 0xb9, 0xd4,
	0x86, 0x03, 0x3f, 0x88, 0x99, 0x0c, 0xb9, 0x8f, 0x25, 0xea, 0x02, 0x18, 0x33, 0x61, 0xaf, 0xd8,
	0x14, 0x9d, 0x02, 0xc7, 0x8a, 0x4d, 0xb1, 0x4c, 0x4f, 0x81, 0x0a, 0x9c, 0x04, 0x51, 0xac, 0x58,
	0xa4, 0xb0, 0x42, 0x04, 0x5d, 0xeb, 0x57, 0x22, 0x89, 0x82, 0xc1, 0x50, 0x61, 0x95, 0x4e, 0xe0,
	0xe9, 0xbe, 0x6f, 0x9b, 0x5f, 0xa3, 0x27, 0x70, 0xb4, 0x8d, 0x5d, 0x32, 0x8f, 0xa9, 0x40, 0x44,
	0x58, 0xa7, 0x2f, 0xe0, 0x84, 0xc9, 0x30, 0x89, 0x15, 0x1f, 0x1b, 0x16, 0xbf, 0x52, 0xa6, 0x92,
	0x24, 0x14, 0x3e, 0xc7, 0xc6, 0xa3, 0x71, 0x53, 0x89, 0x89, 0x1f, 0xe8, 0x72, 0xfb, 0x42, 0x7a,
	0xdc, 0x96, 0xdf, 0x2c, 0xb0, 0x29, 0x1f, 0xdc, 0xdf, 0x01, 0x8a, 0x06, 0x53, 0x07, 0x9a, 0x91,
	0x50, 0x89, 0xe4, 0xcc, 0x9f, 0x62, 0x89, 0x9a, 0x50, 0xb3, 0xa6, 0x43, 0x87, 0xd0, 0x32, 0x66,
	0x62, 0xd8, 0x58, 0xa6, 0x4f, 0x01, 0x6d, 0xbd, 0xfc, 0x2a, 0x50, 0x89, 0x0a, 0x42, 0x2e, 0xb1,
	0xa2, 0x19, 0xb6, 0x71, 0x55, 0xcd, 0x58, 0x9f, 0x6c, 0x1c, 0x35, 0x42, 0x68, 0x5b, 0x46, 0x6f,
	0x3a, 0x66, 0x71, 0x8c, 0x75, 0xf7, 0xaf, 0x32, 0x40, 0x31, 0x15, 0xdd, 0xf8, 0x48, 0x24, 0x6c,
	0xc4, 0x64, 0x88, 0x25, 0xdd, 0x40, 0x1e, 0x29, 0xc9, 0x22, 0x8f, 0x27, 0x3e, 0x1f, 0xb1, 0x75,
	0x15, 0xac, 0x27, 0xa4, 0x5a, 0x3b, 0xca, 0x66, 0x1a, 0x3a, 0x3f, 0xe9, 0x07, 0x92, 0x63, 0x85,
	0x8e, 0xa0, 0x63, 0x71, 0xc8, 0xfd, 0xc0, 0x63, 0x23, 0xac, 0x9a, 0x63, 0x8d, 0x6b, 0x2c, 0x46,
	0x81, 0xc7, 0xb1, 0x56, 0x24, 0xf5, 0x26, 0x72, 0x30, 0x62, 0x12, 0xeb, 0x85, 0x0e, 0x9b, 0x5c,
	0x9d, 0x63, 0x63, 0x0f, 0x5f, 0xd8, 0x36, 0x6e, 0xf1, 0x4b, 0xdb, 0xc6, 0x2d, 0xfe, 0x06, 0x81,
	0x9e, 0xc1, 0x13, 0x8b, 0x3d, 0x26, 0x7b, 0x22, 0x4a, 0x42, 0x11, 0x89, 0xab, 0xc0, 0xe7, 0xd8,
	0xa2, 0x4f, 0xe0, 0xd0, 0x86, 0x78, 0xc8, 0xe5, 0x80, 0x47, 0xde, 0x14, 0xdb, 0x45, 0x51, 0x7d,
	0xc9, 0xf9, 0x6b, 0x8e, 0x1d, 0xdd, 0x7a, 0xeb, 0x19, 0xb0, 0x18, 0xbb, 0xc5, 0x01, 0x43, 0xce,
	0x14, 0x1e, 0x9a, 0x2f, 0x37, 0xf8, 0x15, 0x53, 0x5c, 0x22, 0x9a, 0xbb, 0xb4, 0xfd, 0xf2, 0x24,
	0x9e, 0x8c, 0xb9, 0xbc, 0x0c, 0x62, 0x21, 0xa7, 0x78, 0x64, 0xee, 0x52, 0x11, 0xbb, 0xe4, 0x32,
	0xe8, 0x4f, 0x91, 0xdc, 0xaf, 0xa1, 0xc1, 0xf2, 0x7b, 0xbd, 0x48, 0x1e, 0xfd, 0xd5, 0x11, 0x2a,
	0x8b, 0xdb, 0x99, 0xf9, 0xc1, 0x6b, 0x52, 0x9b, 0xee, 0x1f, 0x65, 0xc0, 0x62, 0x2b, 0x4d, 0x16,
	0x3f, 0xe9, 0x29, 0x11, 0x54, 0x7f, 0x9b, 0xcf, 0xb2, 0x0d, 0x55, 0xdb, 0xf4, 0x3d, 0xd4, 0xf7,
	0xd6, 0xc3, 0xee, 0x7f, 0xf8, 0xb6, 0x80, 0x71, 0xac, 0x97, 0xc3, 0x9a, 0x45, 0x43, 0x80, 0xe5,
	0xc3, 0xf5, 0x66, 0xc5, 0xd8, 0xa5, 0xf0, 0xd5, 0x07, 0x35, 0x1e, 0xae, 0xd7, 0x32, 0xcd, 0xe5,
	0xc6, 0x74, 0x19, 0x40, 0xa1, 0x4f, 0x00, 0xf5, 0x48, 0xc8, 0x90, 0x8d, 0xb0, 0x44, 0x2d, 0x68,
	0x28, 0x29, 0x26, 0xbd, 0x11, 0x47, 0x47, 0x5f, 0xb5, 0xcb, 0x40, 0x8c, 0x98, 0xe2, 0x3e, 0x96,
	0x35, 0xb2, 0x77, 0x92, 0xfb, 0x58, 0x71, 0x7f, 0x80, 0xce, 0x9e, 0xbc, 0x1e, 0xd6, 0x24, 0xf2,
	0x44, 0xd4, 0x0f, 0x06, 0x13, 0x69, 0x96, 0xc2, 0x01, 0x54, 0xc5, 0x98, 0x47, 0xe8, 0x50, 0x03,
	0x2a, 0x5c, 0x8c, 0xb0, 0xac, 0x6f, 0x7e, 0x3c, 0x14, 0x52, 0x61, 0xe5, 0xe2, 0xcf, 0x32, 0x34,
	0xf8, 0xdd, 0x9b, 0x81, 0x1c, 0x7b, 0xf4, 0x9d, 0xe9, 0x79, 0xbc, 0x4a, 0x7f, 0x25, 0xda, 0x5f,
	0x4e, 0x7a, 0x0e, 0x27, 0xef, 0xdb, 0xd5, 0x6e, 0x69, 0xcd, 0x66, 0xbf, 0x7c, 0x1c, 0xfb, 0x5b,
	0xa8, 0xfb, 0xb7, 0xcb, 0x34, 0xbf, 0xff, 0x18, 0xb2, 0xdc, 0x7f, 0x6a, 0x6c, 0xf7, 0xe9, 0x5d,
	0x52, 0xf1, 0x3a, 0x9d, 0x7c, 0xfe, 0x1f, 0x8a, 0xf6, 0x99, 0x7a, 0xe1, 0x50, 0xdf, 0x0e, 0x67,
	0xad, 0xf5, 0xec, 0xd1, 0x01, 0xbf, 0x53, 0xdb, 0xdb, 0xb3, 0x7f, 0xe1, 0x5c, 0xd7, 0xcd, 0xf3,
	0xf9, 0xf2, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x00, 0xa9, 0x59, 0x51, 0x07, 0x00, 0x00,
}
