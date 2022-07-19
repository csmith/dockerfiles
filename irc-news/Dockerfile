# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-news/Dockerfile.gotpl
# BOM: {"github:csmith/irc-news":"v2.0.2","image:base":"bffee9632ec79071de34e1a4a4f6c4563a71ad26d5923062bd9d0f9016f72beb","image:golang":"1387c3f016ef116eb4cc78ca21627774eef5d67c5ce758b775c0bfa7b9901818"}

FROM reg.c5h.io/golang@sha256:1387c3f016ef116eb4cc78ca21627774eef5d67c5ce758b775c0bfa7b9901818 AS build

ARG TAG="v2.0.2"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/irc-news /go/src/github.com/csmith/irc-news; \
    cd /go/src/github.com/csmith/irc-news; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:bffee9632ec79071de34e1a4a4f6c4563a71ad26d5923062bd9d0f9016f72beb

COPY --from=build /go/src/github.com/csmith/irc-news/main /irc-news
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-news"]
