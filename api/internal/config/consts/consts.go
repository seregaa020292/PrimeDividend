package consts

import "time"

const (
	ServerAddr = ":3000"

	TimestampFormat = "02.01.2006 15:04:05"
	Timezone        = "Europe/Moscow"

	TmpLog   = "./tmp/logs/server.log"
	LevelLog = "debug"

	TimeoutShutdown = 5 * time.Second
)
