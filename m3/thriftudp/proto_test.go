package thriftudp

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"

import strconv "strconv"

import strings "strings"
import reflect "reflect"

import (
	"os"
	"runtime/pprof"
	"testing"

	"github.com/stretchr/testify/require"

	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

func BenchmarkFoo(b *testing.B) {
	addr := &Address{}

	data := []struct {
		ip   string
		port uint32
	}{
		{ip: "foo", port: 123},
		{ip: "foo", port: 123},
		{ip: "foo", port: 123},
		{ip: "foo", port: 123},
		{ip: "foo", port: 123},
		{ip: "foo", port: 123},
		{ip: "foo", port: 123},
	}

	buffer := proto.NewBuffer(make([]byte, 0, 1024))
	transport, _ := NewTUDPClientTransport("127.0.0.1:123", "")
	// transport.writeBuf.Grow(1024)
	require.NoError(b, transport.Open())

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, d := range data {
			addr.Ip = d.ip
			addr.Port = d.port

			buffer.Reset()
			require.NoError(b, buffer.EncodeMessage(addr))
			bytes := buffer.Bytes()
			if int(transport.RemainingBytes())-len(bytes) < 0 {
				require.NoError(b, transport.Flush())
			}

			_, err := transport.Write(bytes)
			require.NoError(b, err)
		}
	}
	b.StopTimer()

	require.NoError(b, transport.Flush())

	out, _ := os.Create("/tmp/foo_profile.prof")
	pprof.WriteHeapProfile(out)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ApplicationType int32

const (
	APPLICATION_TYPE_INVALID               ApplicationType = 0
	APPLICATION_TYPE_SERVICE               ApplicationType = 1
	APPLICATION_TYPE_HAPROXY_BACKEND       ApplicationType = 2
	APPLICATION_TYPE_MUTTLEY_TRAFFIC_GROUP ApplicationType = 3
	APPLICATION_TYPE_TOPIC                 ApplicationType = 4
	APPLICATION_TYPE_APPLICATION           ApplicationType = 5
	APPLICATION_TYPE_UNRESOLVED            ApplicationType = 6
)

var ApplicationType_name = map[int32]string{
	0: "APPLICATION_TYPE_INVALID",
	1: "APPLICATION_TYPE_SERVICE",
	2: "APPLICATION_TYPE_HAPROXY_BACKEND",
	3: "APPLICATION_TYPE_MUTTLEY_TRAFFIC_GROUP",
	4: "APPLICATION_TYPE_TOPIC",
	5: "APPLICATION_TYPE_APPLICATION",
	6: "APPLICATION_TYPE_UNRESOLVED",
}
var ApplicationType_value = map[string]int32{
	"APPLICATION_TYPE_INVALID":               0,
	"APPLICATION_TYPE_SERVICE":               1,
	"APPLICATION_TYPE_HAPROXY_BACKEND":       2,
	"APPLICATION_TYPE_MUTTLEY_TRAFFIC_GROUP": 3,
	"APPLICATION_TYPE_TOPIC":                 4,
	"APPLICATION_TYPE_APPLICATION":           5,
	"APPLICATION_TYPE_UNRESOLVED":            6,
}

func (ApplicationType) EnumDescriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{0} }

type EventType int32

const (
	EVENT_TYPE_INVALID  EventType = 0
	EVENT_TYPE_SETSTATE EventType = 1
)

var EventType_name = map[int32]string{
	0: "EVENT_TYPE_INVALID",
	1: "EVENT_TYPE_SETSTATE",
}
var EventType_value = map[string]int32{
	"EVENT_TYPE_INVALID":  0,
	"EVENT_TYPE_SETSTATE": 1,
}

func (EventType) EnumDescriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{1} }

type EdgeState int32

const (
	EDGE_STATE_INVALID     EdgeState = 0
	EDGE_STATE_UNREGULATED EdgeState = 1
	EDGE_STATE_DETECTED    EdgeState = 2
	EDGE_STATE_REVIEW      EdgeState = 3
	EDGE_STATE_HOLD        EdgeState = 4
	EDGE_STATE_REJECTED    EdgeState = 5
	EDGE_STATE_ACCEPTED    EdgeState = 6
)

var EdgeState_name = map[int32]string{
	0: "EDGE_STATE_INVALID",
	1: "EDGE_STATE_UNREGULATED",
	2: "EDGE_STATE_DETECTED",
	3: "EDGE_STATE_REVIEW",
	4: "EDGE_STATE_HOLD",
	5: "EDGE_STATE_REJECTED",
	6: "EDGE_STATE_ACCEPTED",
}
var EdgeState_value = map[string]int32{
	"EDGE_STATE_INVALID":     0,
	"EDGE_STATE_UNREGULATED": 1,
	"EDGE_STATE_DETECTED":    2,
	"EDGE_STATE_REVIEW":      3,
	"EDGE_STATE_HOLD":        4,
	"EDGE_STATE_REJECTED":    5,
	"EDGE_STATE_ACCEPTED":    6,
}

func (EdgeState) EnumDescriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{2} }

type Address struct {
	Ip   string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{0} }

func (m *Address) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Address) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

// AggregatorTopicEvent is a message sent to the m3msg aggregator
// topic.
type AggregatorTopicEvent struct {
	// Types that are valid to be assigned to Event:
	//	*AggregatorTopicEvent_ConnectionEvent
	//	*AggregatorTopicEvent_ResolvedConnectionEvent
	//	*AggregatorTopicEvent_TopicEvent
	Event isAggregatorTopicEvent_Event `protobuf_oneof:"event"`
}

func (m *AggregatorTopicEvent) Reset()                    { *m = AggregatorTopicEvent{} }
func (*AggregatorTopicEvent) ProtoMessage()               {}
func (*AggregatorTopicEvent) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{1} }

type isAggregatorTopicEvent_Event interface {
	isAggregatorTopicEvent_Event()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type AggregatorTopicEvent_ConnectionEvent struct {
	ConnectionEvent *ConnectionEvent `protobuf:"bytes,1,opt,name=connection_event,json=connectionEvent,oneof"`
}
type AggregatorTopicEvent_ResolvedConnectionEvent struct {
	ResolvedConnectionEvent *ResolvedConnectionEvent `protobuf:"bytes,2,opt,name=resolved_connection_event,json=resolvedConnectionEvent,oneof"`
}
type AggregatorTopicEvent_TopicEvent struct {
	TopicEvent *TopicEvent `protobuf:"bytes,3,opt,name=topic_event,json=topicEvent,oneof"`
}

func (*AggregatorTopicEvent_ConnectionEvent) isAggregatorTopicEvent_Event()         {}
func (*AggregatorTopicEvent_ResolvedConnectionEvent) isAggregatorTopicEvent_Event() {}
func (*AggregatorTopicEvent_TopicEvent) isAggregatorTopicEvent_Event()              {}

func (m *AggregatorTopicEvent) GetEvent() isAggregatorTopicEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *AggregatorTopicEvent) GetConnectionEvent() *ConnectionEvent {
	if x, ok := m.GetEvent().(*AggregatorTopicEvent_ConnectionEvent); ok {
		return x.ConnectionEvent
	}
	return nil
}

func (m *AggregatorTopicEvent) GetResolvedConnectionEvent() *ResolvedConnectionEvent {
	if x, ok := m.GetEvent().(*AggregatorTopicEvent_ResolvedConnectionEvent); ok {
		return x.ResolvedConnectionEvent
	}
	return nil
}

func (m *AggregatorTopicEvent) GetTopicEvent() *TopicEvent {
	if x, ok := m.GetEvent().(*AggregatorTopicEvent_TopicEvent); ok {
		return x.TopicEvent
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AggregatorTopicEvent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AggregatorTopicEvent_OneofMarshaler, _AggregatorTopicEvent_OneofUnmarshaler, _AggregatorTopicEvent_OneofSizer, []interface{}{
		(*AggregatorTopicEvent_ConnectionEvent)(nil),
		(*AggregatorTopicEvent_ResolvedConnectionEvent)(nil),
		(*AggregatorTopicEvent_TopicEvent)(nil),
	}
}

func _AggregatorTopicEvent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AggregatorTopicEvent)
	// event
	switch x := m.Event.(type) {
	case *AggregatorTopicEvent_ConnectionEvent:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ConnectionEvent); err != nil {
			return err
		}
	case *AggregatorTopicEvent_ResolvedConnectionEvent:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ResolvedConnectionEvent); err != nil {
			return err
		}
	case *AggregatorTopicEvent_TopicEvent:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TopicEvent); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AggregatorTopicEvent.Event has unexpected type %T", x)
	}
	return nil
}

func _AggregatorTopicEvent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AggregatorTopicEvent)
	switch tag {
	case 1: // event.connection_event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ConnectionEvent)
		err := b.DecodeMessage(msg)
		m.Event = &AggregatorTopicEvent_ConnectionEvent{msg}
		return true, err
	case 2: // event.resolved_connection_event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ResolvedConnectionEvent)
		err := b.DecodeMessage(msg)
		m.Event = &AggregatorTopicEvent_ResolvedConnectionEvent{msg}
		return true, err
	case 3: // event.topic_event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TopicEvent)
		err := b.DecodeMessage(msg)
		m.Event = &AggregatorTopicEvent_TopicEvent{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AggregatorTopicEvent_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AggregatorTopicEvent)
	// event
	switch x := m.Event.(type) {
	case *AggregatorTopicEvent_ConnectionEvent:
		s := proto.Size(x.ConnectionEvent)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AggregatorTopicEvent_ResolvedConnectionEvent:
		s := proto.Size(x.ResolvedConnectionEvent)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AggregatorTopicEvent_TopicEvent:
		s := proto.Size(x.TopicEvent)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ConnectionEvent struct {
	InboundListenAddress *Address `protobuf:"bytes,1,opt,name=inbound_listen_address,json=inboundListenAddress" json:"inbound_listen_address,omitempty"`
	// Types that are valid to be assigned to Event:
	//	*ConnectionEvent_OutboundEvent
	//	*ConnectionEvent_InboundEvent
	Event isConnectionEvent_Event `protobuf_oneof:"event"`
}

func (m *ConnectionEvent) Reset()                    { *m = ConnectionEvent{} }
func (*ConnectionEvent) ProtoMessage()               {}
func (*ConnectionEvent) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{2} }

type isConnectionEvent_Event interface {
	isConnectionEvent_Event()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type ConnectionEvent_OutboundEvent struct {
	OutboundEvent *ConnectionOutboundEvent `protobuf:"bytes,2,opt,name=outbound_event,json=outboundEvent,oneof"`
}
type ConnectionEvent_InboundEvent struct {
	InboundEvent *ConnectionInboundEvent `protobuf:"bytes,3,opt,name=inbound_event,json=inboundEvent,oneof"`
}

func (*ConnectionEvent_OutboundEvent) isConnectionEvent_Event() {}
func (*ConnectionEvent_InboundEvent) isConnectionEvent_Event()  {}

func (m *ConnectionEvent) GetEvent() isConnectionEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *ConnectionEvent) GetInboundListenAddress() *Address {
	if m != nil {
		return m.InboundListenAddress
	}
	return nil
}

func (m *ConnectionEvent) GetOutboundEvent() *ConnectionOutboundEvent {
	if x, ok := m.GetEvent().(*ConnectionEvent_OutboundEvent); ok {
		return x.OutboundEvent
	}
	return nil
}

func (m *ConnectionEvent) GetInboundEvent() *ConnectionInboundEvent {
	if x, ok := m.GetEvent().(*ConnectionEvent_InboundEvent); ok {
		return x.InboundEvent
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ConnectionEvent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ConnectionEvent_OneofMarshaler, _ConnectionEvent_OneofUnmarshaler, _ConnectionEvent_OneofSizer, []interface{}{
		(*ConnectionEvent_OutboundEvent)(nil),
		(*ConnectionEvent_InboundEvent)(nil),
	}
}

func _ConnectionEvent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ConnectionEvent)
	// event
	switch x := m.Event.(type) {
	case *ConnectionEvent_OutboundEvent:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OutboundEvent); err != nil {
			return err
		}
	case *ConnectionEvent_InboundEvent:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.InboundEvent); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ConnectionEvent.Event has unexpected type %T", x)
	}
	return nil
}

func _ConnectionEvent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ConnectionEvent)
	switch tag {
	case 2: // event.outbound_event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ConnectionOutboundEvent)
		err := b.DecodeMessage(msg)
		m.Event = &ConnectionEvent_OutboundEvent{msg}
		return true, err
	case 3: // event.inbound_event
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ConnectionInboundEvent)
		err := b.DecodeMessage(msg)
		m.Event = &ConnectionEvent_InboundEvent{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ConnectionEvent_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ConnectionEvent)
	// event
	switch x := m.Event.(type) {
	case *ConnectionEvent_OutboundEvent:
		s := proto.Size(x.OutboundEvent)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ConnectionEvent_InboundEvent:
		s := proto.Size(x.InboundEvent)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Application is primarily just a string for now but
// may have more attributes in the future when less is known.
type Application struct {
	Type ApplicationType `protobuf:"varint,1,opt,name=type,proto3,enum=connmon.ApplicationType" json:"type,omitempty"`
	Name string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Application) Reset()                    { *m = Application{} }
func (*Application) ProtoMessage()               {}
func (*Application) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{3} }

func (m *Application) GetType() ApplicationType {
	if m != nil {
		return m.Type
	}
	return APPLICATION_TYPE_INVALID
}

func (m *Application) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ConnectionOutboundEvent struct {
	OutboundApplication *Application `protobuf:"bytes,1,opt,name=outbound_application,json=outboundApplication" json:"outbound_application,omitempty"`
	OutboundAddress     *Address     `protobuf:"bytes,2,opt,name=outbound_address,json=outboundAddress" json:"outbound_address,omitempty"`
}

func (m *ConnectionOutboundEvent) Reset()                    { *m = ConnectionOutboundEvent{} }
func (*ConnectionOutboundEvent) ProtoMessage()               {}
func (*ConnectionOutboundEvent) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{4} }

func (m *ConnectionOutboundEvent) GetOutboundApplication() *Application {
	if m != nil {
		return m.OutboundApplication
	}
	return nil
}

func (m *ConnectionOutboundEvent) GetOutboundAddress() *Address {
	if m != nil {
		return m.OutboundAddress
	}
	return nil
}

type ConnectionInboundEvent struct {
	InboundApplication *Application `protobuf:"bytes,1,opt,name=inbound_application,json=inboundApplication" json:"inbound_application,omitempty"`
	InboundAddress     *Address     `protobuf:"bytes,2,opt,name=inbound_address,json=inboundAddress" json:"inbound_address,omitempty"`
}

func (m *ConnectionInboundEvent) Reset()                    { *m = ConnectionInboundEvent{} }
func (*ConnectionInboundEvent) ProtoMessage()               {}
func (*ConnectionInboundEvent) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{5} }

func (m *ConnectionInboundEvent) GetInboundApplication() *Application {
	if m != nil {
		return m.InboundApplication
	}
	return nil
}

func (m *ConnectionInboundEvent) GetInboundAddress() *Address {
	if m != nil {
		return m.InboundAddress
	}
	return nil
}

type ResolvedConnectionEvent struct {
	OutboundApplication *Application `protobuf:"bytes,1,opt,name=outbound_application,json=outboundApplication" json:"outbound_application,omitempty"`
	InboundApplication  *Application `protobuf:"bytes,2,opt,name=inbound_application,json=inboundApplication" json:"inbound_application,omitempty"`
}

func (m *ResolvedConnectionEvent) Reset()                    { *m = ResolvedConnectionEvent{} }
func (*ResolvedConnectionEvent) ProtoMessage()               {}
func (*ResolvedConnectionEvent) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{6} }

func (m *ResolvedConnectionEvent) GetOutboundApplication() *Application {
	if m != nil {
		return m.OutboundApplication
	}
	return nil
}

func (m *ResolvedConnectionEvent) GetInboundApplication() *Application {
	if m != nil {
		return m.InboundApplication
	}
	return nil
}

type TopicEvent struct {
	OutboundAddress *Address `protobuf:"bytes,1,opt,name=outbound_address,json=outboundAddress" json:"outbound_address,omitempty"`
	// Types that are valid to be assigned to Application:
	//	*TopicEvent_OutboundApplication
	//	*TopicEvent_InboundApplication
	Application isTopicEvent_Application `protobuf_oneof:"application"`
}

func (m *TopicEvent) Reset()                    { *m = TopicEvent{} }
func (*TopicEvent) ProtoMessage()               {}
func (*TopicEvent) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{7} }

type isTopicEvent_Application interface {
	isTopicEvent_Application()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type TopicEvent_OutboundApplication struct {
	OutboundApplication *Application `protobuf:"bytes,2,opt,name=outbound_application,json=outboundApplication,oneof"`
}
type TopicEvent_InboundApplication struct {
	InboundApplication *Application `protobuf:"bytes,3,opt,name=inbound_application,json=inboundApplication,oneof"`
}

func (*TopicEvent_OutboundApplication) isTopicEvent_Application() {}
func (*TopicEvent_InboundApplication) isTopicEvent_Application()  {}

func (m *TopicEvent) GetApplication() isTopicEvent_Application {
	if m != nil {
		return m.Application
	}
	return nil
}

func (m *TopicEvent) GetOutboundAddress() *Address {
	if m != nil {
		return m.OutboundAddress
	}
	return nil
}

func (m *TopicEvent) GetOutboundApplication() *Application {
	if x, ok := m.GetApplication().(*TopicEvent_OutboundApplication); ok {
		return x.OutboundApplication
	}
	return nil
}

func (m *TopicEvent) GetInboundApplication() *Application {
	if x, ok := m.GetApplication().(*TopicEvent_InboundApplication); ok {
		return x.InboundApplication
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TopicEvent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TopicEvent_OneofMarshaler, _TopicEvent_OneofUnmarshaler, _TopicEvent_OneofSizer, []interface{}{
		(*TopicEvent_OutboundApplication)(nil),
		(*TopicEvent_InboundApplication)(nil),
	}
}

func _TopicEvent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TopicEvent)
	// application
	switch x := m.Application.(type) {
	case *TopicEvent_OutboundApplication:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OutboundApplication); err != nil {
			return err
		}
	case *TopicEvent_InboundApplication:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.InboundApplication); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TopicEvent.Application has unexpected type %T", x)
	}
	return nil
}

func _TopicEvent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TopicEvent)
	switch tag {
	case 2: // application.outbound_application
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Application)
		err := b.DecodeMessage(msg)
		m.Application = &TopicEvent_OutboundApplication{msg}
		return true, err
	case 3: // application.inbound_application
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Application)
		err := b.DecodeMessage(msg)
		m.Application = &TopicEvent_InboundApplication{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TopicEvent_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TopicEvent)
	// application
	switch x := m.Application.(type) {
	case *TopicEvent_OutboundApplication:
		s := proto.Size(x.OutboundApplication)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TopicEvent_InboundApplication:
		s := proto.Size(x.InboundApplication)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type NodeSignature struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (m *NodeSignature) Reset()                    { *m = NodeSignature{} }
func (*NodeSignature) ProtoMessage()               {}
func (*NodeSignature) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{8} }

func (m *NodeSignature) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeSignature) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type Node struct {
	Uuid     string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type     string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	NoNotify bool   `protobuf:"varint,4,opt,name=no_notify,json=noNotify,proto3" json:"no_notify,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{9} }

func (m *Node) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Node) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Node) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Node) GetNoNotify() bool {
	if m != nil {
		return m.NoNotify
	}
	return false
}

type Edge struct {
	Uuid     string                     `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Sender   *Node                      `protobuf:"bytes,2,opt,name=sender" json:"sender,omitempty"`
	Receiver *Node                      `protobuf:"bytes,3,opt,name=receiver" json:"receiver,omitempty"`
	LastSeen *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=last_seen,json=lastSeen" json:"last_seen,omitempty"`
	State    EdgeState                  `protobuf:"varint,5,opt,name=state,proto3,enum=connmon.EdgeState" json:"state,omitempty"`
}

func (m *Edge) Reset()                    { *m = Edge{} }
func (*Edge) ProtoMessage()               {}
func (*Edge) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{10} }

func (m *Edge) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Edge) GetSender() *Node {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *Edge) GetReceiver() *Node {
	if m != nil {
		return m.Receiver
	}
	return nil
}

func (m *Edge) GetLastSeen() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastSeen
	}
	return nil
}

func (m *Edge) GetState() EdgeState {
	if m != nil {
		return m.State
	}
	return EDGE_STATE_INVALID
}

type Event struct {
	Uuid      string                     `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Username  string                     `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Type      EventType                  `protobuf:"varint,4,opt,name=type,proto3,enum=connmon.EventType" json:"type,omitempty"`
	Value     string                     `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	Meta      string                     `protobuf:"bytes,6,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{11} }

func (m *Event) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Event) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Event) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Event) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EVENT_TYPE_INVALID
}

func (m *Event) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Event) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

type Record struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{12} }

func (m *Record) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Record) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type GetNodesRequest struct {
}

func (m *GetNodesRequest) Reset()                    { *m = GetNodesRequest{} }
func (*GetNodesRequest) ProtoMessage()               {}
func (*GetNodesRequest) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{13} }

type GetNodesResponse struct {
	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *GetNodesResponse) Reset()                    { *m = GetNodesResponse{} }
func (*GetNodesResponse) ProtoMessage()               {}
func (*GetNodesResponse) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{14} }

func (m *GetNodesResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type GetNodeBySignatureRequest struct {
	Signature *NodeSignature `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
}

func (m *GetNodeBySignatureRequest) Reset()      { *m = GetNodeBySignatureRequest{} }
func (*GetNodeBySignatureRequest) ProtoMessage() {}
func (*GetNodeBySignatureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorConnmon, []int{15}
}

func (m *GetNodeBySignatureRequest) GetSignature() *NodeSignature {
	if m != nil {
		return m.Signature
	}
	return nil
}

type GetNodeBySignatureResponse struct {
	Node *Node `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
}

func (m *GetNodeBySignatureResponse) Reset()      { *m = GetNodeBySignatureResponse{} }
func (*GetNodeBySignatureResponse) ProtoMessage() {}
func (*GetNodeBySignatureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorConnmon, []int{16}
}

func (m *GetNodeBySignatureResponse) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

type GetEdgeRequest struct {
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (m *GetEdgeRequest) Reset()                    { *m = GetEdgeRequest{} }
func (*GetEdgeRequest) ProtoMessage()               {}
func (*GetEdgeRequest) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{17} }

func (m *GetEdgeRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type GetEdgeResponse struct {
	Edge *Edge `protobuf:"bytes,1,opt,name=edge" json:"edge,omitempty"`
}

func (m *GetEdgeResponse) Reset()                    { *m = GetEdgeResponse{} }
func (*GetEdgeResponse) ProtoMessage()               {}
func (*GetEdgeResponse) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{18} }

func (m *GetEdgeResponse) GetEdge() *Edge {
	if m != nil {
		return m.Edge
	}
	return nil
}

type GetEdgeByNodesRequest struct {
	Sender   *NodeSignature `protobuf:"bytes,1,opt,name=sender" json:"sender,omitempty"`
	Receiver *NodeSignature `protobuf:"bytes,2,opt,name=receiver" json:"receiver,omitempty"`
}

func (m *GetEdgeByNodesRequest) Reset()                    { *m = GetEdgeByNodesRequest{} }
func (*GetEdgeByNodesRequest) ProtoMessage()               {}
func (*GetEdgeByNodesRequest) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{19} }

func (m *GetEdgeByNodesRequest) GetSender() *NodeSignature {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *GetEdgeByNodesRequest) GetReceiver() *NodeSignature {
	if m != nil {
		return m.Receiver
	}
	return nil
}

type GetEdgeByNodesResponse struct {
	Edge *Edge `protobuf:"bytes,1,opt,name=edge" json:"edge,omitempty"`
}

func (m *GetEdgeByNodesResponse) Reset()                    { *m = GetEdgeByNodesResponse{} }
func (*GetEdgeByNodesResponse) ProtoMessage()               {}
func (*GetEdgeByNodesResponse) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{20} }

func (m *GetEdgeByNodesResponse) GetEdge() *Edge {
	if m != nil {
		return m.Edge
	}
	return nil
}

type GetEdgesRequest struct {
}

func (m *GetEdgesRequest) Reset()                    { *m = GetEdgesRequest{} }
func (*GetEdgesRequest) ProtoMessage()               {}
func (*GetEdgesRequest) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{21} }

type GetEdgesResponse struct {
	Edges []*Edge `protobuf:"bytes,1,rep,name=edges" json:"edges,omitempty"`
}

func (m *GetEdgesResponse) Reset()                    { *m = GetEdgesResponse{} }
func (*GetEdgesResponse) ProtoMessage()               {}
func (*GetEdgesResponse) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{22} }

func (m *GetEdgesResponse) GetEdges() []*Edge {
	if m != nil {
		return m.Edges
	}
	return nil
}

type GetEventLogForEdgeRequest struct {
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (m *GetEventLogForEdgeRequest) Reset()      { *m = GetEventLogForEdgeRequest{} }
func (*GetEventLogForEdgeRequest) ProtoMessage() {}
func (*GetEventLogForEdgeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorConnmon, []int{23}
}

func (m *GetEventLogForEdgeRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type GetEventLogForEdgeResponse struct {
	Events []*Event `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
}

func (m *GetEventLogForEdgeResponse) Reset()      { *m = GetEventLogForEdgeResponse{} }
func (*GetEventLogForEdgeResponse) ProtoMessage() {}
func (*GetEventLogForEdgeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorConnmon, []int{24}
}

func (m *GetEventLogForEdgeResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type SetEdgeStateRequest struct {
	EdgeUuid string    `protobuf:"bytes,1,opt,name=edge_uuid,json=edgeUuid,proto3" json:"edge_uuid,omitempty"`
	State    EdgeState `protobuf:"varint,2,opt,name=state,proto3,enum=connmon.EdgeState" json:"state,omitempty"`
	Comment  string    `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (m *SetEdgeStateRequest) Reset()                    { *m = SetEdgeStateRequest{} }
func (*SetEdgeStateRequest) ProtoMessage()               {}
func (*SetEdgeStateRequest) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{25} }

func (m *SetEdgeStateRequest) GetEdgeUuid() string {
	if m != nil {
		return m.EdgeUuid
	}
	return ""
}

func (m *SetEdgeStateRequest) GetState() EdgeState {
	if m != nil {
		return m.State
	}
	return EDGE_STATE_INVALID
}

func (m *SetEdgeStateRequest) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type SetEdgeStateResponse struct {
}

func (m *SetEdgeStateResponse) Reset()                    { *m = SetEdgeStateResponse{} }
func (*SetEdgeStateResponse) ProtoMessage()               {}
func (*SetEdgeStateResponse) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{26} }

type GetRecordsRequest struct {
}

func (m *GetRecordsRequest) Reset()                    { *m = GetRecordsRequest{} }
func (*GetRecordsRequest) ProtoMessage()               {}
func (*GetRecordsRequest) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{27} }

type GetRecordsResponse struct {
	Records []*Record `protobuf:"bytes,1,rep,name=records" json:"records,omitempty"`
}

func (m *GetRecordsResponse) Reset()                    { *m = GetRecordsResponse{} }
func (*GetRecordsResponse) ProtoMessage()               {}
func (*GetRecordsResponse) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{28} }

func (m *GetRecordsResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type IsRegulatedRequest struct {
	Signature *NodeSignature `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
}

func (m *IsRegulatedRequest) Reset()                    { *m = IsRegulatedRequest{} }
func (*IsRegulatedRequest) ProtoMessage()               {}
func (*IsRegulatedRequest) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{29} }

func (m *IsRegulatedRequest) GetSignature() *NodeSignature {
	if m != nil {
		return m.Signature
	}
	return nil
}

type IsRegulatedResponse struct {
	Regulated bool `protobuf:"varint,1,opt,name=regulated,proto3" json:"regulated,omitempty"`
}

func (m *IsRegulatedResponse) Reset()                    { *m = IsRegulatedResponse{} }
func (*IsRegulatedResponse) ProtoMessage()               {}
func (*IsRegulatedResponse) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{30} }

func (m *IsRegulatedResponse) GetRegulated() bool {
	if m != nil {
		return m.Regulated
	}
	return false
}

type TestDetectEdgeRequest struct {
	Sender   *NodeSignature `protobuf:"bytes,1,opt,name=sender" json:"sender,omitempty"`
	Receiver *NodeSignature `protobuf:"bytes,2,opt,name=receiver" json:"receiver,omitempty"`
}

func (m *TestDetectEdgeRequest) Reset()                    { *m = TestDetectEdgeRequest{} }
func (*TestDetectEdgeRequest) ProtoMessage()               {}
func (*TestDetectEdgeRequest) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{31} }

func (m *TestDetectEdgeRequest) GetSender() *NodeSignature {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *TestDetectEdgeRequest) GetReceiver() *NodeSignature {
	if m != nil {
		return m.Receiver
	}
	return nil
}

type TestDetectEdgeResponse struct {
	Edge *Edge `protobuf:"bytes,1,opt,name=edge" json:"edge,omitempty"`
}

func (m *TestDetectEdgeResponse) Reset()                    { *m = TestDetectEdgeResponse{} }
func (*TestDetectEdgeResponse) ProtoMessage()               {}
func (*TestDetectEdgeResponse) Descriptor() ([]byte, []int) { return fileDescriptorConnmon, []int{32} }

func (m *TestDetectEdgeResponse) GetEdge() *Edge {
	if m != nil {
		return m.Edge
	}
	return nil
}

func init() {
	proto.RegisterType((*Address)(nil), "connmon.Address")
	proto.RegisterType((*AggregatorTopicEvent)(nil), "connmon.AggregatorTopicEvent")
	proto.RegisterType((*ConnectionEvent)(nil), "connmon.ConnectionEvent")
	proto.RegisterType((*Application)(nil), "connmon.Application")
	proto.RegisterType((*ConnectionOutboundEvent)(nil), "connmon.ConnectionOutboundEvent")
	proto.RegisterType((*ConnectionInboundEvent)(nil), "connmon.ConnectionInboundEvent")
	proto.RegisterType((*ResolvedConnectionEvent)(nil), "connmon.ResolvedConnectionEvent")
	proto.RegisterType((*TopicEvent)(nil), "connmon.TopicEvent")
	proto.RegisterType((*NodeSignature)(nil), "connmon.NodeSignature")
	proto.RegisterType((*Node)(nil), "connmon.Node")
	proto.RegisterType((*Edge)(nil), "connmon.Edge")
	proto.RegisterType((*Event)(nil), "connmon.Event")
	proto.RegisterType((*Record)(nil), "connmon.Record")
	proto.RegisterType((*GetNodesRequest)(nil), "connmon.GetNodesRequest")
	proto.RegisterType((*GetNodesResponse)(nil), "connmon.GetNodesResponse")
	proto.RegisterType((*GetNodeBySignatureRequest)(nil), "connmon.GetNodeBySignatureRequest")
	proto.RegisterType((*GetNodeBySignatureResponse)(nil), "connmon.GetNodeBySignatureResponse")
	proto.RegisterType((*GetEdgeRequest)(nil), "connmon.GetEdgeRequest")
	proto.RegisterType((*GetEdgeResponse)(nil), "connmon.GetEdgeResponse")
	proto.RegisterType((*GetEdgeByNodesRequest)(nil), "connmon.GetEdgeByNodesRequest")
	proto.RegisterType((*GetEdgeByNodesResponse)(nil), "connmon.GetEdgeByNodesResponse")
	proto.RegisterType((*GetEdgesRequest)(nil), "connmon.GetEdgesRequest")
	proto.RegisterType((*GetEdgesResponse)(nil), "connmon.GetEdgesResponse")
	proto.RegisterType((*GetEventLogForEdgeRequest)(nil), "connmon.GetEventLogForEdgeRequest")
	proto.RegisterType((*GetEventLogForEdgeResponse)(nil), "connmon.GetEventLogForEdgeResponse")
	proto.RegisterType((*SetEdgeStateRequest)(nil), "connmon.SetEdgeStateRequest")
	proto.RegisterType((*SetEdgeStateResponse)(nil), "connmon.SetEdgeStateResponse")
	proto.RegisterType((*GetRecordsRequest)(nil), "connmon.GetRecordsRequest")
	proto.RegisterType((*GetRecordsResponse)(nil), "connmon.GetRecordsResponse")
	proto.RegisterType((*IsRegulatedRequest)(nil), "connmon.IsRegulatedRequest")
	proto.RegisterType((*IsRegulatedResponse)(nil), "connmon.IsRegulatedResponse")
	proto.RegisterType((*TestDetectEdgeRequest)(nil), "connmon.TestDetectEdgeRequest")
	proto.RegisterType((*TestDetectEdgeResponse)(nil), "connmon.TestDetectEdgeResponse")
	proto.RegisterEnum("connmon.ApplicationType", ApplicationType_name, ApplicationType_value)
	proto.RegisterEnum("connmon.EventType", EventType_name, EventType_value)
	proto.RegisterEnum("connmon.EdgeState", EdgeState_name, EdgeState_value)
}
func (x ApplicationType) String() string {
	s, ok := ApplicationType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x EventType) String() string {
	s, ok := EventType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x EdgeState) String() string {
	s, ok := EdgeState_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *Address) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Address)
	if !ok {
		that2, ok := that.(Address)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Ip != that1.Ip {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	return true
}
func (this *AggregatorTopicEvent) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*AggregatorTopicEvent)
	if !ok {
		that2, ok := that.(AggregatorTopicEvent)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if that1.Event == nil {
		if this.Event != nil {
			return false
		}
	} else if this.Event == nil {
		return false
	} else if !this.Event.Equal(that1.Event) {
		return false
	}
	return true
}
func (this *AggregatorTopicEvent_ConnectionEvent) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*AggregatorTopicEvent_ConnectionEvent)
	if !ok {
		that2, ok := that.(AggregatorTopicEvent_ConnectionEvent)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.ConnectionEvent.Equal(that1.ConnectionEvent) {
		return false
	}
	return true
}
func (this *AggregatorTopicEvent_ResolvedConnectionEvent) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*AggregatorTopicEvent_ResolvedConnectionEvent)
	if !ok {
		that2, ok := that.(AggregatorTopicEvent_ResolvedConnectionEvent)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.ResolvedConnectionEvent.Equal(that1.ResolvedConnectionEvent) {
		return false
	}
	return true
}
func (this *AggregatorTopicEvent_TopicEvent) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*AggregatorTopicEvent_TopicEvent)
	if !ok {
		that2, ok := that.(AggregatorTopicEvent_TopicEvent)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.TopicEvent.Equal(that1.TopicEvent) {
		return false
	}
	return true
}
func (this *ConnectionEvent) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ConnectionEvent)
	if !ok {
		that2, ok := that.(ConnectionEvent)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.InboundListenAddress.Equal(that1.InboundListenAddress) {
		return false
	}
	if that1.Event == nil {
		if this.Event != nil {
			return false
		}
	} else if this.Event == nil {
		return false
	} else if !this.Event.Equal(that1.Event) {
		return false
	}
	return true
}
func (this *ConnectionEvent_OutboundEvent) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ConnectionEvent_OutboundEvent)
	if !ok {
		that2, ok := that.(ConnectionEvent_OutboundEvent)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.OutboundEvent.Equal(that1.OutboundEvent) {
		return false
	}
	return true
}
func (this *ConnectionEvent_InboundEvent) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ConnectionEvent_InboundEvent)
	if !ok {
		that2, ok := that.(ConnectionEvent_InboundEvent)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.InboundEvent.Equal(that1.InboundEvent) {
		return false
	}
	return true
}
func (this *Application) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Application)
	if !ok {
		that2, ok := that.(Application)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	return true
}
func (this *ConnectionOutboundEvent) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ConnectionOutboundEvent)
	if !ok {
		that2, ok := that.(ConnectionOutboundEvent)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.OutboundApplication.Equal(that1.OutboundApplication) {
		return false
	}
	if !this.OutboundAddress.Equal(that1.OutboundAddress) {
		return false
	}
	return true
}
func (this *ConnectionInboundEvent) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ConnectionInboundEvent)
	if !ok {
		that2, ok := that.(ConnectionInboundEvent)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.InboundApplication.Equal(that1.InboundApplication) {
		return false
	}
	if !this.InboundAddress.Equal(that1.InboundAddress) {
		return false
	}
	return true
}
func (this *ResolvedConnectionEvent) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ResolvedConnectionEvent)
	if !ok {
		that2, ok := that.(ResolvedConnectionEvent)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.OutboundApplication.Equal(that1.OutboundApplication) {
		return false
	}
	if !this.InboundApplication.Equal(that1.InboundApplication) {
		return false
	}
	return true
}
func (this *TopicEvent) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*TopicEvent)
	if !ok {
		that2, ok := that.(TopicEvent)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.OutboundAddress.Equal(that1.OutboundAddress) {
		return false
	}
	if that1.Application == nil {
		if this.Application != nil {
			return false
		}
	} else if this.Application == nil {
		return false
	} else if !this.Application.Equal(that1.Application) {
		return false
	}
	return true
}
func (this *TopicEvent_OutboundApplication) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*TopicEvent_OutboundApplication)
	if !ok {
		that2, ok := that.(TopicEvent_OutboundApplication)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.OutboundApplication.Equal(that1.OutboundApplication) {
		return false
	}
	return true
}
func (this *TopicEvent_InboundApplication) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*TopicEvent_InboundApplication)
	if !ok {
		that2, ok := that.(TopicEvent_InboundApplication)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.InboundApplication.Equal(that1.InboundApplication) {
		return false
	}
	return true
}
func (this *NodeSignature) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*NodeSignature)
	if !ok {
		that2, ok := that.(NodeSignature)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	return true
}
func (this *Node) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Node)
	if !ok {
		that2, ok := that.(Node)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Uuid != that1.Uuid {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.NoNotify != that1.NoNotify {
		return false
	}
	return true
}
func (this *Edge) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Edge)
	if !ok {
		that2, ok := that.(Edge)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Uuid != that1.Uuid {
		return false
	}
	if !this.Sender.Equal(that1.Sender) {
		return false
	}
	if !this.Receiver.Equal(that1.Receiver) {
		return false
	}
	if !this.LastSeen.Equal(that1.LastSeen) {
		return false
	}
	if this.State != that1.State {
		return false
	}
	return true
}
func (this *Event) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Event)
	if !ok {
		that2, ok := that.(Event)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Uuid != that1.Uuid {
		return false
	}
	if !this.Timestamp.Equal(that1.Timestamp) {
		return false
	}
	if this.Username != that1.Username {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if this.Meta != that1.Meta {
		return false
	}
	return true
}
func (this *Record) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Record)
	if !ok {
		that2, ok := that.(Record)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	return true
}
func (this *GetNodesRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*GetNodesRequest)
	if !ok {
		that2, ok := that.(GetNodesRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *GetNodesResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*GetNodesResponse)
	if !ok {
		that2, ok := that.(GetNodesResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Nodes) != len(that1.Nodes) {
		return false
	}
	for i := range this.Nodes {
		if !this.Nodes[i].Equal(that1.Nodes[i]) {
			return false
		}
	}
	return true
}
func (this *GetNodeBySignatureRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*GetNodeBySignatureRequest)
	if !ok {
		that2, ok := that.(GetNodeBySignatureRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Signature.Equal(that1.Signature) {
		return false
	}
	return true
}
func (this *GetNodeBySignatureResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*GetNodeBySignatureResponse)
	if !ok {
		that2, ok := that.(GetNodeBySignatureResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Node.Equal(that1.Node) {
		return false
	}
	return true
}
func (this *GetEdgeRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*GetEdgeRequest)
	if !ok {
		that2, ok := that.(GetEdgeRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Uuid != that1.Uuid {
		return false
	}
	return true
}
func (this *GetEdgeResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*GetEdgeResponse)
	if !ok {
		that2, ok := that.(GetEdgeResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Edge.Equal(that1.Edge) {
		return false
	}
	return true
}
func (this *GetEdgeByNodesRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*GetEdgeByNodesRequest)
	if !ok {
		that2, ok := that.(GetEdgeByNodesRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Sender.Equal(that1.Sender) {
		return false
	}
	if !this.Receiver.Equal(that1.Receiver) {
		return false
	}
	return true
}
func (this *GetEdgeByNodesResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*GetEdgeByNodesResponse)
	if !ok {
		that2, ok := that.(GetEdgeByNodesResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Edge.Equal(that1.Edge) {
		return false
	}
	return true
}
func (this *GetEdgesRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*GetEdgesRequest)
	if !ok {
		that2, ok := that.(GetEdgesRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *GetEdgesResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*GetEdgesResponse)
	if !ok {
		that2, ok := that.(GetEdgesResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Edges) != len(that1.Edges) {
		return false
	}
	for i := range this.Edges {
		if !this.Edges[i].Equal(that1.Edges[i]) {
			return false
		}
	}
	return true
}
func (this *GetEventLogForEdgeRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*GetEventLogForEdgeRequest)
	if !ok {
		that2, ok := that.(GetEventLogForEdgeRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Uuid != that1.Uuid {
		return false
	}
	return true
}
func (this *GetEventLogForEdgeResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*GetEventLogForEdgeResponse)
	if !ok {
		that2, ok := that.(GetEventLogForEdgeResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Events) != len(that1.Events) {
		return false
	}
	for i := range this.Events {
		if !this.Events[i].Equal(that1.Events[i]) {
			return false
		}
	}
	return true
}
func (this *SetEdgeStateRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*SetEdgeStateRequest)
	if !ok {
		that2, ok := that.(SetEdgeStateRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.EdgeUuid != that1.EdgeUuid {
		return false
	}
	if this.State != that1.State {
		return false
	}
	if this.Comment != that1.Comment {
		return false
	}
	return true
}
func (this *SetEdgeStateResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*SetEdgeStateResponse)
	if !ok {
		that2, ok := that.(SetEdgeStateResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *GetRecordsRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*GetRecordsRequest)
	if !ok {
		that2, ok := that.(GetRecordsRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	return true
}
func (this *GetRecordsResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*GetRecordsResponse)
	if !ok {
		that2, ok := that.(GetRecordsResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Records) != len(that1.Records) {
		return false
	}
	for i := range this.Records {
		if !this.Records[i].Equal(that1.Records[i]) {
			return false
		}
	}
	return true
}
func (this *IsRegulatedRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*IsRegulatedRequest)
	if !ok {
		that2, ok := that.(IsRegulatedRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Signature.Equal(that1.Signature) {
		return false
	}
	return true
}
func (this *IsRegulatedResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*IsRegulatedResponse)
	if !ok {
		that2, ok := that.(IsRegulatedResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Regulated != that1.Regulated {
		return false
	}
	return true
}
func (this *TestDetectEdgeRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*TestDetectEdgeRequest)
	if !ok {
		that2, ok := that.(TestDetectEdgeRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Sender.Equal(that1.Sender) {
		return false
	}
	if !this.Receiver.Equal(that1.Receiver) {
		return false
	}
	return true
}
func (this *TestDetectEdgeResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*TestDetectEdgeResponse)
	if !ok {
		that2, ok := that.(TestDetectEdgeResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Edge.Equal(that1.Edge) {
		return false
	}
	return true
}
func (this *Address) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&connmonpb.Address{")
	s = append(s, "Ip: "+fmt.Sprintf("%#v", this.Ip)+",\n")
	s = append(s, "Port: "+fmt.Sprintf("%#v", this.Port)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *AggregatorTopicEvent) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&connmonpb.AggregatorTopicEvent{")
	if this.Event != nil {
		s = append(s, "Event: "+fmt.Sprintf("%#v", this.Event)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *AggregatorTopicEvent_ConnectionEvent) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&connmonpb.AggregatorTopicEvent_ConnectionEvent{` +
		`ConnectionEvent:` + fmt.Sprintf("%#v", this.ConnectionEvent) + `}`}, ", ")
	return s
}
func (this *AggregatorTopicEvent_ResolvedConnectionEvent) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&connmonpb.AggregatorTopicEvent_ResolvedConnectionEvent{` +
		`ResolvedConnectionEvent:` + fmt.Sprintf("%#v", this.ResolvedConnectionEvent) + `}`}, ", ")
	return s
}
func (this *AggregatorTopicEvent_TopicEvent) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&connmonpb.AggregatorTopicEvent_TopicEvent{` +
		`TopicEvent:` + fmt.Sprintf("%#v", this.TopicEvent) + `}`}, ", ")
	return s
}
func (this *ConnectionEvent) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&connmonpb.ConnectionEvent{")
	if this.InboundListenAddress != nil {
		s = append(s, "InboundListenAddress: "+fmt.Sprintf("%#v", this.InboundListenAddress)+",\n")
	}
	if this.Event != nil {
		s = append(s, "Event: "+fmt.Sprintf("%#v", this.Event)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ConnectionEvent_OutboundEvent) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&connmonpb.ConnectionEvent_OutboundEvent{` +
		`OutboundEvent:` + fmt.Sprintf("%#v", this.OutboundEvent) + `}`}, ", ")
	return s
}
func (this *ConnectionEvent_InboundEvent) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&connmonpb.ConnectionEvent_InboundEvent{` +
		`InboundEvent:` + fmt.Sprintf("%#v", this.InboundEvent) + `}`}, ", ")
	return s
}
func (this *Application) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&connmonpb.Application{")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ConnectionOutboundEvent) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&connmonpb.ConnectionOutboundEvent{")
	if this.OutboundApplication != nil {
		s = append(s, "OutboundApplication: "+fmt.Sprintf("%#v", this.OutboundApplication)+",\n")
	}
	if this.OutboundAddress != nil {
		s = append(s, "OutboundAddress: "+fmt.Sprintf("%#v", this.OutboundAddress)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ConnectionInboundEvent) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&connmonpb.ConnectionInboundEvent{")
	if this.InboundApplication != nil {
		s = append(s, "InboundApplication: "+fmt.Sprintf("%#v", this.InboundApplication)+",\n")
	}
	if this.InboundAddress != nil {
		s = append(s, "InboundAddress: "+fmt.Sprintf("%#v", this.InboundAddress)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ResolvedConnectionEvent) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&connmonpb.ResolvedConnectionEvent{")
	if this.OutboundApplication != nil {
		s = append(s, "OutboundApplication: "+fmt.Sprintf("%#v", this.OutboundApplication)+",\n")
	}
	if this.InboundApplication != nil {
		s = append(s, "InboundApplication: "+fmt.Sprintf("%#v", this.InboundApplication)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TopicEvent) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&connmonpb.TopicEvent{")
	if this.OutboundAddress != nil {
		s = append(s, "OutboundAddress: "+fmt.Sprintf("%#v", this.OutboundAddress)+",\n")
	}
	if this.Application != nil {
		s = append(s, "Application: "+fmt.Sprintf("%#v", this.Application)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TopicEvent_OutboundApplication) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&connmonpb.TopicEvent_OutboundApplication{` +
		`OutboundApplication:` + fmt.Sprintf("%#v", this.OutboundApplication) + `}`}, ", ")
	return s
}
func (this *TopicEvent_InboundApplication) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&connmonpb.TopicEvent_InboundApplication{` +
		`InboundApplication:` + fmt.Sprintf("%#v", this.InboundApplication) + `}`}, ", ")
	return s
}
func (this *NodeSignature) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&connmonpb.NodeSignature{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Node) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&connmonpb.Node{")
	s = append(s, "Uuid: "+fmt.Sprintf("%#v", this.Uuid)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "NoNotify: "+fmt.Sprintf("%#v", this.NoNotify)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Edge) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&connmonpb.Edge{")
	s = append(s, "Uuid: "+fmt.Sprintf("%#v", this.Uuid)+",\n")
	if this.Sender != nil {
		s = append(s, "Sender: "+fmt.Sprintf("%#v", this.Sender)+",\n")
	}
	if this.Receiver != nil {
		s = append(s, "Receiver: "+fmt.Sprintf("%#v", this.Receiver)+",\n")
	}
	if this.LastSeen != nil {
		s = append(s, "LastSeen: "+fmt.Sprintf("%#v", this.LastSeen)+",\n")
	}
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Event) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 10)
	s = append(s, "&connmonpb.Event{")
	s = append(s, "Uuid: "+fmt.Sprintf("%#v", this.Uuid)+",\n")
	if this.Timestamp != nil {
		s = append(s, "Timestamp: "+fmt.Sprintf("%#v", this.Timestamp)+",\n")
	}
	s = append(s, "Username: "+fmt.Sprintf("%#v", this.Username)+",\n")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "Meta: "+fmt.Sprintf("%#v", this.Meta)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Record) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&connmonpb.Record{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetNodesRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&connmonpb.GetNodesRequest{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetNodesResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&connmonpb.GetNodesResponse{")
	if this.Nodes != nil {
		s = append(s, "Nodes: "+fmt.Sprintf("%#v", this.Nodes)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetNodeBySignatureRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&connmonpb.GetNodeBySignatureRequest{")
	if this.Signature != nil {
		s = append(s, "Signature: "+fmt.Sprintf("%#v", this.Signature)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetNodeBySignatureResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&connmonpb.GetNodeBySignatureResponse{")
	if this.Node != nil {
		s = append(s, "Node: "+fmt.Sprintf("%#v", this.Node)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetEdgeRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&connmonpb.GetEdgeRequest{")
	s = append(s, "Uuid: "+fmt.Sprintf("%#v", this.Uuid)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetEdgeResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&connmonpb.GetEdgeResponse{")
	if this.Edge != nil {
		s = append(s, "Edge: "+fmt.Sprintf("%#v", this.Edge)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetEdgeByNodesRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&connmonpb.GetEdgeByNodesRequest{")
	if this.Sender != nil {
		s = append(s, "Sender: "+fmt.Sprintf("%#v", this.Sender)+",\n")
	}
	if this.Receiver != nil {
		s = append(s, "Receiver: "+fmt.Sprintf("%#v", this.Receiver)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetEdgeByNodesResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&connmonpb.GetEdgeByNodesResponse{")
	if this.Edge != nil {
		s = append(s, "Edge: "+fmt.Sprintf("%#v", this.Edge)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetEdgesRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&connmonpb.GetEdgesRequest{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetEdgesResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&connmonpb.GetEdgesResponse{")
	if this.Edges != nil {
		s = append(s, "Edges: "+fmt.Sprintf("%#v", this.Edges)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetEventLogForEdgeRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&connmonpb.GetEventLogForEdgeRequest{")
	s = append(s, "Uuid: "+fmt.Sprintf("%#v", this.Uuid)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetEventLogForEdgeResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&connmonpb.GetEventLogForEdgeResponse{")
	if this.Events != nil {
		s = append(s, "Events: "+fmt.Sprintf("%#v", this.Events)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SetEdgeStateRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&connmonpb.SetEdgeStateRequest{")
	s = append(s, "EdgeUuid: "+fmt.Sprintf("%#v", this.EdgeUuid)+",\n")
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
	s = append(s, "Comment: "+fmt.Sprintf("%#v", this.Comment)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SetEdgeStateResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&connmonpb.SetEdgeStateResponse{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetRecordsRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&connmonpb.GetRecordsRequest{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetRecordsResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&connmonpb.GetRecordsResponse{")
	if this.Records != nil {
		s = append(s, "Records: "+fmt.Sprintf("%#v", this.Records)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *IsRegulatedRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&connmonpb.IsRegulatedRequest{")
	if this.Signature != nil {
		s = append(s, "Signature: "+fmt.Sprintf("%#v", this.Signature)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *IsRegulatedResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&connmonpb.IsRegulatedResponse{")
	s = append(s, "Regulated: "+fmt.Sprintf("%#v", this.Regulated)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TestDetectEdgeRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&connmonpb.TestDetectEdgeRequest{")
	if this.Sender != nil {
		s = append(s, "Sender: "+fmt.Sprintf("%#v", this.Sender)+",\n")
	}
	if this.Receiver != nil {
		s = append(s, "Receiver: "+fmt.Sprintf("%#v", this.Receiver)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TestDetectEdgeResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&connmonpb.TestDetectEdgeResponse{")
	if this.Edge != nil {
		s = append(s, "Edge: "+fmt.Sprintf("%#v", this.Edge)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringConnmon(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Analyzer service

type AnalyzerClient interface {
	GetNodes(ctx context.Context, in *GetNodesRequest, opts ...grpc.CallOption) (*GetNodesResponse, error)
	GetNodeBySignature(ctx context.Context, in *GetNodeBySignatureRequest, opts ...grpc.CallOption) (*GetNodeBySignatureResponse, error)
	GetEdges(ctx context.Context, in *GetEdgesRequest, opts ...grpc.CallOption) (*GetEdgesResponse, error)
	GetEdge(ctx context.Context, in *GetEdgeRequest, opts ...grpc.CallOption) (*GetEdgeResponse, error)
	GetEdgeByNodes(ctx context.Context, in *GetEdgeByNodesRequest, opts ...grpc.CallOption) (*GetEdgeByNodesResponse, error)
	GetEventLogForEdge(ctx context.Context, in *GetEventLogForEdgeRequest, opts ...grpc.CallOption) (*GetEventLogForEdgeResponse, error)
	SetEdgeState(ctx context.Context, in *SetEdgeStateRequest, opts ...grpc.CallOption) (*SetEdgeStateResponse, error)
	GetRecords(ctx context.Context, in *GetRecordsRequest, opts ...grpc.CallOption) (*GetRecordsResponse, error)
	IsRegulated(ctx context.Context, in *IsRegulatedRequest, opts ...grpc.CallOption) (*IsRegulatedResponse, error)
	TestDetectEdge(ctx context.Context, in *TestDetectEdgeRequest, opts ...grpc.CallOption) (*TestDetectEdgeResponse, error)
}

type analyzerClient struct {
	cc *grpc.ClientConn
}

func NewAnalyzerClient(cc *grpc.ClientConn) AnalyzerClient {
	return &analyzerClient{cc}
}

func (c *analyzerClient) GetNodes(ctx context.Context, in *GetNodesRequest, opts ...grpc.CallOption) (*GetNodesResponse, error) {
	out := new(GetNodesResponse)
	err := grpc.Invoke(ctx, "/connmon.Analyzer/GetNodes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) GetNodeBySignature(ctx context.Context, in *GetNodeBySignatureRequest, opts ...grpc.CallOption) (*GetNodeBySignatureResponse, error) {
	out := new(GetNodeBySignatureResponse)
	err := grpc.Invoke(ctx, "/connmon.Analyzer/GetNodeBySignature", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) GetEdges(ctx context.Context, in *GetEdgesRequest, opts ...grpc.CallOption) (*GetEdgesResponse, error) {
	out := new(GetEdgesResponse)
	err := grpc.Invoke(ctx, "/connmon.Analyzer/GetEdges", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) GetEdge(ctx context.Context, in *GetEdgeRequest, opts ...grpc.CallOption) (*GetEdgeResponse, error) {
	out := new(GetEdgeResponse)
	err := grpc.Invoke(ctx, "/connmon.Analyzer/GetEdge", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) GetEdgeByNodes(ctx context.Context, in *GetEdgeByNodesRequest, opts ...grpc.CallOption) (*GetEdgeByNodesResponse, error) {
	out := new(GetEdgeByNodesResponse)
	err := grpc.Invoke(ctx, "/connmon.Analyzer/GetEdgeByNodes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) GetEventLogForEdge(ctx context.Context, in *GetEventLogForEdgeRequest, opts ...grpc.CallOption) (*GetEventLogForEdgeResponse, error) {
	out := new(GetEventLogForEdgeResponse)
	err := grpc.Invoke(ctx, "/connmon.Analyzer/GetEventLogForEdge", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) SetEdgeState(ctx context.Context, in *SetEdgeStateRequest, opts ...grpc.CallOption) (*SetEdgeStateResponse, error) {
	out := new(SetEdgeStateResponse)
	err := grpc.Invoke(ctx, "/connmon.Analyzer/SetEdgeState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) GetRecords(ctx context.Context, in *GetRecordsRequest, opts ...grpc.CallOption) (*GetRecordsResponse, error) {
	out := new(GetRecordsResponse)
	err := grpc.Invoke(ctx, "/connmon.Analyzer/GetRecords", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) IsRegulated(ctx context.Context, in *IsRegulatedRequest, opts ...grpc.CallOption) (*IsRegulatedResponse, error) {
	out := new(IsRegulatedResponse)
	err := grpc.Invoke(ctx, "/connmon.Analyzer/IsRegulated", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *analyzerClient) TestDetectEdge(ctx context.Context, in *TestDetectEdgeRequest, opts ...grpc.CallOption) (*TestDetectEdgeResponse, error) {
	out := new(TestDetectEdgeResponse)
	err := grpc.Invoke(ctx, "/connmon.Analyzer/TestDetectEdge", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Analyzer service

type AnalyzerServer interface {
	GetNodes(context.Context, *GetNodesRequest) (*GetNodesResponse, error)
	GetNodeBySignature(context.Context, *GetNodeBySignatureRequest) (*GetNodeBySignatureResponse, error)
	GetEdges(context.Context, *GetEdgesRequest) (*GetEdgesResponse, error)
	GetEdge(context.Context, *GetEdgeRequest) (*GetEdgeResponse, error)
	GetEdgeByNodes(context.Context, *GetEdgeByNodesRequest) (*GetEdgeByNodesResponse, error)
	GetEventLogForEdge(context.Context, *GetEventLogForEdgeRequest) (*GetEventLogForEdgeResponse, error)
	SetEdgeState(context.Context, *SetEdgeStateRequest) (*SetEdgeStateResponse, error)
	GetRecords(context.Context, *GetRecordsRequest) (*GetRecordsResponse, error)
	IsRegulated(context.Context, *IsRegulatedRequest) (*IsRegulatedResponse, error)
	TestDetectEdge(context.Context, *TestDetectEdgeRequest) (*TestDetectEdgeResponse, error)
}

func RegisterAnalyzerServer(s *grpc.Server, srv AnalyzerServer) {
	s.RegisterService(&_Analyzer_serviceDesc, srv)
}

func _Analyzer_GetNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).GetNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connmon.Analyzer/GetNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).GetNodes(ctx, req.(*GetNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_GetNodeBySignature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeBySignatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).GetNodeBySignature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connmon.Analyzer/GetNodeBySignature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).GetNodeBySignature(ctx, req.(*GetNodeBySignatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_GetEdges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEdgesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).GetEdges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connmon.Analyzer/GetEdges",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).GetEdges(ctx, req.(*GetEdgesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_GetEdge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEdgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).GetEdge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connmon.Analyzer/GetEdge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).GetEdge(ctx, req.(*GetEdgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_GetEdgeByNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEdgeByNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).GetEdgeByNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connmon.Analyzer/GetEdgeByNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).GetEdgeByNodes(ctx, req.(*GetEdgeByNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_GetEventLogForEdge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventLogForEdgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).GetEventLogForEdge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connmon.Analyzer/GetEventLogForEdge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).GetEventLogForEdge(ctx, req.(*GetEventLogForEdgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_SetEdgeState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetEdgeStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).SetEdgeState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connmon.Analyzer/SetEdgeState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).SetEdgeState(ctx, req.(*SetEdgeStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_GetRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).GetRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connmon.Analyzer/GetRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).GetRecords(ctx, req.(*GetRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_IsRegulated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsRegulatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).IsRegulated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connmon.Analyzer/IsRegulated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).IsRegulated(ctx, req.(*IsRegulatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Analyzer_TestDetectEdge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestDetectEdgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalyzerServer).TestDetectEdge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connmon.Analyzer/TestDetectEdge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalyzerServer).TestDetectEdge(ctx, req.(*TestDetectEdgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Analyzer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "connmon.Analyzer",
	HandlerType: (*AnalyzerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNodes",
			Handler:    _Analyzer_GetNodes_Handler,
		},
		{
			MethodName: "GetNodeBySignature",
			Handler:    _Analyzer_GetNodeBySignature_Handler,
		},
		{
			MethodName: "GetEdges",
			Handler:    _Analyzer_GetEdges_Handler,
		},
		{
			MethodName: "GetEdge",
			Handler:    _Analyzer_GetEdge_Handler,
		},
		{
			MethodName: "GetEdgeByNodes",
			Handler:    _Analyzer_GetEdgeByNodes_Handler,
		},
		{
			MethodName: "GetEventLogForEdge",
			Handler:    _Analyzer_GetEventLogForEdge_Handler,
		},
		{
			MethodName: "SetEdgeState",
			Handler:    _Analyzer_SetEdgeState_Handler,
		},
		{
			MethodName: "GetRecords",
			Handler:    _Analyzer_GetRecords_Handler,
		},
		{
			MethodName: "IsRegulated",
			Handler:    _Analyzer_IsRegulated_Handler,
		},
		{
			MethodName: "TestDetectEdge",
			Handler:    _Analyzer_TestDetectEdge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "infra/connmon/connmon.proto",
}

func (m *Address) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Address) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Ip) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(len(m.Ip)))
		i += copy(dAtA[i:], m.Ip)
	}
	if m.Port != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.Port))
	}
	return i, nil
}

func (m *AggregatorTopicEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AggregatorTopicEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Event != nil {
		nn1, err := m.Event.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *AggregatorTopicEvent_ConnectionEvent) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.ConnectionEvent != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.ConnectionEvent.Size()))
		n2, err := m.ConnectionEvent.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func (m *AggregatorTopicEvent_ResolvedConnectionEvent) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.ResolvedConnectionEvent != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.ResolvedConnectionEvent.Size()))
		n3, err := m.ResolvedConnectionEvent.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func (m *AggregatorTopicEvent_TopicEvent) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.TopicEvent != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.TopicEvent.Size()))
		n4, err := m.TopicEvent.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func (m *ConnectionEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConnectionEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.InboundListenAddress != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.InboundListenAddress.Size()))
		n5, err := m.InboundListenAddress.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.Event != nil {
		nn6, err := m.Event.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn6
	}
	return i, nil
}

func (m *ConnectionEvent_OutboundEvent) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.OutboundEvent != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.OutboundEvent.Size()))
		n7, err := m.OutboundEvent.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}
func (m *ConnectionEvent_InboundEvent) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.InboundEvent != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.InboundEvent.Size()))
		n8, err := m.InboundEvent.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	return i, nil
}
func (m *Application) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Application) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.Type))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func (m *ConnectionOutboundEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConnectionOutboundEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.OutboundApplication != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.OutboundApplication.Size()))
		n9, err := m.OutboundApplication.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n9
	}
	if m.OutboundAddress != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.OutboundAddress.Size()))
		n10, err := m.OutboundAddress.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n10
	}
	return i, nil
}

func (m *ConnectionInboundEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConnectionInboundEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.InboundApplication != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.InboundApplication.Size()))
		n11, err := m.InboundApplication.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n11
	}
	if m.InboundAddress != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.InboundAddress.Size()))
		n12, err := m.InboundAddress.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n12
	}
	return i, nil
}

func (m *ResolvedConnectionEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResolvedConnectionEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.OutboundApplication != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.OutboundApplication.Size()))
		n13, err := m.OutboundApplication.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n13
	}
	if m.InboundApplication != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.InboundApplication.Size()))
		n14, err := m.InboundApplication.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n14
	}
	return i, nil
}

func (m *TopicEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TopicEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.OutboundAddress != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.OutboundAddress.Size()))
		n15, err := m.OutboundAddress.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n15
	}
	if m.Application != nil {
		nn16, err := m.Application.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn16
	}
	return i, nil
}

func (m *TopicEvent_OutboundApplication) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.OutboundApplication != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.OutboundApplication.Size()))
		n17, err := m.OutboundApplication.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n17
	}
	return i, nil
}
func (m *TopicEvent_InboundApplication) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.InboundApplication != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.InboundApplication.Size()))
		n18, err := m.InboundApplication.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n18
	}
	return i, nil
}
func (m *NodeSignature) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeSignature) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Type) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	return i, nil
}

func (m *Node) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Node) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Uuid) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(len(m.Uuid)))
		i += copy(dAtA[i:], m.Uuid)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Type) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if m.NoNotify {
		dAtA[i] = 0x20
		i++
		if m.NoNotify {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *Edge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Edge) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Uuid) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(len(m.Uuid)))
		i += copy(dAtA[i:], m.Uuid)
	}
	if m.Sender != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.Sender.Size()))
		n19, err := m.Sender.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n19
	}
	if m.Receiver != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.Receiver.Size()))
		n20, err := m.Receiver.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n20
	}
	if m.LastSeen != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.LastSeen.Size()))
		n21, err := m.LastSeen.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n21
	}
	if m.State != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.State))
	}
	return i, nil
}

func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Uuid) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(len(m.Uuid)))
		i += copy(dAtA[i:], m.Uuid)
	}
	if m.Timestamp != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.Timestamp.Size()))
		n22, err := m.Timestamp.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n22
	}
	if len(m.Username) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(len(m.Username)))
		i += copy(dAtA[i:], m.Username)
	}
	if m.Type != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.Type))
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	if len(m.Meta) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(len(m.Meta)))
		i += copy(dAtA[i:], m.Meta)
	}
	return i, nil
}

func (m *Record) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Record) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Type) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	return i, nil
}

func (m *GetNodesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetNodesRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *GetNodesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetNodesResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Nodes) > 0 {
		for _, msg := range m.Nodes {
			dAtA[i] = 0xa
			i++
			i = encodeVarintConnmon(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *GetNodeBySignatureRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetNodeBySignatureRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Signature != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.Signature.Size()))
		n23, err := m.Signature.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n23
	}
	return i, nil
}

func (m *GetNodeBySignatureResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetNodeBySignatureResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Node != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.Node.Size()))
		n24, err := m.Node.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n24
	}
	return i, nil
}

func (m *GetEdgeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetEdgeRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Uuid) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(len(m.Uuid)))
		i += copy(dAtA[i:], m.Uuid)
	}
	return i, nil
}

func (m *GetEdgeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetEdgeResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Edge != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.Edge.Size()))
		n25, err := m.Edge.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n25
	}
	return i, nil
}

func (m *GetEdgeByNodesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetEdgeByNodesRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Sender != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.Sender.Size()))
		n26, err := m.Sender.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n26
	}
	if m.Receiver != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.Receiver.Size()))
		n27, err := m.Receiver.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n27
	}
	return i, nil
}

func (m *GetEdgeByNodesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetEdgeByNodesResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Edge != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.Edge.Size()))
		n28, err := m.Edge.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n28
	}
	return i, nil
}

func (m *GetEdgesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetEdgesRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *GetEdgesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetEdgesResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Edges) > 0 {
		for _, msg := range m.Edges {
			dAtA[i] = 0xa
			i++
			i = encodeVarintConnmon(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *GetEventLogForEdgeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetEventLogForEdgeRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Uuid) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(len(m.Uuid)))
		i += copy(dAtA[i:], m.Uuid)
	}
	return i, nil
}

func (m *GetEventLogForEdgeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetEventLogForEdgeResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Events) > 0 {
		for _, msg := range m.Events {
			dAtA[i] = 0xa
			i++
			i = encodeVarintConnmon(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *SetEdgeStateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetEdgeStateRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.EdgeUuid) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(len(m.EdgeUuid)))
		i += copy(dAtA[i:], m.EdgeUuid)
	}
	if m.State != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.State))
	}
	if len(m.Comment) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(len(m.Comment)))
		i += copy(dAtA[i:], m.Comment)
	}
	return i, nil
}

func (m *SetEdgeStateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetEdgeStateResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *GetRecordsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetRecordsRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *GetRecordsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetRecordsResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Records) > 0 {
		for _, msg := range m.Records {
			dAtA[i] = 0xa
			i++
			i = encodeVarintConnmon(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *IsRegulatedRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IsRegulatedRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Signature != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.Signature.Size()))
		n29, err := m.Signature.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n29
	}
	return i, nil
}

func (m *IsRegulatedResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IsRegulatedResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Regulated {
		dAtA[i] = 0x8
		i++
		if m.Regulated {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *TestDetectEdgeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestDetectEdgeRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Sender != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.Sender.Size()))
		n30, err := m.Sender.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n30
	}
	if m.Receiver != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.Receiver.Size()))
		n31, err := m.Receiver.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n31
	}
	return i, nil
}

func (m *TestDetectEdgeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestDetectEdgeResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Edge != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConnmon(dAtA, i, uint64(m.Edge.Size()))
		n32, err := m.Edge.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n32
	}
	return i, nil
}

func encodeVarintConnmon(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Address) Size() (n int) {
	var l int
	_ = l
	l = len(m.Ip)
	if l > 0 {
		n += 1 + l + sovConnmon(uint64(l))
	}
	if m.Port != 0 {
		n += 1 + sovConnmon(uint64(m.Port))
	}
	return n
}

func (m *AggregatorTopicEvent) Size() (n int) {
	var l int
	_ = l
	if m.Event != nil {
		n += m.Event.Size()
	}
	return n
}

func (m *AggregatorTopicEvent_ConnectionEvent) Size() (n int) {
	var l int
	_ = l
	if m.ConnectionEvent != nil {
		l = m.ConnectionEvent.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}
func (m *AggregatorTopicEvent_ResolvedConnectionEvent) Size() (n int) {
	var l int
	_ = l
	if m.ResolvedConnectionEvent != nil {
		l = m.ResolvedConnectionEvent.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}
func (m *AggregatorTopicEvent_TopicEvent) Size() (n int) {
	var l int
	_ = l
	if m.TopicEvent != nil {
		l = m.TopicEvent.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}
func (m *ConnectionEvent) Size() (n int) {
	var l int
	_ = l
	if m.InboundListenAddress != nil {
		l = m.InboundListenAddress.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	if m.Event != nil {
		n += m.Event.Size()
	}
	return n
}

func (m *ConnectionEvent_OutboundEvent) Size() (n int) {
	var l int
	_ = l
	if m.OutboundEvent != nil {
		l = m.OutboundEvent.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}
func (m *ConnectionEvent_InboundEvent) Size() (n int) {
	var l int
	_ = l
	if m.InboundEvent != nil {
		l = m.InboundEvent.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}
func (m *Application) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovConnmon(uint64(m.Type))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}

func (m *ConnectionOutboundEvent) Size() (n int) {
	var l int
	_ = l
	if m.OutboundApplication != nil {
		l = m.OutboundApplication.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	if m.OutboundAddress != nil {
		l = m.OutboundAddress.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}

func (m *ConnectionInboundEvent) Size() (n int) {
	var l int
	_ = l
	if m.InboundApplication != nil {
		l = m.InboundApplication.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	if m.InboundAddress != nil {
		l = m.InboundAddress.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}

func (m *ResolvedConnectionEvent) Size() (n int) {
	var l int
	_ = l
	if m.OutboundApplication != nil {
		l = m.OutboundApplication.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	if m.InboundApplication != nil {
		l = m.InboundApplication.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}

func (m *TopicEvent) Size() (n int) {
	var l int
	_ = l
	if m.OutboundAddress != nil {
		l = m.OutboundAddress.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	if m.Application != nil {
		n += m.Application.Size()
	}
	return n
}

func (m *TopicEvent_OutboundApplication) Size() (n int) {
	var l int
	_ = l
	if m.OutboundApplication != nil {
		l = m.OutboundApplication.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}
func (m *TopicEvent_InboundApplication) Size() (n int) {
	var l int
	_ = l
	if m.InboundApplication != nil {
		l = m.InboundApplication.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}
func (m *NodeSignature) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovConnmon(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}

func (m *Node) Size() (n int) {
	var l int
	_ = l
	l = len(m.Uuid)
	if l > 0 {
		n += 1 + l + sovConnmon(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovConnmon(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovConnmon(uint64(l))
	}
	if m.NoNotify {
		n += 2
	}
	return n
}

func (m *Edge) Size() (n int) {
	var l int
	_ = l
	l = len(m.Uuid)
	if l > 0 {
		n += 1 + l + sovConnmon(uint64(l))
	}
	if m.Sender != nil {
		l = m.Sender.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	if m.Receiver != nil {
		l = m.Receiver.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	if m.LastSeen != nil {
		l = m.LastSeen.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	if m.State != 0 {
		n += 1 + sovConnmon(uint64(m.State))
	}
	return n
}

func (m *Event) Size() (n int) {
	var l int
	_ = l
	l = len(m.Uuid)
	if l > 0 {
		n += 1 + l + sovConnmon(uint64(l))
	}
	if m.Timestamp != nil {
		l = m.Timestamp.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovConnmon(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovConnmon(uint64(m.Type))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovConnmon(uint64(l))
	}
	l = len(m.Meta)
	if l > 0 {
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}

func (m *Record) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovConnmon(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}

func (m *GetNodesRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *GetNodesResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Nodes) > 0 {
		for _, e := range m.Nodes {
			l = e.Size()
			n += 1 + l + sovConnmon(uint64(l))
		}
	}
	return n
}

func (m *GetNodeBySignatureRequest) Size() (n int) {
	var l int
	_ = l
	if m.Signature != nil {
		l = m.Signature.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}

func (m *GetNodeBySignatureResponse) Size() (n int) {
	var l int
	_ = l
	if m.Node != nil {
		l = m.Node.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}

func (m *GetEdgeRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Uuid)
	if l > 0 {
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}

func (m *GetEdgeResponse) Size() (n int) {
	var l int
	_ = l
	if m.Edge != nil {
		l = m.Edge.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}

func (m *GetEdgeByNodesRequest) Size() (n int) {
	var l int
	_ = l
	if m.Sender != nil {
		l = m.Sender.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	if m.Receiver != nil {
		l = m.Receiver.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}

func (m *GetEdgeByNodesResponse) Size() (n int) {
	var l int
	_ = l
	if m.Edge != nil {
		l = m.Edge.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}

func (m *GetEdgesRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *GetEdgesResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Edges) > 0 {
		for _, e := range m.Edges {
			l = e.Size()
			n += 1 + l + sovConnmon(uint64(l))
		}
	}
	return n
}

func (m *GetEventLogForEdgeRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Uuid)
	if l > 0 {
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}

func (m *GetEventLogForEdgeResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovConnmon(uint64(l))
		}
	}
	return n
}

func (m *SetEdgeStateRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.EdgeUuid)
	if l > 0 {
		n += 1 + l + sovConnmon(uint64(l))
	}
	if m.State != 0 {
		n += 1 + sovConnmon(uint64(m.State))
	}
	l = len(m.Comment)
	if l > 0 {
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}

func (m *SetEdgeStateResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *GetRecordsRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *GetRecordsResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Records) > 0 {
		for _, e := range m.Records {
			l = e.Size()
			n += 1 + l + sovConnmon(uint64(l))
		}
	}
	return n
}

func (m *IsRegulatedRequest) Size() (n int) {
	var l int
	_ = l
	if m.Signature != nil {
		l = m.Signature.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}

func (m *IsRegulatedResponse) Size() (n int) {
	var l int
	_ = l
	if m.Regulated {
		n += 2
	}
	return n
}

func (m *TestDetectEdgeRequest) Size() (n int) {
	var l int
	_ = l
	if m.Sender != nil {
		l = m.Sender.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	if m.Receiver != nil {
		l = m.Receiver.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}

func (m *TestDetectEdgeResponse) Size() (n int) {
	var l int
	_ = l
	if m.Edge != nil {
		l = m.Edge.Size()
		n += 1 + l + sovConnmon(uint64(l))
	}
	return n
}

func sovConnmon(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConnmon(x uint64) (n int) {
	return sovConnmon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Address) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Address{`,
		`Ip:` + fmt.Sprintf("%v", this.Ip) + `,`,
		`Port:` + fmt.Sprintf("%v", this.Port) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AggregatorTopicEvent) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AggregatorTopicEvent{`,
		`Event:` + fmt.Sprintf("%v", this.Event) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AggregatorTopicEvent_ConnectionEvent) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AggregatorTopicEvent_ConnectionEvent{`,
		`ConnectionEvent:` + strings.Replace(fmt.Sprintf("%v", this.ConnectionEvent), "ConnectionEvent", "ConnectionEvent", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AggregatorTopicEvent_ResolvedConnectionEvent) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AggregatorTopicEvent_ResolvedConnectionEvent{`,
		`ResolvedConnectionEvent:` + strings.Replace(fmt.Sprintf("%v", this.ResolvedConnectionEvent), "ResolvedConnectionEvent", "ResolvedConnectionEvent", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *AggregatorTopicEvent_TopicEvent) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&AggregatorTopicEvent_TopicEvent{`,
		`TopicEvent:` + strings.Replace(fmt.Sprintf("%v", this.TopicEvent), "TopicEvent", "TopicEvent", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ConnectionEvent) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ConnectionEvent{`,
		`InboundListenAddress:` + strings.Replace(fmt.Sprintf("%v", this.InboundListenAddress), "Address", "Address", 1) + `,`,
		`Event:` + fmt.Sprintf("%v", this.Event) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ConnectionEvent_OutboundEvent) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ConnectionEvent_OutboundEvent{`,
		`OutboundEvent:` + strings.Replace(fmt.Sprintf("%v", this.OutboundEvent), "ConnectionOutboundEvent", "ConnectionOutboundEvent", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ConnectionEvent_InboundEvent) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ConnectionEvent_InboundEvent{`,
		`InboundEvent:` + strings.Replace(fmt.Sprintf("%v", this.InboundEvent), "ConnectionInboundEvent", "ConnectionInboundEvent", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Application) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Application{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ConnectionOutboundEvent) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ConnectionOutboundEvent{`,
		`OutboundApplication:` + strings.Replace(fmt.Sprintf("%v", this.OutboundApplication), "Application", "Application", 1) + `,`,
		`OutboundAddress:` + strings.Replace(fmt.Sprintf("%v", this.OutboundAddress), "Address", "Address", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ConnectionInboundEvent) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ConnectionInboundEvent{`,
		`InboundApplication:` + strings.Replace(fmt.Sprintf("%v", this.InboundApplication), "Application", "Application", 1) + `,`,
		`InboundAddress:` + strings.Replace(fmt.Sprintf("%v", this.InboundAddress), "Address", "Address", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ResolvedConnectionEvent) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ResolvedConnectionEvent{`,
		`OutboundApplication:` + strings.Replace(fmt.Sprintf("%v", this.OutboundApplication), "Application", "Application", 1) + `,`,
		`InboundApplication:` + strings.Replace(fmt.Sprintf("%v", this.InboundApplication), "Application", "Application", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TopicEvent) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TopicEvent{`,
		`OutboundAddress:` + strings.Replace(fmt.Sprintf("%v", this.OutboundAddress), "Address", "Address", 1) + `,`,
		`Application:` + fmt.Sprintf("%v", this.Application) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TopicEvent_OutboundApplication) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TopicEvent_OutboundApplication{`,
		`OutboundApplication:` + strings.Replace(fmt.Sprintf("%v", this.OutboundApplication), "Application", "Application", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TopicEvent_InboundApplication) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TopicEvent_InboundApplication{`,
		`InboundApplication:` + strings.Replace(fmt.Sprintf("%v", this.InboundApplication), "Application", "Application", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NodeSignature) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NodeSignature{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Node) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Node{`,
		`Uuid:` + fmt.Sprintf("%v", this.Uuid) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`NoNotify:` + fmt.Sprintf("%v", this.NoNotify) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Edge) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Edge{`,
		`Uuid:` + fmt.Sprintf("%v", this.Uuid) + `,`,
		`Sender:` + strings.Replace(fmt.Sprintf("%v", this.Sender), "Node", "Node", 1) + `,`,
		`Receiver:` + strings.Replace(fmt.Sprintf("%v", this.Receiver), "Node", "Node", 1) + `,`,
		`LastSeen:` + strings.Replace(fmt.Sprintf("%v", this.LastSeen), "Timestamp", "google_protobuf.Timestamp", 1) + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Event) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Event{`,
		`Uuid:` + fmt.Sprintf("%v", this.Uuid) + `,`,
		`Timestamp:` + strings.Replace(fmt.Sprintf("%v", this.Timestamp), "Timestamp", "google_protobuf.Timestamp", 1) + `,`,
		`Username:` + fmt.Sprintf("%v", this.Username) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`Meta:` + fmt.Sprintf("%v", this.Meta) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Record) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Record{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetNodesRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetNodesRequest{`,
		`}`,
	}, "")
	return s
}
func (this *GetNodesResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetNodesResponse{`,
		`Nodes:` + strings.Replace(fmt.Sprintf("%v", this.Nodes), "Node", "Node", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetNodeBySignatureRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetNodeBySignatureRequest{`,
		`Signature:` + strings.Replace(fmt.Sprintf("%v", this.Signature), "NodeSignature", "NodeSignature", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetNodeBySignatureResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetNodeBySignatureResponse{`,
		`Node:` + strings.Replace(fmt.Sprintf("%v", this.Node), "Node", "Node", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetEdgeRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetEdgeRequest{`,
		`Uuid:` + fmt.Sprintf("%v", this.Uuid) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetEdgeResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetEdgeResponse{`,
		`Edge:` + strings.Replace(fmt.Sprintf("%v", this.Edge), "Edge", "Edge", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetEdgeByNodesRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetEdgeByNodesRequest{`,
		`Sender:` + strings.Replace(fmt.Sprintf("%v", this.Sender), "NodeSignature", "NodeSignature", 1) + `,`,
		`Receiver:` + strings.Replace(fmt.Sprintf("%v", this.Receiver), "NodeSignature", "NodeSignature", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetEdgeByNodesResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetEdgeByNodesResponse{`,
		`Edge:` + strings.Replace(fmt.Sprintf("%v", this.Edge), "Edge", "Edge", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetEdgesRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetEdgesRequest{`,
		`}`,
	}, "")
	return s
}
func (this *GetEdgesResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetEdgesResponse{`,
		`Edges:` + strings.Replace(fmt.Sprintf("%v", this.Edges), "Edge", "Edge", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetEventLogForEdgeRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetEventLogForEdgeRequest{`,
		`Uuid:` + fmt.Sprintf("%v", this.Uuid) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetEventLogForEdgeResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetEventLogForEdgeResponse{`,
		`Events:` + strings.Replace(fmt.Sprintf("%v", this.Events), "Event", "Event", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SetEdgeStateRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SetEdgeStateRequest{`,
		`EdgeUuid:` + fmt.Sprintf("%v", this.EdgeUuid) + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
		`Comment:` + fmt.Sprintf("%v", this.Comment) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SetEdgeStateResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SetEdgeStateResponse{`,
		`}`,
	}, "")
	return s
}
func (this *GetRecordsRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetRecordsRequest{`,
		`}`,
	}, "")
	return s
}
func (this *GetRecordsResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetRecordsResponse{`,
		`Records:` + strings.Replace(fmt.Sprintf("%v", this.Records), "Record", "Record", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *IsRegulatedRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&IsRegulatedRequest{`,
		`Signature:` + strings.Replace(fmt.Sprintf("%v", this.Signature), "NodeSignature", "NodeSignature", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *IsRegulatedResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&IsRegulatedResponse{`,
		`Regulated:` + fmt.Sprintf("%v", this.Regulated) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TestDetectEdgeRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TestDetectEdgeRequest{`,
		`Sender:` + strings.Replace(fmt.Sprintf("%v", this.Sender), "NodeSignature", "NodeSignature", 1) + `,`,
		`Receiver:` + strings.Replace(fmt.Sprintf("%v", this.Receiver), "NodeSignature", "NodeSignature", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TestDetectEdgeResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TestDetectEdgeResponse{`,
		`Edge:` + strings.Replace(fmt.Sprintf("%v", this.Edge), "Edge", "Edge", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringConnmon(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Address) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Address: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Address: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ip", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ip = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AggregatorTopicEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AggregatorTopicEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AggregatorTopicEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionEvent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ConnectionEvent{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Event = &AggregatorTopicEvent_ConnectionEvent{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResolvedConnectionEvent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ResolvedConnectionEvent{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Event = &AggregatorTopicEvent_ResolvedConnectionEvent{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TopicEvent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &TopicEvent{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Event = &AggregatorTopicEvent_TopicEvent{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConnectionEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConnectionEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConnectionEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InboundListenAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.InboundListenAddress == nil {
				m.InboundListenAddress = &Address{}
			}
			if err := m.InboundListenAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutboundEvent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ConnectionOutboundEvent{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Event = &ConnectionEvent_OutboundEvent{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InboundEvent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ConnectionInboundEvent{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Event = &ConnectionEvent_InboundEvent{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Application) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Application: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Application: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (ApplicationType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConnectionOutboundEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConnectionOutboundEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConnectionOutboundEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutboundApplication", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OutboundApplication == nil {
				m.OutboundApplication = &Application{}
			}
			if err := m.OutboundApplication.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutboundAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OutboundAddress == nil {
				m.OutboundAddress = &Address{}
			}
			if err := m.OutboundAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConnectionInboundEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConnectionInboundEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConnectionInboundEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InboundApplication", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.InboundApplication == nil {
				m.InboundApplication = &Application{}
			}
			if err := m.InboundApplication.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InboundAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.InboundAddress == nil {
				m.InboundAddress = &Address{}
			}
			if err := m.InboundAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResolvedConnectionEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResolvedConnectionEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResolvedConnectionEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutboundApplication", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OutboundApplication == nil {
				m.OutboundApplication = &Application{}
			}
			if err := m.OutboundApplication.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InboundApplication", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.InboundApplication == nil {
				m.InboundApplication = &Application{}
			}
			if err := m.InboundApplication.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TopicEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TopicEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TopicEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutboundAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OutboundAddress == nil {
				m.OutboundAddress = &Address{}
			}
			if err := m.OutboundAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutboundApplication", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Application{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Application = &TopicEvent_OutboundApplication{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InboundApplication", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Application{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Application = &TopicEvent_InboundApplication{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NodeSignature) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodeSignature: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeSignature: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Node) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Node: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Node: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uuid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoNotify", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.NoNotify = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Edge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Edge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Edge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uuid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Sender == nil {
				m.Sender = &Node{}
			}
			if err := m.Sender.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Receiver == nil {
				m.Receiver = &Node{}
			}
			if err := m.Receiver.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastSeen", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastSeen == nil {
				m.LastSeen = &google_protobuf.Timestamp{}
			}
			if err := m.LastSeen.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= (EdgeState(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uuid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = &google_protobuf.Timestamp{}
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (EventType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Meta = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Record) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Record: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Record: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetNodesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetNodesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetNodesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetNodesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetNodesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetNodesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nodes = append(m.Nodes, &Node{})
			if err := m.Nodes[len(m.Nodes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetNodeBySignatureRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetNodeBySignatureRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetNodeBySignatureRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Signature == nil {
				m.Signature = &NodeSignature{}
			}
			if err := m.Signature.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetNodeBySignatureResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetNodeBySignatureResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetNodeBySignatureResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Node == nil {
				m.Node = &Node{}
			}
			if err := m.Node.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetEdgeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetEdgeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetEdgeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uuid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetEdgeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetEdgeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetEdgeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Edge", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Edge == nil {
				m.Edge = &Edge{}
			}
			if err := m.Edge.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetEdgeByNodesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetEdgeByNodesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetEdgeByNodesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Sender == nil {
				m.Sender = &NodeSignature{}
			}
			if err := m.Sender.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Receiver == nil {
				m.Receiver = &NodeSignature{}
			}
			if err := m.Receiver.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetEdgeByNodesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetEdgeByNodesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetEdgeByNodesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Edge", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Edge == nil {
				m.Edge = &Edge{}
			}
			if err := m.Edge.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetEdgesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetEdgesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetEdgesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetEdgesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetEdgesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetEdgesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Edges", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Edges = append(m.Edges, &Edge{})
			if err := m.Edges[len(m.Edges)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetEventLogForEdgeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetEventLogForEdgeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetEventLogForEdgeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uuid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetEventLogForEdgeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetEventLogForEdgeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetEventLogForEdgeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, &Event{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SetEdgeStateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SetEdgeStateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetEdgeStateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EdgeUuid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EdgeUuid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= (EdgeState(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Comment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Comment = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SetEdgeStateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SetEdgeStateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetEdgeStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetRecordsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetRecordsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetRecordsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetRecordsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetRecordsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetRecordsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Records", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Records = append(m.Records, &Record{})
			if err := m.Records[len(m.Records)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IsRegulatedRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IsRegulatedRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IsRegulatedRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Signature == nil {
				m.Signature = &NodeSignature{}
			}
			if err := m.Signature.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IsRegulatedResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IsRegulatedResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IsRegulatedResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Regulated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Regulated = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TestDetectEdgeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TestDetectEdgeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestDetectEdgeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Sender == nil {
				m.Sender = &NodeSignature{}
			}
			if err := m.Sender.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Receiver == nil {
				m.Receiver = &NodeSignature{}
			}
			if err := m.Receiver.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TestDetectEdgeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TestDetectEdgeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestDetectEdgeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Edge", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConnmon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Edge == nil {
				m.Edge = &Edge{}
			}
			if err := m.Edge.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnmon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnmon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConnmon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConnmon
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConnmon
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthConnmon
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConnmon
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipConnmon(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthConnmon = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConnmon   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("infra/connmon/connmon.proto", fileDescriptorConnmon) }

var fileDescriptorConnmon = []byte{
	// 1549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0x4f, 0x6f, 0x1b, 0xd5,
	0x16, 0xf7, 0x38, 0xb6, 0x63, 0x9f, 0x34, 0xf1, 0xe4, 0x3a, 0x75, 0x5c, 0x27, 0xcf, 0xc9, 0x9b,
	0xf6, 0x55, 0x69, 0xf4, 0x48, 0x50, 0x5a, 0x51, 0x50, 0x2b, 0x55, 0x8e, 0x3d, 0x71, 0xdc, 0x9a,
	0xd8, 0xbd, 0x9e, 0x04, 0x8a, 0x04, 0x96, 0x63, 0xdf, 0x5a, 0x96, 0xec, 0x19, 0x33, 0x33, 0x8e,
	0x14, 0xd8, 0xf0, 0x01, 0x58, 0xb0, 0x66, 0xc3, 0x16, 0xbe, 0x00, 0x5f, 0x80, 0x0d, 0x3b, 0xba,
	0x64, 0x49, 0xc3, 0x86, 0x05, 0x8b, 0x6e, 0xd8, 0x56, 0xe8, 0xde, 0xb9, 0xf3, 0x7f, 0x9c, 0x06,
	0x81, 0x58, 0x79, 0xe6, 0xfc, 0xf9, 0x9d, 0xdf, 0x39, 0xe7, 0xde, 0x73, 0xc6, 0xb0, 0x36, 0x54,
	0x9f, 0xeb, 0xdd, 0xdd, 0x9e, 0xa6, 0xaa, 0x63, 0x4d, 0xb5, 0x7f, 0x77, 0x26, 0xba, 0x66, 0x6a,
	0x68, 0x9e, 0xbf, 0x16, 0x37, 0x06, 0x9a, 0x36, 0x18, 0x91, 0x5d, 0x26, 0x3e, 0x9d, 0x3e, 0xdf,
	0x35, 0x87, 0x63, 0x62, 0x98, 0xdd, 0xf1, 0xc4, 0xb2, 0x94, 0xde, 0x82, 0xf9, 0x72, 0xbf, 0xaf,
	0x13, 0xc3, 0x40, 0x4b, 0x10, 0x1f, 0x4e, 0x0a, 0xc2, 0xa6, 0xb0, 0x95, 0xc1, 0xf1, 0xe1, 0x04,
	0x21, 0x48, 0x4c, 0x34, 0xdd, 0x2c, 0xc4, 0x37, 0x85, 0xad, 0x45, 0xcc, 0x9e, 0xa5, 0xd7, 0x02,
	0xac, 0x94, 0x07, 0x03, 0x9d, 0x0c, 0xba, 0xa6, 0xa6, 0x2b, 0xda, 0x64, 0xd8, 0x93, 0xcf, 0x88,
	0x6a, 0x22, 0x19, 0x44, 0x1a, 0x93, 0xf4, 0xcc, 0xa1, 0xa6, 0x76, 0x08, 0x95, 0x31, 0xa8, 0x85,
	0xbd, 0xc2, 0x8e, 0xcd, 0xad, 0xe2, 0x18, 0x30, 0x9f, 0xc3, 0x18, 0xce, 0xf6, 0xfc, 0x22, 0xf4,
	0x09, 0xdc, 0xd0, 0x89, 0xa1, 0x8d, 0xce, 0x48, 0xbf, 0x13, 0xc2, 0x8b, 0x33, 0xbc, 0x4d, 0x07,
	0x0f, 0x73, 0xcb, 0x30, 0xee, 0xaa, 0x1e, 0xad, 0x42, 0xef, 0xc0, 0x82, 0x49, 0x49, 0x73, 0xc4,
	0x39, 0x86, 0x98, 0x73, 0x10, 0xdd, 0x84, 0x0e, 0x63, 0x18, 0x4c, 0xe7, 0x6d, 0x7f, 0x1e, 0x92,
	0xcc, 0x43, 0xfa, 0x43, 0x80, 0x6c, 0x10, 0xf4, 0x00, 0xf2, 0x43, 0xf5, 0x54, 0x9b, 0xaa, 0xfd,
	0xce, 0x68, 0x68, 0x98, 0x44, 0xed, 0x74, 0xad, 0x92, 0xf2, 0x0a, 0x88, 0x0e, 0x3e, 0x2f, 0x35,
	0x5e, 0xe1, 0xf6, 0x0d, 0x66, 0x6e, 0x37, 0xa0, 0x0e, 0x4b, 0xda, 0xd4, 0xb4, 0x80, 0xa2, 0x33,
	0x76, 0x23, 0x37, 0xb9, 0xa1, 0x4d, 0x76, 0x51, 0xf3, 0x0a, 0xd0, 0x01, 0x2c, 0xda, 0x94, 0xbc,
	0x99, 0x6e, 0x44, 0x20, 0xd5, 0x55, 0x1f, 0xd0, 0xb5, 0xa1, 0xe7, 0xdd, 0xcd, 0xbb, 0x09, 0x0b,
	0xe5, 0xc9, 0x64, 0x34, 0xec, 0x75, 0xa9, 0x0f, 0xfa, 0x3f, 0x24, 0xcc, 0xf3, 0x09, 0x61, 0x09,
	0x2e, 0x79, 0x5a, 0xec, 0xb1, 0x51, 0xce, 0x27, 0x04, 0x33, 0x2b, 0x7a, 0x92, 0xd4, 0xee, 0x98,
	0xb0, 0x74, 0x32, 0x98, 0x3d, 0x4b, 0xdf, 0x08, 0xb0, 0x3a, 0x23, 0x1d, 0x54, 0x83, 0x15, 0xa7,
	0x10, 0x5d, 0x17, 0x91, 0x97, 0x73, 0x25, 0x2a, 0x1a, 0xce, 0xd9, 0x1e, 0x5e, 0x9a, 0x0f, 0x40,
	0x74, 0x81, 0x78, 0x4f, 0xe2, 0x33, 0x7a, 0x92, 0x75, 0x00, 0x2c, 0x81, 0xf4, 0xb5, 0x00, 0xf9,
	0xe8, 0x32, 0x21, 0x19, 0x72, 0x76, 0x79, 0xaf, 0xca, 0x0f, 0x71, 0x07, 0x2f, 0xbd, 0xf7, 0x20,
	0xeb, 0xc0, 0xbc, 0x81, 0xdd, 0x92, 0xed, 0xce, 0xc9, 0x7d, 0x27, 0xc0, 0xea, 0x8c, 0xf3, 0xff,
	0xcf, 0x95, 0x6f, 0x46, 0x9a, 0xf1, 0xbf, 0x96, 0xa6, 0xf4, 0xbb, 0x00, 0xe0, 0x19, 0x15, 0x51,
	0x4d, 0x11, 0xae, 0xd8, 0x14, 0x54, 0x9f, 0x91, 0xdb, 0x25, 0x9c, 0x0e, 0x63, 0xd1, 0xd9, 0xd5,
	0xa2, 0xb3, 0x9b, 0xbb, 0x14, 0x29, 0x22, 0xbf, 0xfd, 0x45, 0x58, 0xf0, 0x00, 0x48, 0xf7, 0x61,
	0xf1, 0x48, 0xeb, 0x93, 0xf6, 0x70, 0xa0, 0x76, 0xcd, 0xa9, 0xee, 0x1e, 0x7f, 0xc1, 0x3d, 0xfe,
	0x54, 0xc6, 0x2e, 0x10, 0xbf, 0x12, 0xf4, 0x59, 0xea, 0x40, 0x82, 0x3a, 0x52, 0xdd, 0x74, 0x3a,
	0xec, 0xdb, 0xf6, 0xf4, 0x39, 0xea, 0x0a, 0x39, 0x18, 0x73, 0x2e, 0x06, 0x5a, 0x83, 0x8c, 0xaa,
	0x75, 0x54, 0xcd, 0x1c, 0x3e, 0x3f, 0x2f, 0x24, 0x36, 0x85, 0xad, 0x34, 0x4e, 0xab, 0xda, 0x11,
	0x7b, 0x97, 0x7e, 0x12, 0x20, 0x21, 0xf7, 0x07, 0xd1, 0x11, 0xfe, 0x07, 0x29, 0x83, 0xa8, 0x7d,
	0xa2, 0xf3, 0x5a, 0x2e, 0x3a, 0x15, 0xa0, 0xa4, 0x30, 0x57, 0xa2, 0x3b, 0x90, 0xd6, 0x49, 0x8f,
	0x0c, 0xcf, 0x88, 0xce, 0x4b, 0x15, 0x30, 0x74, 0xd4, 0xe8, 0x3e, 0x64, 0x46, 0x5d, 0xc3, 0xec,
	0x18, 0x84, 0xa8, 0x8c, 0xcb, 0xc2, 0x5e, 0x71, 0xc7, 0x5a, 0x48, 0x3b, 0xf6, 0x42, 0xda, 0x51,
	0xec, 0x85, 0x84, 0xd3, 0xd4, 0xb8, 0x4d, 0x88, 0x8a, 0xb6, 0x20, 0x69, 0x98, 0x5d, 0x93, 0x14,
	0x92, 0x6c, 0xbc, 0x20, 0x27, 0x00, 0x25, 0xdf, 0xa6, 0x1a, 0x6c, 0x19, 0x48, 0x3f, 0x08, 0x90,
	0xb4, 0x4e, 0x55, 0x54, 0x4a, 0xef, 0x42, 0xc6, 0xd9, 0x77, 0x3c, 0xab, 0xcb, 0x08, 0xb8, 0xc6,
	0xa8, 0x08, 0xe9, 0xa9, 0x41, 0x74, 0x56, 0x72, 0xab, 0xbc, 0xce, 0x3b, 0xba, 0xcd, 0xcb, 0x9e,
	0x08, 0x92, 0xa3, 0x3c, 0x3c, 0x53, 0x6f, 0x05, 0x92, 0x67, 0xdd, 0xd1, 0xd4, 0xca, 0x22, 0x83,
	0xad, 0x17, 0xca, 0x73, 0x4c, 0xcc, 0x6e, 0x21, 0x65, 0xf1, 0xa4, 0xcf, 0xd2, 0xdb, 0x90, 0xc2,
	0xa4, 0xa7, 0xe9, 0xfd, 0x2b, 0x1f, 0x95, 0x65, 0xc8, 0xd6, 0x88, 0x49, 0xeb, 0x6d, 0x60, 0xf2,
	0xe9, 0x94, 0x18, 0xa6, 0x74, 0x1f, 0x44, 0x57, 0x64, 0x4c, 0x34, 0xd5, 0x20, 0xe8, 0x26, 0x24,
	0x55, 0x2a, 0x28, 0x08, 0x9b, 0x73, 0xe1, 0x4e, 0x59, 0x3a, 0xe9, 0x29, 0xdc, 0xe0, 0x8e, 0xfb,
	0xe7, 0xce, 0xa1, 0xe5, 0xa8, 0xe8, 0x1e, 0x64, 0x0c, 0x5b, 0xc6, 0x6f, 0x69, 0xde, 0x87, 0xe2,
	0x7a, 0xb8, 0x86, 0xd2, 0x23, 0x28, 0x46, 0x41, 0x72, 0x56, 0xff, 0x85, 0x04, 0x8d, 0xcc, 0xe1,
	0x02, 0xa4, 0x98, 0x4a, 0xba, 0x05, 0x4b, 0x35, 0x62, 0xd2, 0x76, 0xdb, 0x44, 0x22, 0xfa, 0x2b,
	0xdd, 0x63, 0x55, 0xb0, 0xac, 0x5c, 0x6c, 0xd2, 0x1f, 0x84, 0xb1, 0x99, 0x11, 0x53, 0x49, 0x9f,
	0xc3, 0x75, 0xee, 0xb5, 0x7f, 0xee, 0xad, 0x20, 0xda, 0x71, 0x6e, 0xc0, 0xe5, 0x89, 0xda, 0x57,
	0x61, 0xcf, 0x73, 0x15, 0xe2, 0x97, 0x7a, 0x38, 0x76, 0xd2, 0x03, 0xc8, 0x07, 0x83, 0x5f, 0x9d,
	0xf9, 0xb2, 0x93, 0x6f, 0xa0, 0xeb, 0x5c, 0xe4, 0x76, 0x9d, 0x9a, 0x87, 0xbb, 0xce, 0xa0, 0x2c,
	0x9d, 0xb4, 0xcb, 0xba, 0xce, 0xce, 0x6c, 0x43, 0x1b, 0x1c, 0x68, 0xfa, 0x9b, 0x8a, 0x5d, 0x65,
	0x3d, 0x0d, 0x39, 0xf0, 0x98, 0xb7, 0x21, 0xc5, 0x3e, 0x14, 0xec, 0xa0, 0x4b, 0xfe, 0x6b, 0x81,
	0xb9, 0x56, 0x3a, 0x83, 0x5c, 0xdb, 0xe2, 0x6b, 0xdd, 0x63, 0x1e, 0x70, 0x0d, 0x32, 0x94, 0x56,
	0xc7, 0x13, 0x35, 0x4d, 0x05, 0xc7, 0xf4, 0x1a, 0x3b, 0xe3, 0x20, 0xfe, 0x86, 0x71, 0x80, 0x0a,
	0x30, 0xdf, 0xd3, 0xc6, 0x63, 0xfb, 0x83, 0x27, 0x83, 0xed, 0x57, 0x29, 0x0f, 0x2b, 0xfe, 0xb8,
	0x16, 0x6f, 0x29, 0x07, 0xcb, 0x35, 0x62, 0x5a, 0xb7, 0xcf, 0x29, 0xea, 0x23, 0x40, 0x5e, 0x21,
	0x4f, 0xf1, 0x0e, 0xcc, 0xeb, 0x96, 0x88, 0xe7, 0x98, 0xf5, 0x7c, 0x89, 0x52, 0x39, 0xb6, 0xf5,
	0xd2, 0x63, 0x40, 0x75, 0x03, 0x93, 0xc1, 0x74, 0xd4, 0x35, 0x49, 0xff, 0xef, 0xdd, 0xa5, 0xbb,
	0x90, 0xf3, 0x61, 0x71, 0x36, 0xeb, 0x90, 0xd1, 0x6d, 0x21, 0x03, 0x4b, 0x63, 0x57, 0x40, 0xcf,
	0xb8, 0x42, 0x0c, 0xb3, 0x4a, 0x4c, 0xd2, 0xf3, 0x5d, 0xa3, 0x7f, 0xe9, 0x8c, 0x07, 0x83, 0x5f,
	0xf9, 0x8c, 0x6f, 0xbf, 0x16, 0x20, 0x1b, 0xf8, 0x8a, 0x44, 0xeb, 0x50, 0x28, 0xb7, 0x5a, 0x8d,
	0x7a, 0xa5, 0xac, 0xd4, 0x9b, 0x47, 0x1d, 0xe5, 0x59, 0x4b, 0xee, 0xd4, 0x8f, 0x4e, 0xca, 0x8d,
	0x7a, 0x55, 0x8c, 0x45, 0x6a, 0xdb, 0x32, 0x3e, 0xa9, 0x57, 0x64, 0x51, 0x40, 0xb7, 0x60, 0x33,
	0xa4, 0x3d, 0x2c, 0xb7, 0x70, 0xf3, 0xc3, 0x67, 0x9d, 0xfd, 0x72, 0xe5, 0x89, 0x7c, 0x54, 0x15,
	0xe3, 0x68, 0x1b, 0x6e, 0x87, 0xac, 0xde, 0x3f, 0x56, 0x94, 0x86, 0xfc, 0xac, 0xa3, 0xe0, 0xf2,
	0xc1, 0x41, 0xbd, 0xd2, 0xa9, 0xe1, 0xe6, 0x71, 0x4b, 0x9c, 0x43, 0x45, 0xc8, 0x87, 0x6c, 0x95,
	0x66, 0xab, 0x5e, 0x11, 0x13, 0x68, 0x13, 0xd6, 0x43, 0x3a, 0x8f, 0x40, 0x4c, 0xa2, 0x0d, 0x58,
	0x0b, 0x59, 0x1c, 0x1f, 0x61, 0xb9, 0xdd, 0x6c, 0x9c, 0xc8, 0x55, 0x31, 0xb5, 0xfd, 0x10, 0x32,
	0xce, 0x26, 0x41, 0x79, 0x40, 0xf2, 0x89, 0x7c, 0xa4, 0x04, 0x73, 0x5e, 0x85, 0x9c, 0x47, 0xde,
	0x96, 0x95, 0xb6, 0x52, 0x56, 0x64, 0x51, 0xd8, 0xfe, 0x5e, 0x80, 0x8c, 0x73, 0xca, 0x99, 0x7b,
	0xb5, 0x26, 0x77, 0x98, 0xd6, 0xe3, 0x5e, 0x84, 0xbc, 0x47, 0x4e, 0xc3, 0xd7, 0x8e, 0x1b, 0x65,
	0x45, 0xae, 0x8a, 0x02, 0x83, 0x76, 0x75, 0x55, 0x59, 0x91, 0x2b, 0x54, 0x11, 0x47, 0xd7, 0x61,
	0xd9, 0xa3, 0xc0, 0xf2, 0x49, 0x5d, 0xfe, 0x40, 0x9c, 0x43, 0x39, 0xc8, 0x7a, 0xc4, 0x87, 0xcd,
	0x46, 0x55, 0x4c, 0x04, 0x40, 0xb0, 0xfc, 0xd8, 0x02, 0x49, 0x06, 0x14, 0xe5, 0x4a, 0x45, 0x6e,
	0x51, 0x45, 0x6a, 0xef, 0xcb, 0x14, 0xa4, 0xcb, 0x6a, 0x77, 0x74, 0xfe, 0x19, 0xd1, 0xd1, 0x23,
	0x48, 0xdb, 0xbb, 0x0c, 0xb9, 0x7f, 0x2e, 0x02, 0x1b, 0xaf, 0x78, 0x23, 0x42, 0xc3, 0x0f, 0xda,
	0xc7, 0xec, 0x06, 0x07, 0x16, 0x10, 0x92, 0x82, 0x0e, 0xe1, 0x85, 0x57, 0xbc, 0x79, 0xa9, 0x0d,
	0x87, 0xb7, 0xf8, 0xb1, 0xa9, 0xeb, 0xe7, 0xe7, 0x9d, 0xcd, 0x7e, 0x7e, 0xfe, 0x11, 0xfd, 0x10,
	0xe6, 0xb9, 0x0c, 0xad, 0x06, 0xad, 0x6c, 0xf7, 0x42, 0x58, 0xc1, 0xbd, 0x9f, 0x3a, 0xdb, 0x91,
	0x2f, 0x11, 0x54, 0x0a, 0xda, 0xfa, 0x57, 0x5b, 0x71, 0x63, 0xa6, 0xde, 0x57, 0xb0, 0xc0, 0x74,
	0xf7, 0x17, 0x2c, 0x7a, 0x57, 0xf8, 0x0b, 0x36, 0x6b, 0x3d, 0x3c, 0x81, 0x6b, 0xde, 0xf1, 0x8b,
	0xd6, 0x1d, 0xa7, 0x88, 0x6d, 0x50, 0xfc, 0xcf, 0x0c, 0x2d, 0x07, 0x93, 0x01, 0xdc, 0xf1, 0x8c,
	0x8a, 0xde, 0xf8, 0xfe, 0x41, 0x5e, 0x5c, 0x8b, 0xd4, 0x71, 0x98, 0x43, 0x58, 0xf0, 0x0c, 0x56,
	0xe4, 0xda, 0x86, 0x47, 0x77, 0x71, 0x3d, 0x5a, 0xe9, 0xf6, 0xc3, 0x3f, 0xf0, 0x3c, 0xfd, 0x88,
	0x1c, 0xc3, 0x9e, 0x7e, 0x44, 0x4f, 0xca, 0x7d, 0xe5, 0xc5, 0xcb, 0x52, 0xec, 0xe7, 0x97, 0xa5,
	0xd8, 0xab, 0x97, 0x25, 0xe1, 0x8b, 0x8b, 0x92, 0xf0, 0xed, 0x45, 0x49, 0xf8, 0xf1, 0xa2, 0x24,
	0xbc, 0xb8, 0x28, 0x09, 0xbf, 0x5c, 0x94, 0x84, 0xdf, 0x2e, 0x4a, 0xb1, 0x57, 0x17, 0x25, 0xe1,
	0xab, 0x5f, 0x4b, 0x31, 0x58, 0xe8, 0x69, 0x63, 0x1b, 0x71, 0xff, 0x5a, 0xc5, 0x7a, 0x68, 0xd1,
	0x2f, 0xdc, 0x96, 0xf0, 0x51, 0x86, 0x2b, 0x26, 0xa7, 0xa7, 0x29, 0xf6, 0xd5, 0x7b, 0xf7, 0xcf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x94, 0x39, 0x7b, 0x3d, 0x12, 0x00, 0x00,
}
