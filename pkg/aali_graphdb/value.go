// Copyright (C) 2025 ANSYS, Inc. and/or its affiliates.
// SPDX-License-Identifier: MIT
//
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package aali_graphdb

import (
	"encoding/json"
	"time"

	"cloud.google.com/go/civil"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Value interface {
	IsKuzuValue()
	MarshalJSON() ([]byte, error)
}

type NullValue struct {
	LogicalType LogicalType
}

func (v NullValue) IsKuzuValue() {}
func (v NullValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"Null",
		v.LogicalType,
	})
}

type BoolValue bool

func (v BoolValue) IsKuzuValue() {}
func (v BoolValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"Bool",
		bool(v),
	})
}

type Int64Value int64

func (v Int64Value) IsKuzuValue() {}
func (v Int64Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"Int64",
		int64(v),
	})
}

type Int32Value int32

func (v Int32Value) IsKuzuValue() {}
func (v Int32Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"Int32",
		int32(v),
	})
}

type Int16Value int16

func (v Int16Value) IsKuzuValue() {}
func (v Int16Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"Int16",
		int16(v),
	})
}

type Int8Value int8

func (v Int8Value) IsKuzuValue() {}

func (v Int8Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"Int8",
		int8(v),
	})
}

type UInt64Value uint64

func (v UInt64Value) IsKuzuValue() {}
func (v UInt64Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"UInt64",
		uint64(v),
	})
}

type UInt32Value uint32

func (v UInt32Value) IsKuzuValue() {}
func (v UInt32Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"UInt32",
		uint32(v),
	})
}

type UInt16Value uint16

func (v UInt16Value) IsKuzuValue() {}
func (v UInt16Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"UInt16",
		uint16(v),
	})
}

type UInt8Value uint8

func (v UInt8Value) IsKuzuValue() {}
func (v UInt8Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"UInt8",
		uint8(v),
	})
}

type Int128Value int64 // no int128 in go, but could still be that type in the DB

func (v Int128Value) IsKuzuValue() {}
func (v Int128Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"Int128",
		int64(v),
	})
}

type DoubleValue float64

func (v DoubleValue) IsKuzuValue() {}

func (v DoubleValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"Double",
		float64(v),
	})
}

type FloatValue float32

func (v FloatValue) IsKuzuValue() {}
func (v FloatValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"Float",
		float32(v),
	})
}

type DateValue civil.Date

func (v DateValue) IsKuzuValue() {}
func (v DateValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"Date",
		civil.Date(v),
	})
}

type IntervalValue time.Duration

func (v IntervalValue) IsKuzuValue() {}
func (v IntervalValue) MarshalJSON() ([]byte, error) {
	// should return an array of [<seconds>, <nanoseconds>]
	val := time.Duration(v)

	roundMultiple, err := time.ParseDuration("1s")
	if err != nil {
		return nil, err
	}
	rounded := val.Round(roundMultiple)
	wholeSecs := int64(rounded.Seconds())
	leftover := val - rounded
	nanos := leftover.Nanoseconds()
	return json.Marshal(ExternallyTagged{
		"Interval",
		[]int64{wholeSecs, nanos},
	})
}

type TimestampValue time.Time

func (v TimestampValue) IsKuzuValue() {}
func (v TimestampValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"Timestamp",
		time.Time(v).Format(time.RFC3339Nano),
	})
}

type TimestampTzValue time.Time

func (v TimestampTzValue) IsKuzuValue() {}
func (v TimestampTzValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"TimestampTz",
		time.Time(v).Format(time.RFC3339Nano),
	})
}

type TimestampNsValue time.Time

func (v TimestampNsValue) IsKuzuValue() {}
func (v TimestampNsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"TimestampNs",
		time.Time(v).Format(time.RFC3339Nano),
	})
}

type TimestampMsValue time.Time

func (v TimestampMsValue) IsKuzuValue() {}
func (v TimestampMsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"TimestampMs",
		time.Time(v).Format(time.RFC3339Nano),
	})
}

type TimestampSecValue time.Time

func (v TimestampSecValue) IsKuzuValue() {}
func (v TimestampSecValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"TimestampSec",
		time.Time(v).Format(time.RFC3339Nano),
	})
}

type InternalID struct {
	TableID uint64 `json:"table_id"`
	Offset  uint64 `json:"offset"`
}

type InternalIDValue struct {
	TableID uint64
	Offset  uint64
}

func (v InternalIDValue) IsKuzuValue() {}
func (v InternalIDValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"InternalID",
		InternalID(v),
	})
}

type StringValue string

func (v StringValue) IsKuzuValue() {}
func (v StringValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"String",
		string(v),
	})
}

type BlobValue []uint8

func (v BlobValue) IsKuzuValue() {}
func (v BlobValue) MarshalJSON() ([]byte, error) {
	inner := make([]uint16, len(v))
	for i, element := range v {
		inner[i] = uint16(element)
	}
	return json.Marshal(ExternallyTagged{
		"Blob",
		inner,
	})
}

type ListValue struct {
	LogicalType LogicalType
	Values      []Value
}

func (v ListValue) IsKuzuValue() {}
func (v ListValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"List",
		[]any{v.LogicalType, v.Values},
	})
}

type ArrayValue struct {
	LogicalType LogicalType
	Values      []Value
}

func (v ArrayValue) IsKuzuValue() {}
func (v ArrayValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"Array",
		[]any{v.LogicalType, v.Values},
	})
}

type StructValue map[string]Value

func (v StructValue) IsKuzuValue() {}
func (v StructValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"Struct",
		twoTupleArrayFromMap(v),
	})

}

type NodeValue struct {
	ID         InternalID
	Label      string
	Properties map[string]Value
}

func (v NodeValue) IsKuzuValue() {}
func (v NodeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"Node",
		map[string]any{"id": v.ID, "label": v.Label, "properties": twoTupleArrayFromMap(v.Properties)},
	})
}

type RelValue struct {
	SrcNode    InternalID
	DstNode    InternalID
	Label      string
	Properties map[string]Value
}

func (v RelValue) IsKuzuValue() {}
func (v RelValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"Rel",
		map[string]any{"src_node": v.SrcNode, "dst_node": v.DstNode, "label": v.Label, "properties": twoTupleArrayFromMap(v.Properties)},
	})
}

type RecursiveRelValue struct {
	Nodes []NodeValue
	Rels  []RelValue
}

func (v RecursiveRelValue) IsKuzuValue() {}
func (v RecursiveRelValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"RecursiveRel",
		map[string]any{"nodes": v.Nodes, "rels": v.Rels},
	})
}

type MapValue struct {
	KeyType   LogicalType
	ValueType LogicalType
	Pairs     map[Value]Value
}

func (v MapValue) IsKuzuValue() {}
func (v MapValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"Map",
		TwoTuple[TwoTuple[LogicalType, LogicalType], []TwoTuple[Value, Value]]{
			TwoTuple[LogicalType, LogicalType]{
				v.KeyType,
				v.ValueType,
			},
			twoTupleArrayFromMap(v.Pairs),
		},
	})
}

type UnionValue struct {
	Types map[string]LogicalType
	Value Value
}

func (v UnionValue) IsKuzuValue() {}
func (v UnionValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"Union",
		map[string]any{
			"types": twoTupleArrayFromMap(v.Types),
			"value": v.Value,
		},
	})
}

type UUIDValue uuid.UUID

func (v UUIDValue) IsKuzuValue() {}
func (v UUIDValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"UUID",
		uuid.UUID(v).String(),
	})
}

type DecimalValue decimal.Decimal

func (v DecimalValue) IsKuzuValue() {}
func (v DecimalValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"Decimal",
		decimal.Decimal(v).String(),
	})
}

// Json marshal helper for externally-tagged types
//
// Mimics the externally taggged serde format in rust: https://serde.rs/enum-representations.html
//
// examples:
// `"Any"`
// `{"Type": "content"}`
type ExternallyTagged struct {
	type_ string
	value any
}

func (exTag ExternallyTagged) MarshalJSON() ([]byte, error) {
	valBytes, err := json.Marshal(exTag.value)
	if err != nil {
		return nil, err
	}
	if len(valBytes) == 0 {
		return json.Marshal(exTag.type_)

	}
	var valJson any
	err = json.Unmarshal(valBytes, &valJson)
	if err != nil {
		return nil, err
	}
	return json.Marshal(map[string]any{exTag.type_: valJson})
}

// helper type for serializing 2-tuples
type TwoTuple[A any, B any] struct {
	a A
	b B
}

func (tup TwoTuple[A, B]) MarshalJSON() ([]byte, error) {
	return json.Marshal([]any{tup.a, tup.b})
}

func twoTupleArrayFromMap[A comparable, B any](m map[A]B) []TwoTuple[A, B] {
	twoTuples := make([]TwoTuple[A, B], len(m))
	i := 0
	for k, v := range m {
		twoTuples[i] = TwoTuple[A, B]{k, v}
		i++
	}
	return twoTuples
}
