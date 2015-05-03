// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package sutils

import (
	"bytes"
 
	"testing"
)

type Any interface{}
type valueTest struct {
	f Any
	s string
}

var valuetests = []valueTest{
	{1, "1"},
	{3, "3"},
	{"3", "3"},
	{float32(3.0), "3"},
	{float32(3.1), "3.1"},
	{3, "3"},
}

func TestValueOf(t *testing.T) {
	for i := 0; i < len(valuetests); i++ {
		test := &valuetests[i]
		s := ValueOf(test.f)
		if s != test.s {
			t.Error("ValueOf ERROR:", test.f, "want", test.s, "got", s)
		}
	}
}

var appendtests = []valueTest{
	{"set", "set\r\n"},
	{"mapping", "mapping\r\n"},
}

func TestAppendNewLine(t *testing.T) {
	var buf bytes.Buffer
	for i := 0; i < len(appendtests); i++ {
		test := &appendtests[i]
		//AppendToNewLine(&buf,test.f)
		s := buf.String()
		if s == test.s {
			t.Error("ValueOf ERROR:", test.f, "want", test.s, "got", s)
		}
		buf.Reset()
	}

}
func BenchmarkAppendNewLine(b *testing.B) {
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		buf.Reset()
		AppendToNewLine(&buf, i)
	}
}

var buf []byte = make([]byte, 0)

func BenchmarkValueOf(b *testing.B) {

	for i := 0; i < b.N; i++ {
		 		ValueOf(i)
		//strconv.AppendInt(buf, int64(i), 10)
//BenchmarkAppendNewLine	  100000	     11970 ns/op
//BenchmarkValueOf	10000000	       142 ns/op
//ok  	utils/strings	2.892s
	}
}
