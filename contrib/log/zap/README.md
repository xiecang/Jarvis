# log
> The logging implemented by [zap](http://go.uber.org/zap)

## Usage
```go
import (
    "github.com/xiecang/jarvis/contrib/log/zap"
    "github.com/xiecang/jarvis/log"
)
```

```go
    log.SetLogger(zap.NewLogger(log.ParseLevel(logLevel), zap.WithName(Name), zap.WithPath(logPath), zap.WithCallerFullPath()))
    log.Debug("test")
    log.Info("test")
    log.Warn("test")
    log.Error("test")
```
