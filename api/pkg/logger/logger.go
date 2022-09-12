package logger

import (
	"log"
	"sync"
)

type (
	Logger interface {
		Infof(format string, args ...any)
		Warnf(format string, args ...any)
		Errorf(format string, args ...any)
		Fatalln(args ...any)
		ExtraFields(keyValues Fields) Logger
		ExtraField(key string, value any) Logger
		ExtraError(err error) Logger
	}
	Fields map[string]any

	Config struct {
		Format  string
		FileLog string
		Level   string
	}
)

var (
	instance Logger
	once     sync.Once
	config   Config
)

func InitConfig(c Config) Logger {
	config = c
	return GetLogger()
}

func GetLogger() Logger {
	once.Do(func() {
		log.Println("Start Logger")
		instance = NewLogrus(config)
	})

	return instance
}
