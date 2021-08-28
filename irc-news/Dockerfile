# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-news/Dockerfile.gotpl
# BOM: {"github:csmith/ircplugins":"v1.0.0","image:reg.c5h.io/base":"sha256:f7f27db7afb58bae23ad902072228f1b090b878af303d9f63bd2c1526b9b4f53","image:reg.c5h.io/golang":"sha256:06b64e8124c836635ec4c678dc29fe73a12ccc245d5fe811888f403f989c52aa"}

FROM reg.c5h.io/golang@sha256:06b64e8124c836635ec4c678dc29fe73a12ccc245d5fe811888f403f989c52aa AS build

ARG TAG="v1.0.0"

RUN set -eux; \
    apk add --no-cache git; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/ircplugins /go/src/github.com/csmith/ircplugins; \
    cd /go/src/github.com/csmith/ircplugins; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/news; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /go/src/github.com/csmith/ircplugins/main;

FROM reg.c5h.io/base@sha256:f7f27db7afb58bae23ad902072228f1b090b878af303d9f63bd2c1526b9b4f53

COPY --from=build /go/src/github.com/csmith/ircplugins/main /irc-news
ENTRYPOINT ["/irc-news"]
