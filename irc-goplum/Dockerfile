# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-goplum/Dockerfile.gotpl
# BOM: {"github:greboid/irc-goplum":"v4.0.2","image:base":"a09ba1c5a9b2ec41293a8fa0cef935f6ab81978fcbb62271f490ebb07bbf6024","image:golang":"52cd0b4e6cae554978001ed6b324b5eebafffb3292bdbe8c5f71425909db7877"}

FROM reg.c5h.io/golang@sha256:52cd0b4e6cae554978001ed6b324b5eebafffb3292bdbe8c5f71425909db7877 AS build

RUN set -eux; \
    git clone --depth=1 -b v4.0.2 --single-branch https://github.com/greboid/irc-goplum /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/goplum; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:a09ba1c5a9b2ec41293a8fa0cef935f6ab81978fcbb62271f490ebb07bbf6024

COPY --from=build /work/out /irc-goplum
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-goplum"]
