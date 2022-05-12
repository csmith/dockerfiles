# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-news/Dockerfile.gotpl
# BOM: {"github:csmith/irc-news":"v2.0.2","image:base":"f85825cc6e579732f339638d1cd16d87f0673759127df47e8bcd3586412acb5d","image:golang":"5993a328d81c3a30ddc6f57cd84bbc116ee76bfb86d1d222ab6cc206ea43a497"}

FROM reg.c5h.io/golang@sha256:5993a328d81c3a30ddc6f57cd84bbc116ee76bfb86d1d222ab6cc206ea43a497 AS build

ARG TAG="v2.0.2"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/irc-news /go/src/github.com/csmith/irc-news; \
    cd /go/src/github.com/csmith/irc-news; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:f85825cc6e579732f339638d1cd16d87f0673759127df47e8bcd3586412acb5d

COPY --from=build /go/src/github.com/csmith/irc-news/main /irc-news
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-news"]
