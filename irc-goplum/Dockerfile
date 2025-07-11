# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-goplum/Dockerfile.gotpl
# BOM: {"github:greboid/irc-goplum":"v4.0.2","image:base":"c97e4a9e2aabe7541fb67a9d165b264bfb9d7cc15d954fc48b760e55156666d8","image:golang":"d51835bb4b9909f172e3d9b495e5aa14564a70a36fabeafba6832cf0d8ae631b"}

FROM reg.c5h.io/golang@sha256:d51835bb4b9909f172e3d9b495e5aa14564a70a36fabeafba6832cf0d8ae631b AS build

RUN set -eux; \
    git clone --depth=1 -b v4.0.2 --single-branch https://github.com/greboid/irc-goplum /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/goplum; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:c97e4a9e2aabe7541fb67a9d165b264bfb9d7cc15d954fc48b760e55156666d8

COPY --from=build /work/out /irc-goplum
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-goplum"]
