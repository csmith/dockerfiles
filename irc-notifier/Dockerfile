# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-notifier/Dockerfile.gotpl
# BOM: {"github:greboid/irc-notifier":"v1.0.4","image:base":"45ece244ca0ace29f890ad27114ff83a36c00e74e6cbec3d1e226b7e62d0de7a","image:golang":"da990a895e18b0c39c5990bfac24933ef58beb9dfef258de0e9bc393274f2f7e"}

FROM reg.c5h.io/golang@sha256:da990a895e18b0c39c5990bfac24933ef58beb9dfef258de0e9bc393274f2f7e AS build

ARG TAG="v1.0.4"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-notifier /go/src/github.com/greboid/irc-notifier; \
    cd /go/src/github.com/greboid/irc-notifier; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/notifier; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:45ece244ca0ace29f890ad27114ff83a36c00e74e6cbec3d1e226b7e62d0de7a

COPY --from=build /go/src/github.com/greboid/irc-notifier/main /irc-notifier
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-notifier"]
