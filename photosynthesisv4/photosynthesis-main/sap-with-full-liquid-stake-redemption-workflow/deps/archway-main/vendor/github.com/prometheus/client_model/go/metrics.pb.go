// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metrics.proto

package io_prometheus_client

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type MetricType int32

const (
	MetricType_COUNTER   MetricType = 0
	MetricType_GAUGE     MetricType = 1
	MetricType_SUMMARY   MetricType = 2
	MetricType_UNTYPED   MetricType = 3
	MetricType_HISTOGRAM MetricType = 4
)

var MetricType_name = map[int32]string{
	0: "COUNTER",
	1: "GAUGE",
	2: "SUMMARY",
	3: "UNTYPED",
	4: "HISTOGRAM",
}

var MetricType_value = map[string]int32{
	"COUNTER":   0,
	"GAUGE":     1,
	"SUMMARY":   2,
	"UNTYPED":   3,
	"HISTOGRAM": 4,
}

func (x MetricType) Enum() *MetricType {
	p := new(MetricType)
	*p = x
	return p
}

func (x MetricType) String() string {
	return proto.EnumName(MetricType_name, int32(x))
}

func (x *MetricType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MetricType_value, data, "MetricType")
	if err != nil {
		return err
	}
	*x = MetricType(value)
	return nil
}

func (MetricType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{0}
}

type LabelPair struct {
	Name                 *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value                *string  `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LabelPair) Reset()         { *m = LabelPair{} }
func (m *LabelPair) String() string { return proto.CompactTextString(m) }
func (*LabelPair) ProtoMessage()    {}
func (*LabelPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{0}
}

func (m *LabelPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelPair.Unmarshal(m, b)
}
func (m *LabelPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelPair.Marshal(b, m, deterministic)
}
func (m *LabelPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelPair.Merge(m, src)
}
func (m *LabelPair) XXX_Size() int {
	return xxx_messageInfo_LabelPair.Size(m)
}
func (m *LabelPair) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelPair.DiscardUnknown(m)
}

var xxx_messageInfo_LabelPair proto.InternalMessageInfo

func (m *LabelPair) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *LabelPair) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type Gauge struct {
	Value                *float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Gauge) Reset()         { *m = Gauge{} }
func (m *Gauge) String() string { return proto.CompactTextString(m) }
func (*Gauge) ProtoMessage()    {}
func (*Gauge) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{1}
}

func (m *Gauge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gauge.Unmarshal(m, b)
}
func (m *Gauge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gauge.Marshal(b, m, deterministic)
}
func (m *Gauge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gauge.Merge(m, src)
}
func (m *Gauge) XXX_Size() int {
	return xxx_messageInfo_Gauge.Size(m)
}
func (m *Gauge) XXX_DiscardUnknown() {
	xxx_messageInfo_Gauge.DiscardUnknown(m)
}

var xxx_messageInfo_Gauge proto.InternalMessageInfo

func (m *Gauge) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type Counter struct {
	Value                *float64  `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	Exemplar             *Exemplar `protobuf:"bytes,2,opt,name=exemplar" json:"exemplar,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Counter) Reset()         { *m = Counter{} }
func (m *Counter) String() string { return proto.CompactTextString(m) }
func (*Counter) ProtoMessage()    {}
func (*Counter) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{2}
}

func (m *Counter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Counter.Unmarshal(m, b)
}
func (m *Counter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Counter.Marshal(b, m, deterministic)
}
func (m *Counter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Counter.Merge(m, src)
}
func (m *Counter) XXX_Size() int {
	return xxx_messageInfo_Counter.Size(m)
}
func (m *Counter) XXX_DiscardUnknown() {
	xxx_messageInfo_Counter.DiscardUnknown(m)
}

var xxx_messageInfo_Counter proto.InternalMessageInfo

func (m *Counter) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *Counter) GetExemplar() *Exemplar {
	if m != nil {
		return m.Exemplar
	}
	return nil
}

type Quantile struct {
	Quantile             *float64 `protobuf:"fixed64,1,opt,name=quantile" json:"quantile,omitempty"`
	Value                *float64 `protobuf:"fixed64,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Quantile) Reset()         { *m = Quantile{} }
func (m *Quantile) String() string { return proto.CompactTextString(m) }
func (*Quantile) ProtoMessage()    {}
func (*Quantile) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{3}
}

func (m *Quantile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Quantile.Unmarshal(m, b)
}
func (m *Quantile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Quantile.Marshal(b, m, deterministic)
}
func (m *Quantile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Quantile.Merge(m, src)
}
func (m *Quantile) XXX_Size() int {
	return xxx_messageInfo_Quantile.Size(m)
}
func (m *Quantile) XXX_DiscardUnknown() {
	xxx_messageInfo_Quantile.DiscardUnknown(m)
}

var xxx_messageInfo_Quantile proto.InternalMessageInfo

func (m *Quantile) GetQuantile() float64 {
	if m != nil && m.Quantile != nil {
		return *m.Quantile
	}
	return 0
}

func (m *Quantile) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type Summary struct {
	SampleCount          *uint64     `protobuf:"varint,1,opt,name=sample_count,json=sampleCount" json:"sample_count,omitempty"`
	SampleSum            *float64    `protobuf:"fixed64,2,opt,name=sample_sum,json=sampleSum" json:"sample_sum,omitempty"`
	Quantile             []*Quantile `protobuf:"bytes,3,rep,name=quantile" json:"quantile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Summary) Reset()         { *m = Summary{} }
func (m *Summary) String() string { return proto.CompactTextString(m) }
func (*Summary) ProtoMessage()    {}
func (*Summary) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{4}
}

func (m *Summary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Summary.Unmarshal(m, b)
}
func (m *Summary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Summary.Marshal(b, m, deterministic)
}
func (m *Summary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Summary.Merge(m, src)
}
func (m *Summary) XXX_Size() int {
	return xxx_messageInfo_Summary.Size(m)
}
func (m *Summary) XXX_DiscardUnknown() {
	xxx_messageInfo_Summary.DiscardUnknown(m)
}

var xxx_messageInfo_Summary proto.InternalMessageInfo

func (m *Summary) GetSampleCount() uint64 {
	if m != nil && m.SampleCount != nil {
		return *m.SampleCount
	}
	return 0
}

func (m *Summary) GetSampleSum() float64 {
	if m != nil && m.SampleSum != nil {
		return *m.SampleSum
	}
	return 0
}

func (m *Summary) GetQuantile() []*Quantile {
	if m != nil {
		return m.Quantile
	}
	return nil
}

type Untyped struct {
	Value                *float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Untyped) Reset()         { *m = Untyped{} }
func (m *Untyped) String() string { return proto.CompactTextString(m) }
func (*Untyped) ProtoMessage()    {}
func (*Untyped) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{5}
}

func (m *Untyped) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Untyped.Unmarshal(m, b)
}
func (m *Untyped) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Untyped.Marshal(b, m, deterministic)
}
func (m *Untyped) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Untyped.Merge(m, src)
}
func (m *Untyped) XXX_Size() int {
	return xxx_messageInfo_Untyped.Size(m)
}
func (m *Untyped) XXX_DiscardUnknown() {
	xxx_messageInfo_Untyped.DiscardUnknown(m)
}

var xxx_messageInfo_Untyped proto.InternalMessageInfo

func (m *Untyped) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type Histogram struct {
	SampleCount          *uint64   `protobuf:"varint,1,opt,name=sample_count,json=sampleCount" json:"sample_count,omitempty"`
	SampleSum            *float64  `protobuf:"fixed64,2,opt,name=sample_sum,json=sampleSum" json:"sample_sum,omitempty"`
	Bucket               []*Bucket `protobuf:"bytes,3,rep,name=bucket" json:"bucket,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Histogram) Reset()         { *m = Histogram{} }
func (m *Histogram) String() string { return proto.CompactTextString(m) }
func (*Histogram) ProtoMessage()    {}
func (*Histogram) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{6}
}

func (m *Histogram) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Histogram.Unmarshal(m, b)
}
func (m *Histogram) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Histogram.Marshal(b, m, deterministic)
}
func (m *Histogram) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Histogram.Merge(m, src)
}
func (m *Histogram) XXX_Size() int {
	return xxx_messageInfo_Histogram.Size(m)
}
func (m *Histogram) XXX_DiscardUnknown() {
	xxx_messageInfo_Histogram.DiscardUnknown(m)
}

var xxx_messageInfo_Histogram proto.InternalMessageInfo

func (m *Histogram) GetSampleCount() uint64 {
	if m != nil && m.SampleCount != nil {
		return *m.SampleCount
	}
	return 0
}

func (m *Histogram) GetSampleSum() float64 {
	if m != nil && m.SampleSum != nil {
		return *m.SampleSum
	}
	return 0
}

func (m *Histogram) GetBucket() []*Bucket {
	if m != nil {
		return m.Bucket
	}
	return nil
}

type Bucket struct {
	CumulativeCount      *uint64   `protobuf:"varint,1,opt,name=cumulative_count,json=cumulativeCount" json:"cumulative_count,omitempty"`
	UpperBound           *float64  `protobuf:"fixed64,2,opt,name=upper_bound,json=upperBound" json:"upper_bound,omitempty"`
	Exemplar             *Exemplar `protobuf:"bytes,3,opt,name=exemplar" json:"exemplar,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Bucket) Reset()         { *m = Bucket{} }
func (m *Bucket) String() string { return proto.CompactTextString(m) }
func (*Bucket) ProtoMessage()    {}
func (*Bucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{7}
}

func (m *Bucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bucket.Unmarshal(m, b)
}
func (m *Bucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bucket.Marshal(b, m, deterministic)
}
func (m *Bucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bucket.Merge(m, src)
}
func (m *Bucket) XXX_Size() int {
	return xxx_messageInfo_Bucket.Size(m)
}
func (m *Bucket) XXX_DiscardUnknown() {
	xxx_messageInfo_Bucket.DiscardUnknown(m)
}

var xxx_messageInfo_Bucket proto.InternalMessageInfo

func (m *Bucket) GetCumulativeCount() uint64 {
	if m != nil && m.CumulativeCount != nil {
		return *m.CumulativeCount
	}
	return 0
}

func (m *Bucket) GetUpperBound() float64 {
	if m != nil && m.UpperBound != nil {
		return *m.UpperBound
	}
	return 0
}

func (m *Bucket) GetExemplar() *Exemplar {
	if m != nil {
		return m.Exemplar
	}
	return nil
}

type Exemplar struct {
	Label                []*LabelPair         `protobuf:"bytes,1,rep,name=label" json:"label,omitempty"`
	Value                *float64             `protobuf:"fixed64,2,opt,name=value" json:"value,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Exemplar) Reset()         { *m = Exemplar{} }
func (m *Exemplar) String() string { return proto.CompactTextString(m) }
func (*Exemplar) ProtoMessage()    {}
func (*Exemplar) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{8}
}

func (m *Exemplar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Exemplar.Unmarshal(m, b)
}
func (m *Exemplar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Exemplar.Marshal(b, m, deterministic)
}
func (m *Exemplar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Exemplar.Merge(m, src)
}
func (m *Exemplar) XXX_Size() int {
	return xxx_messageInfo_Exemplar.Size(m)
}
func (m *Exemplar) XXX_DiscardUnknown() {
	xxx_messageInfo_Exemplar.DiscardUnknown(m)
}

var xxx_messageInfo_Exemplar proto.InternalMessageInfo

func (m *Exemplar) GetLabel() []*LabelPair {
	if m != nil {
		return m.Label
	}
	return nil
}

func (m *Exemplar) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *Exemplar) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type Metric struct {
	Label                []*LabelPair `protobuf:"bytes,1,rep,name=label" json:"label,omitempty"`
	Gauge                *Gauge       `protobuf:"bytes,2,opt,name=gauge" json:"gauge,omitempty"`
	Counter              *Counter     `protobuf:"bytes,3,opt,name=counter" json:"counter,omitempty"`
	Summary              *Summary     `protobuf:"bytes,4,opt,name=summary" json:"summary,omitempty"`
	Untyped              *Untyped     `protobuf:"bytes,5,opt,name=untyped" json:"untyped,omitempty"`
	Histogram            *Histogram   `protobuf:"bytes,7,opt,name=histogram" json:"histogram,omitempty"`
	TimestampMs          *int64       `protobuf:"varint,6,opt,name=timestamp_ms,json=timestampMs" json:"timestamp_ms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{9}
}

func (m *Metric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metric.Unmarshal(m, b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metric.Marshal(b, m, deterministic)
}
func (m *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(m, src)
}
func (m *Metric) XXX_Size() int {
	return xxx_messageInfo_Metric.Size(m)
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

func (m *Metric) GetLabel() []*LabelPair {
	if m != nil {
		return m.Label
	}
	return nil
}

func (m *Metric) GetGauge() *Gauge {
	if m != nil {
		return m.Gauge
	}
	return nil
}

func (m *Metric) GetCounter() *Counter {
	if m != nil {
		return m.Counter
	}
	return nil
}

func (m *Metric) GetSummary() *Summary {
	if m != nil {
		return m.Summary
	}
	return nil
}

func (m *Metric) GetUntyped() *Untyped {
	if m != nil {
		return m.Untyped
	}
	return nil
}

func (m *Metric) GetHistogram() *Histogram {
	if m != nil {
		return m.Histogram
	}
	return nil
}

func (m *Metric) GetTimestampMs() int64 {
	if m != nil && m.TimestampMs != nil {
		return *m.TimestampMs
	}
	return 0
}

type MetricFamily struct {
	Name                 *string     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Help                 *string     `protobuf:"bytes,2,opt,name=help" json:"help,omitempty"`
	Type                 *MetricType `protobuf:"varint,3,opt,name=type,enum=io.prometheus.client.MetricType" json:"type,omitempty"`
	Metric               []*Metric   `protobuf:"bytes,4,rep,name=metric" json:"metric,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *MetricFamily) Reset()         { *m = MetricFamily{} }
func (m *MetricFamily) String() string { return proto.CompactTextString(m) }
func (*MetricFamily) ProtoMessage()    {}
func (*MetricFamily) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{10}
}

func (m *MetricFamily) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricFamily.Unmarshal(m, b)
}
func (m *MetricFamily) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricFamily.Marshal(b, m, deterministic)
}
func (m *MetricFamily) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricFamily.Merge(m, src)
}
func (m *MetricFamily) XXX_Size() int {
	return xxx_messageInfo_MetricFamily.Size(m)
}
func (m *MetricFamily) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricFamily.DiscardUnknown(m)
}

var xxx_messageInfo_MetricFamily proto.InternalMessageInfo

func (m *MetricFamily) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *MetricFamily) GetHelp() string {
	if m != nil && m.Help != nil {
		return *m.Help
	}
	return ""
}

func (m *MetricFamily) GetType() MetricType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return MetricType_COUNTER
}

func (m *MetricFamily) GetMetric() []*Metric {
	if m != nil {
		return m.Metric
	}
	return nil
}

func init() {
	proto.RegisterEnum("io.prometheus.client.MetricType", MetricType_name, MetricType_value)
	proto.RegisterType((*LabelPair)(nil), "io.prometheus.client.LabelPair")
	proto.RegisterType((*Gauge)(nil), "io.prometheus.client.Gauge")
	proto.RegisterType((*Counter)(nil), "io.prometheus.client.Counter")
	proto.RegisterType((*Quantile)(nil), "io.prometheus.client.Quantile")
	proto.RegisterType((*Summary)(nil), "io.prometheus.client.Summary")
	proto.RegisterType((*Untyped)(nil), "io.prometheus.client.Untyped")
	proto.RegisterType((*Histogram)(nil), "io.prometheus.client.Histogram")
	proto.RegisterType((*Bucket)(nil), "io.prometheus.client.Bucket")
	proto.RegisterType((*Exemplar)(nil), "io.prometheus.client.Exemplar")
	proto.RegisterType((*Metric)(nil), "io.prometheus.client.Metric")
	proto.RegisterType((*MetricFamily)(nil), "io.prometheus.client.MetricFamily")
}

func init() { proto.RegisterFile("metrics.proto", fileDescriptor_6039342a2ba47b72) }

var fileDescriptor_6039342a2ba47b72 = []byte{
	// 665 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xfd, 0xdc, 0x38, 0x3f, 0xbe, 0x69, 0x3f, 0xa2, 0x51, 0x17, 0x56, 0xa1, 0x24, 0x78, 0x55,
	0x58, 0x38, 0xa2, 0x6a, 0x05, 0x2a, 0xb0, 0x68, 0x4b, 0x48, 0x91, 0x48, 0x5b, 0x26, 0xc9, 0xa2,
	0xb0, 0x88, 0x1c, 0x77, 0x70, 0x2c, 0x3c, 0xb1, 0xb1, 0x67, 0x2a, 0xb2, 0x66, 0xc1, 0x16, 0x5e,
	0x81, 0x17, 0x05, 0xcd, 0x8f, 0x6d, 0x2a, 0xb9, 0x95, 0x40, 0xec, 0x66, 0xee, 0x3d, 0xe7, 0xfa,
	0xcc, 0xf8, 0x9c, 0x81, 0x0d, 0x4a, 0x58, 0x1a, 0xfa, 0x99, 0x9b, 0xa4, 0x31, 0x8b, 0xd1, 0x66,
	0x18, 0x8b, 0x15, 0x25, 0x6c, 0x41, 0x78, 0xe6, 0xfa, 0x51, 0x48, 0x96, 0x6c, 0xab, 0x1b, 0xc4,
	0x71, 0x10, 0x91, 0xbe, 0xc4, 0xcc, 0xf9, 0x87, 0x3e, 0x0b, 0x29, 0xc9, 0x98, 0x47, 0x13, 0x45,
	0x73, 0xf6, 0xc1, 0x7a, 0xe3, 0xcd, 0x49, 0x74, 0xee, 0x85, 0x29, 0x42, 0x60, 0x2e, 0x3d, 0x4a,
	0x6c, 0xa3, 0x67, 0xec, 0x58, 0x58, 0xae, 0xd1, 0x26, 0xd4, 0xaf, 0xbc, 0x88, 0x13, 0x7b, 0x4d,
	0x16, 0xd5, 0xc6, 0xd9, 0x86, 0xfa, 0xd0, 0xe3, 0xc1, 0x6f, 0x6d, 0xc1, 0x31, 0xf2, 0xf6, 0x7b,
	0x68, 0x1e, 0xc7, 0x7c, 0xc9, 0x48, 0x5a, 0x0d, 0x40, 0x07, 0xd0, 0x22, 0x9f, 0x09, 0x4d, 0x22,
	0x2f, 0x95, 0x83, 0xdb, 0xbb, 0xf7, 0xdd, 0xaa, 0x03, 0xb8, 0x03, 0x8d, 0xc2, 0x05, 0xde, 0x79,
	0x0e, 0xad, 0xb7, 0xdc, 0x5b, 0xb2, 0x30, 0x22, 0x68, 0x0b, 0x5a, 0x9f, 0xf4, 0x5a, 0x7f, 0xa0,
	0xd8, 0x5f, 0x57, 0x5e, 0x48, 0xfb, 0x6a, 0x40, 0x73, 0xcc, 0x29, 0xf5, 0xd2, 0x15, 0x7a, 0x00,
	0xeb, 0x99, 0x47, 0x93, 0x88, 0xcc, 0x7c, 0xa1, 0x56, 0x4e, 0x30, 0x71, 0x5b, 0xd5, 0xe4, 0x01,
	0xd0, 0x36, 0x80, 0x86, 0x64, 0x9c, 0xea, 0x49, 0x96, 0xaa, 0x8c, 0x39, 0x15, 0xe7, 0x28, 0xbe,
	0x5f, 0xeb, 0xd5, 0x6e, 0x3e, 0x47, 0xae, 0xb8, 0xd4, 0xe7, 0x74, 0xa1, 0x39, 0x5d, 0xb2, 0x55,
	0x42, 0x2e, 0x6f, 0xb8, 0xc5, 0x2f, 0x06, 0x58, 0x27, 0x61, 0xc6, 0xe2, 0x20, 0xf5, 0xe8, 0x3f,
	0x10, 0xbb, 0x07, 0x8d, 0x39, 0xf7, 0x3f, 0x12, 0xa6, 0xa5, 0xde, 0xab, 0x96, 0x7a, 0x24, 0x31,
	0x58, 0x63, 0x9d, 0x6f, 0x06, 0x34, 0x54, 0x09, 0x3d, 0x84, 0x8e, 0xcf, 0x29, 0x8f, 0x3c, 0x16,
	0x5e, 0x5d, 0x97, 0x71, 0xa7, 0xac, 0x2b, 0x29, 0x5d, 0x68, 0xf3, 0x24, 0x21, 0xe9, 0x6c, 0x1e,
	0xf3, 0xe5, 0xa5, 0xd6, 0x02, 0xb2, 0x74, 0x24, 0x2a, 0xd7, 0x1c, 0x50, 0xfb, 0x43, 0x07, 0x7c,
	0x37, 0xa0, 0x95, 0x97, 0xd1, 0x3e, 0xd4, 0x23, 0xe1, 0x60, 0xdb, 0x90, 0x87, 0xea, 0x56, 0x4f,
	0x29, 0x4c, 0x8e, 0x15, 0xba, 0xda, 0x1d, 0xe8, 0x29, 0x58, 0x45, 0x42, 0xb4, 0xac, 0x2d, 0x57,
	0x65, 0xc8, 0xcd, 0x33, 0xe4, 0x4e, 0x72, 0x04, 0x2e, 0xc1, 0xce, 0xcf, 0x35, 0x68, 0x8c, 0x64,
	0x22, 0xff, 0x56, 0xd1, 0x63, 0xa8, 0x07, 0x22, 0x53, 0x3a, 0x10, 0x77, 0xab, 0x69, 0x32, 0x76,
	0x58, 0x21, 0xd1, 0x13, 0x68, 0xfa, 0x2a, 0x67, 0x5a, 0xec, 0x76, 0x35, 0x49, 0x87, 0x11, 0xe7,
	0x68, 0x41, 0xcc, 0x54, 0x08, 0x6c, 0xf3, 0x36, 0xa2, 0x4e, 0x0a, 0xce, 0xd1, 0x82, 0xc8, 0x95,
	0x69, 0xed, 0xfa, 0x6d, 0x44, 0xed, 0x6c, 0x9c, 0xa3, 0xd1, 0x0b, 0xb0, 0x16, 0xb9, 0x97, 0xed,
	0xa6, 0xa4, 0xde, 0x70, 0x31, 0x85, 0xe5, 0x71, 0xc9, 0x10, 0xee, 0x2f, 0xee, 0x7a, 0x46, 0x33,
	0xbb, 0xd1, 0x33, 0x76, 0x6a, 0xb8, 0x5d, 0xd4, 0x46, 0x99, 0xf3, 0xc3, 0x80, 0x75, 0xf5, 0x07,
	0x5e, 0x79, 0x34, 0x8c, 0x56, 0x95, 0xcf, 0x19, 0x02, 0x73, 0x41, 0xa2, 0x44, 0xbf, 0x66, 0x72,
	0x8d, 0xf6, 0xc0, 0x14, 0x1a, 0xe5, 0x15, 0xfe, 0xbf, 0xdb, 0xab, 0x56, 0xa5, 0x26, 0x4f, 0x56,
	0x09, 0xc1, 0x12, 0x2d, 0xd2, 0xa4, 0x5e, 0x60, 0xdb, 0xbc, 0x2d, 0x4d, 0x8a, 0x87, 0x35, 0xf6,
	0xd1, 0x08, 0xa0, 0x9c, 0x84, 0xda, 0xd0, 0x3c, 0x3e, 0x9b, 0x9e, 0x4e, 0x06, 0xb8, 0xf3, 0x1f,
	0xb2, 0xa0, 0x3e, 0x3c, 0x9c, 0x0e, 0x07, 0x1d, 0x43, 0xd4, 0xc7, 0xd3, 0xd1, 0xe8, 0x10, 0x5f,
	0x74, 0xd6, 0xc4, 0x66, 0x7a, 0x3a, 0xb9, 0x38, 0x1f, 0xbc, 0xec, 0xd4, 0xd0, 0x06, 0x58, 0x27,
	0xaf, 0xc7, 0x93, 0xb3, 0x21, 0x3e, 0x1c, 0x75, 0xcc, 0x23, 0x0c, 0x95, 0xef, 0xfe, 0xbb, 0x83,
	0x20, 0x64, 0x0b, 0x3e, 0x77, 0xfd, 0x98, 0xf6, 0xcb, 0x6e, 0x5f, 0x75, 0x67, 0x34, 0xbe, 0x24,
	0x51, 0x3f, 0x88, 0x9f, 0x85, 0xf1, 0xac, 0xec, 0xce, 0x54, 0xf7, 0x57, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xd0, 0x84, 0x91, 0x73, 0x59, 0x06, 0x00, 0x00,
}
