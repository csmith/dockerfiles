# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-bot/Dockerfile.gotpl
# BOM: {"github:greboid/irc-bot":"v5.0.6","image:base":"2ad8fb5e47298cb81072a36ee939a2024bc3acb6221b105cb3928c0e6a97c665","image:golang":"dc4eac488c6c21ead1b8d08cf44dfa9e8d33ccec71eb063fe8b3e2ede109d082"}

FROM reg.c5h.io/golang@sha256:dc4eac488c6c21ead1b8d08cf44dfa9e8d33ccec71eb063fe8b3e2ede109d082 AS build

ARG TAG="v5.0.6"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-bot /go/src/github.com/greboid/irc-bot; \
    cd /go/src/github.com/greboid/irc-bot; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/bot; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:2ad8fb5e47298cb81072a36ee939a2024bc3acb6221b105cb3928c0e6a97c665

COPY --from=build /go/src/github.com/greboid/irc-bot/main /irc-bot
COPY --from=build /notices /notices
EXPOSE 8080
ENTRYPOINT ["/irc-bot"]
