package papertrail

import (
	"fmt"
	"log"
	"log/syslog"
	"runtime"
	"strings"
)

var debug bool = false

var w *syslog.Writer

func Init(Host, Applicationname string) {
	var err error
	w, err = syslog.Dial("udp", Host, syslog.LOG_EMERG|syslog.LOG_DAEMON, Applicationname)
	if err != nil {
		log.Fatal("failed to dial syslog")
	}
}

// Info example:
//
// Info("timezone %s", timezone)
//
func Info(msg string, vars ...interface{}) {
	w.Info(fmt.Sprintf(strings.Join([]string{"[INFO ]", msg}, " "), vars...))
}

// Debug example:
//
// Debug("timezone %s", timezone)
//
func Debug(msg string, vars ...interface{}) {
	if debug {
		w.Debug(fmt.Sprintf(strings.Join([]string{"[DEBUG]", msg}, " "), vars...))
	}
}

// Fatal example:
//
// Fatal(errors.New("db timezone must be UTC"))
//
func Fatal(err error) {
	pc, fn, line, _ := runtime.Caller(1)
	// Include function name if debugging
	if debug {
		w.Alert(fmt.Sprintf("[FATAL] %s [%s:%s:%d]", err, runtime.FuncForPC(pc).Name(), fn, line))
	} else {
		w.Alert(fmt.Sprintf("[FATAL] %s [%s:%d]", err, fn, line))
	}
}

// Error example:
//
// Error(errors.Errorf("Invalid timezone %s", timezone))
//
func Error(err error) {
	pc, fn, line, _ := runtime.Caller(1)
	// Include function name if debugging
	if debug {
		w.Err(fmt.Sprintf("[ERROR] %s [%s:%s:%d]", err, runtime.FuncForPC(pc).Name(), fn, line))
	} else {
		w.Err(fmt.Sprintf("[ERROR] %s [%s:%d]", err, fn, line))
	}
}
