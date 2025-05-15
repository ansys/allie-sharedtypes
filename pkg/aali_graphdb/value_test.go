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
	"time"

	"cloud.google.com/go/civil"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type valueTest struct {
	name          string
	value         Value
	expectedOneOf []any // allow for multiple correct options since sometimes order can change
}

var valueTests = []valueTest{
	{
		"Null[Any]",
		NullValue{AnyLogicalType{}},
		[]any{map[string]any{"Null": "Any"}},
	},
	{
		"Null[List[Float]]",
		NullValue{ListLogicalType{FloatLogicalType{}}},
		[]any{map[string]any{"Null": map[string]any{"List": map[string]any{"child_type": "Float"}}}},
	},
	{
		"Bool",
		BoolValue(true),
		[]any{map[string]any{"Bool": true}},
	},
	{
		"Int64",
		Int64Value(82),
		[]any{map[string]any{"Int64": float64(82)}},
	},
	{
		"Int32",
		Int32Value(1),
		[]any{map[string]any{"Int32": float64(1)}},
	},
	{
		"Int16",
		Int16Value(100),
		[]any{map[string]any{"Int16": float64(100)}},
	},
	{
		"Int8",
		Int8Value(-6),
		[]any{map[string]any{"Int8": float64(-6)}},
	},
	{
		"UInt64",
		UInt64Value(0),
		[]any{map[string]any{"UInt64": float64(0)}},
	},
	{
		"UInt32",
		UInt32Value(1001),
		[]any{map[string]any{"UInt32": float64(1001)}},
	},
	{
		"UInt16",
		UInt16Value(212),
		[]any{map[string]any{"UInt16": float64(212)}},
	},
	{
		"UInt8",
		UInt8Value(50),
		[]any{map[string]any{"UInt8": float64(50)}},
	},
	{
		"Int128",
		Int128Value(9009),
		[]any{map[string]any{"Int128": float64(9009)}},
	},
	{
		"Double",
		DoubleValue(-56.1234),
		[]any{map[string]any{"Double": -56.1234}},
	},
	{
		"Float",
		FloatValue(90.0),
		[]any{map[string]any{"Float": 90.0}},
	},
	{
		"InternalID",
		InternalIDValue{0, 0},
		[]any{map[string]any{"InternalID": map[string]any{"table_id": float64(0), "offset": float64(0)}}},
	},
	{
		"String",
		StringValue("Hello"),
		[]any{map[string]any{"String": "Hello"}},
	},
	{
		"Blob",
		BlobValue([]uint8{0, 1, 2, 3, 4}),
		[]any{map[string]any{"Blob": []any{float64(0), float64(1), float64(2), float64(3), float64(4)}}},
	},
	{
		"List",
		ListValue{UInt64LogicalType{}, []Value{UInt64Value(0), UInt64Value(12)}},
		[]any{map[string]any{"List": []any{"UInt64", []any{map[string]any{"UInt64": float64(0)}, map[string]any{"UInt64": float64(12)}}}}},
	},
	{
		"Array",
		ArrayValue{BoolLogicalType{}, []Value{BoolValue(true), BoolValue(false)}},
		[]any{map[string]any{"Array": []any{"Bool", []any{map[string]any{"Bool": true}, map[string]any{"Bool": false}}}}},
	},
	{
		"Struct",
		StructValue(map[string]Value{"a": BoolValue(false), "name": StringValue("Joe")}),
		[]any{
			map[string]any{"Struct": []any{[]any{"a", map[string]any{"Bool": false}}, []any{"name", map[string]any{"String": "Joe"}}}},
			map[string]any{"Struct": []any{[]any{"name", map[string]any{"String": "Joe"}}, []any{"a", map[string]any{"Bool": false}}}},
		},
	},
	{
		"Node",
		NodeValue{InternalID{1, 10}, "my-label", map[string]Value{}},
		[]any{
			map[string]any{"Node": map[string]any{"id": map[string]any{"table_id": float64(1), "offset": float64(10)}, "label": "my-label", "properties": []any{}}},
		},
	},
	{
		"Rel",
		RelValue{InternalID{4, 1}, InternalID{6, 0}, "lab", map[string]Value{}},
		[]any{
			map[string]any{"Rel": map[string]any{"src_node": map[string]any{"table_id": float64(4), "offset": float64(1)}, "dst_node": map[string]any{"table_id": float64(6), "offset": float64(0)}, "label": "lab", "properties": []any{}}},
		},
	},
	{
		"Map",
		MapValue{UInt64LogicalType{}, BoolLogicalType{}, map[Value]Value{UInt64Value(4): BoolValue(false)}},
		[]any{
			map[string]any{"Map": []any{[]any{"UInt64", "Bool"}, []any{[]any{map[string]any{"UInt64": float64(4)}, map[string]any{"Bool": false}}}}},
		},
	},
	{
		"Union",
		UnionValue{map[string]LogicalType{"num": Int64LogicalType{}, "str": StringLogicalType{}}, Int64Value(1)},
		[]any{
			map[string]any{"Union": map[string]any{"types": []any{[]any{"num", "Int64"}, []any{"str", "String"}}, "value": map[string]any{"Int64": float64(1)}}},
			map[string]any{"Union": map[string]any{"types": []any{[]any{"str", "String"}, []any{"num", "Int64"}}, "value": map[string]any{"Int64": float64(1)}}},
		},
	},
	{
		"UUID[zeros]",
		UUIDValue(uuid.MustParse("00000000-0000-0000-0000-ffff00000000")),
		[]any{map[string]any{"UUID": "00000000-0000-0000-0000-ffff00000000"}},
	},
	{
		"UUID",
		UUIDValue(uuid.MustParse("8f914bce-df4e-4244-9cd4-ea96bf0c58d4")),
		[]any{map[string]any{"UUID": "8f914bce-df4e-4244-9cd4-ea96bf0c58d4"}},
	},
	{
		"Decimal[small]",
		DecimalValue(decimal.RequireFromString("12.34")),
		[]any{map[string]any{"Decimal": "12.34"}},
	},
	{
		"Decimal[big]",
		DecimalValue(decimal.RequireFromString("12.3456789")),
		[]any{map[string]any{"Decimal": "12.3456789"}},
	},
	{
		"Date",
		DateValue(civil.Date{Year: 2025, Month: time.April, Day: 23}),
		[]any{map[string]any{"Date": "2025-04-23"}},
	},
	{
		"Timestamp",
		TimestampValue(time.Date(2025, time.April, 23, 13, 26, 21, 123450000, time.UTC)),
		[]any{map[string]any{"Timestamp": "2025-04-23T13:26:21.12345Z"}},
	},
	{
		"TimestampTz",
		TimestampTzValue(time.Date(2025, time.April, 23, 13, 26, 21, 123450000, time.UTC)),
		[]any{map[string]any{"TimestampTz": "2025-04-23T13:26:21.12345Z"}},
	},
	{
		"TimestampNs",
		TimestampNsValue(time.Date(2025, time.April, 23, 13, 26, 21, 123450000, time.UTC)),
		[]any{map[string]any{"TimestampNs": "2025-04-23T13:26:21.12345Z"}},
	},
	{
		"TimestampMs",
		TimestampMsValue(time.Date(2025, time.April, 23, 13, 26, 21, 123450000, time.UTC)),
		[]any{map[string]any{"TimestampMs": "2025-04-23T13:26:21.12345Z"}},
	},
	{
		"TimestampSec",
		TimestampSecValue(time.Date(2025, time.April, 23, 13, 26, 21, 123450000, time.UTC)),
		[]any{map[string]any{"TimestampSec": "2025-04-23T13:26:21.12345Z"}},
	},
	{
		"Interval",
		IntervalValue(23 * 24 * time.Hour),
		[]any{map[string]any{"Interval": []any{float64(1987200), float64(0)}}},
	},
	{
		"Interval[ns]",
		IntervalValue(23*24*time.Hour + 456*time.Nanosecond),
		[]any{map[string]any{"Interval": []any{float64(1987200), float64(456)}}},
	},
}

func TestValuesMarshal(t *testing.T) {
	for _, test := range valueTests {
		t.Run(test.name, func(t *testing.T) {
			actualBytes, err := json.Marshal(test.value)
			require.NoError(t, err)
			var actualJson any
			err = json.Unmarshal(actualBytes, &actualJson)
			require.NoError(t, err)
			assert.Contains(t, test.expectedOneOf, actualJson)
		})
	}
}
