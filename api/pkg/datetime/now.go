package datetime

import (
	"log"
	"sync"
	"time"

	"github.com/jinzhu/now"

	"primedividend/api/internal/config/consts"
)

var (
	instance *now.Now
	once     sync.Once
)

func GetNow() *now.Now {
	once.Do(func() {
		location, err := time.LoadLocation(consts.Timezone)
		if err != nil {
			log.Fatal(err)
		}

		nowConfig := &now.Config{
			WeekStartDay: time.Monday,
			TimeLocation: location,
		}

		instance = nowConfig.With(time.Now())
	})

	return instance
}
