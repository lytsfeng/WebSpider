package base

import "logging"

func NewLogger() logging.Logger  {
	return logging.NewSimpleLogger()
}
