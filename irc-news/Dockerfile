# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-news/Dockerfile.gotpl
# BOM: {"github:csmith/irc-news":"v2.0.2","image:base":"4061e85f242c026f31ae923214310d31cf4a3c2b112953fd9307fe649b3b1d49","image:golang":"7f7f0260925d34f1c0734ccc9eb6106f220ff3f285e2d5bdf07e5fd4231d94e1"}

FROM reg.c5h.io/golang@sha256:7f7f0260925d34f1c0734ccc9eb6106f220ff3f285e2d5bdf07e5fd4231d94e1 AS build

ARG TAG="v2.0.2"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/irc-news /go/src/github.com/csmith/irc-news; \
    cd /go/src/github.com/csmith/irc-news; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:4061e85f242c026f31ae923214310d31cf4a3c2b112953fd9307fe649b3b1d49

COPY --from=build /go/src/github.com/csmith/irc-news/main /irc-news
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-news"]
