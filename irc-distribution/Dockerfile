# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-distribution/Dockerfile.gotpl
# BOM: {"github:csmith/irc-distribution":"v1.0.0","image:base":"d9260eae728da42310e4b14cf5ca84500c2cb480d1c04fbf7f574b7e6b3e9743","image:golang":"9776708de22db709823ea6880f1d6e4a3af1e8bc09801f7a9dfd1276c8adae13"}

FROM reg.c5h.io/golang@sha256:9776708de22db709823ea6880f1d6e4a3af1e8bc09801f7a9dfd1276c8adae13 AS build

ARG TAG="v1.0.0"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/irc-distribution /go/src/github.com/csmith/irc-distribution; \
    cd /go/src/github.com/csmith/irc-distribution; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/distribution; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:d9260eae728da42310e4b14cf5ca84500c2cb480d1c04fbf7f574b7e6b3e9743

COPY --from=build /go/src/github.com/csmith/irc-distribution/main /irc-distribution
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-distribution"]
