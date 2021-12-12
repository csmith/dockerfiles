# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-goplum/Dockerfile.gotpl
# BOM: {"github:greboid/irc-goplum":"v4.0.1","image:base":"e3b3d5f2e5fa7e5cbe763581a8da6b56c4498cb2107b9e385acc8f13d4de5903","image:golang":"72adae738352aa1e5dc312952446651fc767661e2b9bdb006be122bfde7fd03a"}

FROM reg.c5h.io/golang@sha256:72adae738352aa1e5dc312952446651fc767661e2b9bdb006be122bfde7fd03a AS build

ARG TAG="v4.0.1"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-goplum /go/src/github.com/greboid/irc-goplum; \
    cd /go/src/github.com/greboid/irc-goplum; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/goplum; \
    go run github.com/google/go-licenses@latest save ./... --save_path=/notices; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /go/src/github.com/greboid/irc-goplum/main;

FROM reg.c5h.io/base@sha256:e3b3d5f2e5fa7e5cbe763581a8da6b56c4498cb2107b9e385acc8f13d4de5903

COPY --from=build /go/src/github.com/greboid/irc-goplum/main /irc-goplum
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-goplum"]
