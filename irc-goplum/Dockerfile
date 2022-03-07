# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-goplum/Dockerfile.gotpl
# BOM: {"github:greboid/irc-goplum":"v4.0.1","image:base":"f47ed6eddc8491fd9683c0bb66751d16b0dd6f9e920d96c9d028ec0001ae5968","image:golang":"04b65b8002f53c66b0ab66d124fc81bdb06a8a24c69c23f8b58d777c2d21bae4"}

FROM reg.c5h.io/golang@sha256:04b65b8002f53c66b0ab66d124fc81bdb06a8a24c69c23f8b58d777c2d21bae4 AS build

ARG TAG="v4.0.1"

RUN set -eux; \
    git clone --depth=1 -b $TAG --single-branch https://github.com/greboid/irc-goplum /go/src/github.com/greboid/irc-goplum; \
    cd /go/src/github.com/greboid/irc-goplum; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main ./cmd/goplum; \
    go run github.com/google/go-licenses@latest save ./... --save_path=/notices; \
    # Clobber all the timestamps to make the build more reproducible
    touch --date=@0 /go/src/github.com/greboid/irc-goplum/main;

FROM reg.c5h.io/base@sha256:f47ed6eddc8491fd9683c0bb66751d16b0dd6f9e920d96c9d028ec0001ae5968

COPY --from=build /go/src/github.com/greboid/irc-goplum/main /irc-goplum
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-goplum"]
