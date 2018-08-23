package colog

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

type outputTest struct {
	in  string
	out string
}

var outputTests = []outputTest{
	{"trace: Trace should be plain %s", "[ trace ] Trace should be plain %s\n"},
	{"debug: Debug should be blue %s", "[ debug ] Debug should be blue %s\n"},
	{"info: Info should be green %s", "[  info ] Info should be green %s\n"},
	{"warning: Warning should be yellow %s", "[  warn ] Warning should be yellow %s\n"},
	{"error: Error should be red %s", "[ error ] Error should be red %s\n"},
	{"alert: Alert should be white on red %s", "[ alert ] Alert should be white on red %s\n"},
}

func TestColors(t *testing.T) {

	log.SetFlags(log.LstdFlags)
	Register()

	SetMinLevel(LTrace)
	SetDefaultLevel(LTrace)

	for _, tt := range outputTests {
		tt.in = fmt.Sprintf(tt.in, "")
		log.Println(tt.in)
	}
}

func TestDefaultLevel(t *testing.T) {

	tw := new(mockWriter)
	log.SetFlags(0)
	Register()
	SetOutput(tw)
	SetFormatter(&StdFormatter{Colors: false})
	SetDefaultLevel(LDebug)

	log.Println("no prefix text")
	if "[ debug ] no prefix text\n" != tw.String() {
		t.Fatalf("Default level failed: %s", tw.String())
	}
}

func TestMinLevel(t *testing.T) {

	tw := new(mockWriter)
	log.SetFlags(0)
	Register()
	SetOutput(tw)
	SetFormatter(&StdFormatter{Colors: false})
	SetMinLevel(LInfo)

	log.Println("debug: should not print")
	if "" != tw.String() {
		t.Fatal("Min level failed")
	}
}

func TestMinDefaultLevel(t *testing.T) {

	tw := new(mockWriter)
	log.SetFlags(0)
	Register()
	SetOutput(tw)
	SetFormatter(&StdFormatter{Colors: false})

	SetMinLevel(LDebug)
	SetDefaultLevel(LDebug)
	log.Println("no prefix text")
	if "[ debug ] no prefix text\n" != tw.String() {
		t.Fatalf("Default level failed: %s", tw.String())
	}

	SetMinLevel(LInfo)
	log.Println("should not print")
	if "[ debug ] no prefix text\n" != tw.String() {
		t.Fatalf("Default level failed: %s", tw.String())
	}
}

func TestPrefix(t *testing.T) {

	tw := new(mockWriter)
	cl := NewCoLog(tw, "abc", 0)
	cl.SetFormatter(&StdFormatter{Colors: false})

	logger := cl.NewLogger()
	logger.Println("debug: some text")
	if "[ debug ] abcsome text\n" != tw.String() {
		t.Fatalf("Prefix output failed: %s", tw.String())
	}
}

func TestSimpleOutput(t *testing.T) {

	tw := new(mockWriter)
	cl := NewCoLog(tw, "", 0)
	cl.SetFormatter(&StdFormatter{Colors: false})
	logger := cl.NewLogger()
	for k, tt := range outputTests {

		seq := randSeq(k)
		tt.in = fmt.Sprintf(tt.in, seq)
		tt.out = fmt.Sprintf(tt.out, seq)

		logger.Println(tt.in)
		if tt.out != tw.String() {
			t.Fatalf("Simple output not found:\n %s\n %s", tt.out, string(tw.Data))
		}
	}
}

func TestCustomHeaders(t *testing.T) {

	customHeaders := HeaderMap{
		"(meh) ":    LDebug,
		"watchout ": LWarning,
	}

	tw := new(mockWriter)
	cl := NewCoLog(tw, "", 0)
	cl.SetFormatter(&StdFormatter{Colors: false})
	cl.SetHeaders(customHeaders)
	cl.AddHeader("oh no! ", LAlert)

	logger := cl.NewLogger()

	logger.Println("(meh) nothing important")
	if "[ debug ] nothing important\n" != tw.String() {
		t.Errorf("message malformed after custom header %s", tw.String())
	}

	logger.Println("watchout look at this")
	if "[  warn ] look at this\n" != tw.String() {
		t.Errorf("message malformed after custom header %s", tw.String())
	}

	logger.Println("oh no! something terrible happened")
	if "[ alert ] something terrible happened\n" != tw.String() {
		t.Errorf("message malformed after custom header %s", tw.String())
	}
}

func TestOutputRace(t *testing.T) {

	wg := sync.WaitGroup{}
	wg.Add(1)

	go testStdLoggerOutput(t, &wg)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go testNewLoggerOutput(t, &wg)
	}

	wg.Wait()
}

func testStdLoggerOutput(t *testing.T, wg *sync.WaitGroup) {

	tb := new(mockBufferWriter)
	tb.Data = make(map[string][]byte, len(outputTests))
	log.SetFlags(0)
	Register()
	SetOutput(tb)
	SetMinLevel(LTrace)
	SetDefaultLevel(LTrace)
	SetFormatter(&StdFormatter{Colors: false})

	for k, tt := range outputTests {
		wg.Add(1)
		go func(tt outputTest, k int) {

			seq := randSeq(k)
			tt.in = fmt.Sprintf(tt.in, seq)
			tt.out = fmt.Sprintf(tt.out, seq)

			log.Println(tt.in)
			if !tb.IsWritten(tt.out) {
				t.Errorf("Raced std output not found: %s", tt.out)
			}
			wg.Done()
		}(tt, k)
	}

	wg.Done()
}

func testNewLoggerOutput(t *testing.T, wg *sync.WaitGroup) {

	tb := new(mockBufferWriter)
	tb.Data = make(map[string][]byte, len(outputTests))
	cl := NewCoLog(tb, "", 0)
	cl.SetFormatter(&StdFormatter{Colors: false})
	logger := cl.NewLogger()

	for k, tt := range outputTests {
		wg.Add(1)
		go func(tt outputTest, k int) {

			seq := randSeq(k)
			tt.in = fmt.Sprintf(tt.in, seq)
			tt.out = fmt.Sprintf(tt.out, seq)

			logger.Println(tt.in)
			if !tb.IsWritten(tt.out) {
				t.Errorf("Raced logger output not found: %s", tt.out)
			}
			wg.Done()
		}(tt, k)
	}

	wg.Done()
}

func TestHooks(t *testing.T) {

	h1 := &mockHook{levels: []Level{LDebug, LInfo}}
	h2 := &mockHook{levels: []Level{LInfo, LWarning}}
	tw := new(mockWriter)
	cl := NewCoLog(tw, "", 0)
	cl.ParseFields(true)
	cl.AddHook(h1)
	cl.AddHook(h2)
	logger := cl.NewLogger()

	logger.Println("debug: debug message foo=d")
	if h1.entry == nil {
		t.Fatal("h1 not written for debug")
	}

	if string(h1.entry.Message) != "debug message" || h1.entry.Fields["foo"] != "d" {
		t.Fatal("h1 hook got incorrect data")
	}

	if h2.entry != nil {
		t.Fatal("h2 written for debug")
	}

	h1.entry = nil
	h2.entry = nil

	logger.Println("info: info message foo=i")
	if h1.entry == nil {
		t.Fatal("h1 not written for info")
	}

	if string(h1.entry.Message) != "info message" || h1.entry.Fields["foo"] != "i" {
		t.Fatal("h1 hook got incorrect data")
	}

	if h2.entry == nil {
		t.Fatal("h2 not written for info")
	}

	if string(h2.entry.Message) != "info message" || h2.entry.Fields["foo"] != "i" {
		t.Fatal("h2 hook got incorrect data")
	}

	if h1.entry.Time != h2.entry.Time {
		t.Fatal("h1 and h2 saw different timestamps")
	}

	if h1.entry.Time.After(time.Now()) {
		t.Fatal("received invalid future timestamp")
	}

	if h1.entry.Time.Before(time.Now().Add(-1 * time.Millisecond)) {
		t.Fatal("received timestamp older than 1 millisecond")
	}

	h1.entry = nil
	h2.entry = nil

	logger.Println("warning: warning message foo=w")
	if h1.entry != nil {
		t.Fatal("h1 written for warning")
	}

	if h2.entry == nil {
		t.Fatal("h2 no written for warning")
	}

	if string(h2.entry.Message) != "warning message" || h2.entry.Fields["foo"] != "w" {
		t.Fatal("h2 hook got incorrect data")
	}

	h1.entry = nil
	h2.entry = nil

	logger.Println("error: error message foo=e")
	if h1.entry != nil {
		t.Fatal("h1 written for error")
	}

	if h2.entry != nil {
		t.Fatal("h2 written for error")
	}
}

func TestFixedValues(t *testing.T) {

	hook := &mockHook{levels: []Level{LDebug, LInfo}}
	tw := new(mockWriter)
	cl := NewCoLog(tw, "", 0)
	cl.ParseFields(true)
	cl.FixedValue("key", "val")
	cl.FixedValue("foo", "bar")
	cl.AddHook(hook)
	logger := cl.NewLogger()

	logger.Println("debug: debug message")
	if hook.entry == nil {
		t.Fatal("h1 not written for debug")
	}

	if len(hook.entry.Fields) != 2 {
		t.Errorf("Wrong number of fixed values %d", len(hook.entry.Fields))
	}

	if string(hook.entry.Message) != "debug message" || hook.entry.Fields["key"] != "val" {
		t.Fatal("Missing or incorrect fixed value foo")
	}

	if string(hook.entry.Message) != "debug message" || hook.entry.Fields["foo"] != "bar" {
		t.Fatal("Missing or incorrect fixed value foo")
	}
}

type mockWriter struct {
	mux  sync.Mutex
	Data []byte
}

func (tw *mockWriter) Write(p []byte) (n int, err error) {
	tw.mux.Lock()
	defer tw.mux.Unlock()
	tw.Data = p
	return len(p), nil
}

func (tw *mockWriter) String() string {
	tw.mux.Lock()
	defer tw.mux.Unlock()
	return string(tw.Data)
}

type mockBufferWriter struct {
	mux  sync.Mutex
	Data map[string][]byte
}

func (tb *mockBufferWriter) Write(p []byte) (n int, err error) {
	tb.mux.Lock()
	defer tb.mux.Unlock()
	tb.Data[string(p)] = p
	return len(p), nil
}

func (tb *mockBufferWriter) IsWritten(s string) bool {
	tb.mux.Lock()
	defer tb.mux.Unlock()
	_, ok := tb.Data[s]
	return ok
}

type mockHook struct {
	entry  *Entry
	levels []Level
}

func (h *mockHook) Levels() []Level {
	return h.levels
}

func (h *mockHook) Fire(e *Entry) error {
	h.entry = e
	return nil
}

func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
