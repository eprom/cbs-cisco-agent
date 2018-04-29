package config

import (
	"crypto/md5"
	"github.com/go-ini/ini"
	"github.com/jakhog/cbs-cisco-agent/log"
	"io"
	"os"
	"strings"
)

var configFilePath string

func init() {
	// Check if we got the config file path as an ENV
	configFilePath = os.Getenv("CAF_APP_CONFIG_FILE")
	if strings.TrimSpace(configFilePath) == "" {
		configFilePath = CONFIG_FILE
	}
}

type ReaderTask struct {
	lastHash []byte
	iniFile  *ini.File
}

func NewReaderTask() *ReaderTask {
	return &ReaderTask{}
}

func (*ReaderTask) Name() string {
	return "Config Reader"
}

func (rt *ReaderTask) Run(logger *log.Logger) {
	if rt.lastHash == nil {
		/* --- Initial setup --- */
		// Make sure we can load the file
		file, err := ini.Load(configFilePath)
		if err != nil {
			panic(err)
		}
		rt.iniFile = file
		// Calculate the hash
		rt.lastHash = calculateHash()
		// Load the config from file
		err = file.MapTo(GetConfig())
		if err != nil {
			panic(err)
		}
		logger.Info("Initial config loaded")
	} else {
		/* --- Later runs --- */
		// Check if the file contents have changed
		newHash := calculateHash()
		sameHash := true
		for i := range rt.lastHash {
			if rt.lastHash[i] != newHash[i] {
				sameHash = false
				break
			}
		}
		// If so, reload the config
		if !sameHash {
			rt.lastHash = newHash
			err := rt.iniFile.Reload()
			if err != nil {
				panic(err)
			}
			err = rt.iniFile.MapTo(GetConfig())
			if err != nil {
				panic(err)
			}
			logger.Info("Config reloaded")
		}
	}
}

func calculateHash() []byte {
	file, err := os.Open(configFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	hasher := md5.New()
	_, err = io.Copy(hasher, file)
	if err != nil {
		panic(err)
	}
	return hasher.Sum(nil)
}
