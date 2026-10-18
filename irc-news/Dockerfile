# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-news/Dockerfile.gotpl
# BOM: {"github:https://github.com/csmith/irc-news":"v2.0.2","image:base":"fbcd09463f4b0824d709a18598637c2f7f7266d0591ca7b61311b2a12cca6f11","image:golang":"2a596927c391e5688e97229dbddd5d65429b1a7459a530477522f8fd1d8d29bd"}

FROM reg.c5h.io/golang@sha256:2a596927c391e5688e97229dbddd5d65429b1a7459a530477522f8fd1d8d29bd AS build

RUN set -eux; \
    git clone --depth=1 -b v2.0.2 --single-branch https://github.com/csmith/irc-news /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:fbcd09463f4b0824d709a18598637c2f7f7266d0591ca7b61311b2a12cca6f11

COPY --from=build /work/out /irc-news
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-news"]
