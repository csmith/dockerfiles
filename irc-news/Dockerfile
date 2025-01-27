# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-news/Dockerfile.gotpl
# BOM: {"github:csmith/irc-news":"v2.0.2","image:base":"27b64203bd464e43e2314afe482fa80c5c8a3258cada24e4c97b40ab9be1bbd4","image:golang":"b6640cf9e9db2e2132ed4722ea3221bad060651fe0ed20803a17752d88c18c6a"}

FROM reg.c5h.io/golang@sha256:b6640cf9e9db2e2132ed4722ea3221bad060651fe0ed20803a17752d88c18c6a AS build

RUN set -eux; \
    git clone --depth=1 -b v2.0.2 --single-branch https://github.com/csmith/irc-news /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:27b64203bd464e43e2314afe482fa80c5c8a3258cada24e4c97b40ab9be1bbd4

COPY --from=build /work/out /irc-news
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-news"]
