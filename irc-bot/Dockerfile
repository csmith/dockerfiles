# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-bot/Dockerfile.gotpl
# BOM: {"github:greboid/irc-bot":"v5.0.8","image:base":"5eae401fa2451fc9f557a2ecd950180c3cc0c43425fa6e9af5cd063e7a8cae38","image:golang":"919b3de54087ebdfb600d8dfea6f1022eefc0f591bdc99ea1e57993c30301693"}

FROM reg.c5h.io/golang@sha256:919b3de54087ebdfb600d8dfea6f1022eefc0f591bdc99ea1e57993c30301693 AS build

RUN set -eux; \
    git clone --depth=1 -b v5.0.8 --single-branch https://github.com/greboid/irc-bot /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/bot; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:5eae401fa2451fc9f557a2ecd950180c3cc0c43425fa6e9af5cd063e7a8cae38

COPY --from=build /work/out /irc-bot
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-bot"]


EXPOSE 8080
