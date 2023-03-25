# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-github/Dockerfile.gotpl
# BOM: {"github:greboid/irc-github":"v4.0.1","image:base":"c41c52eb39d01d876edd44db743514a484b09b61793b6ada6f064315eb18327d","image:golang":"da990a895e18b0c39c5990bfac24933ef58beb9dfef258de0e9bc393274f2f7e"}

FROM reg.c5h.io/golang@sha256:da990a895e18b0c39c5990bfac24933ef58beb9dfef258de0e9bc393274f2f7e AS build

ARG TAG="v4.0.1"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-github /go/src/github.com/greboid/irc-github; \
    cd /go/src/github.com/greboid/irc-github; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/github; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:c41c52eb39d01d876edd44db743514a484b09b61793b6ada6f064315eb18327d

COPY --from=build /go/src/github.com/greboid/irc-github/main /irc-github
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-github"]
