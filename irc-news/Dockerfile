# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-news/Dockerfile.gotpl
# BOM: {"github:csmith/irc-news":"v2.0.2","image:base":"d97e687affbcacbd2173a2a58381ecf270adc8d37d409258cfbc04b0d0a8b092","image:golang":"062d03490d9c0dfe31de1fd91620f2e3cc9b5c2b9c8a8ac7cc34f7d820511e40"}

FROM reg.c5h.io/golang@sha256:062d03490d9c0dfe31de1fd91620f2e3cc9b5c2b9c8a8ac7cc34f7d820511e40 AS build

ARG TAG="v2.0.2"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/irc-news /go/src/github.com/csmith/irc-news; \
    cd /go/src/github.com/csmith/irc-news; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:d97e687affbcacbd2173a2a58381ecf270adc8d37d409258cfbc04b0d0a8b092

COPY --from=build /go/src/github.com/csmith/irc-news/main /irc-news
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-news"]
