package logger

import (
	"fmt"
	"io"
	"log"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"

	"primedivident/pkg/utils"
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
		PrettyPrint:      true,
		DisableTimestamp: true,
		DataKey:          "detail",
	})

	outputFile, err := utils.OpenOrCreateFile(config.FileLog)
	if err != nil {
		log.Fatal(err)
	}

	logNew.SetLevel(level)
	logNew.SetOutput(io.MultiWriter( /*os.Stdout,*/ outputFile))

	logEntry := &Logrus{logrus.NewEntry(logNew)}

	return logEntry
}

func (l *Logrus) Infof(format string, args ...any) {
	fields := callerPrettyfier(runtime.Caller(1))
	l.Entry.WithFields(fields).Infof(format, args...)
}

func (l *Logrus) Warnf(format string, args ...any) {
	fields := callerPrettyfier(runtime.Caller(1))
	l.Entry.WithFields(fields).Warnf(format, args...)
}

func (l *Logrus) Errorf(format string, args ...any) {
	fields := callerPrettyfier(runtime.Caller(1))
	l.Entry.WithFields(fields).Errorf(format, args...)
}

func (l *Logrus) Fatalln(args ...any) {
	fields := callerPrettyfier(runtime.Caller(1))
	l.Entry.WithFields(fields).Fatalln(args...)
}

func (l *Logrus) ExtraFields(fields Fields) Logger {
	return &Logrus{l.Entry.WithFields(logrus.Fields(fields))}
}

func (l *Logrus) ExtraField(key string, value any) Logger {
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
