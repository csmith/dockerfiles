# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-bot/Dockerfile.gotpl
# BOM: {"github:greboid/irc-bot":"v5.0.4","image:base":"e3b3d5f2e5fa7e5cbe763581a8da6b56c4498cb2107b9e385acc8f13d4de5903","image:golang":"17bd5548beba9a661f8646117b490dfcd314a3856053a035f7811bee640b6430"}

FROM reg.c5h.io/golang@sha256:17bd5548beba9a661f8646117b490dfcd314a3856053a035f7811bee640b6430 AS build

ARG TAG="v5.0.4"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-bot /go/src/github.com/greboid/irc-bot; \
    cd /go/src/github.com/greboid/irc-bot; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/bot; \
    go run github.com/google/go-licenses@latest save ./... --save_path=/notices; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /go/src/github.com/greboid/irc-bot/main;

FROM reg.c5h.io/base@sha256:e3b3d5f2e5fa7e5cbe763581a8da6b56c4498cb2107b9e385acc8f13d4de5903

COPY --from=build /go/src/github.com/greboid/irc-bot/main /irc-bot
COPY --from=build /notices /notices
EXPOSE 8080
ENTRYPOINT ["/irc-bot"]
