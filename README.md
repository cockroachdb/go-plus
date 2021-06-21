# go-plus

**go-plus** (Go+) is an API library that exposes **CRL's custom Go
extensions** to Go programs.

The following two extensions are provided:

- **proc** provides Go programs with finer grain control over the
  Go runtime system. This makes it possible to build more "advanced"
  concurrency data structures such as a custom `sync.Pool` implementation.

  It also provides a new runtime feature called "Inheritable Goroutine
  ID" (IGID).

  API doc link: https://pkg.go.dev/github.com/cockroachdb/go-plus/proc

- **taskgroup** provides Go program with a new runtime abstraction called
  “task group” (discussed [here](https://github.com/cockroachdb/cockroach/pull/60589)),
  to provide aggregate accounting of resource usage to fleets of related
  goroutines.

  API doc link: https://pkg.go.dev/github.com/cockroachdb/go-plus/taskgroup

## How to use in Go programs

When using this library with a “vanilla” Go runtime, the `Supported()`
API calls in each package return `false`, to indicate the extension is
not available.

This makes it possible to build Go programs using this API
even with a “vanilla” Go compiler and runtime.

To actually enable the extension:

1. clone the repository at https://github.com/cockroachdb/go
2. check out the `crdb-fixes` branch
3. build the runtime: `bash src/make.bash`
4. build the Go application with:
   - `GOROOT` set to the custom go source tree
   - `PATH` including `$GOROOT/bin` at the start
   - add the tag `goplus` to the build command line,
     for example: `go build -tags goplus`

Possible build configurations:

|                                       | “vanilla” Go runtime             | Built with `crdb-fixes` runtime extension |
|---------------------------------------|----------------------------------|-------------------------------------------|
| Program compiled without `goplus` tag | `Supported() == false`           | `Supported() == false`                    |
| Program compiled with `goplus` tag    | Build error in `go-plus` library | `Supported() == true`                     |

