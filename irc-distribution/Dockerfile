# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-distribution/Dockerfile.gotpl
# BOM: {"github:csmith/irc-distribution":"v1.0.0","image:base":"c4936e066851061b7fd136c869de1dd56413038f3fc44545e606b637a074fd8d","image:golang":"99f92c3368d01fa05c26a46857d73679f8b1cce3ba3606a7ac072f6baad69eae"}

FROM reg.c5h.io/golang@sha256:99f92c3368d01fa05c26a46857d73679f8b1cce3ba3606a7ac072f6baad69eae AS build

ARG TAG="v1.0.0"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/csmith/irc-distribution /go/src/github.com/csmith/irc-distribution; \
    cd /go/src/github.com/csmith/irc-distribution; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/distribution; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:c4936e066851061b7fd136c869de1dd56413038f3fc44545e606b637a074fd8d

COPY --from=build /go/src/github.com/csmith/irc-distribution/main /irc-distribution
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-distribution"]
