// strings_utils_test
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
		AppendToNewLine(&buf,test.f)
		s := buf.String()
		if s != test.s {
			t.Error("ValueOf ERROR:", test.f, "want", test.s, "got", s)
		}
		buf.Reset()
	}

}
