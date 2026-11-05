# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-notifier/Dockerfile.gotpl
# BOM: {"github:https://github.com/greboid/irc-notifier":"v1.0.6","image:base":"fb20443cb04c169a893b5d736e3f466a42ed299664a12d0dcf68aa37acdc62e9","image:golang":"4a8d5555821344f4f81945c99f58174d10913f63c25568befd2c2b58b43f237b"}

FROM reg.c5h.io/golang@sha256:4a8d5555821344f4f81945c99f58174d10913f63c25568befd2c2b58b43f237b AS build

RUN set -eux; \
    git clone --depth=1 -b v1.0.6 --single-branch https://github.com/greboid/irc-notifier /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/notifier; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:fb20443cb04c169a893b5d736e3f466a42ed299664a12d0dcf68aa37acdc62e9

COPY --from=build /work/out /irc-notifier
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-notifier"]
