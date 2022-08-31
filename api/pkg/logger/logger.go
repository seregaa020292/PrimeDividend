package logger

import (
	"log"
	"sync"

	"primedivident/internal/config/consts"
)

type (
	Logger interface {
		Infof(format string, args ...interface{})
		Warnf(format string, args ...interface{})
		Errorf(format string, args ...interface{})
		Fatalln(args ...interface{})
		ExtraFields(keyValues Fields) Logger
	}
	Fields map[string]interface{}

	Config struct {
		Format  string
		FileLog string
		Level   string
	}
)

var (
	instance Logger
	once     sync.Once
)

func GetLogger() Logger {
	once.Do(func() {
		log.Println("Start Logger")

		instance = NewLogrus(Config{
			Format:  consts.TimestampFormat,
			FileLog: consts.TmpLog,
			Level:   consts.LevelLog,
		})
	})

	return instance
}
