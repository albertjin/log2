package log2

import (
    "fmt"
)

// Convert array of any object to string with fmt.Sprintf(). The major purpose is to delay function evaluations when logs might be skipped.
// Logger functions are `func() string` and `func(calldepth int) string`
func Strings(a []interface{}, skip int) (ss []string) {
    skip++
    if l := len(a); l > 0 {
        ss = make([]string, l)
        for i, x := range a {
            switch fn := x.(type) {
            case func(int) string:
                ss[i] = fn(skip)
            case func() string:
                ss[i] = fn()
            default:
                ss[i] = fmt.Sprintf("%v", x)
            }
        }
    }
    return
}
