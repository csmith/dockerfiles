# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-goplum/Dockerfile.gotpl
# BOM: {"github:greboid/irc-goplum":"v4.0.2","image:base":"fa9cdd83581affde03cb2234170235ff287588e4da3786133d72846838210621","image:golang":"6cc5093fb6259e4e12d7a964725c55295eed06383ecfcab4c2a96220cb39e5da"}

FROM reg.c5h.io/golang@sha256:6cc5093fb6259e4e12d7a964725c55295eed06383ecfcab4c2a96220cb39e5da AS build

ARG TAG="v4.0.2"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-goplum /go/src/github.com/greboid/irc-goplum; \
    cd /go/src/github.com/greboid/irc-goplum; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/goplum; \
    go-licenses save ./... --save_path=/notices;

FROM reg.c5h.io/base@sha256:fa9cdd83581affde03cb2234170235ff287588e4da3786133d72846838210621

COPY --from=build /go/src/github.com/greboid/irc-goplum/main /irc-goplum
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-goplum"]
