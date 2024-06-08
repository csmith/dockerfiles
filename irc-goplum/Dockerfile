# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-goplum/Dockerfile.gotpl
# BOM: {"github:greboid/irc-goplum":"v4.0.2","image:base":"b53f7bd15a51e9e9e1c9d1ddea46a2f62223ed7cd8d8fa320c0d6b323f9f3f99","image:golang":"74d19beff0dc3e4d62308c0bff99714eb6fcdb90a7590f6d0aac3a841bfc54ef"}

FROM reg.c5h.io/golang@sha256:74d19beff0dc3e4d62308c0bff99714eb6fcdb90a7590f6d0aac3a841bfc54ef AS build

ARG TAG="v4.0.2"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-goplum /go/src/github.com/greboid/irc-goplum; \
    cd /go/src/github.com/greboid/irc-goplum; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/goplum; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:b53f7bd15a51e9e9e1c9d1ddea46a2f62223ed7cd8d8fa320c0d6b323f9f3f99

COPY --from=build /go/src/github.com/greboid/irc-goplum/main /irc-goplum
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-goplum"]
