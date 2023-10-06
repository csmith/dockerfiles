# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-notifier/Dockerfile.gotpl
# BOM: {"github:greboid/irc-notifier":"v1.0.4","image:base":"8e5d45c90d7bcc56a15f1770933c8052a79773c1c631b96f36b5e25ef5606d5f","image:golang":"a62ef18994033df8f458aa6b4fc3b4291332e35859c5c3bfe957eb2ec9a7f932"}

FROM reg.c5h.io/golang@sha256:a62ef18994033df8f458aa6b4fc3b4291332e35859c5c3bfe957eb2ec9a7f932 AS build

ARG TAG="v1.0.4"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-notifier /go/src/github.com/greboid/irc-notifier; \
    cd /go/src/github.com/greboid/irc-notifier; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/notifier; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:8e5d45c90d7bcc56a15f1770933c8052a79773c1c631b96f36b5e25ef5606d5f

COPY --from=build /go/src/github.com/greboid/irc-notifier/main /irc-notifier
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-notifier"]
