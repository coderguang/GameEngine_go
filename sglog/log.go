package sglog

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/coderguang/GameEngine_go/sgdef"

	"github.com/coderguang/GameEngine_go/sgfile"
	"github.com/coderguang/GameEngine_go/sgtime"
)

var globalLogger *logger

func init() {

}

//==============export function======

func NewLogger(strLevel string, pathname string, flag int, isConsole bool) error {

	if globalLogger != nil {
		return errors.New("logger already exist,don't not init second times")
	}

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
		return errors.New("unknow level:" + strLevel)
	}

	//new
	globalLogger = new(logger)
	globalLogger.level = level
	globalLogger.pathname = pathname
	globalLogger.flag = flag
	globalLogger.chanBuff = make(chan *logData, 1000)
	globalLogger.status = sgdef.DefServerStatusInit
	globalLogger.console = isConsole

	now := sgtime.New()
	globalLogger.createNewFile(now)

	go globalLogger.loopWriteLog()

	time.Sleep(time.Duration(100) * time.Millisecond)

	return nil

}

func (logger *logger) close() {
	logger.status = sgdef.DefServerStatusStop

	for {
		if len(logger.chanBuff) == 0 {
			logger.onceCloseFunc.Do(func() {
				close(logger.chanBuff)
				if logger.baseFile != nil {
					logger.baseFile.Close()
				}
				logger.baseFile = nil
			})
			break
		}
	}
}

func IsStop() bool {
	if globalLogger == nil {
		return true
	}
	if globalLogger.status == sgdef.DefServerStatusStop {
		return true
	}
	return false
}

func IsRunning() bool {
	if globalLogger == nil {
		return false
	}
	if globalLogger.status != sgdef.DefServerStatusRunning {
		return false
	}
	return true
}

func CloseGlobalLogger() {
	if globalLogger != nil {
		globalLogger.close()
	}
}

func Debug(a ...interface{}) {
	if !IsRunning() {
		log.Println("sglog.Debug", a)
		return
	}
	logData := newLogData(debugLevel, a...)
	globalLogger.addLogData(logData)
}

func Info(a ...interface{}) {
	if !IsRunning() {
		log.Println("sglog.Info", a)
		return
	}
	logData := newLogData(infoLevel, a...)
	globalLogger.addLogData(logData)
}

func Error(a ...interface{}) {
	if !IsRunning() {
		log.Println("sglog.Error", a)
		return
	}
	logData := newLogData(errorLevel, a...)
	globalLogger.addLogData(logData)
}

func Fatal(a ...interface{}) {
	if !IsRunning() {
		log.Println("sglog.Fatal", a)
		return
	}
	logData := newLogData(fatalLevel, a...)
	globalLogger.addLogData(logData)
}

//==========inner function

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

type logData struct {
	level int
	dt    *sgtime.DateTime
	data  string
}

type logger struct {
	level         int
	baseFile      *os.File
	console       bool
	dt            *sgtime.DateTime
	pathname      string
	flag          int
	chanBuff      chan *logData
	onceCloseFunc sync.Once
	status        sgdef.DefServerStatus
}

func newLogData(lv int, a ...interface{}) *logData {
	logData := new(logData)
	logData.level = lv
	logData.dt = sgtime.New()
	logData.data = fmt.Sprintln(a...)
	return logData
}

func (logger *logger) addLogData(logData *logData) {
	logger.chanBuff <- logData
}

func getLevelStr(level int) string {
	str := "unkonw"
	switch level {
	case debugLevel:
		str = yellow(printDebugLevel)
	case infoLevel:
		str = green(printInfoLevel)
	case errorLevel:
		str = red(printErrorLevel)
	case fatalLevel:
		str = red(printFatalLevel)
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

func getFileName() string {
	now := sgtime.New()
	str := sgtime.YearString(now) + sgtime.MonthString(now) + sgtime.DayString(now)
	str += ".log"
	return str
}

func (logger *logger) write(logData *logData) {
	if logData.level < logger.level {
		return
	}
	str := sgtime.LogString(logData.dt) + " " + getLevelStrRaw(logData.level) + " " + logData.data
	logger.baseFile.WriteString(str)
	if logger.console {
		strEx := sgtime.LogString(logData.dt) + " " + getLevelStr(logData.level) + " " + logData.data
		fmt.Print(strEx)
	}
}

func (logger *logger) checkAndSwapLogger() {
	now := sgtime.New()
	if sgtime.GetTotalDay(logger.dt) != sgtime.GetTotalDay(now) {
		logger.createNewFile(now)
	}
}

func (logger *logger) createNewFile(now *sgtime.DateTime) {
	if logger.baseFile != nil {
		logger.baseFile.Close()
	}
	logger.baseFile = nil
	filename := getFileName()
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

func (logger *logger) loopWriteLog() {
	logger.status = sgdef.DefServerStatusRunning

	for logData := range logger.chanBuff {
		logger.checkAndSwapLogger()
		logger.write(logData)
	}
}
