package colog

import (
	"log"
	"testing"
	"time"
)

type formatTest struct {
	entry  Entry
	prefix string
	flags  int
	width  int
	colors bool
	output string
}

// TTime is the fixed point in time for all formatting tests
var TTime = time.Date(2015, time.August, 1, 20, 45, 30, 9999, time.UTC)

var formatterTests = []formatTest{
	{
		entry: Entry{
			Level:   LInfo,
			Message: []byte("some message"),
			Fields:  map[string]interface{}{"foo": "bar"},
		},
		output: "[  info ] some message  foo=bar\n",
	},
	{
		entry: Entry{
			Time:    TTime,
			Level:   LDebug,
			Message: []byte("some message"),
			Fields:  map[string]interface{}{"foo": "bar"},
		},
		flags:  log.Ldate,
		output: "[ debug ] 2015/08/01 some message  foo=bar\n",
	},
	{
		entry: Entry{
			Time:    TTime,
			Level:   LAlert,
			Message: []byte("some message"),
			Fields:  map[string]interface{}{"foo": "bar"},
		},
		flags:  log.Ldate | log.Lmicroseconds,
		output: "[ alert ] 2015/08/01 20:45:30.000009 some message  foo=bar\n",
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
		output: "[ trace ] 2015/08/01 20:45:30.000009 file.go:42: some message  foo=bar\n",
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
		flags:  log.Ldate | log.Llongfile,
		output: "[  warn ] 2015/08/01 /src/file.go:142: some message  foo=bar\n",
	},
	{
		entry: Entry{
			Time:    TTime,
			Level:   LWarning,
			File:    "/src/file.go",
			Line:    142,
			Prefix:  "worker1 ",
			Message: []byte("some message"),
			Fields:  map[string]interface{}{"foo": "bar"},
		},
		flags:  log.Ldate | log.Llongfile,
		output: "[  warn ] worker1 2015/08/01 /src/file.go:142: some message  foo=bar\n",
	},
	{
		entry: Entry{
			Time:    TTime,
			Level:   LDebug,
			Message: []byte("some message"),
			Fields:  map[string]interface{}{"foo": "bar"},
		},
		colors: true,
		width:  40,
		flags:  log.Ldate,
		output: "[ \x1b[0;36mdebug\x1b[0m ] 2015/08/01 some message  \033[1mfoo\033[0m=bar\n",
	},
	{
		entry: Entry{
			Time:    TTime,
			Level:   LDebug,
			Message: []byte("some message"),
			Fields:  map[string]interface{}{"foo": "bar"},
		},
		colors: true,
		width:  140,
		flags:  log.Ldate,
		output: "[ \x1b[0;36mdebug\x1b[0m ] 2015/08/01 some message                                             \033[1mfoo\033[0m=bar\n",
	},
	{
		entry: Entry{
			Time:    TTime,
			Level:   LDebug,
			File:    "/src/file.go",
			Line:    142,
			Message: []byte("some message"),
			Fields:  map[string]interface{}{"foo": "bar"},
		},
		colors: true,
		width:  140,
		flags:  log.Llongfile,
		output: "[ \x1b[0;36mdebug\x1b[0m ] \x1b[1;30m/src/file.go:142:\x1b[0m some message                           \033[1mfoo\033[0m=bar\n",
	},
}

func TestStdFormatter(t *testing.T) {
	for _, tt := range formatterTests {
		f := StdFormatter{
			Flag:   tt.flags,
			Colors: tt.colors,
		}

		if tt.width > 0 {
			terminalWidth = fixedWidthTerminal(tt.width)
		}

		b, err := f.Format(&tt.entry)
		if err != nil {
			t.Fatal(err)
		}

		if string(b) != tt.output {
			t.Errorf("Unexpected formatter output: %s", b)
		}
	}
}

// stub for terminal width function
func fixedWidthTerminal(width int) func(int) int {
	return func(fd int) int {
		return width
	}
}
