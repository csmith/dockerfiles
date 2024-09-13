# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-distribution/Dockerfile.gotpl
# BOM: {"github:csmith/irc-distribution":"v1.0.0","image:base":"fa9cdd83581affde03cb2234170235ff287588e4da3786133d72846838210621","image:golang":"4d5308af1e133bec7aabbc9f267e4719a991904ba8bb9fbad15ce4a76e3e6498"}

FROM reg.c5h.io/golang@sha256:4d5308af1e133bec7aabbc9f267e4719a991904ba8bb9fbad15ce4a76e3e6498 AS build

ARG TAG="v1.0.0"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/irc-distribution /go/src/github.com/csmith/irc-distribution; \
    cd /go/src/github.com/csmith/irc-distribution; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/distribution; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:fa9cdd83581affde03cb2234170235ff287588e4da3786133d72846838210621

COPY --from=build /go/src/github.com/csmith/irc-distribution/main /irc-distribution
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-distribution"]
