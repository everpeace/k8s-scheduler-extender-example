package colog

import (
	"testing"
)

type extractInOut struct {
	in      string
	outMsg  string
	outData map[string]interface{}
}

var extractorTests = []extractInOut{
	{
		"some message foo=bar",
		"some message",
		map[string]interface{}{"foo": "bar"},
	},
	{
		"some message foo=42 foo2='other bar'",
		"some message",
		map[string]interface{}{
			"foo":  "42",
			"foo2": "other bar",
		},
	},
	{
		"some foo=42 otherfoo='other bar' mixed text",
		"some   mixed text",
		map[string]interface{}{
			"foo":      "42",
			"otherfoo": "other bar",
		},
	},
}

func TestStdExtractor(t *testing.T) {
	stde := StdExtractor{}
	for _, tt := range extractorTests {
		e := &Entry{
			Message: []byte(tt.in),
			Fields:  make(Fields),
		}
		stde.Extract(e)
		if string(e.Message) != tt.outMsg {
			t.Errorf("Extract error:\n %s\n %s", string(e.Message), tt.outMsg)
		}
		for k, v := range tt.outData {
			if e.Fields[k] != v {
				t.Errorf("Invalid data value %s\n %s vs %s", k, v, e.Fields[k])
			}
		}
	}
}
