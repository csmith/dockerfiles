# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-bot/Dockerfile.gotpl
# BOM: {"github:greboid/irc-bot":"v5.0.5","image:base":"f85825cc6e579732f339638d1cd16d87f0673759127df47e8bcd3586412acb5d","image:golang":"1693248c51599cc0c0f021fd911b8c2d2cd1d911c53996eab10525ede4ad9afa"}

FROM reg.c5h.io/golang@sha256:1693248c51599cc0c0f021fd911b8c2d2cd1d911c53996eab10525ede4ad9afa AS build

ARG TAG="v5.0.5"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-bot /go/src/github.com/greboid/irc-bot; \
    cd /go/src/github.com/greboid/irc-bot; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/bot; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:f85825cc6e579732f339638d1cd16d87f0673759127df47e8bcd3586412acb5d

COPY --from=build /go/src/github.com/greboid/irc-bot/main /irc-bot
COPY --from=build /notices /notices
EXPOSE 8080
ENTRYPOINT ["/irc-bot"]
