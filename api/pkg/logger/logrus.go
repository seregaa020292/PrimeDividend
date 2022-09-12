package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

type Logrus struct {
	*logrus.Entry
}

func NewLogrus(config Config) Logger {
	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		log.Fatalln(err)
	}

	logNew := logrus.New()

	logNew.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint:     true,
		TimestampFormat: config.Format,
	})

	logNew.SetLevel(level)
	logNew.SetOutput(io.MultiWriter( /*os.Stdout,*/ outputFile(config.FileLog)))

	logEntry := &Logrus{logrus.NewEntry(logNew)}

	return logEntry
}

func (l *Logrus) Infof(format string, args ...interface{}) {
	fields := callerPrettyfier(runtime.Caller(1))
	l.Entry.WithFields(fields).Infof(format, args...)
}

func (l *Logrus) Warnf(format string, args ...interface{}) {
	fields := callerPrettyfier(runtime.Caller(1))
	l.Entry.WithFields(fields).Warnf(format, args...)
}

func (l *Logrus) Errorf(format string, args ...interface{}) {
	fields := callerPrettyfier(runtime.Caller(1))
	l.Entry.WithFields(fields).Errorf(format, args...)
}

func (l *Logrus) Fatalln(args ...interface{}) {
	fields := callerPrettyfier(runtime.Caller(1))
	l.Entry.WithFields(fields).Fatalln(args...)
}

func (l *Logrus) ExtraFields(fields Fields) Logger {
	return &Logrus{l.Entry.WithFields(convertToFields(fields))}
}

func (l *Logrus) ExtraField(key string, value interface{}) Logger {
	return &Logrus{l.Entry.WithField(key, value)}
}

func (l *Logrus) ExtraError(err error) Logger {
	return &Logrus{l.Entry.WithError(err)}
}

func callerPrettyfier(pc uintptr, file string, line int, ok bool) logrus.Fields {
	if !ok {
		return logrus.Fields{}
	}

	return logrus.Fields{
		"file": fmt.Sprintf("%s:%d", path.Clean(file), line),
		"func": fmt.Sprintf("%s()", runtime.FuncForPC(pc).Name()),
	}
}

func convertToFields(fields Fields) logrus.Fields {
	return logrus.Fields(fields)
}

func outputFile(fileName string) *os.File {
	if err := os.MkdirAll(path.Dir(fileName), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
