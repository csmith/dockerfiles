# Base

* Upstream: _n/a_
* Dependencies: alpine

A minimal image to use as a base for binary services, inspired heavily by
[Distroless](https://github.com/GoogleContainerTools/distroless) (but buildable without bazel!)

Creates and uses a "nonroot" user (uid: `65532`), and provides the `tzdata`, `ca-certificates` and
`musl` packages from Alpine.
