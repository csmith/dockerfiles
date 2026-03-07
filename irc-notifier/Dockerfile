# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-notifier/Dockerfile.gotpl
# BOM: {"github:greboid/irc-notifier":"v1.0.6","image:base":"84a319cda607326ba5aff905ae01191c4a3c793dbad0e058d9a816e3fb29036f","image:golang":"d5a6ed5cf7aaf950b1976155217f488356e39ed334b93a14418812f169e3fc69"}

FROM reg.c5h.io/golang@sha256:d5a6ed5cf7aaf950b1976155217f488356e39ed334b93a14418812f169e3fc69 AS build

RUN set -eux; \
    git clone --depth=1 -b v1.0.6 --single-branch https://github.com/greboid/irc-notifier /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/notifier; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:84a319cda607326ba5aff905ae01191c4a3c793dbad0e058d9a816e3fb29036f

COPY --from=build /work/out /irc-notifier
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-notifier"]
