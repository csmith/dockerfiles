# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-bot/Dockerfile.gotpl
# BOM: {"github:greboid/irc-bot":"v5.0.6","image:base":"d97e687affbcacbd2173a2a58381ecf270adc8d37d409258cfbc04b0d0a8b092","image:golang":"15cf850e7cd8674aa791e1f2fe31fdd8cb6ef07dfce430f840ab30fa46b38260"}

FROM reg.c5h.io/golang@sha256:15cf850e7cd8674aa791e1f2fe31fdd8cb6ef07dfce430f840ab30fa46b38260 AS build

ARG TAG="v5.0.6"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-bot /go/src/github.com/greboid/irc-bot; \
    cd /go/src/github.com/greboid/irc-bot; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/bot; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:d97e687affbcacbd2173a2a58381ecf270adc8d37d409258cfbc04b0d0a8b092

COPY --from=build /go/src/github.com/greboid/irc-bot/main /irc-bot
COPY --from=build /notices /notices
EXPOSE 8080
ENTRYPOINT ["/irc-bot"]
