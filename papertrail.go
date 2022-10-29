package papertrail

import (
	"fmt"
	"log"
	"log/syslog"
	"runtime"
	"time"
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

func console(s string) {
	fmt.Println(time + " " + s)
}

// Info example:
//
// Info("timezone %s", timezone)
//
func Info(msg string, vars ...interface{}) {
	s := fmt.Sprintf(strings.Join([]string{"[INFO ]", msg}, " "), vars...)
	w.Info(s)
	console(s)
	
}

// Debug example:
//
// Debug("timezone %s", timezone)
//
func Debug(msg string, vars ...interface{}) {
	if debug {
		s := fmt.Sprintf(strings.Join([]string{"[DEBUG]", msg}, " "), vars...)
		w.Debug(s)
		console(s)
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
		s := fmt.Sprintf("[FATAL] %s [%s:%s:%d]", err, runtime.FuncForPC(pc).Name(),fn, line)
		w.Alert(s)
	} else {
		s := fmt.Sprintf("[FATAL] %s [%s:%d]", err, fn, line)
		w.Alert(s)
		console(s)
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
		s := fmt.Sprintf("[ERROR] %s [%s:%s:%d]", err, runtime.FuncForPC(pc).Name(), fn, line)
		w.Err(s)
		console(s)
	} else {
		s:= fmt.Sprintf("[ERROR] %s [%s:%d]", err, fn, line)
		w.Err(s)
		console(s)
	}
}
