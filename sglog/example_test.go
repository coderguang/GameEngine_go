package sglog

import (
	"log"
	"testing"
)

func Test_NewLog(t *testing.T) {
	logData := NewLogData(debugLevel, "hello", "hi")
	if logData != nil {
		t.Log("new log data ok")
	} else {
		t.Error("new log data error")
	}
}

func Test_LogPackage(t *testing.T) {

	logger, err := NewLogger("debug", "../../log/", log.LstdFlags, true)
	if err != nil {
		t.Error("init logger failed")
	}
	Swap(logger)
	go LoopLogServer()
	Debug("hello,debug")
	// Info("hello info")
	// Error("hello error")
	// Fatal("hello fatal")
	CloseGlobalLogger()
	t.Log("logger test ok")
}

func BenchmarkLogger(b *testing.B) {
	logger, err := NewLogger("debug", "../../log/", log.LstdFlags, true)
	if err != nil {
		return
	}

	defer func() {
		CloseGlobalLogger()
	}()

	Swap(logger)
	go LoopLogServer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Debug("hello,debug", i)
	}
}
