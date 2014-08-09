Standardlog
===========

The [go 1 compatibility promise](http://golang.org/doc/go1compat) prevents the
standard library from changing `log.Logger` to be an interface.

This library has one single goal: to provide an interface that matches the
functions in `log.Logger` so that you can substitute the standard Logger with a
replacement if needed.

## Usage

TBD
