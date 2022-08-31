package datetime

import (
	"sync"
	"time"

	"github.com/jinzhu/now"

	"primedivident/internal/config"
)

var (
	instance *now.Now
	once     sync.Once
)

func GetNow() *now.Now {
	once.Do(func() {
		location, _ := time.LoadLocation(config.Timezone)

		nowConfig := &now.Config{
			WeekStartDay: time.Monday,
			TimeLocation: location,
		}

		instance = nowConfig.With(time.Now())
	})

	return instance
}
