package colog

import (
	"log"
	"testing"
	"time"
)

type JSONFormatTest struct {
	entry  Entry
	prefix string
	flags  int
	tfmt   string
	lnum   bool
	output string
}

var JSONFormatTests = []JSONFormatTest{
	{
		entry: Entry{
			Level:   LInfo,
			Message: []byte("some message"),
			Fields:  map[string]interface{}{"foo": "bar"},
		},
		output: `{"level":"info","message":"some message","fields":{"foo":"bar"}}` + "\n",
	},
	{
		entry: Entry{
			Time:    TTime,
			Level:   LDebug,
			Message: []byte("some message"),
			Fields:  map[string]interface{}{"foo": "bar"},
		},
		flags:  log.Ldate,
		output: `{"level":"debug","time":"2015/8/1","message":"some message","fields":{"foo":"bar"}}` + "\n",
	},
	{
		entry: Entry{
			Time:    TTime,
			Level:   LAlert,
			Message: []byte("some message"),
			Fields:  map[string]interface{}{"foo": "bar"},
		},
		flags:  log.Ldate | log.Lmicroseconds,
		output: `{"level":"alert","time":"2015/8/1 20:45:30.9999","message":"some message","fields":{"foo":"bar"}}` + "\n",
	},
	{
		entry: Entry{
			Time:    TTime,
			Level:   LTrace,
			File:    "/src/file.go",
			Line:    42,
			Message: []byte("some message"),
			Fields:  map[string]interface{}{"foo": "bar"},
		},
		flags:  log.Ldate | log.Lmicroseconds | log.Lshortfile,
		output: `{"level":"trace","time":"2015/8/1 20:45:30.9999","file":"file.go","line":42,"message":"some message","fields":{"foo":"bar"}}` + "\n",
	},
	{
		entry: Entry{
			Time:    TTime,
			Level:   LWarning,
			File:    "/src/file.go",
			Line:    142,
			Message: []byte("some message"),
			Fields:  map[string]interface{}{"foo": "bar"},
		},
		tfmt:   time.RFC3339Nano,
		flags:  log.Ldate | log.Llongfile,
		output: `{"level":"warning","time":"2015-08-01T20:45:30.000009999Z","file":"/src/file.go","line":142,"message":"some message","fields":{"foo":"bar"}}` + "\n",
	},
	{
		entry: Entry{
			Time:    TTime,
			Level:   LError,
			File:    "/src/file.go",
			Line:    142,
			Prefix:  "worker1 ",
			Message: []byte("some message"),
			Fields:  map[string]interface{}{"foo": "bar"},
		},
		flags:  log.Ldate | log.Llongfile,
		output: `{"level":"error","time":"2015/8/1","prefix":"worker1 ","file":"/src/file.go","line":142,"message":"some message","fields":{"foo":"bar"}}` + "\n",
	},
	{
		entry: Entry{
			Time:    TTime,
			Level:   LDebug,
			Message: []byte("some message"),
			Fields:  map[string]interface{}{"foo": "bar"},
		},
		lnum:   true,
		flags:  log.Ldate,
		output: `{"level":"2","time":"2015/8/1","message":"some message","fields":{"foo":"bar"}}` + "\n",
	},
	{
		entry: Entry{
			Time:    TTime,
			Level:   LAlert,
			File:    "/src/file.go",
			Line:    142,
			Message: []byte("some message"),
			Fields:  map[string]interface{}{"foo": "bar"},
		},
		lnum:   true,
		tfmt:   time.RFC822Z,
		flags:  log.Llongfile,
		output: `{"level":"6","time":"01 Aug 15 20:45 +0000","file":"/src/file.go","line":142,"message":"some message","fields":{"foo":"bar"}}` + "\n",
	},
}

func TestJSONFormatter(t *testing.T) {
	for _, tt := range JSONFormatTests {
		f := JSONFormatter{
			Flag:       tt.flags,
			LevelAsNum: tt.lnum,
			TimeFormat: tt.tfmt,
		}

		b, err := f.Format(&tt.entry)
		if err != nil {
			t.Fatal(err)
		}

		if string(b) != tt.output {
			t.Errorf("Unexpected JSON formatter output: %s", b)
		}
	}
}
