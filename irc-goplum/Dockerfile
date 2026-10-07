# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-goplum/Dockerfile.gotpl
# BOM: {"github:https://github.com/greboid/irc-goplum":"v4.0.4","image:base":"fbcd09463f4b0824d709a18598637c2f7f7266d0591ca7b61311b2a12cca6f11","image:golang":"0635ad12ab31fa3ef67aa7be29d89d0ec89f0783e3e58e631e9ce244bf187f47"}

FROM reg.c5h.io/golang@sha256:0635ad12ab31fa3ef67aa7be29d89d0ec89f0783e3e58e631e9ce244bf187f47 AS build

RUN set -eux; \
    git clone --depth=1 -b v4.0.4 --single-branch https://github.com/greboid/irc-goplum /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/goplum; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:fbcd09463f4b0824d709a18598637c2f7f7266d0591ca7b61311b2a12cca6f11

COPY --from=build /work/out /irc-goplum
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-goplum"]
