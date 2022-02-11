package dependency

import "log"

func NewLogger() *log.Logger {
	return log.Default()
}
