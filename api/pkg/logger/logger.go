package logger

import (
	"log"
	"sync"
)

type (
	Logger interface {
		Infof(format string, args ...interface{})
		Warnf(format string, args ...interface{})
		Errorf(format string, args ...interface{})
		Fatalln(args ...interface{})
		ExtraFields(keyValues Fields) Logger
		ExtraField(key string, value interface{}) Logger
		ExtraError(err error) Logger
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
	config   Config
)

func SetConfig(c Config) {
	log.Println("Config Logger")
	config = c
}

func GetLogger() Logger {
	once.Do(func() {
		log.Println("Start Logger")
		instance = NewLogrus(config)
	})

	return instance
}
