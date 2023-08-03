# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-distribution/Dockerfile.gotpl
# BOM: {"github:csmith/irc-distribution":"v1.0.0","image:base":"8e5d45c90d7bcc56a15f1770933c8052a79773c1c631b96f36b5e25ef5606d5f","image:golang":"59d70d137edf684ebef3dccd7bbcacc7ad173c4f2bae1fcd8b3a15e827e9a891"}

FROM reg.c5h.io/golang@sha256:59d70d137edf684ebef3dccd7bbcacc7ad173c4f2bae1fcd8b3a15e827e9a891 AS build

ARG TAG="v1.0.0"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/irc-distribution /go/src/github.com/csmith/irc-distribution; \
    cd /go/src/github.com/csmith/irc-distribution; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/distribution; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:8e5d45c90d7bcc56a15f1770933c8052a79773c1c631b96f36b5e25ef5606d5f

COPY --from=build /go/src/github.com/csmith/irc-distribution/main /irc-distribution
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-distribution"]
