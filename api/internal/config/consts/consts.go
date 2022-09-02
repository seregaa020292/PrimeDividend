package consts

import "time"

const (
	ServerAddr = ":3000"

	RequestLimit = 100
	WindowLength = 1 * time.Minute

	TimestampFormat = "02.01.2006 15:04:05"
	Timezone        = "Europe/Moscow"

	TmpLog = "./tmp/logs/server.log"

	TimeoutShutdown = 10 * time.Second
)
