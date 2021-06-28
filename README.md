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

Note: the way this works is that the customized go runtime implicitly
provides an extra build tag `goplus`. One can check whether the custom
runtime is active with e.g. `go test -v ./proc` inside the library
(the provided `TestSupported` reports the build tags and support
status).
