# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-news/Dockerfile.gotpl
# BOM: {"github:csmith/irc-news":"v2.0.2","image:base":"524e2a72907aa209ab61ecec7762951bd727b2c31592c5d3c1b2acce5f33c83f","image:golang":"1960f6b4b2b1e2a683843f94da41dec035948604b17c7baa08a91c39c373e636"}

FROM reg.c5h.io/golang@sha256:1960f6b4b2b1e2a683843f94da41dec035948604b17c7baa08a91c39c373e636 AS build

RUN set -eux; \
    git clone --depth=1 -b v2.0.2 --single-branch https://github.com/csmith/irc-news /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:524e2a72907aa209ab61ecec7762951bd727b2c31592c5d3c1b2acce5f33c83f

COPY --from=build /work/out /irc-news
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-news"]
