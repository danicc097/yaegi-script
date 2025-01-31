# yaegi-script

An example wrapper for `https://github.com/bitfield/script` using
`https://github.com/traefik/yaegi`.

Build the cli with `go build -o yaegi-script main.go`
and add to your `$PATH`.

Usage:

- Direct call: `yaegi-script scripts/echo.go`
- With shebang `#!/usr/bin/env yaegi-script`
