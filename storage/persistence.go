package storage

import (
	"encoding/json"
	"github.com/jakhog/cbs-cisco-agent/log"
	"io/ioutil"
	"os"
	"path"
	"sync"
	"time"
)

const STORAGE_FILE = "storage.json"

var storageFilePath string

func init() {
	storageDir := os.Getenv("CAF_APP_PERSISTENT_DIR")
	storageFilePath = path.Join(storageDir, STORAGE_FILE)
}

var persistedStorage *Storage

var inMemoryStorage *Storage
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

func (t *Task) Run(logger *log.Logger) {
	storageLock.Lock()
	defer storageLock.Unlock()

	/* ---------- Initial setup ---------- */
	if persistedStorage == nil {
		if _, err := os.Stat(storageFilePath); !os.IsNotExist(err) {
			// File already exists, try to parse it
			bytes, err := ioutil.ReadFile(storageFilePath)
			if err != nil {
				panic(err)
			}
			cleanPersited := &Storage{}
			err = json.Unmarshal(bytes, cleanPersited)
			if err == nil {
				// Contains good data, so lets use it
				cleanMemory := &Storage{}
				err = json.Unmarshal(bytes, cleanMemory)
				if err != nil {
					panic(err)
				}
				// Set the freshly read storages to current storage
				persistedStorage = cleanPersited
				inMemoryStorage = cleanMemory
				inMemoryStorageDirty = false
				logger.Info("Perstisted storage read")
				return
			} else {
				// The data is wrong, so let's move it somewhere else and continue
				backupName := "storage-" + time.Now().Format("2006-01-02T15-04-05") + ".json.bak"
				backupDir := os.Getenv("CAF_APP_PERSISTENT_DIR")
				backupPath := path.Join(backupDir, backupName)
				err := os.Rename(storageFilePath, backupPath)
				if err != nil {
					panic(err)
				}
				logger.Warning("Persisted storage contained bad data and was moved")
			}
		}
		// File doesn't exist, or contained bad data and was moved to a backup
		// So we need to create a fresh one and store that
		cleanPersited := &Storage{}
		bytes, cleanMemory, err := t.marshalAndClone(cleanPersited)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(storageFilePath, bytes, os.ModePerm)
		if err != nil {
			panic(err)
		}
		// Set the current storage to the freshly created one
		persistedStorage = cleanPersited
		inMemoryStorage = cleanMemory
		logger.Info("A clean persistant storage was created")
		return
	}

	/* ---------- New updates ---------- */
	if inMemoryStorageDirty {
		bytes, updatedPersisted, err := t.marshalAndClone(inMemoryStorage)
		if err != nil {
			panic(err)
		}
		// Try to write the file to a temporary location to not destroy old data if something goes wrong
		tempFilePath := storageFilePath + ".tmp"
		err = ioutil.WriteFile(tempFilePath, bytes, os.ModePerm)
		if err != nil {
			panic(err)
		}
		// Now overwrite the old data
		err = os.Rename(tempFilePath, storageFilePath)
		if err != nil {
			panic(err)
		}
		// New data successfully written to disk, reflect that in memory
		persistedStorage = updatedPersisted
		inMemoryStorageDirty = false
		logger.Info("Persistent storage updated")
	}
}

func (*Task) marshalAndClone(original *Storage) (bytes []byte, clone *Storage, err error) {
	bytes, err = json.MarshalIndent(original, "", "  ")
	if err != nil {
		return
	}
	clone = &Storage{}
	err = json.Unmarshal(bytes, clone)
	return
}
