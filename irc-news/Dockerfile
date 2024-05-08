# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-news/Dockerfile.gotpl
# BOM: {"github:csmith/irc-news":"v2.0.2","image:base":"c4936e066851061b7fd136c869de1dd56413038f3fc44545e606b637a074fd8d","image:golang":"700442935b6dca3bc16ce009fd0e140ee319f2b7d40a8ddd6484bcf089b1df01"}

FROM reg.c5h.io/golang@sha256:700442935b6dca3bc16ce009fd0e140ee319f2b7d40a8ddd6484bcf089b1df01 AS build

ARG TAG="v2.0.2"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/irc-news /go/src/github.com/csmith/irc-news; \
    cd /go/src/github.com/csmith/irc-news; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:c4936e066851061b7fd136c869de1dd56413038f3fc44545e606b637a074fd8d

COPY --from=build /go/src/github.com/csmith/irc-news/main /irc-news
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-news"]
