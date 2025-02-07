# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-notifier/Dockerfile.gotpl
# BOM: {"github:greboid/irc-notifier":"v1.0.5","image:base":"27b64203bd464e43e2314afe482fa80c5c8a3258cada24e4c97b40ab9be1bbd4","image:golang":"f3419d02565a6b5e84ff26770fd0c5a40616c78450dd1a3344016c6e42f7658b"}

FROM reg.c5h.io/golang@sha256:f3419d02565a6b5e84ff26770fd0c5a40616c78450dd1a3344016c6e42f7658b AS build

RUN set -eux; \
    git clone --depth=1 -b v1.0.5 --single-branch https://github.com/greboid/irc-notifier /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/notifier; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:27b64203bd464e43e2314afe482fa80c5c8a3258cada24e4c97b40ab9be1bbd4

COPY --from=build /work/out /irc-notifier
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-notifier"]
