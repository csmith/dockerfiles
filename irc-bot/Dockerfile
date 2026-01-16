# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-bot/Dockerfile.gotpl
# BOM: {"github:greboid/irc-bot":"v5.0.8","image:base":"df41e52c0ffc5527d6042026b2d11605ff293dd09f4e088b4b6fb2d54e17a18a","image:golang":"a52db640e1f7200df3b5d9e40efd0ba639a526aad05e9d6427031ea23ad8c019"}

FROM reg.c5h.io/golang@sha256:a52db640e1f7200df3b5d9e40efd0ba639a526aad05e9d6427031ea23ad8c019 AS build

RUN set -eux; \
    git clone --depth=1 -b v5.0.8 --single-branch https://github.com/greboid/irc-bot /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/bot; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:df41e52c0ffc5527d6042026b2d11605ff293dd09f4e088b4b6fb2d54e17a18a

COPY --from=build /work/out /irc-bot
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-bot"]


EXPOSE 8080
