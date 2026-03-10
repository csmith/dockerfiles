# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-bot/Dockerfile.gotpl
# BOM: {"github:greboid/irc-bot":"v5.0.8","image:base":"84a319cda607326ba5aff905ae01191c4a3c793dbad0e058d9a816e3fb29036f","image:golang":"1c1866ed33a6140348e9afd6b32a2d0e328280f2a3c34e90242a715b714d5e15"}

FROM reg.c5h.io/golang@sha256:1c1866ed33a6140348e9afd6b32a2d0e328280f2a3c34e90242a715b714d5e15 AS build

RUN set -eux; \
    git clone --depth=1 -b v5.0.8 --single-branch https://github.com/greboid/irc-bot /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/bot; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:84a319cda607326ba5aff905ae01191c4a3c793dbad0e058d9a816e3fb29036f

COPY --from=build /work/out /irc-bot
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-bot"]


EXPOSE 8080
