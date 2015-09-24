package log2

import (
    "errors"
    "log"
    "os"
    "runtime"
    "strings"
)

// The critical flag to indicate the log request is something should not be muted.
var Critical = errors.New("[critical]")

// Interface for output.
type Logger interface {
    // skip is to skip call depth.
    Output(debug bool, skip int, a []interface{})
}

// Implementation of Logger using log.Logger from the standard library.
type stdLogger struct {
    logger *log.Logger
}

func (sl *stdLogger) Output(debug bool, skip int, a []interface{}) {
    if debug || ((len(a) > 0) && (a[0] == Critical)) {
        skip++
        sl.logger.Output(skip+1, strings.Join(Strings(a, skip), " "))
    }
}

// New Logger with stdLogger.
func NewStdLogger(prefix interface{}) Logger {
    if prefix == nil {
        pcs := make([]uintptr, 1)
        if runtime.Callers(2, pcs) == 1 {
            fn := runtime.FuncForPC(pcs[0])
            s := fn.Name()
            prefix = "[" + s[:len(s)-5] + "]"
        }
    }

    p, ok := prefix.(string)
    if !ok {
        if q, ok := prefix.(interface{String() string}); ok {
            p = q.String()
        }
    }
    return &stdLogger{log.New(os.Stdout, p, log.Ldate | log.Ltime | log.Lshortfile)}
}
