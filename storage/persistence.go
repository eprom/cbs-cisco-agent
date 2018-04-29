package storage

import (
	"github.com/jakhog/cbs-cisco-agent/log"
	"sync"
)

var persistedStorage Storage

var inMemoryStorage Storage
var inMemoryStorageDirty bool

var storageLock sync.Mutex

type Task struct {
}

func NewTask() *Task {
	return &Task{}
}

func (*Task) Name() string {
	return "Storage"
}

func (*Task) Run(logger *log.Logger) {
	storageLock.Lock()
	defer storageLock.Unlock()

	logger.Info(inMemoryStorage)
}
