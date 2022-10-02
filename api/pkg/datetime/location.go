package datetime

import (
	"log"
	"time"

	"primedivident/internal/config/consts"
)

func InitLocation() {
	location, err := time.LoadLocation(consts.Timezone)
	if err != nil {
		log.Fatal(err)
	}

	time.Local = location
}
