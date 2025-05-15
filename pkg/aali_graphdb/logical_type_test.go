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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type logicalTypesTest struct {
	name        string
	logicalType LogicalType
	expected    any
}

var logicalTypesTests = []logicalTypesTest{
	{
		"Any",
		AnyLogicalType{},
		"Any",
	},
	{
		"Bool",
		BoolLogicalType{},
		"Bool",
	},
	{
		"Serial",
		SerialLogicalType{},
		"Serial",
	},
	{
		"Int64",
		Int64LogicalType{},
		"Int64",
	},
	{
		"Int32",
		Int32LogicalType{},
		"Int32",
	},
	{
		"Int16",
		Int16LogicalType{},
		"Int16",
	},
	{
		"Int8",
		Int8LogicalType{},
		"Int8",
	},
	{
		"UInt64",
		UInt64LogicalType{},
		"UInt64",
	},
	{
		"UInt32",
		UInt32LogicalType{},
		"UInt32",
	},
	{
		"UInt16",
		UInt16LogicalType{},
		"UInt16",
	},
	{
		"UInt8",
		UInt8LogicalType{},
		"UInt8",
	},
	{
		"Int128",
		Int128LogicalType{},
		"Int128",
	},
	{
		"Double",
		DoubleLogicalType{},
		"Double",
	},
	{
		"Float",
		FloatLogicalType{},
		"Float",
	},
	{
		"Date",
		DateLogicalType{},
		"Date",
	},
	{
		"Interval",
		IntervalLogicalType{},
		"Interval",
	},
	{
		"Timestamp",
		TimestampLogicalType{},
		"Timestamp",
	},
	{
		"TimestampTz",
		TimestampTzLogicalType{},
		"TimestampTz",
	},
	{
		"TimestampNs",
		TimestampNsLogicalType{},
		"TimestampNs",
	},
	{
		"TimestampSec",
		TimestampSecLogicalType{},
		"TimestampSec",
	},
	{
		"InternalID",
		InternalIDTypeLogicalType{},
		"InternalID",
	},
	{
		"String",
		StringLogicalType{},
		"String",
	},
	{
		"Blob",
		BlobLogicalType{},
		"Blob",
	},
	{
		"List[Bool]",
		ListLogicalType{BoolLogicalType{}},
		map[string]any{"List": map[string]any{"child_type": "Bool"}},
	},
	{
		"Array",
		ArrayLogicalType{UInt16LogicalType{}, 12},
		map[string]any{"Array": map[string]any{"child_type": "UInt16", "num_elements": float64(12)}},
	},
	{
		"Struct",
		StructLogicalType{[]TwoTuple[string, LogicalType]{
			{
				"name",
				StringLogicalType{},
			},
			{
				"age",
				UInt32LogicalType{},
			},
			{
				"items",
				ListLogicalType{StringLogicalType{}},
			},
		}},
		map[string]any{"Struct": map[string]any{
			"fields": []any{
				[]any{"name", "String"},
				[]any{"age", "UInt32"},
				[]any{"items", map[string]any{"List": map[string]any{"child_type": "String"}}},
			}}},
	},
	{
		"Node",
		NodeLogicalType{},
		"Node",
	},
	{
		"Rel",
		RelLogicalType{},
		"Rel",
	},
	{
		"RecursiveRel",
		RecursiveRelLogicalType{},
		"RecursiveRel",
	},
	{
		"Map",
		MapLogicalType{StringLogicalType{}, Int8LogicalType{}},
		map[string]any{"Map": map[string]any{"key_type": "String", "value_type": "Int8"}},
	},
	{
		"Union",
		UnionLogicalType{[]TwoTuple[string, LogicalType]{
			{
				"name",
				StringLogicalType{},
			},
			{
				"age",
				UInt32LogicalType{},
			},
			{
				"items",
				ListLogicalType{StringLogicalType{}},
			},
		}},
		map[string]any{"Union": map[string]any{
			"fields": []any{
				[]any{"name", "String"},
				[]any{"age", "UInt32"},
				[]any{"items", map[string]any{"List": map[string]any{"child_type": "String"}}},
			}}},
	},
	{
		"UUID",
		UUIDLogicalType{},
		"UUID",
	},
	{
		"Decimal",
		DecimalLogicalType{5, 3},
		map[string]any{"Decimal": map[string]any{"precision": float64(5), "scale": float64(3)}},
	},
}

func TestLogicalTypesMarshal(t *testing.T) {
	for _, test := range logicalTypesTests {
		t.Run(test.name, func(t *testing.T) {
			actualBytes, err := json.Marshal(test.logicalType)
			require.NoError(t, err)
			var actualJson any
			err = json.Unmarshal(actualBytes, &actualJson)
			require.NoError(t, err)
			assert.Equal(t, test.expected, actualJson)

		})
	}
}
