package logger

import (
	"os"
	"sync"
	"testing"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var isWriteLogToFiles = false

func enableWriteLogToFiles() {
	isWriteLogToFiles = true
}

func disableWriteLogToFiles() {
	isWriteLogToFiles = false
}

func newEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func newLogWriter() zapcore.WriteSyncer {
	syncConsole := zapcore.AddSync(os.Stderr)
	if !isWriteLogToFiles {
		return zapcore.AddSync(syncConsole)
	}
	rollingFilesCfg := &lumberjack.Logger{
		Filename:   "./log/test.log",
		MaxSize:    1,
		MaxBackups: 2,
		MaxAge:     1,
		Compress:   true,
	}
	return zapcore.NewMultiWriteSyncer(syncConsole, zapcore.AddSync(rollingFilesCfg))
}

func Test_Logger_LogLevel(t *testing.T) {
	atom := zap.NewAtomicLevel()
	core := zapcore.NewCore(newEncoder(), newLogWriter(), atom)
	logger := zap.New(core, zap.AddCaller()).Sugar().Named("Collector")

	atom.SetLevel(zapcore.ErrorLevel)
	logger.Debugf("a:%v, %v ", 1, "b")
	logger.Infof("a:%v, %v ", 1, "b")
	logger.Warnf("a:%v, %v ", 1, "b")
	logger.Errorf("a:%v, %v ", 1, "b")

	atom.SetLevel(zapcore.DebugLevel)
	logger.Debugf("a:%v, %v ", 1, "b")
	logger.Infof("a:%v, %v ", 1, "b")
	logger.Warnf("a:%v, %v ", 1, "b")
	logger.Errorf("a:%v, %v ", 1, "b")
}

func Test_Logger_WriteToFiles(t *testing.T) {
	enableWriteLogToFiles()
	defer disableWriteLogToFiles()

	atom := zap.NewAtomicLevel()

	core := zapcore.NewCore(newEncoder(), newLogWriter(), atom)
	logger := zap.New(core, zap.AddCaller()).Sugar().Named("Collector")

	for i := 0; i <= 5000; i++ {
		logger.Debugf("a:%v, %v ", 1, "b")
		logger.Infof("a:%v, %v ", 1, "b")
		logger.Warnf("a:%v, %v ", 1, "b")
		logger.Errorf("a:%v, %v ", 1, "b")
	}
}

type ID struct {
	id int
	sync.Mutex
}

func (i *ID) GetID() int {
	i.Mutex.Lock()
	res := i.id
	i.id++
	i.Mutex.Unlock()
	return res
}

func Test_Logger_Concurrent(t *testing.T) {
	enableWriteLogToFiles()
	defer disableWriteLogToFiles()

	atom := zap.NewAtomicLevel()

	core := zapcore.NewCore(newEncoder(), newLogWriter(), atom)
	logger := zap.New(core, zap.AddCaller()).Sugar().Named("Collector")

	wg := &sync.WaitGroup{}
	id := &ID{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			pid := id.GetID()
			logger.Debugf("pid: %v, msg: %v", pid, "aaa")
			logger.Infof("pid: %v, msg: %v", pid, "aaa")
			logger.Warnf("pid: %v, msg: %v", pid, "aaa")
			logger.Errorf("pid: %v, msg: %v", pid, "aaa")
		}()
	}
	wg.Wait()
}
