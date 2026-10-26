# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-bot/Dockerfile.gotpl
# BOM: {"github:https://github.com/greboid/irc-bot":"v5.1.0","image:base":"091cf7586a6b9d750c6e29255c7231896ce1ec4ef852f351068072ca00216b47","image:golang":"c04403d382d3f01c5081af9cf8b27648cf75d06c3111ca5310304a8b3e8e454f"}

FROM reg.c5h.io/golang@sha256:c04403d382d3f01c5081af9cf8b27648cf75d06c3111ca5310304a8b3e8e454f AS build

RUN set -eux; \
    git clone --depth=1 -b v5.1.0 --single-branch https://github.com/greboid/irc-bot /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/bot; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:091cf7586a6b9d750c6e29255c7231896ce1ec4ef852f351068072ca00216b47

COPY --from=build /work/out /irc-bot
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-bot"]


EXPOSE 8080
