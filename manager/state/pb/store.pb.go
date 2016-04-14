// Code generated by protoc-gen-gogo.
// source: store.proto
// DO NOT EDIT!

/*
	Package pb is a generated protocol buffer package.

	It is generated from these files:
		store.proto

	It has these top-level messages:
		StoreSnapshot
		ClusterSnapshot
		Snapshot
*/
package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import docker_cluster_api3 "github.com/docker/swarm-v2/api"
import docker_cluster_api4 "github.com/docker/swarm-v2/api"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

type Snapshot_Version int32

const (
	// V0 is the initial version of the StoreSnapshot message.
	Snapshot_V0 Snapshot_Version = 0
)

var Snapshot_Version_name = map[int32]string{
	0: "V0",
}
var Snapshot_Version_value = map[string]int32{
	"V0": 0,
}

func (x Snapshot_Version) String() string {
	return proto.EnumName(Snapshot_Version_name, int32(x))
}
func (Snapshot_Version) EnumDescriptor() ([]byte, []int) { return fileDescriptorStore, []int{2, 0} }

// StoreSnapshot is used to store snapshots of the store.
type StoreSnapshot struct {
	Nodes    []*docker_cluster_api3.Node    `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
	Services []*docker_cluster_api3.Service `protobuf:"bytes,2,rep,name=services" json:"services,omitempty"`
	Networks []*docker_cluster_api3.Network `protobuf:"bytes,3,rep,name=networks" json:"networks,omitempty"`
	Tasks    []*docker_cluster_api3.Task    `protobuf:"bytes,4,rep,name=tasks" json:"tasks,omitempty"`
	Volumes  []*docker_cluster_api3.Volume  `protobuf:"bytes,5,rep,name=volumes" json:"volumes,omitempty"`
}

func (m *StoreSnapshot) Reset()                    { *m = StoreSnapshot{} }
func (*StoreSnapshot) ProtoMessage()               {}
func (*StoreSnapshot) Descriptor() ([]byte, []int) { return fileDescriptorStore, []int{0} }

// ClusterSnapshot stores cluster membership information in snapshots.
type ClusterSnapshot struct {
	Members []*docker_cluster_api4.RaftNode `protobuf:"bytes,1,rep,name=members" json:"members,omitempty"`
	Removed []uint64                        `protobuf:"varint,2,rep,name=removed" json:"removed,omitempty"`
}

func (m *ClusterSnapshot) Reset()                    { *m = ClusterSnapshot{} }
func (*ClusterSnapshot) ProtoMessage()               {}
func (*ClusterSnapshot) Descriptor() ([]byte, []int) { return fileDescriptorStore, []int{1} }

type Snapshot struct {
	Version    Snapshot_Version `protobuf:"varint,1,opt,name=version,proto3,enum=pb.Snapshot_Version" json:"version,omitempty"`
	Membership ClusterSnapshot  `protobuf:"bytes,2,opt,name=membership" json:"membership"`
	Store      StoreSnapshot    `protobuf:"bytes,3,opt,name=store" json:"store"`
}

func (m *Snapshot) Reset()                    { *m = Snapshot{} }
func (*Snapshot) ProtoMessage()               {}
func (*Snapshot) Descriptor() ([]byte, []int) { return fileDescriptorStore, []int{2} }

func init() {
	proto.RegisterType((*StoreSnapshot)(nil), "pb.StoreSnapshot")
	proto.RegisterType((*ClusterSnapshot)(nil), "pb.ClusterSnapshot")
	proto.RegisterType((*Snapshot)(nil), "pb.Snapshot")
	proto.RegisterEnum("pb.Snapshot_Version", Snapshot_Version_name, Snapshot_Version_value)
}
func (this *StoreSnapshot) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&pb.StoreSnapshot{")
	if this.Nodes != nil {
		s = append(s, "Nodes: "+fmt.Sprintf("%#v", this.Nodes)+",\n")
	}
	if this.Services != nil {
		s = append(s, "Services: "+fmt.Sprintf("%#v", this.Services)+",\n")
	}
	if this.Networks != nil {
		s = append(s, "Networks: "+fmt.Sprintf("%#v", this.Networks)+",\n")
	}
	if this.Tasks != nil {
		s = append(s, "Tasks: "+fmt.Sprintf("%#v", this.Tasks)+",\n")
	}
	if this.Volumes != nil {
		s = append(s, "Volumes: "+fmt.Sprintf("%#v", this.Volumes)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ClusterSnapshot) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pb.ClusterSnapshot{")
	if this.Members != nil {
		s = append(s, "Members: "+fmt.Sprintf("%#v", this.Members)+",\n")
	}
	s = append(s, "Removed: "+fmt.Sprintf("%#v", this.Removed)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Snapshot) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&pb.Snapshot{")
	s = append(s, "Version: "+fmt.Sprintf("%#v", this.Version)+",\n")
	s = append(s, "Membership: "+strings.Replace(this.Membership.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "Store: "+strings.Replace(this.Store.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringStore(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringStore(e map[int32]github_com_gogo_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
func (m *StoreSnapshot) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *StoreSnapshot) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Nodes) > 0 {
		for _, msg := range m.Nodes {
			data[i] = 0xa
			i++
			i = encodeVarintStore(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Services) > 0 {
		for _, msg := range m.Services {
			data[i] = 0x12
			i++
			i = encodeVarintStore(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Networks) > 0 {
		for _, msg := range m.Networks {
			data[i] = 0x1a
			i++
			i = encodeVarintStore(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Tasks) > 0 {
		for _, msg := range m.Tasks {
			data[i] = 0x22
			i++
			i = encodeVarintStore(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Volumes) > 0 {
		for _, msg := range m.Volumes {
			data[i] = 0x2a
			i++
			i = encodeVarintStore(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ClusterSnapshot) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ClusterSnapshot) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Members) > 0 {
		for _, msg := range m.Members {
			data[i] = 0xa
			i++
			i = encodeVarintStore(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Removed) > 0 {
		for _, num := range m.Removed {
			data[i] = 0x10
			i++
			i = encodeVarintStore(data, i, uint64(num))
		}
	}
	return i, nil
}

func (m *Snapshot) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Snapshot) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintStore(data, i, uint64(m.Version))
	}
	data[i] = 0x12
	i++
	i = encodeVarintStore(data, i, uint64(m.Membership.Size()))
	n1, err := m.Membership.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	data[i] = 0x1a
	i++
	i = encodeVarintStore(data, i, uint64(m.Store.Size()))
	n2, err := m.Store.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func encodeFixed64Store(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Store(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintStore(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *StoreSnapshot) Size() (n int) {
	var l int
	_ = l
	if len(m.Nodes) > 0 {
		for _, e := range m.Nodes {
			l = e.Size()
			n += 1 + l + sovStore(uint64(l))
		}
	}
	if len(m.Services) > 0 {
		for _, e := range m.Services {
			l = e.Size()
			n += 1 + l + sovStore(uint64(l))
		}
	}
	if len(m.Networks) > 0 {
		for _, e := range m.Networks {
			l = e.Size()
			n += 1 + l + sovStore(uint64(l))
		}
	}
	if len(m.Tasks) > 0 {
		for _, e := range m.Tasks {
			l = e.Size()
			n += 1 + l + sovStore(uint64(l))
		}
	}
	if len(m.Volumes) > 0 {
		for _, e := range m.Volumes {
			l = e.Size()
			n += 1 + l + sovStore(uint64(l))
		}
	}
	return n
}

func (m *ClusterSnapshot) Size() (n int) {
	var l int
	_ = l
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovStore(uint64(l))
		}
	}
	if len(m.Removed) > 0 {
		for _, e := range m.Removed {
			n += 1 + sovStore(uint64(e))
		}
	}
	return n
}

func (m *Snapshot) Size() (n int) {
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovStore(uint64(m.Version))
	}
	l = m.Membership.Size()
	n += 1 + l + sovStore(uint64(l))
	l = m.Store.Size()
	n += 1 + l + sovStore(uint64(l))
	return n
}

func sovStore(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStore(x uint64) (n int) {
	return sovStore(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *StoreSnapshot) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StoreSnapshot{`,
		`Nodes:` + strings.Replace(fmt.Sprintf("%v", this.Nodes), "Node", "docker_cluster_api3.Node", 1) + `,`,
		`Services:` + strings.Replace(fmt.Sprintf("%v", this.Services), "Service", "docker_cluster_api3.Service", 1) + `,`,
		`Networks:` + strings.Replace(fmt.Sprintf("%v", this.Networks), "Network", "docker_cluster_api3.Network", 1) + `,`,
		`Tasks:` + strings.Replace(fmt.Sprintf("%v", this.Tasks), "Task", "docker_cluster_api3.Task", 1) + `,`,
		`Volumes:` + strings.Replace(fmt.Sprintf("%v", this.Volumes), "Volume", "docker_cluster_api3.Volume", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ClusterSnapshot) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ClusterSnapshot{`,
		`Members:` + strings.Replace(fmt.Sprintf("%v", this.Members), "RaftNode", "docker_cluster_api4.RaftNode", 1) + `,`,
		`Removed:` + fmt.Sprintf("%v", this.Removed) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Snapshot) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Snapshot{`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`Membership:` + strings.Replace(strings.Replace(this.Membership.String(), "ClusterSnapshot", "ClusterSnapshot", 1), `&`, ``, 1) + `,`,
		`Store:` + strings.Replace(strings.Replace(this.Store.String(), "StoreSnapshot", "StoreSnapshot", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringStore(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *StoreSnapshot) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StoreSnapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreSnapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nodes = append(m.Nodes, &docker_cluster_api3.Node{})
			if err := m.Nodes[len(m.Nodes)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Services", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Services = append(m.Services, &docker_cluster_api3.Service{})
			if err := m.Services[len(m.Services)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Networks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Networks = append(m.Networks, &docker_cluster_api3.Network{})
			if err := m.Networks[len(m.Networks)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tasks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tasks = append(m.Tasks, &docker_cluster_api3.Task{})
			if err := m.Tasks[len(m.Tasks)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Volumes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Volumes = append(m.Volumes, &docker_cluster_api3.Volume{})
			if err := m.Volumes[len(m.Volumes)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStore(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStore
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
func (m *ClusterSnapshot) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClusterSnapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterSnapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, &docker_cluster_api4.RaftNode{})
			if err := m.Members[len(m.Members)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Removed", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Removed = append(m.Removed, v)
		default:
			iNdEx = preIndex
			skippy, err := skipStore(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStore
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
func (m *Snapshot) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStore
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Snapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Snapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Version |= (Snapshot_Version(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Membership", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Membership.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Store", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStore
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStore
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Store.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStore(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStore
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
func skipStore(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStore
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowStore
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowStore
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthStore
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStore
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipStore(data[start:])
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
	ErrInvalidLengthStore = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStore   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorStore = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0xbf, 0x8e, 0xda, 0x40,
	0x10, 0xc6, 0xb1, 0xf9, 0x63, 0xb4, 0x28, 0x7f, 0xd8, 0x50, 0x58, 0x04, 0x91, 0xc8, 0x15, 0x45,
	0x58, 0x47, 0x24, 0x4a, 0x94, 0x96, 0xf4, 0x14, 0x4b, 0x44, 0x6f, 0x9b, 0x8d, 0x71, 0xc0, 0x5e,
	0x6b, 0x77, 0x31, 0x6d, 0xde, 0x22, 0x8f, 0x92, 0x57, 0xa0, 0xbc, 0xf2, 0xaa, 0xd3, 0x71, 0x4f,
	0x70, 0x8f, 0x70, 0xe3, 0x5d, 0x8c, 0xd0, 0x09, 0xa4, 0x2b, 0x46, 0xf6, 0x6a, 0x7e, 0xdf, 0xcc,
	0xa7, 0x4f, 0x83, 0x3a, 0x52, 0x71, 0xc1, 0x48, 0x2e, 0xb8, 0xe2, 0xd8, 0xce, 0xc3, 0x7e, 0x2f,
	0xe6, 0x31, 0xd7, 0x4f, 0xbf, 0xfc, 0x33, 0x9d, 0xfe, 0xa7, 0x38, 0x51, 0xab, 0x6d, 0x48, 0x22,
	0x9e, 0xfa, 0x4b, 0x1e, 0xad, 0x99, 0xf0, 0xe5, 0x2e, 0x10, 0xe9, 0xb8, 0x98, 0xf8, 0x41, 0x9e,
	0xf8, 0x3c, 0xfc, 0xc3, 0x22, 0x25, 0x5f, 0x48, 0xa7, 0x41, 0x16, 0xc4, 0x4c, 0x18, 0xda, 0xfb,
	0x67, 0xa3, 0x57, 0xf3, 0xd2, 0xc5, 0x3c, 0x0b, 0x72, 0xb9, 0xe2, 0x0a, 0x13, 0xd4, 0xcc, 0xf8,
	0x92, 0x49, 0xd7, 0xfa, 0x58, 0x1f, 0x75, 0x26, 0x2e, 0x31, 0x43, 0x48, 0xb4, 0xd9, 0x4a, 0x05,
	0x5f, 0x98, 0x41, 0x66, 0x00, 0x50, 0x83, 0xe1, 0xef, 0xa8, 0x2d, 0x99, 0x28, 0x92, 0x08, 0x24,
	0xb6, 0x96, 0xbc, 0xbf, 0x24, 0x99, 0x1b, 0x86, 0x9e, 0xe0, 0x52, 0x98, 0x31, 0xb5, 0xe3, 0x62,
	0x2d, 0xdd, 0xfa, 0x75, 0xe1, 0xcc, 0x30, 0xf4, 0x04, 0x97, 0x0e, 0x55, 0x20, 0x41, 0xd5, 0xb8,
	0xee, 0xf0, 0x17, 0x00, 0xd4, 0x60, 0xf8, 0x2b, 0x72, 0x0a, 0xbe, 0xd9, 0xa6, 0x60, 0xb0, 0xa9,
	0x15, 0xfd, 0x4b, 0x8a, 0x85, 0x46, 0x68, 0x85, 0x7a, 0x11, 0x7a, 0xf3, 0xd3, 0xb4, 0x4f, 0xd1,
	0x7c, 0x43, 0x4e, 0xca, 0xd2, 0x90, 0x89, 0x2a, 0x9c, 0xc1, 0xa5, 0x41, 0x34, 0xf8, 0xad, 0x74,
	0x40, 0x15, 0x8c, 0x5d, 0xe4, 0x08, 0x96, 0xf2, 0x82, 0x2d, 0x75, 0x42, 0x0d, 0x5a, 0x3d, 0xbd,
	0xff, 0x16, 0x6a, 0x9f, 0x25, 0xef, 0x14, 0x80, 0x27, 0x3c, 0x83, 0xf1, 0xd6, 0xe8, 0xf5, 0xa4,
	0x47, 0xf2, 0x90, 0x54, 0x6d, 0xb2, 0x30, 0x3d, 0x5a, 0x41, 0xf8, 0x07, 0x42, 0xc7, 0x0d, 0xab,
	0x24, 0x87, 0xc9, 0x16, 0x38, 0x7a, 0x57, 0x4a, 0x9e, 0xf9, 0x9e, 0x36, 0xf6, 0x77, 0x1f, 0x6a,
	0xf4, 0x0c, 0xc6, 0x63, 0xd4, 0xd4, 0xb7, 0x07, 0xc1, 0x97, 0xaa, 0xae, 0x5e, 0x74, 0x7e, 0x06,
	0x47, 0x8d, 0xa1, 0xbc, 0x2e, 0x72, 0x8e, 0xdb, 0x71, 0x0b, 0xd9, 0x8b, 0xcf, 0x6f, 0x6b, 0xd3,
	0xc1, 0xfe, 0x30, 0xac, 0xdd, 0x42, 0x3d, 0x1e, 0x86, 0xd6, 0xdf, 0x87, 0xa1, 0xb5, 0x87, 0xba,
	0x81, 0xba, 0x87, 0x0a, 0x5b, 0xfa, 0xba, 0xbe, 0x3c, 0x05, 0x00, 0x00, 0xff, 0xff, 0x28, 0x36,
	0xeb, 0xf3, 0xe2, 0x02, 0x00, 0x00,
}
