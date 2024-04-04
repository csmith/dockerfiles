# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-notifier/Dockerfile.gotpl
# BOM: {"github:greboid/irc-notifier":"v1.0.5","image:base":"c4936e066851061b7fd136c869de1dd56413038f3fc44545e606b637a074fd8d","image:golang":"ce380808fdf4cc939295924866b2f58eec6b17b0d028308c9d9e7eaaea2a7fe3"}

FROM reg.c5h.io/golang@sha256:ce380808fdf4cc939295924866b2f58eec6b17b0d028308c9d9e7eaaea2a7fe3 AS build

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
