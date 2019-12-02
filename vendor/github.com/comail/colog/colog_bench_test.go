package colog

import (
	"log"
	"testing"
)

func BenchmarkLoggerNoFlags(b *testing.B) {
	logger := log.New(new(nilWriter), "", 0)
	for i := 0; i <= b.N; i++ {
		logger.Println("error: benchmark")
	}
}

func BenchmarkLoggerWithFlags(b *testing.B) {
	logger := log.New(new(nilWriter), "", log.Llongfile|log.Ldate)
	for i := 0; i <= b.N; i++ {
		logger.Println("error: benchmark")
	}
}

func BenchmarkCoLogPlainNoFlags(b *testing.B) {
	doBench(b, "error: benchmark", 0, false, false)
}

func BenchmarkCoLogPlainWithFlags(b *testing.B) {
	doBench(b, "error: benchmark", log.Llongfile|log.Ldate, false, false)
}

func BenchmarkCoLogPlainWithFlagsAndFields(b *testing.B) {
	doBench(b, "error: benchmark KeyName1='Key Value1'", log.Llongfile|log.Ldate, true, false)
}

func BenchmarkCoLogColorNoFlags(b *testing.B) {
	doBench(b, "error: benchmark", 0, false, true)
}

func BenchmarkCoLogColorWithFlags(b *testing.B) {
	doBench(b, "error: benchmark", log.Llongfile|log.Ldate, false, true)
}

func BenchmarkCoLogColorWithFlagsAndFields(b *testing.B) {
	doBench(b, "error: benchmark KeyName1='Key Value1'", log.Llongfile|log.Ldate, true, true)
}

func doBench(b *testing.B, message string, flag int, fields bool, colors bool) {
	cl := NewCoLog(new(nilWriter), "", flag)
	cl.ParseFields(fields)
	cl.SetFormatter(&StdFormatter{Flag: flag, Colors: colors})
	//	fmt.Printf("flags %d \n", cl.formatter.Flags())
	logger := cl.NewLogger()
	for i := 0; i <= b.N; i++ {
		logger.Println(message)
	}
}

type nilWriter struct{}

func (nw *nilWriter) Write(p []byte) (n int, err error) { return 0, nil }
