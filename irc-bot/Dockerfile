# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-bot/Dockerfile.gotpl
# BOM: {"github:greboid/irc-bot":"v5.0.4","image:base":"sha256:74295276a5989bf84e9bc17d2a44b826a4be6aa2fcfa8e53a6b33404c6f37a02","image:golang":"sha256:2495578a7ef96327cfc4b2971f4264b0006c18d2494a015a03ab03476f86ca70"}

FROM reg.c5h.io/golang@sha256:2495578a7ef96327cfc4b2971f4264b0006c18d2494a015a03ab03476f86ca70 AS build

ARG TAG="v5.0.4"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-bot /go/src/github.com/greboid/irc-bot; \
    cd /go/src/github.com/greboid/irc-bot; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/bot; \
    go run github.com/google/go-licenses@latest save ./... --save_path=/notices; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /go/src/github.com/greboid/irc-bot/main;

FROM reg.c5h.io/base@sha256:74295276a5989bf84e9bc17d2a44b826a4be6aa2fcfa8e53a6b33404c6f37a02

COPY --from=build /go/src/github.com/greboid/irc-bot/main /irc-bot
COPY --from=build /notices /notices
EXPOSE 8080
ENTRYPOINT ["/irc-bot"]
