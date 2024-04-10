# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-notifier/Dockerfile.gotpl
# BOM: {"github:greboid/irc-notifier":"v1.0.5","image:base":"c4936e066851061b7fd136c869de1dd56413038f3fc44545e606b637a074fd8d","image:golang":"00405f110fbb03f3a820fb90077f26c41b754718998a9f643ccadc778cc3522c"}

FROM reg.c5h.io/golang@sha256:00405f110fbb03f3a820fb90077f26c41b754718998a9f643ccadc778cc3522c AS build

ARG TAG="v1.0.5"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-notifier /go/src/github.com/greboid/irc-notifier; \
    cd /go/src/github.com/greboid/irc-notifier; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/notifier; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:c4936e066851061b7fd136c869de1dd56413038f3fc44545e606b637a074fd8d

COPY --from=build /go/src/github.com/greboid/irc-notifier/main /irc-notifier
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-notifier"]
