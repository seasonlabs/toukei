package checker

import (
	"time"
)

func Check(path string) {
	for {
		publish(process(path))
		time.Sleep(20 * time.Second)
	}
}