# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-notifier/Dockerfile.gotpl
# BOM: {"github:greboid/irc-notifier":"v1.0.4","image:base":"06668724f8daa1866a72d07c44d9c0fe5c42bfd277d47fecc8e4ed037a5d0e84","image:golang":"8cca8e81ef263f66855ba7ad711e884e03e7670276e1cd2011d8f0f3f0eec7e0"}

FROM reg.c5h.io/golang@sha256:8cca8e81ef263f66855ba7ad711e884e03e7670276e1cd2011d8f0f3f0eec7e0 AS build

ARG TAG="v1.0.4"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-notifier /go/src/github.com/greboid/irc-notifier; \
    cd /go/src/github.com/greboid/irc-notifier; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/notifier; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:06668724f8daa1866a72d07c44d9c0fe5c42bfd277d47fecc8e4ed037a5d0e84

COPY --from=build /go/src/github.com/greboid/irc-notifier/main /irc-notifier
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-notifier"]
