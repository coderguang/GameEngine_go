package sglog

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/coderguang/GameEngine_go/sgerror"
	"github.com/coderguang/GameEngine_go/sgfile"
	"github.com/coderguang/GameEngine_go/sgtime"
)

// levels

func init() {

}

const (
	debugLevel = 0
	infoLevel  = 1
	errorLevel = 2
	fatalLevel = 3
)

const (
	printDebugLevel = "[debug] "
	printInfoLevel  = "[info ] "
	printErrorLevel = "[error] "
	printFatalLevel = "[fatal] "
)

type LogData struct {
	level int
	dt    *sgtime.DateTime
	data  string
}

func NewLogData(lv int, format string, a ...interface{}) *LogData {
	logData := new(LogData)
	logData.level = lv
	logData.dt = sgtime.New()
	logData.data = fmt.Sprintf(format, a...)
	return logData
}

type Logger struct {
	level     int
	baseFile  *os.File
	console   bool
	dt        *sgtime.DateTime
	pathname  string
	flag      int
	buff      chan *LogData
	onceClose sync.Once
	isStop    bool
}

func NewLogger(strLevel string, pathname string, flag int, isConsole bool) (*Logger, error) {
	var level int
	switch strings.ToLower(strLevel) {
	case "debug":
		level = debugLevel
	case "info":
		level = infoLevel
	case "error":
		level = errorLevel
	case "fatal":
		level = fatalLevel
	default:
		return nil, errors.New("unknow level:" + strLevel)
	}
	var baseFile *os.File

	now := sgtime.New()
	filename := fmt.Sprintf("%d%02d%02d",
		now.Year(),
		now.Month(),
		now.Day())
	filename += ".log"
	pathExist, _ := sgfile.PathExists(pathname)
	if !pathExist {
		sgfile.MkdirAll(pathname, os.ModePerm)
	}

	file, err := sgfile.Create(path.Join(pathname, filename))
	if err != nil {
		return nil, err
	}
	baseFile = file

	//new
	logger := new(Logger)
	logger.level = level
	logger.baseFile = baseFile
	logger.dt = sgtime.New()
	logger.pathname = pathname
	logger.flag = flag
	logger.InitChan()
	logger.isStop = false
	logger.console = isConsole
	return logger, nil

}

func (logger *Logger) InitChan() {
	logger.buff = make(chan *LogData, 1000)
}

func LoopLogServer() {
	for {
		logData := <-globalLogger.buff
		checkAndSwapLogger(globalLogger)
		globalLogger.Write(logData)
		if globalLogger.isStop && len(globalLogger.buff) <= 0 {
			globalLogger.onceClose.Do(func() {
				close(globalLogger.buff)
			})
			return
		}
	}
}

func (logger *Logger) AddData(logData *LogData) {
	logger.buff <- logData
}

func (logger *Logger) Close() {
	logger.isStop = true
}

func getLevelStr(level int) string {
	str := "unkonw"
	switch level {
	case debugLevel:
		str = Yellow(printDebugLevel)
	case infoLevel:
		str = Green(printInfoLevel)
	case errorLevel:
		str = Red(printErrorLevel)
	case fatalLevel:
		str = Red(printFatalLevel)
	}
	return str
}
func getLevelStrRaw(level int) string {
	str := "unkonw"
	switch level {
	case debugLevel:
		str = printDebugLevel
	case infoLevel:
		str = printInfoLevel
	case errorLevel:
		str = printErrorLevel
	case fatalLevel:
		str = printFatalLevel
	}
	return str
}

func (logger *Logger) Write(logData *LogData) {
	if logData.level < logger.level {
		return
	}
	str := sgtime.LogString(logData.dt) + " " + getLevelStrRaw(logData.level) + " " + logData.data + "\n"
	logger.baseFile.WriteString(str)
	if logger.console {
		strEx := sgtime.LogString(logData.dt) + " " + getLevelStr(logData.level) + " " + logData.data + "\n"
		fmt.Print(strEx)
	}
	if logData.level == fatalLevel {
		os.Exit(sgerror.GameEngine_Fatal_Log)
	}
}

var globalLogger *Logger

func Swap(logger *Logger) {
	if logger != nil {
		globalLogger = logger
	}
}

func CloseGlobalLogger() {
	if globalLogger != nil {
		globalLogger.Close()
	}
}

func checkAndSwapLogger(logger *Logger) {
	now := sgtime.New()
	if sgtime.GetTotalDay(logger.dt) != sgtime.GetTotalDay(now) {
		if globalLogger.baseFile != nil {
			globalLogger.baseFile.Close()
		}
		globalLogger.baseFile = nil
		filename := fmt.Sprintf("%d%02d%02d",
			now.Year(),
			now.Month(),
			now.Day())
		filename += ".log"
		pathExist, _ := sgfile.PathExists(logger.pathname)
		if !pathExist {
			sgfile.MkdirAll(logger.pathname, os.ModePerm)
		}

		file, err := sgfile.Create(path.Join(logger.pathname, filename))
		if err != nil {
			return
		}
		logger.baseFile = file
		logger.dt = now
	}
}

func Debug(format string, a ...interface{}) {
	if globalLogger.isStop {
		return
	}
	logData := NewLogData(debugLevel, format, a...)
	globalLogger.AddData(logData)
}

func Info(format string, a ...interface{}) {
	if globalLogger.isStop {
		return
	}
	logData := NewLogData(infoLevel, format, a...)
	globalLogger.AddData(logData)
}

func Error(format string, a ...interface{}) {
	if globalLogger.isStop {
		return
	}
	logData := NewLogData(errorLevel, format, a...)
	globalLogger.AddData(logData)
}

func Fatal(format string, a ...interface{}) {
	if globalLogger.isStop {
		return
	}
	logData := NewLogData(fatalLevel, format, a...)
	globalLogger.AddData(logData)
}
