# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-github/Dockerfile.gotpl
# BOM: {"github:greboid/irc-github":"v4.0.2","image:base":"c4936e066851061b7fd136c869de1dd56413038f3fc44545e606b637a074fd8d","image:golang":"0304828501e1f9414244d2b381c5a8276ade33eb0a6c89bbce9802fcfab2ce4c"}

FROM reg.c5h.io/golang@sha256:0304828501e1f9414244d2b381c5a8276ade33eb0a6c89bbce9802fcfab2ce4c AS build

ARG TAG="v4.0.2"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-github /go/src/github.com/greboid/irc-github; \
    cd /go/src/github.com/greboid/irc-github; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/github; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:c4936e066851061b7fd136c869de1dd56413038f3fc44545e606b637a074fd8d

COPY --from=build /go/src/github.com/greboid/irc-github/main /irc-github
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-github"]
