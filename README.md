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
