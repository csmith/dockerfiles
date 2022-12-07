# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-news/Dockerfile.gotpl
# BOM: {"github:csmith/irc-news":"v2.0.2","image:base":"d9260eae728da42310e4b14cf5ca84500c2cb480d1c04fbf7f574b7e6b3e9743","image:golang":"06bd9df2c9f04a24cf605e73052fa1983be5a1d525dcbecdfee4a0c407d318f5"}

FROM reg.c5h.io/golang@sha256:06bd9df2c9f04a24cf605e73052fa1983be5a1d525dcbecdfee4a0c407d318f5 AS build

ARG TAG="v2.0.2"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/irc-news /go/src/github.com/csmith/irc-news; \
    cd /go/src/github.com/csmith/irc-news; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:d9260eae728da42310e4b14cf5ca84500c2cb480d1c04fbf7f574b7e6b3e9743

COPY --from=build /go/src/github.com/csmith/irc-news/main /irc-news
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-news"]
