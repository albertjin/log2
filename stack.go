package log2

import (
    "fmt"
    "runtime"
    "strings"
)

// Helper function to generate stack trace.
func Stack(skip int, prefix string) (ss []string) {
    pcs := make([]uintptr, 64)
    n := runtime.Callers(skip + 2, pcs)
    for _, pc := range pcs[:n - 2] {
        if fn := runtime.FuncForPC(pc); fn != nil {
            filename, line := fn.FileLine(pc)
            ss = append(ss, fmt.Sprintf("%v%v:%v: %v", prefix, filename, line, fn.Name()))
        }
    }
    return ss
}

// Shortcut implementation of logger function.
func StackLog(calldepth int) string {
    return "\n" + strings.Join(Stack(calldepth+1, "  "), "\n")
}
