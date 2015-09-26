# log2 for Go

To use this package. Create a `log.go` file as follows.

```go
package __package name__

import (
    "github.com/albertjin/log2"
)

var critical = log2.Critical

var Logger = log2.NewStdLogger(nil)
var LogDebug = false

func log(a... interface{}) {
    Logger.Output(LogDebug, 1, a)
}

var stack = log2.StackLog
```

Example of logging code,

```go
func hello() {
    log("hello") // This a simple logging. When LogDebug is false, the log is muted.
    log(critical, stack) // Regardless LogDebug or not, this is always logged.
    log(func() string {
        return "hello" // Some complex computation can be delayed or skipped.
    })
}
```