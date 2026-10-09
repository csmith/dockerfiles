# Generated from https://github.com/csmith/dockerfiles/blob/master/irc-github/Dockerfile.gotpl
# BOM: {"github:https://github.com/greboid/irc-github":"v4.0.4","image:base":"fbcd09463f4b0824d709a18598637c2f7f7266d0591ca7b61311b2a12cca6f11","image:golang":"f8a9cd7b2ec72536db2748d3d16cffd187062d81b1ce28bddd76f4a9c02f55be"}

FROM reg.c5h.io/golang@sha256:f8a9cd7b2ec72536db2748d3d16cffd187062d81b1ce28bddd76f4a9c02f55be AS build

RUN set -eux; \
    git clone --depth=1 -b v4.0.4 --single-branch https://github.com/greboid/irc-github /work

WORKDIR /work


RUN set -eux; \
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-buildid= -extldflags '-static' " -trimpath  -o out ./cmd/github; \
    go-licenses save ./... --save_path=/notices;




FROM reg.c5h.io/base@sha256:fbcd09463f4b0824d709a18598637c2f7f7266d0591ca7b61311b2a12cca6f11

COPY --from=build /work/out /irc-github
COPY --from=build /notices /notices
ENTRYPOINT ["/irc-github"]
