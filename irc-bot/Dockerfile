# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-bot/Dockerfile.gotpl
# BOM: {"github:greboid/irc-bot":"v5.0.6","image:base":"893c39e6f5fbf0a003c8fdb3c55787d6c9e6c79c60e665bf791f612d35fbcff1","image:golang":"16ea2b79650271d5de2adce54f127133681ae36fabd15ec4e1bb69c696243f0f"}

FROM reg.c5h.io/golang@sha256:16ea2b79650271d5de2adce54f127133681ae36fabd15ec4e1bb69c696243f0f AS build

ARG TAG="v5.0.6"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-bot /go/src/github.com/greboid/irc-bot; \
    cd /go/src/github.com/greboid/irc-bot; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/bot; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:893c39e6f5fbf0a003c8fdb3c55787d6c9e6c79c60e665bf791f612d35fbcff1

COPY --from=build /go/src/github.com/greboid/irc-bot/main /irc-bot
COPY --from=build /notices /notices
EXPOSE 8080
ENTRYPOINT ["/irc-bot"]
