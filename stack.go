package log2

import (
    "fmt"
    "runtime"
    "strings"
)

// Helper function to generate stack trace.
func Stack(skip int, prefix string) (ss []string) {
    skip++
    pcs := make([]uintptr, 64)
    n := runtime.Callers(skip+1, pcs)
    for _, pc := range pcs[:n] {
        if fn := runtime.FuncForPC(pc); fn != nil {
            filename, line := fn.FileLine(pc)
            ss = append(ss, fmt.Sprintf("%v%v:%v: %v", prefix, filename, line, fn.Name()))
        }
    }
    return ss
}

// Shortcut implementation of logger function.
func StackLog(skip int) string {
    skip++
    return "\n" + strings.Join(Stack(skip, "  "), "\n")
}
