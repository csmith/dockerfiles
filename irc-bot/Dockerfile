# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-bot/Dockerfile.gotpl
# BOM: {"github:greboid/irc-bot":"v5.0.8","image:base":"66af67568b123f6ec0ff7360f9f2f08e274753f885288e5732d6e0c087c1c1dd","image:golang":"d364769391d3a61e0c701264d331028c3249cb84cc8316dfffe5d882339f1feb"}

FROM reg.c5h.io/golang@sha256:d364769391d3a61e0c701264d331028c3249cb84cc8316dfffe5d882339f1feb AS build

RUN set -eux; \
    git clone --depth=1 -b v5.0.8 --single-branch https://github.com/greboid/irc-bot /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/bot; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:66af67568b123f6ec0ff7360f9f2f08e274753f885288e5732d6e0c087c1c1dd

COPY --from=build /work/out /irc-bot
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-bot"]


EXPOSE 8080
