# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-news/Dockerfile.gotpl
# BOM: {"github:csmith/irc-news":"v2.0.2","image:base":"84a319cda607326ba5aff905ae01191c4a3c793dbad0e058d9a816e3fb29036f","image:golang":"e76894c73b00fe5e379c1a2c5de8a721160d08ac76476b567e0ef9dccdc7d494"}

FROM reg.c5h.io/golang@sha256:e76894c73b00fe5e379c1a2c5de8a721160d08ac76476b567e0ef9dccdc7d494 AS build

RUN set -eux; \
    git clone --depth=1 -b v2.0.2 --single-branch https://github.com/csmith/irc-news /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:84a319cda607326ba5aff905ae01191c4a3c793dbad0e058d9a816e3fb29036f

COPY --from=build /work/out /irc-news
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-news"]
