# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-goplum/Dockerfile.gotpl
# BOM: {"github:greboid/irc-goplum":"v4.0.1","image:base":"b2e73bfd5ae1ce12341dd36f24573d1316bf5172809c65dd669ff2720582967f","image:golang":"e00edff1e392ceae323564d2f925f77a8a4401b70fefdf4e7608042c32725b2a"}

FROM reg.c5h.io/golang@sha256:e00edff1e392ceae323564d2f925f77a8a4401b70fefdf4e7608042c32725b2a AS build

ARG TAG="v4.0.1"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-goplum /go/src/github.com/greboid/irc-goplum; \
    cd /go/src/github.com/greboid/irc-goplum; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/goplum; \
    go run github.com/google/go-licenses@latest save ./... --save_path=/notices; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /go/src/github.com/greboid/irc-goplum/main;

FROM reg.c5h.io/base@sha256:b2e73bfd5ae1ce12341dd36f24573d1316bf5172809c65dd669ff2720582967f

COPY --from=build /go/src/github.com/greboid/irc-goplum/main /irc-goplum
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-goplum"]
