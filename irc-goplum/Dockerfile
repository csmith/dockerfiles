# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-goplum/Dockerfile.gotpl
# BOM: {"github:greboid/irc-goplum":"v4.0.2","image:base":"27b64203bd464e43e2314afe482fa80c5c8a3258cada24e4c97b40ab9be1bbd4","image:golang":"626b5a07ae82946548537dd5f89eda905d9b1b3902adb685b518ce2fee400264"}

FROM reg.c5h.io/golang@sha256:626b5a07ae82946548537dd5f89eda905d9b1b3902adb685b518ce2fee400264 AS build

RUN set -eux; \
    git clone --depth=1 -b v4.0.2 --single-branch https://github.com/greboid/irc-goplum /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/goplum; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:27b64203bd464e43e2314afe482fa80c5c8a3258cada24e4c97b40ab9be1bbd4

COPY --from=build /work/out /irc-goplum
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-goplum"]
