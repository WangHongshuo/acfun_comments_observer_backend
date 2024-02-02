package logger

import (
	"sync"
	"testing"

	_ "github.com/WangHongshuo/acfun_comments_observer_backend/cfg"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Test_Logger_LogLevel(t *testing.T) {
	atom := zap.NewAtomicLevel()
	core := zapcore.NewCore(newEncoder(), newLogWriter(false), atom)
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

	atom := zap.NewAtomicLevel()

	core := zapcore.NewCore(newEncoder(), newLogWriter(true), atom)
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

	atom := zap.NewAtomicLevel()

	core := zapcore.NewCore(newEncoder(), newLogWriter(false), atom)
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
