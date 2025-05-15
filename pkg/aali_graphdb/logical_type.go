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
)

type LogicalType interface {
	IsKuzuLogicalType()
	MarshalJSON() ([]byte, error)
}

type AnyLogicalType struct{}

func (lt AnyLogicalType) IsKuzuLogicalType() {}
func (lt AnyLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("Any")
}

type BoolLogicalType struct{}

func (lt BoolLogicalType) IsKuzuLogicalType() {}
func (lt BoolLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("Bool")
}

type SerialLogicalType struct{}

func (lt SerialLogicalType) IsKuzuLogicalType() {}
func (lt SerialLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("Serial")
}

type Int64LogicalType struct{}

func (lt Int64LogicalType) IsKuzuLogicalType() {}
func (lt Int64LogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("Int64")
}

type Int32LogicalType struct{}

func (lt Int32LogicalType) IsKuzuLogicalType() {}
func (lt Int32LogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("Int32")
}

type Int16LogicalType struct{}

func (lt Int16LogicalType) IsKuzuLogicalType() {}
func (lt Int16LogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("Int16")
}

type Int8LogicalType struct{}

func (lt Int8LogicalType) IsKuzuLogicalType() {}
func (lt Int8LogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("Int8")
}

type UInt64LogicalType struct{}

func (lt UInt64LogicalType) IsKuzuLogicalType() {}
func (lt UInt64LogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("UInt64")
}

type UInt32LogicalType struct{}

func (lt UInt32LogicalType) IsKuzuLogicalType() {}
func (lt UInt32LogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("UInt32")
}

type UInt16LogicalType struct{}

func (lt UInt16LogicalType) IsKuzuLogicalType() {}
func (lt UInt16LogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("UInt16")
}

type UInt8LogicalType struct{}

func (lt UInt8LogicalType) IsKuzuLogicalType() {}
func (lt UInt8LogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("UInt8")
}

type Int128LogicalType struct{}

func (lt Int128LogicalType) IsKuzuLogicalType() {}
func (lt Int128LogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("Int128")
}

type DoubleLogicalType struct{}

func (lt DoubleLogicalType) IsKuzuLogicalType() {}
func (lt DoubleLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("Double")
}

type FloatLogicalType struct{}

func (lt FloatLogicalType) IsKuzuLogicalType() {}
func (lt FloatLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("Float")
}

type DateLogicalType struct{}

func (lt DateLogicalType) IsKuzuLogicalType() {}
func (lt DateLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("Date")
}

type IntervalLogicalType struct{}

func (lt IntervalLogicalType) IsKuzuLogicalType() {}
func (lt IntervalLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("Interval")
}

type TimestampLogicalType struct{}

func (lt TimestampLogicalType) IsKuzuLogicalType() {}
func (lt TimestampLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("Timestamp")
}

type TimestampTzLogicalType struct{}

func (lt TimestampTzLogicalType) IsKuzuLogicalType() {}
func (lt TimestampTzLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("TimestampTz")
}

type TimestampNsLogicalType struct{}

func (lt TimestampNsLogicalType) IsKuzuLogicalType() {}
func (lt TimestampNsLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("TimestampNs")
}

type TimestampMsLogicalType struct{}

func (lt TimestampMsLogicalType) IsKuzuLogicalType() {}
func (lt TimestampMsLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("TimestampMs")
}

type TimestampSecLogicalType struct{}

func (lt TimestampSecLogicalType) IsKuzuLogicalType() {}
func (lt TimestampSecLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("TimestampSec")
}

type InternalIDTypeLogicalType struct{}

func (lt InternalIDTypeLogicalType) IsKuzuLogicalType() {}
func (lt InternalIDTypeLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("InternalID")
}

type StringLogicalType struct{}

func (lt StringLogicalType) IsKuzuLogicalType() {}
func (lt StringLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("String")
}

type BlobLogicalType struct{}

func (lt BlobLogicalType) IsKuzuLogicalType() {}
func (lt BlobLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("Blob")
}

type ListLogicalType struct {
	ChildType LogicalType
}

func (lt ListLogicalType) IsKuzuLogicalType() {}
func (lt ListLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"List",
		map[string]LogicalType{"child_type": lt.ChildType},
	})
}

type ArrayLogicalType struct {
	ChildType   LogicalType
	NumElements uint64
}

func (lt ArrayLogicalType) IsKuzuLogicalType() {}
func (lt ArrayLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"Array",
		map[string]any{"child_type": lt.ChildType, "num_elements": lt.NumElements},
	})
}

type StructLogicalType struct {
	Fields []TwoTuple[string, LogicalType] `json:"fields"`
}

func (lt StructLogicalType) IsKuzuLogicalType() {}
func (lt StructLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"Struct",
		map[string]any{"fields": lt.Fields},
	})
}

type NodeLogicalType struct{}

func (lt NodeLogicalType) IsKuzuLogicalType() {}
func (lt NodeLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("Node")
}

type RelLogicalType struct{}

func (lt RelLogicalType) IsKuzuLogicalType() {}
func (lt RelLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("Rel")
}

type RecursiveRelLogicalType struct{}

func (lt RecursiveRelLogicalType) IsKuzuLogicalType() {}
func (lt RecursiveRelLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("RecursiveRel")
}

type MapLogicalType struct {
	KeyType   LogicalType
	ValueType LogicalType
}

func (lt MapLogicalType) IsKuzuLogicalType() {}
func (lt MapLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"Map",
		map[string]any{"key_type": lt.KeyType, "value_type": lt.ValueType},
	})
}

type UnionLogicalType struct {
	Fields []TwoTuple[string, LogicalType]
}

func (lt UnionLogicalType) IsKuzuLogicalType() {}
func (lt UnionLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"Union",
		map[string]any{"fields": lt.Fields},
	})
}

type UUIDLogicalType struct{}

func (lt UUIDLogicalType) IsKuzuLogicalType() {}
func (lt UUIDLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("UUID")
}

type DecimalLogicalType struct {
	Precision uint32
	Scale     uint32
}

func (lt DecimalLogicalType) IsKuzuLogicalType() {}
func (lt DecimalLogicalType) MarshalJSON() ([]byte, error) {
	return json.Marshal(ExternallyTagged{
		"Decimal",
		map[string]uint32{"precision": lt.Precision, "scale": lt.Scale},
	})
}
