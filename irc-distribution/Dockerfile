# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-distribution/Dockerfile.gotpl
# BOM: {"github:csmith/irc-distribution":"v1.0.0","image:base":"e211f9db3dc7b27fdb9ce4c8155f980370daf4f68b7411a144e82f9d44f16022","image:golang":"c8273d3333b2ccddb6fcea5c25e93d7af5b03009dfebee7635bd94b8f9f26b00"}

FROM reg.c5h.io/golang@sha256:c8273d3333b2ccddb6fcea5c25e93d7af5b03009dfebee7635bd94b8f9f26b00 AS build

ARG TAG="v1.0.0"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/irc-distribution /go/src/github.com/csmith/irc-distribution; \
    cd /go/src/github.com/csmith/irc-distribution; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/distribution; \
    go run github.com/google/go-licenses@latest save ./... --save_path=/notices; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /go/src/github.com/csmith/irc-distribution/main;

FROM reg.c5h.io/base@sha256:e211f9db3dc7b27fdb9ce4c8155f980370daf4f68b7411a144e82f9d44f16022

COPY --from=build /go/src/github.com/csmith/irc-distribution/main /irc-distribution
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-distribution"]
