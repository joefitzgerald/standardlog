Standardlog
===========

[![wercker status](https://app.wercker.com/status/f5ccdd31a32c541bca3271b382909151/m "wercker status")](https://app.wercker.com/project/bykey/f5ccdd31a32c541bca3271b382909151)

The [go 1 compatibility promise](http://golang.org/doc/go1compat) prevents the
standard library from changing `log.Logger` to be an interface.

This library has one single goal: to provide an interface that matches the
functions in `log.Logger` so that you can substitute the standard Logger with a
replacement if needed.

## Usage

`go get -u github.com/joefitzgerald/standardlog`

```go
import (
	"os"

	"github.com/joefitzgerald/standardlog"
)

func main() {
  var l standardlog.Logger
  l = log.New(os.Stdout, "", 0)
  // Use l here...
}
```
