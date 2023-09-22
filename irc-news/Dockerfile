# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-news/Dockerfile.gotpl
# BOM: {"github:csmith/irc-news":"v2.0.2","image:base":"8e5d45c90d7bcc56a15f1770933c8052a79773c1c631b96f36b5e25ef5606d5f","image:golang":"3073df7a16320b2bb862d08160cf5d53a3146655efc6cca2b1d84665274074ce"}

FROM reg.c5h.io/golang@sha256:3073df7a16320b2bb862d08160cf5d53a3146655efc6cca2b1d84665274074ce AS build

ARG TAG="v2.0.2"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/irc-news /go/src/github.com/csmith/irc-news; \
    cd /go/src/github.com/csmith/irc-news; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:8e5d45c90d7bcc56a15f1770933c8052a79773c1c631b96f36b5e25ef5606d5f

COPY --from=build /go/src/github.com/csmith/irc-news/main /irc-news
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-news"]
