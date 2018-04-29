package log

import (
	"io"
	"log"
	"os"
	"path"
)

const LOGFILE = "agent.log"

func init() {
	// Open the specified logfile so we can write to that as well
	logdir := os.Getenv("CAF_APP_LOG_DIR")
	logpath := path.Join(logdir, LOGFILE)
	logfile, err := os.OpenFile(logpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Printf("[WARN] Logger : Couldn't open file %s for logging\n", logpath)
	} else {
		// Make the default logger write to both stdout and the logfile
		writer := io.MultiWriter(os.Stdout, logfile)
		log.SetOutput(writer)
	}
	log.Println("[INFO] Logger : Started logging")
}

func NewLogger(prefix string) *Logger {
	return &Logger{
		prefix: prefix,
	}
}

type Logger struct {
	prefix string
}

func (l *Logger) Info(v ...interface{}) {
	log.Println(append([]interface{}{"[INFO] " + l.prefix + " :"}, v...)...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	log.Printf("[INFO] "+l.prefix+" : "+format, v...)
}

func (l *Logger) Warning(v ...interface{}) {
	log.Println(append([]interface{}{"[WARN] " + l.prefix + " :"}, v...)...)
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	log.Printf("[WARN] "+l.prefix+" : "+format, v...)
}

func (l *Logger) Error(v ...interface{}) {
	log.Println(append([]interface{}{"[ERR] " + l.prefix + " :"}, v...)...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	log.Printf("[ERR] "+l.prefix+" : "+format, v...)
}
