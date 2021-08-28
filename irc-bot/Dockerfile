# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-bot/Dockerfile.gotpl
# BOM: {"image:reg.c5h.io/base":"sha256:f7f27db7afb58bae23ad902072228f1b090b878af303d9f63bd2c1526b9b4f53","image:reg.c5h.io/golang":"sha256:be107b2ab6659ebdf1a47b12e154bc84a46b2312abe85c6fc055ef8b478c5a8a"}

FROM reg.c5h.io/golang@sha256:be107b2ab6659ebdf1a47b12e154bc84a46b2312abe85c6fc055ef8b478c5a8a AS build

ARG TAG="v5.0.0"

RUN set -eux; \
    apk add --no-cache git; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-bot /go/src/github.com/greboid/irc-bot; \
    cd /go/src/github.com/greboid/irc-bot; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/bot; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /go/src/github.com/greboid/irc-bot/main;

FROM reg.c5h.io/base@sha256:f7f27db7afb58bae23ad902072228f1b090b878af303d9f63bd2c1526b9b4f53

COPY --from=build /go/src/github.com/greboid/irc-bot/main /irc-bot
EXPOSE 8080
ENTRYPOINT ["/irc-bot"]
