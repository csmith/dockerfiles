# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-bot/Dockerfile.gotpl
# BOM: {"github:greboid/irc-bot":"v5.0.7","image:base":"cfa1d5f949b495a6e077bdddb1dbbc3d9a9d55036ab1267f4168a6a5c7e7fdce","image:golang":"6d352e89f9ee9e49ffec893024fbcc4996b60c9ff5d48d84ff9197c8c5eb9ea6"}

FROM reg.c5h.io/golang@sha256:6d352e89f9ee9e49ffec893024fbcc4996b60c9ff5d48d84ff9197c8c5eb9ea6 AS build

ARG TAG="v5.0.7"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-bot /go/src/github.com/greboid/irc-bot; \
    cd /go/src/github.com/greboid/irc-bot; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/bot; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:cfa1d5f949b495a6e077bdddb1dbbc3d9a9d55036ab1267f4168a6a5c7e7fdce

COPY --from=build /go/src/github.com/greboid/irc-bot/main /irc-bot
COPY --from=build /notices /notices
EXPOSE 8080
ENTRYPOINT ["/irc-bot"]
