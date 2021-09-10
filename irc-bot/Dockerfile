# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-bot/Dockerfile.gotpl
# BOM: {"github:greboid/irc-bot":"v5.0.2","image:base":"sha256:f7f27db7afb58bae23ad902072228f1b090b878af303d9f63bd2c1526b9b4f53","image:golang":"sha256:1ca61b1147a7b01928b1bb14ca0ca6dc26d38cc06a3ee3af60eaff623eef4a11"}

FROM reg.c5h.io/golang@sha256:1ca61b1147a7b01928b1bb14ca0ca6dc26d38cc06a3ee3af60eaff623eef4a11 AS build

ARG TAG="v5.0.2"

RUN set -eux; \
    apk add --no-cache git; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-bot /go/src/github.com/greboid/irc-bot; \
    cd /go/src/github.com/greboid/irc-bot; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/bot; \
    go run github.com/google/go-licenses@latest save ./... --save_path=/notices; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /go/src/github.com/greboid/irc-bot/main;

FROM reg.c5h.io/base@sha256:f7f27db7afb58bae23ad902072228f1b090b878af303d9f63bd2c1526b9b4f53

COPY --from=build /go/src/github.com/greboid/irc-bot/main /irc-bot
COPY --from=build /notices /notices
EXPOSE 8080
ENTRYPOINT ["/irc-bot"]
