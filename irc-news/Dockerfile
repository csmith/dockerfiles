# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-news/Dockerfile.gotpl
# BOM: {"github:csmith/irc-news":"v2.0.2","image:base":"996f481af28d13a3c5069644629fbb02aace36a81a912b688db0771db8669424","image:golang":"051dcbc104b7513206661a70518fe9d361a3761f05bb356fde27e50671b346ef"}

FROM reg.c5h.io/golang@sha256:051dcbc104b7513206661a70518fe9d361a3761f05bb356fde27e50671b346ef AS build

RUN set -eux; \
    git clone --depth=1 -b v2.0.2 --single-branch https://github.com/csmith/irc-news /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:996f481af28d13a3c5069644629fbb02aace36a81a912b688db0771db8669424

COPY --from=build /work/out /irc-news
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-news"]
