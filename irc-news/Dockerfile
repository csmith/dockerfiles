# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-news/Dockerfile.gotpl
# BOM: {"github:csmith/irc-news":"v2.0.2","image:base":"893c39e6f5fbf0a003c8fdb3c55787d6c9e6c79c60e665bf791f612d35fbcff1","image:golang":"16ea2b79650271d5de2adce54f127133681ae36fabd15ec4e1bb69c696243f0f"}

FROM reg.c5h.io/golang@sha256:16ea2b79650271d5de2adce54f127133681ae36fabd15ec4e1bb69c696243f0f AS build

ARG TAG="v2.0.2"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/irc-news /go/src/github.com/csmith/irc-news; \
    cd /go/src/github.com/csmith/irc-news; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:893c39e6f5fbf0a003c8fdb3c55787d6c9e6c79c60e665bf791f612d35fbcff1

COPY --from=build /go/src/github.com/csmith/irc-news/main /irc-news
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-news"]
