package task

import (
	"github.com/jakhog/cbs-cisco-agent/log"
)

type Task interface {
	Name() string
	Run(*log.Logger)
}
