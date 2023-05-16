# buggy-log, a simple Go logging package

This is a simple and minimal logging utility for the Go language. We use this utility for several i2c drivers but might shift away to the standard `log` utility. Either way, this is a nice way to get familiar with publishing Go modules.

## Installation

```bash
go get github.com/MrBuggy-Amsterdam/buggy-log/
```

## Usage

As you would expect from a logging utility. By default, debug level prints are hidden.

```Go
package main

import (
    log "github.com/MrBuggy-Amsterdam/buggy-log/"
)

func main() {
    log.SetVerbosity(log.DebugLevel)
    log.Debug("Hey, can you read this?")
    log.Debug("Of course, include %s as well", "formatting")
}
```

