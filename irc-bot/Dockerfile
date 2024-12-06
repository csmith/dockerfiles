# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-bot/Dockerfile.gotpl
# BOM: {"github:greboid/irc-bot":"v5.0.7","image:base":"3d9bc81e06db6829d63423bb674f142586ed9bd48613aab2849ada8f7b431a99","image:golang":"2d5332dbd96c2017fc9f481e608f709ca1b3c9a5c5725a66f9dba26ba5432021"}

FROM reg.c5h.io/golang@sha256:2d5332dbd96c2017fc9f481e608f709ca1b3c9a5c5725a66f9dba26ba5432021 AS build

ARG TAG="v5.0.7"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-bot /go/src/github.com/greboid/irc-bot; \
    cd /go/src/github.com/greboid/irc-bot; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/bot; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:3d9bc81e06db6829d63423bb674f142586ed9bd48613aab2849ada8f7b431a99

COPY --from=build /go/src/github.com/greboid/irc-bot/main /irc-bot
COPY --from=build /notices /notices
EXPOSE 8080
ENTRYPOINT ["/irc-bot"]
